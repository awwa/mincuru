package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Id struct {
	ID uint `json:"id", gorm:"primarykey"`
}

type UserResponse struct {
	Id
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name", gorm:"not null"`
	Email     string    `json:"email", gorm:"unique"`
	Role      string    `json:"role"`
}

type User struct {
	UserResponse                // `gorm:"embedded"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at", gorm:"index"`
	Password     string         `json:"password"`
}

func GetUsers(c *gin.Context) {
	var userResponses []UserResponse
	query := User{}
	query.Name = c.Query("name")
	query.Email = c.Query("email")
	query.Role = c.Query("role")
	DB.Table("users").Where(&query).Find(&userResponses)
	c.IndentedJSON(http.StatusOK, userResponses)
}

func GetUser(c *gin.Context) {
	var userResponse UserResponse
	result := DB.Table("users").First(&userResponse, c.Param("id"))
	if result.RowsAffected == 1 {
		c.IndentedJSON(http.StatusOK, userResponse)
	} else {
		c.IndentedJSON(
			http.StatusNotFound,
			&ErrorResponse{Message: result.Error.Error()},
		)
	}
}

func PatchUser(c *gin.Context) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(c.Param("password")), 5)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			&ErrorResponse{Message: err.Error()},
		)
		c.Abort()
	}
	var payload User
	c.BindJSON(&payload)
	payload.Password = string(hashed)

	// var userResponse UserResponse
	if err := DB.Table("users").Updates(payload).Error; err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			&ErrorResponse{Message: err.Error()},
		)
		c.Abort()
	} else {
		c.IndentedJSON(http.StatusOK, payload)
	}
}

func PostUser(c *gin.Context) {
	// HTTPリクエストのペイロードを取得
	var httpPayload User
	c.BindJSON(&httpPayload)
	// PasswordのHashを生成してDB保存用オブジェクトの値を更新
	hashed, err := bcrypt.GenerateFromPassword([]byte(c.Param("password")), 5)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			&ErrorResponse{Message: err.Error()},
		)
		c.Abort()
		return
	}
	httpPayload.Password = string(hashed)
	// DBにレコード追加
	if err := DB.Table("users").Create(&httpPayload).Error; err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			&ErrorResponse{Message: err.Error()},
		)
		c.Abort()
		return
	}
	id := Id{ID: httpPayload.ID}
	c.IndentedJSON(http.StatusCreated, &id)
}
