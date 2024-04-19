package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"structured-api/api/router"
	"structured-api/internal/config"
	"structured-api/internal/server"
	"structured-api/util/logger"
	"syscall"
	"time"
)

func main() {
	// Setup configuration
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	// Setup logger
	l := logger.New(c.IsDebug)

	// Setup router
	r := router.New(l)

	// Setup HTTP Server
	s := server.NewServer(c, r)

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, shutdownCancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer shutdownCancel()

		// Log graceful shutdown timeout
		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				l.Fatal().Msg("graceful shutdown timed out... forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		l.Info().Msg("exit signal received. initiating graceful shutdown.")
		err := s.HttpServer.Shutdown(shutdownCtx)

		// Log in case of error shutting down server
		if err != nil {
			l.Fatal().Msg(err.Error())
		}

		// Stop Server Context
		serverStopCtx()
	}()

	// Start HTTP Server
	l.Info().Msgf("listening on: http://localhost:%d 🚀", config.GetConfig().Port)
	err = s.HttpServer.ListenAndServe()

	// Log in case of error starting server
	if err != http.ErrServerClosed {
		l.Fatal().Msgf("failed to start server: %s", err.Error())
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
	l.Info().Msg("server stopped")
}
