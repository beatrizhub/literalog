package user

import (
	"database/sql"

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

func (s *Service) CreateUser(user User) (User, error) {

	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id"

	_, err := s.Db.Exec(query, user.Username, user.Password, user.Email)
	if err != nil {
		return User{}, err
	}

	return user, nil

}

func (s *Service) CreateUsers(users []User) ([]User, error) {

	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id"

	for i := range users {
		_, err := s.Db.Exec(query, users[i].Username, users[i].Password, users[i].Email)
		if err != nil {
			return users, err
		}
	}

	return users, nil

}

func (s *Service) UpdateUser(id int, user User) error {

	query := "UPDATE users SET username = $1, password = $2, email = $3 WHERE id = $4"

	_, err := s.Db.Exec(query, user.Username, user.Password, user.Email, id)
	if err != nil {
		return err
	}

	return nil

}

func (s *Service) DeleteUser(id int) error {

	query := "DELETE FROM users WHERE id = $1"

	_, err := s.Db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}

func (s *Service) GetUsers() ([]User, error) {

	users := make([]User, 0)

	query := "SELECT * FROM users"

	rows, err := s.Db.Query(query)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil

}

func (s *Service) GetUserByID(id int) (User, error) {

	query := "SELECT * FROM users WHERE id = $1"

	var user User
	err := s.Db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return User{}, err
	}

	return user, nil

}

func (s *Service) GetUserByUsername(username string) (User, error) {

	query := "SELECT * FROM users WHERE username = $1"

	var user User
	err := s.Db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return User{}, err
	}

	return user, nil

}
