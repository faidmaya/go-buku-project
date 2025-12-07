package controllers

import (
	"net/http"
	"strconv"
	"time"

	"go-buku-project/repositories"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	cats, err := repositories.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cats)
}

func CreateCategory(c *gin.Context) {
	var body struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBy := c.GetString("user") // bisa default "Admin" kalau middleware JWT belum set
	if createdBy == "" {
		createdBy = "Admin"
	}

	id, createdAt, err := repositories.CreateCategory(body.Name, createdBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         id,
		"name":       body.Name,
		"created_at": createdAt,
		"created_by": createdBy,
	})
}

func GetCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cat, err := repositories.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	c.JSON(http.StatusOK, cat)
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repositories.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}

func UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var body struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	// Ambil modified_by dari JWT middleware, default "Admin"
	modifiedBy := c.GetString("user")
	if modifiedBy == "" {
		modifiedBy = "Admin"
	}
	modifiedAt := time.Now()

	// Update kategori lewat repository
	err := repositories.UpdateCategory(id, body.Name, modifiedAt, modifiedBy)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          id,
		"name":        body.Name,
		"modified_at": modifiedAt,
		"modified_by": modifiedBy,
	})
}
