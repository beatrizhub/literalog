package models

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/patrickmn/go-cache"
)

type Server struct {
	Db     *sql.DB
	Router chi.Router
	Cache  *cache.Cache
}

func NewServer(db *sql.DB) *Server {

	router := chi.NewRouter()
	cache := cache.New(cache.NoExpiration, cache.NoExpiration)

	s := &Server{
		Db:     db,
		Router: router,
		Cache:  cache,
	}

	router.Get("/books", s.GetBooks)
	router.Get("/books/{id}", s.GetBook)
	router.Get("/books/genre/{genre}", s.GetBooksByGenre)
	router.Get("/books/author/{author}", s.GetBooksByAuthor)
	router.Post("/books", s.CreateBook)
	router.Put("/books/{id}", s.UpdateBook)
	router.Delete("/books/{id}", s.DeleteBook)

	router.Get("/users", s.GetUsers)
	router.Get("/users/{id}", s.GetUser)
	router.Post("/users", s.CreateUser)
	router.Put("/users/{id}", s.UpdateUser)
	router.Delete("/users/{id}", s.DeleteUser)

	router.Get("/users/read", s.GetReadBooks)
	router.Get("/users/read/{id}", s.GetReadBook)
	router.Post("/users/read", s.CreateReadBook)
	router.Put("/users/read/{id}", s.UpdateReadBook)
	router.Delete("/users/read/{id}", s.DeleteReadBook)

	return s

}
