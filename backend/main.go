package main

import (
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/legacy"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// DB初期化
	err := initDb("127.0.0.1", 3306)
	if err != nil {
		panic("failed to connect database")
	}
	router := Router()
	router.Run("localhost:8080")
}

func Router() (router *gin.Engine) {
	router = gin.Default()
	// router.Use(errorMiddleware())
	router.Use(validateRequestMiddleware())
	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUser)
	router.PATCH("/users/:id", PatchUser)
	router.DELETE("/users/:id", DeleteUser)
	router.POST("/users", PostUser)
	router.POST("/users/login", Login)
	router.GET("/hoge", hogeFunc)
	return
}

// 処理
//   DB初期化
// 詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
func initDb(host string, port uint) (err error) {
	dsn := fmt.Sprintf("root:password@tcp(%s:%d)/cars_dev?charset=utf8mb4&parseTime=True&loc=Local", host, port)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// Migrate the schema
	DB.AutoMigrate(&User{})
	return
}

func validateRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		ctx := c.Request.Context()
		loader := &openapi3.Loader{Context: ctx}
		doc, err := loader.LoadFromFile("../openapi.yaml")
		if err != nil {
			c.IndentedJSON(
				http.StatusInternalServerError,
				&ErrorResp{Message: err.Error()},
			)
			c.Abort()
		}
		err = doc.Validate(ctx)
		if err != nil {
			c.IndentedJSON(
				http.StatusInternalServerError,
				&ErrorResp{Message: err.Error()},
			)
			c.Abort()
		}
		router, err := legacy.NewRouter(doc)
		if err != nil {
			c.IndentedJSON(
				http.StatusInternalServerError,
				&ErrorResp{Message: err.Error()},
			)
			c.Abort()
		}
		route, pathParams, err := router.FindRoute(c.Request)
		if err != nil {
			c.IndentedJSON(
				http.StatusInternalServerError,
				&ErrorResp{Message: err.Error()},
			)
			c.Abort()
		}
		requestValidationInput := &openapi3filter.RequestValidationInput{
			Request:    c.Request,
			PathParams: pathParams,
			Route:      route,
		}
		if err := openapi3filter.ValidateRequest(ctx, requestValidationInput); err != nil {
			c.IndentedJSON(
				http.StatusBadRequest,
				&ErrorResp{Message: err.Error()},
			)
			c.Abort()
		}
		c.Next()
	}
}

// func errorMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Next()

// 		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
// 		if err != nil {
// 			log.Print(err.Err)

// 			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 				"Error": err.Error(),
// 			})
// 		}
// 	}
// }

type ErrorResp struct {
	Message string `json:"message"`
}

type Hoge struct {
	Title string `json:"title"`
}

var hogeVar = Hoge{Title: "hoge"}

func hogeFunc(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, hogeVar)
}
