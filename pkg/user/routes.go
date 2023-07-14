package user

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

	r.Get("/", s.GetUsersRoute)
	r.Get("/{id}", s.GetUserByIDRoute)
	r.Get("/username/{username}", s.GetUserByUsernameRoute)
	r.Post("/", s.CreateUserRoute)
	r.Post("/bulk", s.CreateUsersRoute)
	r.Put("/{id}", s.UpdateUserRoute)
	r.Delete("/{id}", s.DeleteUserRoute)

	return r

}

func (s *Server) CreateUserRoute(w http.ResponseWriter, r *http.Request) {

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	if isUsernameTaken(s.Service.Db, user.Username, 0) {
		http.Error(w, "Username is taken", http.StatusBadRequest)
		return
	}

	if isEmailTaken(s.Service.Db, user.Email, 0) {
		http.Error(w, "Email is taken", http.StatusBadRequest)
		return
	}

	if !isEmailValid(user.Email) {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	user, err = s.Service.CreateUser(user)
	if err != nil {
		http.Error(w, "Failed inserting user in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func (s *Server) CreateUsersRoute(w http.ResponseWriter, r *http.Request) {

	var users []User
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	for i := range users {

		if isUsernameTaken(s.Service.Db, users[i].Username, 0) {
			http.Error(w, "Username is taken", http.StatusBadRequest)
			return
		}

		if isEmailTaken(s.Service.Db, users[i].Email, 0) {
			http.Error(w, "Email is taken", http.StatusBadRequest)
			return
		}

		if !isEmailValid(users[i].Email) {
			http.Error(w, "Invalid email", http.StatusBadRequest)
			return
		}

	}

	users, err = s.Service.CreateUsers(users)
	if err != nil {
		http.Error(w, "Failed inserting users in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)

}

func (s *Server) UpdateUserRoute(w http.ResponseWriter, r *http.Request) {

	userId := chi.URLParam(r, "id")
	if userId == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed decoding", http.StatusBadRequest)
		return
	}

	if isUsernameTaken(s.Service.Db, user.Username, user.ID) {
		http.Error(w, "Username is taken", http.StatusBadRequest)
		return
	}

	if isEmailTaken(s.Service.Db, user.Email, user.ID) {
		http.Error(w, "Email is taken", http.StatusBadRequest)
		return
	}

	if !isEmailValid(user.Email) {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	err = s.Service.UpdateUser(userIdInt, user)
	if err != nil {
		http.Error(w, "Failed updating user in database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func (s *Server) DeleteUserRoute(w http.ResponseWriter, r *http.Request) {

	userId := chi.URLParam(r, "id")
	if userId == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = s.Service.DeleteUser(userIdInt)
	if err != nil {
		http.Error(w, "Failed deleting user from database", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User deleted"))

}

func (s *Server) GetUsersRoute(w http.ResponseWriter, r *http.Request) {

	users, err := s.Service.GetUsers()
	if err != nil {
		http.Error(w, "Failed getting users from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)

}

func (s *Server) GetUserByIDRoute(w http.ResponseWriter, r *http.Request) {

	userId := chi.URLParam(r, "id")
	if userId == "" {
		http.Error(w, "Missing user ID", http.StatusBadRequest)
		return
	}

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := s.Service.GetUserByID(userIdInt)
	if err != nil {
		http.Error(w, "Failed getting user from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func (s *Server) GetUserByUsernameRoute(w http.ResponseWriter, r *http.Request) {

	username := chi.URLParam(r, "username")
	if username == "" {
		http.Error(w, "Missing username", http.StatusBadRequest)
		return
	}

	user, err := s.Service.GetUserByUsername(username)
	if err != nil {
		http.Error(w, "Failed getting user from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}
