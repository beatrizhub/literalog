package api

import (
	"log"
	"net/http"

	"github.com/literalog/library/internal/app/domain/author"
	"github.com/literalog/library/internal/app/domain/book"
	"github.com/literalog/library/internal/app/domain/genre"
	"github.com/literalog/library/internal/app/domain/series"
	"github.com/literalog/library/internal/app/gateways/database/mongodb"

	"github.com/gorilla/mux"
)

type Server struct {
	port     string
	logLevel int
	router   *mux.Router
}

func NewServer(port string) Server {
	s := Server{
		port:     port,
		logLevel: 1,
		router:   mux.NewRouter(),
	}

	authorRepository := mongodb.NewAuthorRepository(nil)
	authorService := author.NewService(authorRepository)
	authorHandler := author.NewHandler(authorService)

	seriesRepository := mongodb.NewSeriesRepository(nil)
	seriesService := series.NewService(seriesRepository)
	seriesHandler := series.NewHandler(seriesService)

	genreRepository := mongodb.NewGenreRepository(nil)
	genreService := genre.NewService(genreRepository)
	genreHandler := genre.NewHandler(genreService)

	bookRepository := mongodb.NewBookRepository(nil)
	bookService := book.NewService(bookRepository)
	bookHandler := book.NewHandler(bookService)

	s.router.PathPrefix("/authors").Handler(authorHandler.Routes())
	s.router.PathPrefix("/series").Handler(seriesHandler.Routes())
	s.router.PathPrefix("/genres").Handler(genreHandler.Routes())
	s.router.PathPrefix("/books").Handler(bookHandler.Routes())

	return s
}

func (s *Server) ServeHttp() error {
	log.Println("Server listening on", s.port)
	return http.ListenAndServe(s.port, s.router)
}
