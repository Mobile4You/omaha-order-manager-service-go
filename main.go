package main

import (
	"fmt"

	"github.com/arthurstockler/omaha-order-manager-service-go/config"
	"github.com/arthurstockler/omaha-order-manager-service-go/db"
	"github.com/arthurstockler/omaha-order-manager-service-go/routes"
	"github.com/arthurstockler/omaha-order-manager-service-go/usecase"
)

func main() {

	fmt.Println("Loading configuration file")
	envdb, rhttp := config.LoadingConfig("config.yml")

	fmt.Println("Opening connection to database")
	conn := db.Open(envdb)

	r := &routes.Router{
		Use: &usecase.UseCase{
			DB: &usecase.Store{Conn: conn},
		},
		Rhttp: rhttp,
	}
	r.ServerStart()
}
