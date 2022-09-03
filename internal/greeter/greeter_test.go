package greeter_test

import (
	"testing"

	"github.com/makasim/go-app-template/internal/greeter"
	"github.com/stretchr/testify/require"
)

func TestGreeter_Greet(main *testing.T) {
	// Test suit setup here, setUp(t *testing.T) function, structs and so on.

	main.Run("Hi", func(t *testing.T) {
		g := greeter.New("Hi")

		require.Equal(t, `Hi John!`, g.Greet("John"))
	})

	main.Run("Hello", func(t *testing.T) {
		g := greeter.New("Hello")

		require.Equal(t, `Hello John!`, g.Greet("John"))
	})
}
