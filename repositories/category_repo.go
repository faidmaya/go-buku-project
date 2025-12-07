package repositories

import (
	"database/sql"
	"errors"
	"time"

	"go-buku-project/database"
	"go-buku-project/models"
)

func GetAllCategories() ([]models.Category, error) {
	rows, err := database.DB.Query("SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cats []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt, &c.CreatedBy, &c.ModifiedAt, &c.ModifiedBy); err != nil {
			return nil, err
		}
		cats = append(cats, c)
	}
	return cats, nil
}

func CreateCategory(name, createdBy string) (int, time.Time, error) {
	createdAt := time.Now()
	var id int
	err := database.DB.QueryRow(
		`INSERT INTO categories(name, created_at, created_by) 
		 VALUES($1,$2,$3) RETURNING id`,
		name, createdAt, createdBy,
	).Scan(&id)
	return id, createdAt, err
}

func GetCategoryByID(id int) (models.Category, error) {
	var c models.Category
	err := database.DB.QueryRow(`SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id=$1`, id).
		Scan(&c.ID, &c.Name, &c.CreatedAt, &c.CreatedBy, &c.ModifiedAt, &c.ModifiedBy)
	if err == sql.ErrNoRows {
		return c, errors.New("category not found")
	}
	return c, err
}

func DeleteCategory(id int) error {
	res, err := database.DB.Exec("DELETE FROM categories WHERE id=$1", id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("category not found")
	}
	return nil
}

func UpdateCategory(id int, name string, modifiedAt time.Time, modifiedBy string) error {
	res, err := database.DB.Exec(`
		UPDATE categories
		SET name=$1, modified_at=$2, modified_by=$3
		WHERE id=$4`,
		name, modifiedAt, modifiedBy, id,
	)
	if err != nil {
		return err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("category not found")
	}
	return nil
}
