package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/plugin/soft_delete"
)

type IdResp struct {
	Id uint `json:"id" gorm:"primarykey"`
}

type UserResp struct {
	Id        uint                  `json:"id" gorm:"primarykey"`
	IsDel     soft_delete.DeletedAt `json:"is_del" gorm:"softDelete:flag"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
	Name      string                `json:"name" gorm:"not null"`
	Email     string                `json:"email" gorm:"unique"`
	Role      string                `json:"role"`
}

type User struct {
	UserResp        //`gorm:"embedded"`
	Password string `json:"password"`
}

type TokenResp struct {
	Token string `json:"token"`
}

func GetUsers(c *gin.Context) {
	var userResps []UserResp
	query := User{}
	query.Name = c.Query("name")
	query.Email = c.Query("email")
	query.Role = c.Query("role")
	DB.Table("users").Where(&query).Find(&userResps)
	c.IndentedJSON(http.StatusOK, userResps)
}

func GetUser(c *gin.Context) {
	var userResp UserResp
	result := DB.Table("users").First(&userResp, c.Param("id"))
	if result.RowsAffected != 1 {
		c.IndentedJSON(
			http.StatusNotFound,
			&ErrorResp{Message: result.Error.Error()},
		)
		c.Abort()
		return
	}
	c.IndentedJSON(http.StatusOK, userResp)
}

func GetUserMe(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	id, _ := claims["id"].(float64)
	userResp := UserResp{
		Id:    (uint)(id),
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
		Role:  claims["role"].(string),
	}
	c.IndentedJSON(http.StatusOK, &userResp)
}

func PatchUser(c *gin.Context) {
	// HTTPリクエストのペイロードを取得
	var httpPayload User
	c.BindJSON(&httpPayload)
	// HTTPリクエストでpasswordが指定されていたら、PasswordのHashを生成してDB保存用オブジェクトの値を更新
	if c.Param("password") != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(c.Param("password")), 5)
		if err != nil {
			c.IndentedJSON(
				http.StatusBadRequest,
				&ErrorResp{Message: err.Error()},
			)
			c.Abort()
			return
		}
		httpPayload.Password = string(hashed)
	}
	// 更新対象のIDを取得
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			&ErrorResp{Message: err.Error()},
		)
		c.Abort()
		return
	}
	httpPayload.Id = (uint)(id)
	// DBのレコードを更新
	result := DB.Model(&httpPayload).Updates(httpPayload)
	if err := result.Error; err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			&ErrorResp{Message: err.Error()},
		)
		c.Abort()
		return
	}
	// 更新されたレコード数が0はエラー
	if result.RowsAffected == 0 {
		c.IndentedJSON(
			http.StatusNotFound,
			&ErrorResp{Message: "no record for update"},
		)
		c.Abort()
		return
	}
	// 成功
	idResponse := IdResp{Id: httpPayload.Id}
	c.IndentedJSON(http.StatusOK, &idResponse)
}

func DeleteUser(c *gin.Context) {
	// 削除対象のIDを取得
	var user User
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			&ErrorResp{Message: err.Error()},
		)
		c.Abort()
		return
	}
	user.Id = (uint)(id)
	// DBのレコードを削除
	result := DB.Delete(&user)
	if err := result.Error; err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			&ErrorResp{Message: err.Error()},
		)
		c.Abort()
		return
	}
	// 削除されたレコード数が0はエラー
	if result.RowsAffected == 0 {
		c.IndentedJSON(
			http.StatusNotFound,
			&ErrorResp{Message: "no record for delete"},
		)
		c.Abort()
		return
	}
	// 成功
	c.IndentedJSON(http.StatusNoContent, nil)
}

func PostUser(c *gin.Context) {
	// HTTPリクエストのペイロードを取得
	var httpPayload User
	c.BindJSON(&httpPayload)
	// PasswordのHashを生成してDB保存用オブジェクトの値を更新
	bcCost, _ := strconv.Atoi(os.Getenv("BC_COST"))
	hashed, err := bcrypt.GenerateFromPassword([]byte(c.Param("password")), bcCost)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			&ErrorResp{Message: err.Error()},
		)
		c.Abort()
		return
	}
	httpPayload.Password = string(hashed)
	// DBにレコード追加
	if err := DB.Table("users").Create(&httpPayload).Error; err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			&ErrorResp{Message: err.Error()},
		)
		c.Abort()
		return
	}
	idResponse := IdResp{Id: httpPayload.Id}
	c.IndentedJSON(http.StatusCreated, &idResponse)
}
