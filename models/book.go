package models

import "time"

type Book struct {
	ID          int        `json:"id" db:"id"`
	Title       string     `json:"title" db:"title" binding:"required"`
	Description string     `json:"description" db:"description"`
	ImageURL    string     `json:"image_url" db:"image_url"`
	ReleaseYear int        `json:"release_year" db:"release_year"`
	Price       int        `json:"price" db:"price"`
	TotalPage   int        `json:"total_page" db:"total_page"`
	Thickness   string     `json:"thickness" db:"thickness"`
	CategoryID  *int       `json:"category_id" db:"category_id"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	CreatedBy   string     `json:"created_by" db:"created_by"`
	ModifiedAt  *time.Time `json:"modified_at" db:"modified_at"`
	ModifiedBy  *string    `json:"modified_by" db:"modified_by"`
}
