package rest

import "github.com/nickapopolus/waystone-api-gateway/internal/data"

type Router struct {
	Models *data.Models
}

func NewRouter(models *data.Models) *Router {
	return &Router{
		Models: models,
	}
}
