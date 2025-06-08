package main

import (
	"fmt"
	"log"

	"github.com/bulkashmak/gator-cli/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("bulkashmak")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	fmt.Printf("Read config again: %+v\n", cfg)
}
