package main

import (
	"esb/server"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(`.env`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	server := server.NewServer()
	server.Middlewares()
	server.Routes()
	server.Run()
}
