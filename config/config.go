package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/arthurstockler/omaha-order-manager-service-go/db"
	"gopkg.in/yaml.v2"
)

// LoadingConfig exported
func LoadingConfig(file string) db.YmlEnvironment {

	if _, err := os.Stat(file); os.IsNotExist(err) {
		panic(fmt.Errorf("Database file config not found for the environment %v", os.Getenv("GO_ENV")))
	}

	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var configuration db.YmlEnvironment
	err = yaml.Unmarshal(configFile, &configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}
