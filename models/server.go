package models

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/patrickmn/go-cache"
)

type Server struct {
	db     *sql.DB
	Router chi.Router
	cache  *cache.Cache
}

func NewServer(db *sql.DB) *Server {

	router := chi.NewRouter()
	cache := cache.New(cache.NoExpiration, cache.NoExpiration)

	s := &Server{
		db:     db,
		Router: router,
		cache:  cache,
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

	return s

}
