package book

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/literalog/cerrors"
	"github.com/literalog/library/pkg/models"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Routes() *mux.Router
}

type handler struct {
	service Service
	router  *mux.Router
}

func NewHandler(s Service) Handler {
	h := &handler{
		service: s,
		router:  mux.NewRouter(),
	}

	h.setupRoutes()

	return h
}

func (h *handler) setupRoutes() {
	h.router.HandleFunc("/", h.Create).Methods(http.MethodPost)
	h.router.HandleFunc("/", h.Update).Methods(http.MethodPut)
	h.router.HandleFunc("/{id}", h.Delete).Methods(http.MethodDelete)
	h.router.HandleFunc("/{id}", h.GetById).Methods(http.MethodGet)
	h.router.HandleFunc("/", h.GetAll).Methods(http.MethodGet)
}

func (h handler) Routes() *mux.Router {
	return h.router
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := new(models.BookRequest)
	json.NewDecoder(r.Body).Decode(&req)

	b := models.NewBook(*req)
	if err := h.service.Create(ctx, b); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(b)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := new(models.BookRequest)
	json.NewDecoder(r.Body).Decode(&req)

	b := models.NewBook(*req)
	if err := h.service.Update(ctx, b); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(b)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]

	if err := h.service.Delete(ctx, id); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	bb, err := h.service.GetAll(ctx)
	if err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bb)
}

func (h *handler) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]

	b, err := h.service.GetById(ctx, id)
	if err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}
