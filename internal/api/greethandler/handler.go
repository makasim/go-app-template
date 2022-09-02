package greethandler

import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

type greeter interface {
	Greet(name string) string
}

type Handler struct {
	g greeter
	l *zap.Logger
}

func New(g greeter, l *zap.Logger) *Handler {
	return &Handler{
		g: g,
		l: l,
	}
}

func (h *Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	h.l.Info("handling request")
	defer h.l.Info("handled request")

	b, err := io.ReadAll(req.Body)
	if err != nil {
		h.l.Error("read request body failed", zap.Error(err))
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	greet := []byte(h.g.Greet(string(b)))

	resp.WriteHeader(http.StatusOK)
	if _, err := resp.Write(greet); err != nil {
		h.l.Error("greater: greet failed", zap.Error(err))
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
}
