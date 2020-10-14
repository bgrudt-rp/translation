package main

import (
	"log"
	"translation/api"
	"translation/tools/postgres"
)

func main() {

	err := postgres.BuildDB()

	if err != nil {
		log.Fatal(err)
	}

	api.StartAPI()

	if err != nil {
		log.Fatal(err)
	}
}
