package server

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"github.com/wincentrtz/gobase/config"
	"github.com/wincentrtz/gobase/route"
)

func Serve() {
	db := config.InitDb()
	defer db.Close()
	serverPort := ":" + viper.GetString(`server.port`)
	serverHost := "http://" + viper.GetString(`server.host`) + serverPort

	fmt.Println("Starting at " + serverHost)
	http.Handle("/", route.Router(db))
	http.ListenAndServe(serverPort, nil)
}
