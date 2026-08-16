package main

import (
	"clothing_store_api/internal/config"
)

func main() {
	router := config.InitializeEverything()
	router.Run(":8888")
}
