package main

import (
	"log"

	"github.com/tehrelt/volkswagen-reference-api/internal/api"
)

func main() {

	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
}