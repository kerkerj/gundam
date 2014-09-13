package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		// This is my device port, change to yours
		port = "/dev/tty.Sphero-RGG-AMP-SPP"
	}

	s := NewSphero("Gundam", port)
	defer s.Stop()

	go func() {
		api := NewApi(s)
		http.ListenAndServe(":3000", api.Handler())
	}()

	s.Start()
}
