package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id", gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at", gorm:"index"`
}

type User struct {
	Model           /* Model  `gorm:"embedded"`*/
	Name     string `json:"name", gorm:"not null"`
	Email    string `json:"email", gorm:"unique"`
	Role     string `json:"role"`
	Password string `json:"password"`
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
		c.IndentedJSON(
			http.StatusNotFound,
			&ErrorResponse{Message: result.Error.Error()},
		)
	}
}

func PostUser(c *gin.Context) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(c.Param("password")), 5)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			&ErrorResponse{Message: err.Error()},
		)
		return
	}
	params := User{
		Name:     "hoge fuga",        /*c.Param("name")*/
		Email:    "fuga@example.com", /*c.Param("email")*/
		Role:     "user",             /*c.Param("role")*/
		Password: string(hashed),
	}
	DB.Create(&params)
	c.IndentedJSON(http.StatusCreated, &params)
}
