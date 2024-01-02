package book

import (
	"github.com/gorilla/mux"
)

type Handler interface {
	// 	Create(w http.ResponseWriter, r *http.Request) error
	// 	Update(w http.ResponseWriter, r *http.Request) error
	// 	Delete(w http.ResponseWriter, r *http.Request) error
	// 	GetAll(w http.ResponseWriter, r *http.Request) error
	// 	GetById(w http.ResponseWriter, r *http.Request) error
	Routes() *mux.Router
}

type handler struct {
	service Service
	router  *mux.Router
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) setupRoutes() {
}

func (h handler) Routes() *mux.Router {
	return h.router
}
