package httplogging

import (
	"github.com/gorilla/handlers"
	"github.com/pkg/errors"
	"net/http"
	"os"
)

type ApacheLoggingBehaviour struct {
	Filename string
}

func (c *ApacheLoggingBehaviour) Wrap(wrapped http.Handler) (http.Handler, error) {
	writer, err := os.OpenFile(c.Filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return handlers.CombinedLoggingHandler(writer, wrapped), nil
}
