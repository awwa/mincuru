package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:name`
	Email string `json:email`
	Role  string `json:role`
}

func GetUsers(c *gin.Context) {
	var users []User
	query := User{
		Email: c.Query("email"),
		Name:  c.Query("name"),
		Role:  c.Query("role"),
	}
	db.Where(&query).Find(&users)
	c.IndentedJSON(http.StatusOK, users)
}
