package server

import (
	"books/pkg/books"
	"books/pkg/user"
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/patrickmn/go-cache"
)

type Server struct {
	Db         *sql.DB
	Router     chi.Router
	Cache      *cache.Cache
	BookServer *books.Server
	UserServer *user.Server
}

func NewServer(db *sql.DB) *Server {

	r := chi.NewRouter()
	cache := cache.New(cache.NoExpiration, cache.NoExpiration)

	bookService := books.NewService(db)
	bookServer := books.NewServer(bookService)

	userService := user.NewService(db)
	userServer := user.NewServer(userService)

	s := &Server{
		Db:         db,
		Router:     r,
		Cache:      cache,
		BookServer: bookServer,
		UserServer: userServer,
	}

	r.Mount("/books", bookServer.Routes())
	r.Mount("/users", userServer.Routes())

	return s

}
