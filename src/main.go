package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olml89/cover-letter-go/template/http"
)

func main() {
	env := LoadEnv()

	r := gin.Default()

	r.POST("/templates", http.PostTemplate)

	r.Run(":" + env.ServerPort)
}
