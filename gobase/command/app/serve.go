package app

import (
	"fmt"
	"github.com/wincentrtz/gobase/route"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Serve() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serverPort := ":" + os.Getenv("APP_PORT")
	serverHost := "http://" + os.Getenv("APP_HOST") + serverPort
	fmt.Println("Starting at " + serverHost)
	http.Handle("/", route.Router())
	http.ListenAndServe(serverPort, nil)
}
