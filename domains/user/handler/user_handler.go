package handler

import (
	"net/http"
	"strconv"

	"github.com/wincentrtz/gobase/domains/user/repository"
	"github.com/wincentrtz/gobase/domains/user/usecase"
	"github.com/wincentrtz/gobase/models/responses"

	"github.com/gorilla/mux"
	"github.com/wincentrtz/gobase/domains/user"
	"github.com/wincentrtz/gobase/gobase/infrastructures/db"
	"github.com/wincentrtz/gobase/gobase/utils"
)

type UserHandler struct {
	UserUsecase user.Usecase
}

func NewUserHandler(r *mux.Router) {
	db := db.Postgres()
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
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		er := responses.BaseResponse{
			Message: "Data Not Found",
			Code:    http.StatusNotFound,
		}
		utils.WriteResponse(w, er)
	} else {
		w.WriteHeader(http.StatusOK)
		resp := &responses.BaseResponse{
			Message: "Success",
			Code:    http.StatusOK,
			Data:    user,
		}
		utils.WriteResponse(w, resp)
	}
}
