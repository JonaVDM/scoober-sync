package gotify_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/jonavdm/scoober-sync/internal/gotify"
	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	t.Run("Give an error when missing values", func(t *testing.T) {
		gf := Gotify{}

		gf.Send("Foo", "Lorem Ipsum", 1)
	})

	t.Run("Sends a POST Request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, "/message", req.URL.String())
			assert.Equal(t, "POST", req.Method)
		}))
		defer server.Close()

		gf := Gotify{
			URL:   server.URL,
			Token: "Foo Bar Token",
		}

		err := gf.Send("Foo", "Lorem Ipsum", 1)
		assert.Nil(t, err)
	})
}
