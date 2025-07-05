package main

import (
	"fmt"
	"github.com/nickapopolus/waystone-api-gateway/internal/data"
	"github.com/nickapopolus/waystone-api-gateway/internal/rest"
	v1 "github.com/nickapopolus/waystone-api-gateway/proto/urlservice/v1"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

const WEBPORT = 8080

func main() {

	app := Config{}
	app.Models = data.NewModels()
	fmt.Println(os.Getenv("URL_SERVICE_ADDR"))
	urlconn, err := grpc.Dial(os.Getenv("URL_SERVICE_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to connect to URL service:", err)
	}
	fmt.Println("Connected to URL GRPC service")
	defer urlconn.Close()
	URLClient := v1.NewURLServiceClient(urlconn)
	app.Router = rest.NewRouter(URLClient)

	fmt.Println("Starting API Gateway on port 8080")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", WEBPORT),
		Handler: app.Router.Routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
