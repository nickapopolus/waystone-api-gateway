package rest

import v1 "github.com/nickapopolus/waystone-api-gateway/proto/urlservice/v1"

type Router struct {
	URLClient v1.URLServiceClient
}

func NewRouter(u v1.URLServiceClient) *Router {
	return &Router{
		URLClient: u,
	}
}
