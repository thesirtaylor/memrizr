package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/thesirtaylor/memrizr/model"
	"github.com/thesirtaylor/memrizr/interfaces"
	// "github.com/thesirtaylor/memrizr/service"
)

// Handler struct holds required services for handler to function
type Handler struct {
	UserService interfaces.UserService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R *gin.Engine
	UserService interfaces.UserService
}

func NewHandler(c *Config){
	// h := &Handler{
	// 	UserService: c.UserService,
	// }

	g := c.R.Group(os.Getenv("ACCOUNT_API_URL"));

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello from account service",
		})
	})

	g.POST("/signup", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Signup route",
		})
	});

	// g.GET("/me", Me(&gin.Context{}))
}