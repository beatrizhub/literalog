package models

import (
	"books/books"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/patrickmn/go-cache"
)

type Server struct {
	Db         *sql.DB
	Router     chi.Router
	Cache      *cache.Cache
	BookServer *books.Server
}

func NewServer(db *sql.DB) *Server {

	router := chi.NewRouter()
	cache := cache.New(cache.NoExpiration, cache.NoExpiration)

	bookService := books.NewService(db)
	bookServer := books.NewServer(bookService)

	s := &Server{
		Db:         db,
		Router:     router,
		Cache:      cache,
		BookServer: bookServer,
	}

	router.Mount("/", bookServer.Routes())

	router.Get("/users", s.GetUsers)
	router.Get("/users/{id}", s.GetUser)
	router.Post("/users", s.CreateUser)
	router.Put("/users/{id}", s.UpdateUser)
	router.Delete("/users/{id}", s.DeleteUser)

	return s

}
