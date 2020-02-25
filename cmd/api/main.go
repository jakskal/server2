package main

import (
	"log"
	"os"
	"path/filepath"

	"server2/cmd/router"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	pathDir := "cd ../../"
	err := godotenv.Load(filepath.Join(pathDir, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}
}

func main() {
	serverPort := os.Getenv("SERVER_PORT")

	handler := initHandler()
	router.API(*handler, serverPort)
}
