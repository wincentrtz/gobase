package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/wincentrtz/gobase/route"
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
