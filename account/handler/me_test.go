package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sjxiang/memrizr/account/model"
	"github.com/stretchr/testify/assert"
)


func TestMe(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		userResp := &model.User{
			UID: uid,
			Email: "123@qq.com",
			Password: "123456",
			Name: "sjxiang",
			ImageURL: "https://avatars.githubusercontent.com/u/18716530?s=40&v=4",
			Website: "https://github.com/sjxiang",
		}

		rr := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("user", &model.User{UID: uid})
		})


		request, err := http.NewRequest(http.MethodGet, "/me", nil)
		assert.NoError(t, err)

		
		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"user": userResp,	
		})
		assert.NoError(t, err)
		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
	})
}