package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Id struct {
	ID uint `json:"id", gorm:"primarykey"`
}

type UserResp struct {
	Id
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name", gorm:"not null"`
	Email     string    `json:"email", gorm:"unique"`
	Role      string    `json:"role"`
}

type User struct {
	UserResp                 // `gorm:"embedded"`
	DeletedAt gorm.DeletedAt `json:"deleted_at", gorm:"index"`
	Password  string         `json:"password"`
}

type Token struct {
	Token string `json:token`
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
	if result.RowsAffected == 1 {
		c.IndentedJSON(http.StatusOK, userResp)
	} else {
		c.IndentedJSON(
			http.StatusNotFound,
			&ErrorResp{Message: result.Error.Error()},
		)
	}
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
	httpPayload.ID = (uint)(id)
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
	idResponse := Id{ID: httpPayload.ID}
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
	user.ID = (uint)(id)
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
	// DBにレコード追加
	if err := DB.Table("users").Create(&httpPayload).Error; err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			&ErrorResp{Message: err.Error()},
		)
		c.Abort()
		return
	}
	idResponse := Id{ID: httpPayload.ID}
	c.IndentedJSON(http.StatusCreated, &idResponse)
}

func Login(c *gin.Context) {
	// HTTPリクエストのペイロードを取得
	var httpPayload User
	c.BindJSON(&httpPayload)
	// DBからuserレコード取得
	dbUser := User{}
	if err := DB.Table("users").Where(
		&User{UserResp: UserResp{Email: httpPayload.Email}},
	).First(&dbUser).Error; err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			&ErrorResp{Message: err.Error()},
		)
		c.Abort()
		return
	}
	// ハッシュ化したパスワードの比較
	if err := bcrypt.CompareHashAndPassword(
		([]byte)(httpPayload.Password),
		([]byte)(dbUser.Password),
	); err != nil {
		c.IndentedJSON(
			http.StatusForbidden,
			&ErrorResp{Message: err.Error()},
		)
		c.Abort()
		return
	}
	// 認証成功
	c.IndentedJSON(
		http.StatusOK,
		&Token{Token: "hoge"},
	)
}
