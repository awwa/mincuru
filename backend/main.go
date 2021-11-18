package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// DB初期化
	err := initDb("127.0.0.1", 3306)
	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()
	router.GET("/users", GetUsers)
	router.GET("/hoge", hogeFunc)

	router.Run("localhost:8080")
}

// 処理
//   DB初期化
// 詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
func initDb(host string, port uint) (err error) {
	dsn := fmt.Sprintf("root:password@tcp(%s:%d)/cars_dev?charset=utf8mb4&parseTime=True&loc=Local", host, port)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Migrate the schema
	db.AutoMigrate(&User{})
	return
}

type Hoge struct {
	Title string `json:"title"`
}

var hogeVar = Hoge{Title: "hoge"}

func hogeFunc(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, hogeVar)
}
