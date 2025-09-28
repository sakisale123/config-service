package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sakisale123/config-service/internal/config"
	"github.com/sakisale123/config-service/internal/middleware"
	"golang.org/x/time/rate"
)

func main() {
	router := mux.NewRouter()
	configService := config.NewConfigService()
	configHandler := config.NewConfigHandler(configService)
	limiter := rate.NewLimiter(1, 3)

	router.HandleFunc("/configs", configHandler.CreateConfigurationHandler).Methods("POST")
	router.HandleFunc("/configs/{id}/versions/{version}", configHandler.GetConfigurationHandler).Methods("GET")
	router.HandleFunc("/configs/{id}/versions/{version}", configHandler.DeleteConfigurationHandler).Methods("DELETE")
	router.HandleFunc("/search/configs", configHandler.SearchConfigurationsHandler).Methods("GET") // nova ruta

	router.HandleFunc("/groups", configHandler.CreateConfigurationGroupHandler).Methods("POST")
	router.HandleFunc("/groups/{id}/versions/{version}", configHandler.GetConfigurationGroupHandler).Methods("GET")
	router.HandleFunc("/groups/{id}/versions/{version}", configHandler.DeleteConfigurationGroupHandler).Methods("DELETE")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: middleware.RateLimitMiddleware(limiter)(router),
	}

	go func() {
		log.Println("Server startovan na portu 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Greška pri pokretanju servera: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server se gasi...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server je nasilno ugašen:", err)
	}

	log.Println("Server je uspešno ugašen")
}
