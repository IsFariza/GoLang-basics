package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/FarizaIsmagambetova/Assignment2/internal/api"
	"github.com/FarizaIsmagambetova/Assignment2/internal/model"
	"github.com/FarizaIsmagambetova/Assignment2/internal/store"
	"github.com/FarizaIsmagambetova/Assignment2/internal/worker"
	"github.com/gorilla/mux"
)

func main() {

	repo := store.NewRepository[string, *model.Task]()
	wp := worker.NewWorkerPool(repo, 3, 10)
	taskHandler := api.NewTaskHandler(repo, wp)
	statsHandler := api.NewStatsHandler(repo)

	//router tells server how to handle HTTP requests
	r := mux.NewRouter()
	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", taskHandler.GetByID).Methods("GET")
	r.HandleFunc("/stats", statsHandler.GetStats).Methods("GET")

	//HTTP server configuration
	server := &http.Server{
		Addr:    ":8080", //server listens on this port
		Handler: r,       //requests passed to router
	}

	//monitoring goroutine
	stopMonitoring := make(chan struct{})
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				statusCount := repo.StatusCount
				log.Printf("MONITORING:")
				log.Printf("PENDING=%d, IN_PROGRESS=%d, COMPLETED=%d",
					statusCount["PENDING"], statusCount["IN_PROGRESS"], statusCount["COMPLETED"])
			case <-stopMonitoring: //when channel is closed (not blocked anymore)
				return
			}
		}
	}()

	//shutdown receives os signals, buffered channel, receives 1 signal and shuts down the program
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt) //send SIGINT into channel

	//start HTTP server in goroutine
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	//waits until shutdown channel gets a signal
	<-shutdown
	log.Println("Shutting down server...")
	//context that lasts 5 sec,
	//used to limit the time of shutdown to finish other requests before stopping
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()              //cleans the context
	err := server.Shutdown(ctx) //actual shutdown
	if err != nil {
		log.Println(err)
	}

	wp.Stop()             //stop worker pool
	close(stopMonitoring) //stop monitoring goroutine
	log.Println("Server stopped")
}
