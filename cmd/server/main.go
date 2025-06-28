package main

import (
	"fmt"
	"github.com/nickapopolus/waystone-api-gateway/internal/data"
	"github.com/nickapopolus/waystone-api-gateway/internal/rest"
	"log"
	"net/http"
)

const WEBPORT = 8080

func main() {

	app := Config{}
	app.Models = data.NewModels()
	app.Router = rest.NewRouter()

	fmt.Println("Starting API Gateway on port 8080")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", WEBPORT),
		Handler: app.Router.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
