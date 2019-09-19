package app

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"github.com/wincentrtz/gobase/route"
)

func Serve() {
	serverPort := ":" + viper.GetString(`server.port`)
	serverHost := "http://" + viper.GetString(`server.host`) + serverPort
	fmt.Println("Starting at " + serverHost)
	http.Handle("/", route.Router())
	http.ListenAndServe(serverPort, nil)
}
