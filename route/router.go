package route

import (
	"database/sql"

	userHandler "github.com/wincentrtz/gobase/domains/user/handler"
	userRepository "github.com/wincentrtz/gobase/domains/user/repository"
	userUsecase "github.com/wincentrtz/gobase/domains/user/usecase"

	"github.com/gorilla/mux"
)

func Router(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	registerUserHandler(r, db)
	return r
}

func registerUserHandler(r *mux.Router, db *sql.DB) {
	ur := userRepository.NewUserRepository(db)
	us := userUsecase.NewUserUsecase(ur)
	userHandler.NewUserHandler(r, us)
}
