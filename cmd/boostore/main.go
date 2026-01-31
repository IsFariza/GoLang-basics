package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BlackHole55/software-store-final/config"
	delivery "github.com/BlackHole55/software-store-final/internal/delivery/http"
	"github.com/BlackHole55/software-store-final/internal/repositories/mongodb"
	"github.com/BlackHole55/software-store-final/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	client := config.ConnectDB()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	gameRepo := mongodb.NewGameRepository(client)
	companyRepo := mongodb.NewCompanyRepository(client)

	gameUC := usecase.NewGameUseCase(gameRepo, companyRepo)
	companyUC := usecase.NewCompanyUsecase(companyRepo)

	gameHandler := delivery.NewGameHandler(gameUC)
	companyHandler := delivery.NewCompanyHandler(companyUC)

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	delivery.RegisterRoutes(router, gameHandler, companyHandler)

	port := "8080"
	log.Printf("Server starting on port %s", port)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down gracefully...")

	disconnectCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Disconnect(disconnectCtx); err != nil {
		log.Fatalf("MongoDB Disconnect Error: %v", err)
	}

	log.Println("Database connection closed. Goodbye!")

}
