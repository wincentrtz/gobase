package route

import (
	userHandler "github.com/wincentrtz/gobase/domains/user/handler"
	"github.com/wincentrtz/gobase/configs/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	registerMiddleware(r)
	registerHandler(r)
	return r
}

func registerMiddleware(r *mux.Router) {
	r.Use(middleware.LoggingMiddleware)
}

func registerHandler(r *mux.Router) {
	userHandler.NewUserHandler(r)
}
