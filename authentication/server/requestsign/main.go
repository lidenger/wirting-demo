package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"requestsign/config"
	"requestsign/router"
	"requestsign/service"
	"requestsign/store"
)

func main() {
	conf := config.Initialize()
	db := store.Initialize(conf)
	service.Initialize(db)
	port := flag.Int("port", conf.Server.Port, "http server port")
	g := router.Initialize()
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
		Handler: g,
	}
	fmt.Printf("server starting on port %d\n", *port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
