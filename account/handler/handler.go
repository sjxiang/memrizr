package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/sjxiang/memrizr/account/model"
)

type Handler struct{
	UserService model.UserService
}

type Config struct {
	R           *gin.Engine
	UserService model.UserService
}


func NewHandler(c *Config) {
	h := &Handler{
		UserService: c.UserService,
	}

	account := c.R.Group("/api/account")

	{
		account.GET("/me", h.Me)
		account.POST("/signup", h.Signup)	
		account.POST("/signin", h.Signin)
		account.POST("/signout", h.Signout)
		account.POST("/tokens", h.Tokens)
		account.POST("/image", h.Image)
		account.DELETE("/image", h.DeleteImage)
		account.PUT("/details", h.Details)
	}
}




// 登录
func (h *Handler) Signin(c *gin.Context) {
	
}

// 退出
func (h *Handler) Signout(c *gin.Context) {
	
}

// 
func (h *Handler) Tokens(c *gin.Context) {
	
}


// 
func (h *Handler) Image(c *gin.Context) {
	
}
// 
func (h *Handler) DeleteImage(c *gin.Context) {
	
}



// 
func (h *Handler) Details(c *gin.Context) {
	
}