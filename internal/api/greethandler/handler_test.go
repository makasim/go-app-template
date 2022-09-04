package greethandler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/makasim/go-app-template/internal/api/greethandler"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestHandler_ServeHTTP(main *testing.T) {
	setUp := func(t *testing.T) (*greethandler.Handler, *greeterMock) {
		g := &greeterMock{}

		l := zap.NewNop()

		h := greethandler.New(g, l)

		t.Cleanup(func() {
			g.AssertExpectations(t)
		})

		return h, g
	}

	main.Run("OK", func(t *testing.T) {
		h, g := setUp(t)

		g.On("Greet", "John").Return("Hi John!").Once()

		r := httptest.NewRequest("POST", "/foo", strings.NewReader(`John`))
		rw := httptest.NewRecorder()

		h.ServeHTTP(rw, r)

		resp := rw.Result()
		require.Equal(t, http.StatusOK, resp.StatusCode)

		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, `Hi John!`, string(b))
	})
}

type greeterMock struct {
	mock.Mock
}

func (m *greeterMock) Greet(name string) string {
	args := m.Called(name)
	return args.Get(0).(string)
}
