package user

import (
	"database/sql"
	"log"
	"regexp"
)

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

func isEmailValid(email string) bool {

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)

}
