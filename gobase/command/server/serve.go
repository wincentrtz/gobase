package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"github.com/wincentrtz/gobase/route"
)

var db *sql.DB

func Serve() {
	serverPort := ":" + viper.GetString(`server.port`)
	serverHost := "http://" + viper.GetString(`server.host`) + serverPort
	fmt.Println("Starting at " + serverHost)
	http.Handle("/", route.Router())
	http.ListenAndServe(serverPort, nil)
}
