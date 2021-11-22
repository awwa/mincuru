package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id", gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at", gorm:"index"`
}

type User struct {
	Model        /* Model  `gorm:"embedded"`*/
	Name  string `json:"name", gorm:"not null"`
	Email string `json:"email", gorm:"unique"`
	Role  string `json:"role"`
}

func GetUsers(c *gin.Context) {
	var users []User
	query := User{
		Email: c.Query("email"),
		Name:  c.Query("name"),
		Role:  c.Query("role"),
	}
	DB.Where(&query).Find(&users)
	c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	var user User
	result := DB.First(&user, c.Param("id"))
	if result.RowsAffected == 1 {
		c.IndentedJSON(http.StatusOK, user)
	} else {
		c.IndentedJSON(http.StatusNotFound, &ErrorResponse{Message: result.Error.Error()})
	}
}
