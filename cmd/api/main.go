package main

import (
	"log"

	"github.com/ameernormie/go-api-template/pkg/api"
	"github.com/ameernormie/go-api-template/pkg/models"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := models.Connect()

	if err != nil {
		log.Fatal(err)
	}

	app := &api.App{
		Db: db,
	}

	app.Start()
}
