package main

import (
	"github.com/NeverlnadMJ/GRUD/api-gateway/api"
	"github.com/NeverlnadMJ/GRUD/api-gateway/service"
)

const (
	collecterURL = "localhost:8001"
	grudURL      = "localhost:8002"
)

func main() {
	h := api.NewHandler(service.NewService(collecterURL, grudURL))

	router := api.NewRouter(h)

	router.Run()
}
