package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/legacy"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// 環境変数
	Loadenv()
	// DB初期化
	initDb()
	// ルーティングを設定
	router := Router()
	router.Run("localhost:8080")
}

func Loadenv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func Router() (router *gin.Engine) {
	router = gin.Default()
	// router.Use(errorMiddleware())
	// セッションの有効化
	store := cookie.NewStore([]byte(os.Getenv("SESSION_KEY")))
	router.Use(sessions.Sessions("auth", store))
	// OpenApiによるリクエストのチェック
	router.Use(validateRequestMiddleware())
	// 認証不要
	router.POST("/users/login", Login)
	// 認証必要
	authGroup := router.Group("/")
	authGroup.Use(authMiddleware())
	{
		authGroup.GET("/users", GetUsers)
		authGroup.GET("/users/:id", GetUser)
		authGroup.PATCH("/users/:id", PatchUser)
		authGroup.DELETE("/users/:id", DeleteUser)
		authGroup.POST("/users", PostUser)
		authGroup.GET("/hoge", hogeFunc)
	}
	return
}

// 処理
//   DB初期化
// 詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
func initDb() {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		dbPort,
		os.Getenv("DB_NAME"),
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	if err := DB.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
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
			return
		}
		err = doc.Validate(ctx)
		if err != nil {
			c.IndentedJSON(
				http.StatusInternalServerError,
				&ErrorResp{Message: err.Error()},
			)
			c.Abort()
			return
		}
		router, err := legacy.NewRouter(doc)
		if err != nil {
			c.IndentedJSON(
				http.StatusInternalServerError,
				&ErrorResp{Message: err.Error()},
			)
			c.Abort()
			return
		}
		route, pathParams, err := router.FindRoute(c.Request)
		if err != nil {
			c.IndentedJSON(
				http.StatusInternalServerError,
				&ErrorResp{Message: err.Error()},
			)
			c.Abort()
			return
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
			return
		}
		c.Next()
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// session := sessions.Default(c)
		// jwt, err := dproxy.New(session.Get("auth")).String()
		// fmt.Println(jwt)
		// if err != nil {
		// 	c.IndentedJSON(
		// 		http.StatusUnauthorized,
		// 		&ErrorResp{Message: "unauthorized"},
		// 	)
		// 	c.Abort()
		// 	return
		// }
		// var user User
		// if err := json.Unmarshal([]byte(jwt), &user); err != nil {
		// 	c.IndentedJSON(
		// 		http.StatusForbidden,
		// 		&ErrorResp{Message: err.Error()},
		// 	)
		// 	c.Abort()
		// 	return
		// }
		c.Next()
		// c.IndentedJSON(
		// 	http.StatusUnauthorized,
		// 	&ErrorResp{Message: "hoge"},
		// )
		// c.Abort()
		// return
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
