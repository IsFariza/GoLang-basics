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
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/mongo/mongodriver"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	mongoV1 "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	client := config.ConnectDB()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	router := gin.Default()

	mongoURI := os.Getenv("MONGODB_URI")

	clientV1, _ := mongoV1.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	sessionCollV1 := clientV1.Database("softwarestore").Collection("sessions")
	store := mongodriver.NewStore(sessionCollV1, 3600, true, []byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	gameRepo := mongodb.NewGameRepository(client)
	companyRepo := mongodb.NewCompanyRepository(client)
	emulationRepo := mongodb.NewEmulationRepository(client)
	userRepo := mongodb.NewUserRepository(client)
	reviewRepo := mongodb.NewReviewRepository(client)
	purchaseRepo := mongodb.NewPurchaseRepo(client)

	gameUC := usecase.NewGameUseCase(gameRepo, companyRepo, emulationRepo, reviewRepo, userRepo)
	companyUC := usecase.NewCompanyUsecase(companyRepo)
	emulationUC := usecase.NewEmulationUsecase(emulationRepo)
	userUC := usecase.NewUserUseCase(userRepo)
	reviewUC := usecase.NewReviewUsecase(reviewRepo)
	purchaseUC := usecase.NewPurchaseUsecase(purchaseRepo)

	gameHandler := delivery.NewGameHandler(gameUC)
	companyHandler := delivery.NewCompanyHandler(companyUC)
	emulationHandler := delivery.NewEmulationHandler(emulationUC)
	userHandler := delivery.NewUserHandler(userUC)
	reviewHandler := delivery.NewReviewHandler(reviewUC)
	purchaseHandler := delivery.NewPurchaseHandler(purchaseUC)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	delivery.RegisterRoutes(router, gameHandler, companyHandler, emulationHandler, userHandler, purchaseHandler, reviewHandler)

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
