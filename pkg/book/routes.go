package book

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Server struct {
	Service Handler
}

func NewServer(s Handler) *Server {

	return &Server{
		Service: s,
	}

}

func (s *Server) Routes() chi.Router {

	r := chi.NewRouter()

	r.Get("/", s.GetBooksRoute)
	r.Get("/{id}", s.GetBookByIDRoute)
	r.Get("/genre/{genre}", s.GetBooksByGenreRoute)
	r.Get("/author/{author}", s.GetBooksByAuthorRoute)
	r.Get("/title/{title}", s.GetBooksByTitleRoute)
	r.Get("/recommendations/{userId}", s.GetBooksRecommendationsRoute)
	r.Post("/", s.CreateBookRoute)
	r.Put("/{id}", s.UpdateBookRoute)
	r.Delete("/{id}", s.DeleteBookRoute)

	r.Get("/read", s.GetReadBooksRoute)
	r.Get("/read/{id}", s.GetReadBookByIDRoute)
	r.Get("/read/user/{userId}", s.GetReadBooksByUserRoute)
	r.Post("/read", s.CreateReadBookRoute)
	r.Put("/read/{id}", s.UpdateReadBookRoute)
	r.Delete("/read/{id}", s.DeleteReadBookRoute)

	r.Get("/toberead", s.GetToBeReadBooksRoute)
	r.Get("/toberead/user/{userId}", s.GetToBeReadBooksByUserIDRoute)
	r.Post("/toberead", s.CreateToBeReadBookRoute)
	r.Delete("/toberead/{id}", s.DeleteToBeReadBookRoute)

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

func (s *Server) CreateToBeReadBookRoute(w http.ResponseWriter, r *http.Request) {

	var toBeReadBook ToBeReadBook

	err := json.NewDecoder(r.Body).Decode(&toBeReadBook)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	toBeReadBook, err = s.Service.CreateToBeReadBook(toBeReadBook)
	if err != nil {
		http.Error(w, "Failed inserting to be read book in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(toBeReadBook)

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

func (s *Server) DeleteToBeReadBookRoute(w http.ResponseWriter, r *http.Request) {

	toBeReadBookID := chi.URLParam(r, "id")

	toBeReadBookIDint, err := strconv.Atoi(toBeReadBookID)
	if err != nil {
		http.Error(w, "Failed converting toBeReadBookID to int", http.StatusInternalServerError)
		return
	}

	err = s.Service.DeleteToBeReadBook(toBeReadBookIDint)
	if err != nil {
		http.Error(w, "Failed deleting to be read book in database", http.StatusInternalServerError)
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

func (s *Server) GetToBeReadBooksRoute(w http.ResponseWriter, r *http.Request) {

	toBeReadBooks, err := s.Service.GetToBeReadBooks()
	if err != nil {
		http.Error(w, "Failed getting to be read books from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(toBeReadBooks)

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

	userId := chi.URLParam(r, "userId")

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Failed converting userID to int", http.StatusInternalServerError)
		return
	}

	readBooks, err := s.Service.GetReadBooksByUserID(userIdInt)
	if err != nil {
		http.Error(w, "Failed getting read books from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(readBooks)

}

func (s *Server) GetToBeReadBooksByUserIDRoute(w http.ResponseWriter, r *http.Request) {

	userId := chi.URLParam(r, "userId")

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Failed converting userID to int", http.StatusInternalServerError)
		return
	}

	toBeReadBooks, err := s.Service.GetToBeReadBooksByUserID(userIdInt)
	if err != nil {
		http.Error(w, "Failed getting to be read books from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(toBeReadBooks)

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

func (s *Server) GetBooksByTitleRoute(w http.ResponseWriter, r *http.Request) {

	title := chi.URLParam(r, "title")

	books, err := s.Service.GetBooksByTitle(title)
	if err != nil {
		http.Error(w, "Failed getting book from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)

}

func (s *Server) GetBooksRecommendationsRoute(w http.ResponseWriter, r *http.Request) {

	userId := chi.URLParam(r, "userId")

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Failed converting userID to int", http.StatusInternalServerError)
		return
	}

	books, err := s.Service.GetBooksRecommendations(userIdInt)
	if err != nil {
		http.Error(w, "Failed getting book from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)

}
