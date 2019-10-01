package route

import (
	userHandler "github.com/wincentrtz/gobase/domains/user/handler"
	"github.com/wincentrtz/gobase/infrastructures/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	registerHandler(r)
	registerMiddleware(r)
	return r
}

func registerMiddleware(r *mux.Router) {
	r.Use(middleware.LoggingMiddleware)
}

func registerHandler(r *mux.Router) {
	userHandler.NewUserHandler(r)
}
