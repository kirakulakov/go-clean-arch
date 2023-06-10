package main

import (
	"log"

	"github.com/kirakulakov/go-clean-arch/config"
	"github.com/kirakulakov/go-clean-arch/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)

}
