package main

import (
	"context"
	"flag"
	"github.com/MatthiasKunnen/goprolocal"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var addressFlag = flag.String("address", ":80", "The address to listen on. "+
	"If a port other than 80 is used, do not forget that HTTP runs on port 80 and the GoPro "+
	"will perform its connectivity check to http://ip:80 so port forwarding will be required.")

func main() {
	flag.Parse()
	log.Printf("Binding on %s", *addressFlag)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	errChan := make(chan error)
	srv := goprolocal.Start(*addressFlag, errChan)

	select {
	case err := <-errChan:
		log.Fatal(err)
	case sig := <-stop:
		log.Printf("Received %s, shutting down...\n", sig.String())
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down server: %v\n", err)
		}
	}
}
