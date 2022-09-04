package testapp

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/makasim/go-app-template/internal/app"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type Testapp struct {
	Config app.Config
}

func New() *Testapp {
	return &Testapp{
		Config: app.Config{
			Greet:      "Hi",
			ServerAddr: "127.0.0.1:9999",
			Now: func() time.Time {
				return time.Unix(123, 0)
			},
		},
	}
}

func (ta *Testapp) Start(t *testing.T) func() {
	ctx, ctxCancel := context.WithCancel(context.Background())

	runRes := make(chan error, 1)
	go func() {
		a := app.New(ta.Config, zap.NewNop())
		runRes <- a.Run(ctx)
	}()

	ta.ready(t)

	return func() {
		ctxCancel()
		require.NoError(t, <-runRes)
	}
}

func (ta *Testapp) ready(t *testing.T) {
	try := time.NewTicker(time.Millisecond * 200)
	defer try.Stop()

	timeout := time.NewTimer(time.Second * 2)
	defer timeout.Stop()

	for {
		select {
		case <-try.C:
			conn, err := net.Dial("tcp", ta.Config.ServerAddr)
			if err != nil {
				continue
			}
			conn.Close()
			return
		case <-timeout.C:
			t.Fatal("app not ready")
		}
	}
}
