package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/GabriellaAmah/go-url-shortner/config"
	"github.com/GabriellaAmah/go-url-shortner/router"
	"github.com/GabriellaAmah/go-url-shortner/setup"
	"github.com/gorilla/mux"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	setup.AppConnectionsSetUp()

	r := mux.NewRouter()

	router.RegisterRouters(r)

	srv := &http.Server{
		Addr:         config.EnvData.PORT,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 10,
		Handler:      r,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

	fmt.Println("server is on")
}
