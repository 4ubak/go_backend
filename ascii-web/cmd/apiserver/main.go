package main

import (
	// "fmt"
	// "net/http"
	"log"
	"flag"
	"https://github.com/BurntSushi/toml"
	"https://github.com/4ubak/go_backend/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}
func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}