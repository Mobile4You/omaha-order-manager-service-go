package db

import (
	"fmt"
	"os"

	"github.com/arthurstockler/omaha-order-manager-service-go/config"

	"github.com/jinzhu/gorm"
	//
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Open exported
func Open(config *config.YmlDatabase) *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Dbname)

	conn, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	conn.DB().SetMaxOpenConns(5)
	conn.LogMode(true)
	if os.Getenv("DEBUG") == "true" {
		conn.LogMode(true)
	}

	return conn
}
