package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// YmlHTTP exported
type YmlHTTP struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

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
	HTTP     YmlHTTP     `yaml:"http"`
}

// YmlEnvironment exported
type YmlEnvironment struct {
	Development YmlConfig `yaml:"development"`
	Production  YmlConfig `yaml:"production"`
}

// LoadingConfig exported
func LoadingConfig(file string) (*YmlDatabase, *YmlHTTP) {

	if _, err := os.Stat(file); os.IsNotExist(err) {
		panic(fmt.Errorf("Database file config not found for the environment %v", os.Getenv("GO_ENV")))
	}

	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var env YmlEnvironment
	err = yaml.Unmarshal(configFile, &env)
	if err != nil {
		panic(err)
	}

	switch os := os.Getenv("GO_ENV"); strings.ToLower(os) {
	case "development":
		return &env.Development.Database, &env.Development.HTTP
	case "production":
		return &env.Production.Database, &env.Production.HTTP
	default:
		panic(fmt.Errorf("Environment %v not found YML config", os))
	}
}
