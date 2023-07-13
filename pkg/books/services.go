package books

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/lib/pq"
	"github.com/patrickmn/go-cache"
)

type Service struct {
	Db    *sql.DB
	Cache *cache.Cache
}

func NewService(db *sql.DB) *Service {

	cache := cache.New(cache.NoExpiration, cache.NoExpiration)

	s := &Service{
		Db:    db,
		Cache: cache,
	}

	return s

}

func (s *Service) CreateBook(book Book) (Book, error) {

	query := "INSERT INTO books (title, author, genre) VALUES ($1, $2, $3) RETURNING id"

	_, err := s.Db.Exec(query, book.Title, book.Author, pq.Array(book.Genre))
	if err != nil {
		return Book{}, err
	}

	return book, nil

}

func (s *Service) CreateReadBook(readBook ReadBook) (ReadBook, error) {

	query := "INSERT INTO read_books (user_id, book_id, read_date, rating, review) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	_, err := s.Db.Exec(query, readBook.UserID, readBook.BookID, readBook.ReadDate, readBook.Rating, readBook.Review)
	if err != nil {
		return ReadBook{}, err
	}

	return readBook, nil

}

func (s *Service) UpdateBook(id int, book Book) error {

	query := "UPDATE books SET title = $1, author = $2, genre = $3 WHERE id = $4"

	_, err := s.Db.Exec(query, book.Title, book.Author, pq.Array(book.Genre), id)
	if err != nil {
		return err
	}

	return nil

}

func (s *Service) UpdateReadBook(id int, readBook ReadBook) error {

	query := "UPDATE read_books SET user_id = $1, book_id = $2, read_date = $3, rating = $4, review = $5 WHERE id = $6"

	_, err := s.Db.Exec(query, readBook.UserID, readBook.BookID, readBook.ReadDate, readBook.Rating, readBook.Review, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

func (s *Service) DeleteBook(id int) error {

	query := "DELETE FROM books WHERE id = $1"

	_, err := s.Db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}

func (s *Service) DeleteReadBook(id int) error {

	query := "DELETE FROM read_books WHERE id = $1"

	_, err := s.Db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}

func (s *Service) GetBooks() ([]Book, error) {

	books := make([]Book, 0)

	query := "SELECT * FROM books"

	rows, err := s.Db.Query(query)
	if err != nil {
		return books, err
	}

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, pq.Array(&book.Genre))
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}

	return books, nil

}

func (s *Service) GetReadBooks() ([]ReadBook, error) {

	readBooks := make([]ReadBook, 0)

	query := "SELECT * FROM read_books"

	rows, err := s.Db.Query(query)
	if err != nil {
		return readBooks, err
	}

	for rows.Next() {
		var readBook ReadBook
		err := rows.Scan(&readBook.ID, &readBook.UserID, &readBook.BookID, &readBook.ReadDate, &readBook.Rating, &readBook.Review)
		if err != nil {
			return readBooks, err
		}
		readBooks = append(readBooks, readBook)
	}

	return readBooks, nil

}

func (s *Service) GetBookByID(id int) (Book, error) {

	if cachedBooks, found := s.Cache.Get("book-id-" + strconv.Itoa(id)); found {
		return cachedBooks.(Book), nil
	}

	var book Book
	query := "SELECT * FROM books WHERE id = $1"

	err := s.Db.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, pq.Array(&book.Genre))
	if err != nil {
		return Book{}, err
	}

	s.Cache.Set("book-id-"+strconv.Itoa(id), book, time.Minute*5)

	return book, nil

}

func (s *Service) GetReadBookByID(id int) (ReadBook, error) {

	if cachedBooks, found := s.Cache.Get("read-book-id-" + strconv.Itoa(id)); found {
		return cachedBooks.(ReadBook), nil
	}

	var readBook ReadBook

	query := "SELECT * FROM read_books WHERE id = $1"

	err := s.Db.QueryRow(query, id).Scan(&readBook.ID, &readBook.UserID, &readBook.BookID, &readBook.ReadDate, &readBook.Rating, &readBook.Review)
	if err != nil {
		return ReadBook{}, err
	}

	s.Cache.Set("read-book-id-"+strconv.Itoa(id), readBook, time.Minute*5)

	return readBook, nil

}

func (s *Service) GetBooksByGenre(genre string) ([]Book, error) {

	if cachedBooks, found := s.Cache.Get("books-genre-" + genre); found {
		return cachedBooks.([]Book), nil
	}

	query := "SELECT * FROM books WHERE $1 = ANY(genre)"

	rows, err := s.Db.Query(query, genre)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, pq.Array(&book.Genre))
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	s.Cache.Set("books-genre-"+genre, books, time.Minute*5)

	return books, nil

}

func (s *Service) GetBooksByAuthor(author string) ([]Book, error) {

	if cachedBooks, found := s.Cache.Get("books-author-" + author); found {
		return cachedBooks.([]Book), nil
	}

	query := "SELECT * FROM books WHERE author = $1"

	rows, err := s.Db.Query(query, author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, pq.Array(&book.Genre))
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	s.Cache.Set("books-author-"+author, books, time.Minute*5)

	return books, nil

}

func (s *Service) GetReadBooksByUser(id int) ([]ReadBook, error) {

	if cachedReadBooks, found := s.Cache.Get("read-books-user-id-" + strconv.Itoa(id)); found {
		return cachedReadBooks.([]ReadBook), nil
	}

	query := "SELECT * FROM read_books WHERE user_id = $1"

	rows, err := s.Db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var readBooks []ReadBook
	for rows.Next() {
		var readBook ReadBook
		err := rows.Scan(&readBook.ID, &readBook.UserID, &readBook.BookID, &readBook.ReadDate, &readBook.Rating, &readBook.Review)
		if err != nil {
			return nil, err
		}
		readBooks = append(readBooks, readBook)
	}

	s.Cache.Set("read-books-user-id-"+strconv.Itoa(id), readBooks, time.Minute*5)

	return readBooks, nil

}

func GetBooksRecommendation(w http.ResponseWriter, r *http.Request) {

}