package main

import (
	"github.com/rs/zerolog"
	"log"
	"os"
	"os/signal"
	"syscall"
	"user/internal/app"
	"user/internal/app/config"
	"user/internal/database"
	"user/internal/grpcserver"
	"user/internal/version"
)

func main() {
	//Print version and commit sha
	log.Println("Loading User - v", version.Version, "| Commit:", version.Commit)

	// Parse all configs form env
	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	level, err := zerolog.ParseLevel(cfg.Logger.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	logger := zerolog.New(os.Stdout).Level(level)

	zl := &logger

	db, err := database.NewAdapter(zl, cfg.DB)
	if err != nil {
		zl.Fatal().Err(err).Msg("Database init")
	}

	service := app.NewService(zl, cfg.User, db)
	grpcServer := grpcserver.NewAdapter(zl, cfg.GRPCServer, service)

	// Channels for errors and os signals
	stop := make(chan error, 1)
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, syscall.SIGINT, syscall.SIGTERM)

	// Receive errors form start bot func into error channel
	go func(stop chan<- error) {
		stop <- grpcServer.ListenAndServe()
	}(stop)

	// Blocking select
	select {
	case sig := <-osSig:
		zl.Info().Msgf("Received os syscall signal %v", sig)
	case err := <-stop:
		zl.Error().Err(err).Msg("Received Error signal")
	}

	// Shutdown code
	zl.Info().Msg("Shutting down...")

	grpcServer.Shutdown()

	zl.Info().Msg("Shutdown - success")
}
