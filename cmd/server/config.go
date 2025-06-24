package main

import (
	"github.com/nickapopolus/waystone-api-gateway/internal/data"
	"github.com/nickapopolus/waystone-api-gateway/internal/rest"
)

type Config struct {
	Router *rest.Router
	Models *data.Models
}
