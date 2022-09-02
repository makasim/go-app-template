package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"

	"github.com/makasim/go-app-template/internal/app"
)

func main() {
	time.Local = time.UTC

	l, _ := zap.NewProduction()
	defer l.Sync()
	l = l.With(zap.String("app", "greeter"))

	cfg := app.Config{
		Now: time.Now,
	}

	if err := envconfig.Process("", &cfg); err != nil {
		l.Error("envconfig process failed", zap.Error(err))
		l.Sync()
		os.Exit(1)
	}

	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	go handleSignals(ctxCancel, l)

	if err := app.New(cfg, l).Run(ctx); err != nil {
		l.Error("app run failed", zap.Error(err))
		l.Sync()
		os.Exit(1)
	}
}

func handleSignals(cancel context.CancelFunc, l *zap.Logger) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	<-signals
	l.Info("got signal; canceling context")
	cancel()

	<-signals
	l.Warn("got second signal; force exiting")
	l.Sync()
	os.Exit(1)
}