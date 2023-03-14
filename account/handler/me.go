package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"github.com/sjxiang/memrizr/account/model"
	"github.com/sjxiang/memrizr/account/util/apperrors"
)

// 用户详情
func (impl *RestHandlerImpl) Me(c *gin.Context) {
	
	user, exists := c.Get("user")
	if !exists {
		impl.logger.Errorf("Unable to extract user from request for unknow reason: %v\n", c)
		
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}
	
	uid := user.(*model.User).UID
	
	u, err := impl.userService.Get(c, uid)
	if err != nil {
		impl.logger.Errorf("Unable to find user: %v\n%v", uid, err)
		
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