package repositories

import (
	"errors"
	"time"

	"go-buku-project/database"
	"go-buku-project/models"
)

func CreateUser(username, password, createdBy string) (int, error) {
	now := time.Now()
	var id int
	err := database.DB.QueryRow(
		`INSERT INTO users(username, password, created_at, created_by) 
         VALUES($1,$2,$3,$4) RETURNING id`,
		username, password, now, createdBy,
	).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetUserByUsername(username string) (models.User, error) {
	var u models.User
	err := database.DB.QueryRow(`SELECT id, username, password, created_at, created_by, modified_at, modified_by FROM users WHERE username=$1`, username).
		Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt, &u.CreatedBy, &u.ModifiedAt, &u.ModifiedBy)
	if err != nil {
		return u, errors.New("user not found")
	}
	return u, nil
}
