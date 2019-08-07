package template

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type TemplateHandler struct {
	templateUsecase TemplateUsecase
}

func NewUserHandler(r *mux.Router) {
	handler := &TemplateHandler{
		templateUsecase: NewTemplateUsecase(),
	}
	r.HandleFunc("/api/template/{id}", handler.FindById).Methods("GET")
}

func (uh *TemplateHandler) FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("template")
}
