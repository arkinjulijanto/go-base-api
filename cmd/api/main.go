package main

import (
	"log"

	"github.com/arkinjulijanto/go-base-api/boot"
)

func main() {
	err := boot.InitApi()
	if err != nil {
		log.Fatalf("failed to initialize API: %v", err)
	}
}
