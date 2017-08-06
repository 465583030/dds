package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/riclava/distribute-download-system/master/config"
	"github.com/riclava/distribute-download-system/master/handler"
)

func main() {

	cfg := config.Config{}
	cfg.Addr = "0.0.0.0"
	cfg.Port = 8080
	cfg.Token = "rictina"

	apiHandler, err := handler.CreateAPIHandler(cfg)
	if err != nil {
		panic(err)
	}

	http.Handle("/api/", apiHandler)

	addr := fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)
	go log.Fatal(http.ListenAndServe(addr, nil))

	select {}
}
