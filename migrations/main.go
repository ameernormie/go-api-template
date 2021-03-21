package main

import (
	"log"
	"os"

	"github.com/ameernormie/go-api-template/pkg/models"
	"github.com/go-pg/pg/v10"
	migrations "github.com/robinjoseph08/go-pg-migrations/v3"
)

const directory = "migrations"

func main() {
	log.Println("In the migrations")
	User, Password, Host, Port, Database := models.GetDbInfo()

	log.Println("Host: ", Host)

	db := pg.Connect(&pg.Options{
		Addr:     Host + ":" + Port,
		User:     User,
		Database: Database,
		Password: Password,
	})

	err := migrations.Run(db, directory, os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
