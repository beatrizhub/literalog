package models

import (
	"encoding/json"
	"net/http"
)

type ReadBook struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	BookID   int    `json:"book_id"`
	ReadDate string `json:"read_date"`
	Rating   int    `json:"rating"`
	Review   string `json:"review"`
}

func (s *Server) CreateReadBook(w http.ResponseWriter, r *http.Request) {

	var readBook ReadBook
	err := json.NewDecoder(r.Body).Decode(&readBook)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO read_books (user_id, book_id, read_date, rating, review) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	err = s.Db.QueryRow(query, readBook.UserID, readBook.BookID, readBook.ReadDate, readBook.Rating, readBook.Review).Scan(&readBook.ID)
	if err != nil {
		http.Error(w, "Failed inserting read book in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(readBook)

}

func (s *Server) UpdateReadBook(w http.ResponseWriter, r *http.Request) {

	var readBook ReadBook
	err := json.NewDecoder(r.Body).Decode(&readBook)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	query := "UPDATE read_books SET user_id = $1, book_id = $2, read_date = $3, rating = $4, review = $5 WHERE id = $6"
	err = s.Db.QueryRow(query, readBook.UserID, readBook.BookID, readBook.ReadDate, readBook.Rating, readBook.Review, readBook.ID).Scan(&readBook.ID)
	if err != nil {
		http.Error(w, "Failed updating read book in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(readBook)

}

func (s *Server) DeleteReadBook(w http.ResponseWriter, r *http.Request) {

	var readBook ReadBook
	err := json.NewDecoder(r.Body).Decode(&readBook)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM read_books WHERE id = $1"
	err = s.Db.QueryRow(query, readBook.ID).Scan(&readBook.ID)
	if err != nil {
		http.Error(w, "Failed deleting read book in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(readBook)

}

func (s *Server) GetReadBooks(w http.ResponseWriter, r *http.Request) {

	var readBooks []ReadBook

	query := "SELECT * FROM read_books"
	rows, err := s.Db.Query(query)
	if err != nil {
		http.Error(w, "Failed querying read books in database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var readBook ReadBook
		err := rows.Scan(&readBook.ID, &readBook.UserID, &readBook.BookID, &readBook.ReadDate, &readBook.Rating, &readBook.Review)
		if err != nil {
			http.Error(w, "Failed scanning read books in database", http.StatusInternalServerError)
			return
		}
		readBooks = append(readBooks, readBook)
	}

	json.NewEncoder(w).Encode(readBooks)

}

func (s *Server) GetReadBook(w http.ResponseWriter, r *http.Request) {

	var readBook ReadBook

	id := r.URL.Query().Get("id")

	query := "SELECT * FROM read_books WHERE id = $1"
	err := s.Db.QueryRow(query, id).Scan(&readBook.ID, &readBook.UserID, &readBook.BookID, &readBook.ReadDate, &readBook.Rating, &readBook.Review)
	if err != nil {
		http.Error(w, "Failed querying read book in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(readBook)

}
