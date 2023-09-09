package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AkashGit21/mta-hosting-optimizer/internal/api"
	"github.com/AkashGit21/mta-hosting-optimizer/internal/utilities"
	"github.com/joho/godotenv"
)

// Initiate/Configure the server as per your desired specifications from env file
func NewServer() (*http.Server, error) {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	api, err := api.New()
	if err != nil {
		return nil, err
	}

	srvHost := utilities.GetEnvValue("APP_HOST", "localhost")
	srvPort := utilities.GetEnvValue("APP_PORT", "8081")
	srvAddress := fmt.Sprintf("%s:%v", srvHost, srvPort)
	log.Println("Configuring Server at address ", srvAddress)
	srv := http.Server{
		Addr:    srvAddress,
		Handler: http.TimeoutHandler(api, 1*time.Second, "request timed out"),
		// Read and Write will Timeout after 2s for a single request.
		ReadTimeout:  time.Duration(2 * time.Second),
		WriteTimeout: time.Duration(2 * time.Second),
	}

	return &srv, nil
}

// Starts the application server
func StartServer(srv *http.Server) {
	utilities.DebugLog("Starting Server...")

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		log.Println("Shutting down the server gracefully...")
		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			utilities.ErrorLog("HTTP server Shutdown: ", err)
			return
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		utilities.ErrorLog("HTTP server ListenAndServe: ", err)
		return
	}

	<-idleConnsClosed
}
