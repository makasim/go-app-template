//go:build func_test

package api_test

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/makasim/go-app-template/tests/pkg/testapp"
	"github.com/stretchr/testify/require"
)

func TestGreetHandler(main *testing.T) {
	main.Run("OK", func(t *testing.T) {
		app := testapp.New()
		app.Config.Greet = `TestHello`

		defer app.Start(t)()
		req, err := http.NewRequest(
			http.MethodPost,
			fmt.Sprintf("http://%s/", app.Config.ServerAddr),
			strings.NewReader(`John`),
		)
		require.NoError(t, err)

		resp, err := (&http.Client{}).Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Equal(t, `TestHello John!`, string(b))
	})
}
