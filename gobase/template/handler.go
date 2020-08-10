package template

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wincentrtz/gobase/gobase/infrastructures/db"
	"github.com/wincentrtz/gobase/gobase/utils"
	"github.com/wincentrtz/gobase/models/dto/responses"
)

type TemplateHandler struct {
	templateUsecase TemplateUsecase
}

func NewTemplateHandler(r *mux.Router) {
	db := db.Postgres()
	ur := NewTemplateRepository(db)
	us := NewTemplateUsecase(ur)
	handler := &TemplateHandler{
		templateUsecase: us,
	}
	r.HandleFunc("/api/template/{id}", handler.FindById).Methods("GET")
}

func (uh *TemplateHandler) FindById(w http.ResponseWriter, r *http.Request) {
	resp := &responses.BaseResponse{
		Message: http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Data:    "template",
	}
	utils.WriteResponse(w, resp, http.StatusOK)
}
