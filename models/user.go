package models

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/go-chi/chi"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	if isUsernameTaken(s.Db, user.Username, 0) {
		http.Error(w, "Username already taken", http.StatusBadRequest)
		return
	}

	if isEmailTaken(s.Db, user.Email, 0) {
		http.Error(w, "Email already taken", http.StatusBadRequest)
		return
	}

	if !isValidEmail(user.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id"
	err = s.Db.QueryRow(query, user.Username, user.Password, user.Email).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Failed inserting user in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func isUsernameTaken(db *sql.DB, username string, id int) bool {

	query := "SELECT COUNT(username) FROM users WHERE username = $1 AND id != $2"

	var userCount int
	err := db.QueryRow(query, username, id).Scan(&userCount)
	if err != nil {
		log.Fatal(err)
	}

	return userCount > 0

}

func isEmailTaken(db *sql.DB, email string, id int) bool {

	query := "SELECT COUNT(email) FROM users WHERE email = $1 AND id != $2"

	var emailCount int
	err := db.QueryRow(query, email, id).Scan(&emailCount)
	if err != nil {
		log.Fatal(err)
	}

	return emailCount > 0

}

func isValidEmail(email string) bool {

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)

}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	userID := chi.URLParam(r, "id")

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	user.ID, err = strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Failed converting user id", http.StatusBadRequest)
		return
	}

	if isUsernameTaken(s.Db, user.Username, user.ID) {
		http.Error(w, "Username already taken", http.StatusBadRequest)
		return
	}

	if isEmailTaken(s.Db, user.Email, user.ID) {
		http.Error(w, "Email already taken", http.StatusBadRequest)
		return
	}

	if !isValidEmail(user.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	query := "UPDATE users SET username = $1, password = $2, email = $3 WHERE id = $4"

	_, err = s.Db.Exec(query, user.Username, user.Password, user.Email, userID)
	if err != nil {
		http.Error(w, "Failed updating user in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

	userID := chi.URLParam(r, "id")

	query := "DELETE FROM users WHERE id = $1"

	_, err := s.Db.Exec(query, userID)
	if err != nil {
		http.Error(w, "Failed deleting user in database", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User deleted"))

}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	userID := chi.URLParam(r, "id")

	var user User

	query := "SELECT * FROM users WHERE id = $1"

	err := s.Db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		http.Error(w, "Failed getting user in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []User

	query := "SELECT * FROM users"

	rows, err := s.Db.Query(query)
	if err != nil {
		http.Error(w, "Failed getting users in database", http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		if err != nil {
			http.Error(w, "Failed getting users in database", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)

}

func (s *Server) GetUserByUsername(w http.ResponseWriter, r *http.Request) {

	username := chi.URLParam(r, "username")

	var user User

	query := "SELECT * FROM users WHERE username = $1"

	err := s.Db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		http.Error(w, "Failed getting user in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func (s *Server) GetUserByEmail(w http.ResponseWriter, r *http.Request) {

	email := chi.URLParam(r, "email")

	var user User

	query := "SELECT * FROM users WHERE email = $1"

	err := s.Db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		http.Error(w, "Failed getting user in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func (s *Server) GetUserByUsernameOrEmailAndPassword(w http.ResponseWriter, r *http.Request) {

	username := chi.URLParam(r, "username")
	email := chi.URLParam(r, "email")
	password := chi.URLParam(r, "password")

	var user User

	query := "SELECT * FROM users WHERE (username = $1 OR email = $2) AND password = $3"

	err := s.Db.QueryRow(query, username, email, password).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		http.Error(w, "Failed getting user in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}
