package route

import (
	userHandler "github.com/wincentrtz/gobase/domains/user/handler"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	registerHandler(r)
	return r
}

func registerHandler(r *mux.Router) {
	userHandler.NewUserHandler(r)
}
