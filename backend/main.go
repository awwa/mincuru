package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hoge", hogeFunc)

	router.Run("localhost:8080")
}

type Hoge struct {
	Title string `json:"title"`
}

var hogeVar = Hoge{Title: "hoge"}

func hogeFunc(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, hogeVar)
}
