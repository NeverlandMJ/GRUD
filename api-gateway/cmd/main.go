package main

import (
	"github.com/NeverlnadMJ/GRUD/api-gateway/api"
	"github.com/NeverlnadMJ/GRUD/api-gateway/service"
)

const (
	collecterURL = "localhost:8001"
)

func main() {
	h := api.NewHandler(service.NewService(collecterURL))

	router := api.NewRouter(h)

	router.Run()
}
