package main

import (
	"goserverapi/config/db"
	"goserverapi/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.Init()

	defer db.DB.Close()

	handleRouter := router.Router()

	sitePort := os.Getenv("SITE_PORT")

	handleRouter.Run(":" + sitePort)

}
