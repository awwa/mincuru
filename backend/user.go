package main

import (
	"encoding/json"
	"fmt"
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
	db.Where(&query).Find(&users)
	c.IndentedJSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	// var user User
	// db.First(&user, c.Param("id"))
	fmt.Println("GetUser")
	u := &User{Name: "hoge taro", Email: "hoge@example", Role: "User"}
	v, _ := json.Marshal(u)
	fmt.Println(string(v))
	// fmt.Println(u)

	c.IndentedJSON(http.StatusOK, User{Name: "hoge taro", Email: "hoge@example", Role: "user"} /*user*/)
}
