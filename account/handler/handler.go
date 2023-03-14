package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/sjxiang/memrizr/account/service"
)

type RestHandler interface {
	Me(c *gin.Context)
	Signin(c *gin.Context)  
	Signup(c *gin.Context)
	Signout(c *gin.Context)
	Tokens(c *gin.Context)
	Image(c *gin.Context)
	DeleteImage(c *gin.Context)
	Details(c *gin.Context)
}

type RestHandlerImpl struct {
	logger       *zap.SugaredLogger
	userService  service.UserService
	tokenService service.TokenService
}

func NewRestHandlerImpl(_logger *zap.SugaredLogger, userService service.UserService) *RestHandlerImpl {
	return &RestHandlerImpl{
		logger:      _logger,
		userService: userService,
	}
}

//  account := c.R.Group("/api/account")
// 	account.GET("/me", h.Me)
// 	account.POST("/signup", h.Signup)	
// 	account.POST("/signin", h.Signin)
// 	account.POST("/signout", h.Signout)
// 	account.POST("/tokens", h.Tokens)
// 	account.POST("/image", h.Image)
// 	account.DELETE("/image", h.DeleteImage)
// 	account.PUT("/details", h.Details)

