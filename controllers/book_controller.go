package controllers

import (
	"net/http"
	"strconv"
	"time"

	"go-buku-project/models"
	"go-buku-project/repositories"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	books, err := repositories.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi release_year
	if input.ReleaseYear < 1980 || input.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year must be between 1980 and 2024"})
		return
	}

	// Hitung thickness berdasarkan total_page
	if input.TotalPage > 100 {
		input.Thickness = "tebal"
	} else {
		input.Thickness = "tipis"
	}

	// Ambil created_by dari JWT middleware, default "Admin"
	createdBy := c.GetString("user")
	if createdBy == "" {
		createdBy = "Admin"
	}
	input.CreatedBy = createdBy
	input.CreatedAt = time.Now()

	// Simpan ke DB melalui repository
	id, err := repositories.CreateBook(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Response lengkap
	c.JSON(http.StatusCreated, gin.H{
		"id":           id,
		"title":        input.Title,
		"description":  input.Description,
		"image_url":    input.ImageURL,
		"release_year": input.ReleaseYear,
		"price":        input.Price,
		"total_page":   input.TotalPage,
		"thickness":    input.Thickness,
		"category_id":  input.CategoryID,
		"created_at":   input.CreatedAt,
		"created_by":   input.CreatedBy,
		"modified_at":  input.ModifiedAt, // nil kalau belum diupdate
		"modified_by":  input.ModifiedBy, // nil kalau belum diupdate
	})
}

func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := repositories.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repositories.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
}

func GetBooksByCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	books, err := repositories.GetBooksByCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		ReleaseYear int    `json:"release_year"`
		Price       int    `json:"price"`
		TotalPage   int    `json:"total_page"`
		CategoryID  int    `json:"category_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi release_year
	if input.ReleaseYear < 1980 || input.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year must be between 1980 and 2024"})
		return
	}

	// Hitung thickness
	thickness := "tipis"
	if input.TotalPage > 100 {
		thickness = "tebal"
	}

	// Ambil modified_by dari JWT
	modifiedBy := c.GetString("user")
	if modifiedBy == "" {
		modifiedBy = "Admin"
	}
	modifiedAt := time.Now()

	book := models.Book{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		Thickness:   thickness,
		CategoryID:  &input.CategoryID,
		ModifiedAt:  &modifiedAt,
		ModifiedBy:  &modifiedBy,
	}

	if err := repositories.UpdateBook(id, book); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":           id,
		"title":        book.Title,
		"description":  book.Description,
		"image_url":    book.ImageURL,
		"release_year": book.ReleaseYear,
		"price":        book.Price,
		"total_page":   book.TotalPage,
		"thickness":    book.Thickness,
		"category_id":  book.CategoryID,
		"modified_at":  book.ModifiedAt,
		"modified_by":  book.ModifiedBy,
	})
}
