package models

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/lib/pq"
)

type Book struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Author string   `json:"author"`
	Genre  []string `json:"genre"`
}

func (s *Server) CreateBook(w http.ResponseWriter, r *http.Request) {

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO books (title, author, genre) VALUES ($1, $2, $3) RETURNING id"

	err = s.db.QueryRow(query, book.Title, book.Author, pq.Array(book.Genre)).Scan(&book.ID)
	if err != nil {
		http.Error(w, "Failed inserting book in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)

}

func (s *Server) UpdateBook(w http.ResponseWriter, r *http.Request) {

	bookID := chi.URLParam(r, "id")

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	query := "UPDATE books SET title = $1, author = $2, genre = $3 WHERE id = $4"

	_, err = s.db.Exec(query, book.Title, book.Author, pq.Array(book.Genre), bookID)
	if err != nil {
		http.Error(w, "Failed updating book in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)

}

func (s *Server) DeleteBook(w http.ResponseWriter, r *http.Request) {

	bookID := chi.URLParam(r, "id")

	query := "DELETE FROM books WHERE id = $1"

	_, err := s.db.Exec(query, bookID)
	if err != nil {
		http.Error(w, "Failed deleting book in database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (s *Server) GetBooks(w http.ResponseWriter, r *http.Request) {

	var books []Book

	query := "SELECT * FROM books"

	rows, err := s.db.Query(query)
	if err != nil {
		http.Error(w, "Failed querying database", http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, pq.Array(&book.Genre))
		if err != nil {
			http.Error(w, "Failed scanning row", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)

}

func (s *Server) GetBook(w http.ResponseWriter, r *http.Request) {

	bookID := chi.URLParam(r, "id")

	var book Book
	query := "SELECT * FROM books WHERE id = $1"

	err := s.db.QueryRow(query, bookID).Scan(&book.ID, &book.Title, &book.Author, pq.Array(&book.Genre))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(book)

}

func (s *Server) GetBooksByGenre(w http.ResponseWriter, r *http.Request) {

	genre := chi.URLParam(r, "genre")

	if cachedBooks, found := s.cache.Get("books-" + genre); found {
		json.NewEncoder(w).Encode(cachedBooks)
		log.Println("Cached")
		return
	}

	query := "SELECT * FROM books WHERE $1 = ANY(genre)"

	rows, err := s.db.Query(query, genre)
	if err != nil {
		http.Error(w, "Failed querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, pq.Array(&book.Genre))
		if err != nil {
			http.Error(w, "Failed scanning row", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	s.cache.Set("books-"+genre, books, time.Minute*5)
	json.NewEncoder(w).Encode(books)
}

func (s *Server) GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {

	author := chi.URLParam(r, "author")

	if cachedBooks, found := s.cache.Get("books-" + author); found {
		json.NewEncoder(w).Encode(cachedBooks)
		log.Println("Cached")
		return
	}

	query := "SELECT * FROM books WHERE author = $1"

	rows, err := s.db.Query(query, author)
	if err != nil {
		http.Error(w, "Failed querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, pq.Array(&book.Genre))
		if err != nil {
			http.Error(w, "Failed scanning row", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	s.cache.Set("books-"+author, books, time.Minute*5)
	json.NewEncoder(w).Encode(books)
}

func GetBooksRecommendation(w http.ResponseWriter, r *http.Request) {

}
