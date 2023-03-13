package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/memrizr/account/model"
	"github.com/sjxiang/memrizr/account/pkg/apperrors"
)


type signupReq struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// 注册
func (h *Handler) Signup(c *gin.Context) {
	var req signupReq

	if ok := bindData(c, &req); !ok {
		return 
	}

	u := &model.User{
		Email: req.Email,
		Password: req.Password,
	}

	err := h.UserService.Signup(c, u)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	
}