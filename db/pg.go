package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	//
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// YmlDatabase exported
type YmlDatabase struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

// YmlConfig exported
type YmlConfig struct {
	Database YmlDatabase `yaml:"database"`
}

// YmlEnvironment exported
type YmlEnvironment struct {
	Development YmlConfig `yaml:"development"`
	Production  YmlConfig `yaml:"production"`
}

// Open exported
func Open(config *YmlDatabase) *gorm.DB {

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
