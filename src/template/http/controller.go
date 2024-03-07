package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostTemplate(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, gin.H{"Hello": "World"})
}
