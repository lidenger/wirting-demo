package main

import (
	"flag"
	"fmt"
	"jwtenhance/config"
	"jwtenhance/router"
	"jwtenhance/service"
	"jwtenhance/store"
	"log"
	"net/http"
)

func main() {
	conf := config.Initialize()
	db := store.Initialize(conf)
	service.Initialize(db)
	service.ServerSvcIns.LoadAll()
	service.JwtSvcIns.LoadAllSecret()
	service.JwtSvcIns.LoadAllDeactiveJwt()
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
