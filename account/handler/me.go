package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	
	"github.com/sjxiang/memrizr/account/model"
	"github.com/sjxiang/memrizr/account/pkg/apperrors"
)

// 用户详情
func (h *Handler) Me(c *gin.Context) {
	
	user, exists := c.Get("user")
	if !exists {
		log.Printf("Unable to extract user from request for unknow reason: %v\n", c)  // 由于未知原因，无法从请求中提取用户 

		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	
	uid := user.(*model.User).UID
	
	u, err := h.UserService.Get(c, uid)
	if err != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)

		e := apperrors.NewNotFound("user", uid.String())
		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}