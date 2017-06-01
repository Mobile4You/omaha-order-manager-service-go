package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/arthurstockler/omaha-order-manager-service-go/config"
	"github.com/arthurstockler/omaha-order-manager-service-go/db"
	"github.com/arthurstockler/omaha-order-manager-service-go/routes"
	"github.com/arthurstockler/omaha-order-manager-service-go/usecase"
)

func main() {

	// carregando arquivo de config
	fmt.Println("Carregando arquivo de configuração")
	configDB := config.LoadingConfig("database.yml")

	//abrindo conn
	fmt.Println("Abrindo conn com banco dados")
	conn := db.Open(configDatabase(configDB))

	fmt.Println("Iniciando http server")
	r := &routes.Router{Use: &usecase.UseCase{DB: conn}}
	r.ServerStart()
}

func configDatabase(env db.YmlEnvironment) *db.YmlDatabase {
	switch os := os.Getenv("GO_ENV"); strings.ToLower(os) {
	case "development":
		return &env.Development.Database
	case "production":
		return &env.Production.Database
	default:
		panic(fmt.Errorf("Environment %v not found YML config", os))
	}
}
