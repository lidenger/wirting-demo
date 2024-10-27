package main

import (
	"flag"
	"fmt"
	"jwtbase/config"
	"jwtbase/router"
	"jwtbase/service"
	"jwtbase/store"
	"log"
	"net/http"
)

func main() {
	conf := config.Initialize()
	db := store.Initialize(conf)
	service.Initialize(db)
	service.JwtSvcIns.LoadAllSecret()
	service.AkSkSvcIns.LoadAllAkSk()
	port := flag.Int("port", conf.Server.Port, "http server port")
	g := router.Initialize()
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
		Handler: g,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
