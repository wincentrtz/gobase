package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wincentrtz/gobase/domains/user"
)

type UserHandler struct {
	UserUsecase user.Usecase
}

func NewUserHandler(r *mux.Router, us user.Usecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}
	r.HandleFunc("/api/users/{id}", handler.FindById).Methods("GET")
}

func (uh *UserHandler) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic("Error converting string to integer")
	}
	user, err := uh.UserUsecase.FetchUserById(i)
	if err != nil {
		panic("Can't find any user data")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
