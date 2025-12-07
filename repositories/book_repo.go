package repositories

import (
	"database/sql"
	"errors"

	"go-buku-project/database"
	"go-buku-project/models"
)

func GetAllBooks() ([]models.Book, error) {
	rows, err := database.DB.Query(`
		SELECT id, title, description, image_url, release_year, price, total_page, thickness,
		       category_id, created_at, created_by, modified_at, modified_by
		FROM books ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(
			&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price,
			&b.TotalPage, &b.Thickness, &b.CategoryID, &b.CreatedAt, &b.CreatedBy,
			&b.ModifiedAt, &b.ModifiedBy,
		); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func CreateBook(b models.Book) (int, error) {
	var id int
	err := database.DB.QueryRow(`
		INSERT INTO books(title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
		RETURNING id`,
		b.Title, b.Description, b.ImageURL, b.ReleaseYear, b.Price, b.TotalPage, b.Thickness, b.CategoryID, b.CreatedAt, b.CreatedBy, b.ModifiedAt, b.ModifiedBy,
	).Scan(&id)
	return id, err
}

func GetBookByID(id int) (models.Book, error) {
	var b models.Book
	err := database.DB.QueryRow(`
		SELECT id, title, description, image_url, release_year, price, total_page,
		       thickness, category_id, created_at, created_by, modified_at, modified_by
		FROM books WHERE id=$1`, id,
	).Scan(
		&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price,
		&b.TotalPage, &b.Thickness, &b.CategoryID, &b.CreatedAt, &b.CreatedBy,
		&b.ModifiedAt, &b.ModifiedBy,
	)
	if err == sql.ErrNoRows {
		return b, errors.New("book not found")
	}
	return b, err
}

func DeleteBook(id int) error {
	res, err := database.DB.Exec("DELETE FROM books WHERE id=$1", id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("book not found")
	}
	return nil
}

func GetBooksByCategory(categoryID int) ([]models.Book, error) {
	rows, err := database.DB.Query(`
		SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by
		FROM books WHERE category_id=$1 ORDER BY id`, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Description, &b.ImageURL, &b.ReleaseYear, &b.Price, &b.TotalPage, &b.Thickness, &b.CategoryID, &b.CreatedAt, &b.CreatedBy, &b.ModifiedAt, &b.ModifiedBy); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func UpdateBook(id int, b models.Book) error {
	res, err := database.DB.Exec(`
		UPDATE books
		SET title=$1, description=$2, image_url=$3, release_year=$4,
		    price=$5, total_page=$6, thickness=$7, category_id=$8,
		    modified_at=$9, modified_by=$10
		WHERE id=$11`,
		b.Title, b.Description, b.ImageURL, b.ReleaseYear,
		b.Price, b.TotalPage, b.Thickness, b.CategoryID,
		b.ModifiedAt, b.ModifiedBy, id,
	)
	if err != nil {
		return err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("book not found")
	}
	return nil
}
