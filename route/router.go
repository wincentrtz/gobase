package route

import (
	"database/sql"

	userHandler "github.com/wincentrtz/gobase/domains/user/handler"

	"github.com/gorilla/mux"
)

func Router(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	registerHandler(r, db)
	return r
}

func registerHandler(r *mux.Router, db *sql.DB) {
	userHandler.NewUserHandler(r, db)
}
