package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"user/database"
	"user/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	fmt.Println("CreateUser")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a number"})
		return
	}
	// proceed...

	var user models.User
	result := database.DB.First(&user, id) // auto uses `id = ?`

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	fmt.Println("GetUser")
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetOrCreateUser(db *gorm.DB, input models.User) (models.User, error) {
	var user models.User

	result := db.Where("username = ?", input.Username).First(&user)
	if result.Error == nil {
		return user, nil
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, result.Error
	}

	// User not found, create one
	user = input
	if err := db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetOrCreateUserHandler(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := GetOrCreateUser(database.DB, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
