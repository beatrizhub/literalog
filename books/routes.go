package books

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Server struct {
	Service *Service
}

func NewServer(service *Service) *Server {

	return &Server{
		Service: service,
	}

}

func (s *Server) Routes() chi.Router {

	r := chi.NewRouter()

	r.Get("/books", s.GetBooksRoute)
	r.Get("/books/{id}", s.GetBookByIDRoute)
	r.Get("/books/genre/{genre}", s.GetBooksByGenreRoute)
	r.Get("/books/author/{author}", s.GetBooksByAuthorRoute)
	r.Post("/books", s.CreateBookRoute)
	r.Put("/books/{id}", s.UpdateBookRoute)
	r.Delete("/books/{id}", s.DeleteBookRoute)

	r.Get("/users/read", s.GetReadBooksRoute)
	r.Get("/users/read/{id}", s.GetReadBookByIDRoute)
	r.Get("/users/{userID}/read/{bookID}", s.GetReadBooksByUserRoute)
	r.Post("/users/read", s.CreateReadBookRoute)
	r.Put("/users/read/{id}", s.UpdateReadBookRoute)
	r.Delete("/users/read/{id}", s.DeleteReadBookRoute)

	return r

}
func (s *Server) CreateBookRoute(w http.ResponseWriter, r *http.Request) {

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	book, err = s.Service.CreateBook(book)
	if err != nil {
		http.Error(w, "Failed inserting book in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)

}

func (s *Server) CreateReadBookRoute(w http.ResponseWriter, r *http.Request) {

	var readBook ReadBook
	err := json.NewDecoder(r.Body).Decode(&readBook)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	readBook, err = s.Service.CreateReadBook(readBook)
	if err != nil {
		http.Error(w, "Failed inserting read book in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(readBook)

}

func (s *Server) UpdateBookRoute(w http.ResponseWriter, r *http.Request) {

	bookID := chi.URLParam(r, "id")

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	bookIDint, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, "Failed converting bookID to int", http.StatusInternalServerError)
		return
	}

	err = s.Service.UpdateBook(bookIDint, book)
	if err != nil {
		http.Error(w, "Failed updating book in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)

}

func (s *Server) UpdateReadBookRoute(w http.ResponseWriter, r *http.Request) {

	readBookID := chi.URLParam(r, "id")

	var readBook ReadBook

	err := json.NewDecoder(r.Body).Decode(&readBook)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	readBookIDint, err := strconv.Atoi(readBookID)
	if err != nil {
		http.Error(w, "Failed converting readBookID to int", http.StatusInternalServerError)
		return
	}

	err = s.Service.UpdateReadBook(readBookIDint, readBook)
	if err != nil {
		http.Error(w, "Failed updating read book in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(readBook)

}

func (s *Server) DeleteBookRoute(w http.ResponseWriter, r *http.Request) {

	bookID := chi.URLParam(r, "id")

	bookIDint, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, "Failed converting bookID to int", http.StatusInternalServerError)
		return
	}

	err = s.Service.DeleteBook(bookIDint)
	if err != nil {
		http.Error(w, "Failed deleting book in database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (s *Server) DeleteReadBookRoute(w http.ResponseWriter, r *http.Request) {

	readBookID := chi.URLParam(r, "id")

	readBookIDint, err := strconv.Atoi(readBookID)
	if err != nil {
		http.Error(w, "Failed converting readBookID to int", http.StatusInternalServerError)
		return
	}

	err = s.Service.DeleteReadBook(readBookIDint)
	if err != nil {
		http.Error(w, "Failed deleting read book in database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (s *Server) GetBooksRoute(w http.ResponseWriter, r *http.Request) {

	books, err := s.Service.GetBooks()
	if err != nil {
		http.Error(w, "Failed getting books from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)

}

func (s *Server) GetReadBooksRoute(w http.ResponseWriter, r *http.Request) {

	readBooks, err := s.Service.GetReadBooks()
	if err != nil {
		http.Error(w, "Failed getting read books from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(readBooks)

}

func (s *Server) GetBookByIDRoute(w http.ResponseWriter, r *http.Request) {

	bookID := chi.URLParam(r, "id")

	bookIDint, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, "Failed converting bookID to int", http.StatusInternalServerError)
		return
	}

	book, err := s.Service.GetBookByID(bookIDint)
	if err != nil {
		http.Error(w, "Failed getting book from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)

}

func (s *Server) GetReadBookByIDRoute(w http.ResponseWriter, r *http.Request) {

	readBookID := chi.URLParam(r, "id")

	readBookIDint, err := strconv.Atoi(readBookID)
	if err != nil {
		http.Error(w, "Failed converting readBookID to int", http.StatusInternalServerError)
		return
	}

	readBook, err := s.Service.GetReadBookByID(readBookIDint)
	if err != nil {
		http.Error(w, "Failed getting read book from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(readBook)

}

func (s *Server) GetBooksByGenreRoute(w http.ResponseWriter, r *http.Request) {

	genre := chi.URLParam(r, "genre")

	books, err := s.Service.GetBooksByGenre(genre)
	if err != nil {
		http.Error(w, "Failed getting book from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func (s *Server) GetReadBooksByUserRoute(w http.ResponseWriter, r *http.Request) {

	userID := chi.URLParam(r, "userID")

	userIDint, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Failed converting userID to int", http.StatusInternalServerError)
		return
	}

	readBooks, err := s.Service.GetReadBooksByUser(userIDint)
	if err != nil {
		http.Error(w, "Failed getting read books from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(readBooks)

}

func (s *Server) GetBooksByAuthorRoute(w http.ResponseWriter, r *http.Request) {

	author := chi.URLParam(r, "author")

	books, err := s.Service.GetBooksByAuthor(author)
	if err != nil {
		http.Error(w, "Failed getting book from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)

}

func GetBooksRecommendationRoute(w http.ResponseWriter, r *http.Request) {

}
