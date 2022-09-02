package app

import (
	"context"
	"net/http"
	"time"

	"github.com/makasim/go-app-template/internal/api/greethandler"
	"github.com/makasim/go-app-template/internal/greeter"
	"go.uber.org/zap"
)

type Config struct {
	Greet      string `envconfig:"GREET" default:"Hello"`
	ServerAddr string `envconfig:"SERVER_ADDR" required:"true"`

	Now func() time.Time `ignored:"true"`
}

type App struct {
	cfg Config
	l   *zap.Logger
}

func New(cfg Config, l *zap.Logger) *App {
	return &App{
		cfg: cfg,
		l:   l,
	}
}

func (app *App) Run(ctx context.Context) error {
	app.l.Info("starting")
	defer app.l.Info("stopped")

	g := greeter.New(app.cfg.Greet)

	s := &http.Server{
		Addr: app.cfg.ServerAddr,
	}
	s.Handler = greethandler.New(g, app.l)

	go func() {
		if err := s.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
			app.l.Error("server serve failed", zap.Error(err))
		}
	}()

	app.l.Info("started")
	<-ctx.Done()
	app.l.Info("stopping")

	sCtx, sCtxCancel := context.WithTimeout(context.Background(), time.Second*10)
	defer sCtxCancel()

	if err := s.Shutdown(sCtx); err != nil {
		return err
	}

	// gracefull shutdown

	return nil
}
