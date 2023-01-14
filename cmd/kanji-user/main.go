package main

import (
	"context"
	"fmt"
	"github.com/Hanekawa-chan/kanji-user/internal/app"
	"github.com/Hanekawa-chan/kanji-user/internal/database"
	"github.com/Hanekawa-chan/kanji-user/internal/httpserver"
	"github.com/Hanekawa-chan/kanji-user/internal/services/auth"
	"github.com/Hanekawa-chan/kanji-user/internal/version"
	"github.com/rs/zerolog"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//Print version and commit sha
	log.Println("Loading Mailing - v", version.Version, "| Commit:", version.Commit)

	// Parse all configs form env
	cfg, err := app.Parse()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	level, err := zerolog.ParseLevel(cfg.Logger.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	logger := zerolog.New(os.Stdout).Level(level)

	zl := &logger

	db, err := database.NewAdapter(zl, cfg)
	if err != nil {
		zl.Fatal().Err(err).Msg("Database init")
	}

	authClient := auth.NewAuthClient(zl, cfg)

	service := app.NewService(zl, cfg, db, authClient)
	httpServerAdapter := httpserver.NewAdapter(zl, cfg, service)

	// Channels for errors and os signals
	stop := make(chan error, 1)
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, syscall.SIGINT, syscall.SIGTERM)

	// Receive errors form start bot func into error channel
	go func(stop chan<- error) {
		stop <- httpServerAdapter.ListenAndServe()
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if err := httpServerAdapter.Shutdown(ctx); err != nil {
		zl.Error().Err(err).Msg("Error shutting down the HTTP server!")
	}

	time.Sleep(time.Second * 2)

	zl.Info().Msg("Shutdown - success")
}
