package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	userHandler "github.com/wincentrtz/gobase/domains/user/handler"
	userRepository "github.com/wincentrtz/gobase/domains/user/repository"
	userUsecase "github.com/wincentrtz/gobase/domains/user/usecase"

	"github.com/gorilla/mux"
	"github.com/wincentrtz/gobase/config"
)

func main() {
	db := config.InitDb()
	serverPort := ":" + viper.GetString(`server.port`)
	serverHost := "http://" + viper.GetString(`server.host`) + serverPort
	defer db.Close()
	r := mux.NewRouter()
	registerUserHandler(r, db)

	fmt.Println("Starting at " + serverHost)
	http.Handle("/", r)
	http.ListenAndServe(serverPort, nil)
}

func registerUserHandler(r *mux.Router, db *sql.DB) {
	ur := userRepository.NewUserRepository(db)
	us := userUsecase.NewUserUsecase(ur)
	userHandler.NewUserHandler(r, us)
}
