package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"mtls/config"
	"mtls/router"
	"mtls/service"
	"mtls/store"
	"net/http"
	"os"
)

func main() {
	conf := config.Initialize()
	db := store.Initialize(conf)
	service.Initialize(db)
	port := flag.Int("port", conf.Server.Port, "http server port")
	g := router.Initialize()

	pool := x509.NewCertPool()
	cert, err := os.ReadFile("./config/crt/clientroot.crt")
	if err != nil {
		panic(err)
	}
	pool.AppendCertsFromPEM(cert)
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
		Handler: g,
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}
	err = server.ListenAndServeTLS("./config/crt/s2.crt", "./config/crt/s2.key")
	if err != nil {
		log.Fatal(err)
	}
}
