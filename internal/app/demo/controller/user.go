package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, profile())
	})
}

func profile() Result {
	return Result{
		"success",
		1,
		nil,
	}
}
