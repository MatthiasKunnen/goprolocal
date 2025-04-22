package goprolocal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type Shutdown interface {
	Shutdown(ctx context.Context) error
}

// Start starts the GoPro connectivity handler.
// It will start a goroutine to respond to connectivity checks at the given address.
// Examples of address; :80, 192.168.23.1:80.
// The errorChan will be written to when the connectivity handler fails.
func Start(address string, errorChan chan<- error) Shutdown {
	server := &http.Server{Addr: address}

	http.HandleFunc("/v1/hello.html", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Success")
	})

	go func() {
		err := server.ListenAndServe()
		switch {
		case errors.Is(err, http.ErrServerClosed):
			return
		default:
			errorChan <- err
		}
	}()

	return server
}
