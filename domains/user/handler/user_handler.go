package handler

import (
	"net/http"
	"strconv"

	"github.com/wincentrtz/gobase/domains/user/repository"
	"github.com/wincentrtz/gobase/domains/user/usecase"
	"github.com/wincentrtz/gobase/models/dto/responses"

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
	r.HandleFunc("/api/users/{id}", handler.FindById).Methods(http.MethodGet)
}

func (uh *UserHandler) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, _ := strconv.Atoi(vars["id"])
	user, err := uh.UserUsecase.FetchUserById(i)
	if err != nil {
		er := responses.BaseResponse{
			Message: http.StatusText(http.StatusNotFound),
			Code:    http.StatusNotFound,
		}
		utils.WriteResponse(w, er, http.StatusNotFound)
	} else {
		resp := &responses.BaseResponse{
			Message: http.StatusText(http.StatusOK),
			Code:    http.StatusOK,
			Data:    user,
		}
		utils.WriteResponse(w, resp, http.StatusOK)
	}
}
