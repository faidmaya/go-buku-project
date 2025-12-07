package models

import "time"

type User struct {
	ID         int        `json:"id" db:"id"`
	Username   string     `json:"username" db:"username"`
	Password   string     `json:"-" db:"password"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	CreatedBy  string     `json:"created_by" db:"created_by"`
	ModifiedAt *time.Time `json:"modified_at" db:"modified_at"`
	ModifiedBy *string    `json:"modified_by" db:"modified_by"`
}
