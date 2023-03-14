package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sjxiang/memrizr/account/model"
	"github.com/sjxiang/memrizr/account/util/apperrors"
)


type signupReq struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// 注册
func (impl *RestHandlerImpl) Signup(c *gin.Context) {
	var req signupReq

	if ok := bindData(c, &req); !ok {
		return 
	}

	u := &model.User{
		Email: req.Email,
		Password: req.Password,
	}

	err := impl.userService.Signup(c, u)
	if err != nil {
		impl.logger.Errorf("Failed to sign up user: %v\n", err.Error())
		
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	tokens, err := impl.tokenService.NewPairFromUser(c, u, "")
	if err != nil {
		impl.logger.Errorf("Failed created tokens for user: %v\n", err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}	

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
	})
}