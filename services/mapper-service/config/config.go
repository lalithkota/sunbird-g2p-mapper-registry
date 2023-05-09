package config

import "github.com/jinzhu/configor"

var Config = struct {
	Registry struct {
		Url string `default:"http://localhost:8081/" required:"true" env:"REGISTRY_URL"`
	}
	CallbackService struct {
		Url string `default:"http://localhost:3001/v0.1.0/mapper/" required:"true" env:"CALLBACK_SERVICE_URL"`
	}
}{}

func init() {
	err := configor.Load(&Config, "./config/default.yml") //"config/application.yml"
	if err != nil {
		panic("Unable to read configurations")
	}
}
