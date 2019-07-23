package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	userUsecase "github.com/wincentrtz/gobase/domains/user/usecase"

	"github.com/gorilla/mux"
	"github.com/wincentrtz/gobase/domains/user"
	"github.com/wincentrtz/gobase/models/responses"
)

type UserHandler struct {
	UserUsecase user.Usecase
}

func NewUserHandler(r *mux.Router) {
	handler := &UserHandler{
		UserUsecase: userUsecase.NewUserUsecase(),
	}
	r.HandleFunc("/api/users/{id}", handler.FindById).Methods("GET")
}

func (uh *UserHandler) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ := strconv.Atoi(vars["id"])
	user, err := uh.UserUsecase.FetchUserById(i)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		er := responses.ErrorResponse{
			Message: "Data Not Found",
		}
		json.NewEncoder(w).Encode(er)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}

}
