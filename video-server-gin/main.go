package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/video-server-gin/server/routes"
)

func main() {
	os.Setenv("PORT", "8081")
	addr := ":" + os.Getenv("PORT")
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":2112", nil))
	}()
	log.Printf("listening on %s", addr)
	log.Fatal(routes.Init().Run(addr))
}
