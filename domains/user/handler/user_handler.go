package handler

import (
	"net/http"
	"strconv"

	"github.com/wincentrtz/gobase/domains/user/repository"
	"github.com/wincentrtz/gobase/domains/user/usecase"
	"github.com/wincentrtz/gobase/models/responses"

	"github.com/gorilla/mux"
	"github.com/wincentrtz/gobase/domains/user"
	"github.com/wincentrtz/gobase/gobase/config"
	"github.com/wincentrtz/gobase/gobase/utils"
)

type UserHandler struct {
	UserUsecase user.Usecase
}

func NewUserHandler(r *mux.Router) {
	db := config.InitDb()
	ur := repository.NewUserRepository(db)
	us := usecase.NewUserUsecase(ur)
	handler := &UserHandler{
		UserUsecase: us,
	}
	r.HandleFunc("/api/users/{id}", handler.FindById).Methods("GET")
}

func (uh *UserHandler) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ := strconv.Atoi(vars["id"])
	user, err := uh.UserUsecase.FetchUserById(i)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		er := responses.ErrorResponse{
			Message: "Data Not Found",
		}
		utils.ErrorResponse(w, er)
	} else {
		utils.SuccessResponse(w, user)
	}
}
