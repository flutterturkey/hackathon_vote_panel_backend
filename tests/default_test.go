package tests

import (
	"goBoilterplate/config"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.Database()
	if err != nil {
		log.Fatal("Db connection error")
	}
	defer db.Close()

	config.Redis()

	os.Exit(m.Run())
}
