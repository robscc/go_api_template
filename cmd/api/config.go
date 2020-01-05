package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	BindAddr string `default:"0.0.0.0:8082" split_words:"true"`

	Debug bool `default:"false" split_words:"true"`
}

var cfg *Config

//
func ParseEnvConfig() error {
	cfg = new(Config)
	err := envconfig.Process("API", cfg)
	fmt.Printf("get config %+v\n", *cfg)
	return err
}
