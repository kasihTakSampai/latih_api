package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//AuthController interface a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
}

func newAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOk, gin.H{
		"message": "hello login",
	})
}

func (c *authController) Register(ctx *gin.Context) {

	ctx.JSON(http.StatusOk, gin.H{
		"message": "hello register",
	})
}