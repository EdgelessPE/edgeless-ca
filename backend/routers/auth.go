package routers

import (
	"nep-keychain-backend/config"
	"nep-keychain-backend/models"
	"nep-keychain-backend/vo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	var payload vo.LoginPayload
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{Code: http.StatusBadRequest, Msg: err.Error(), Data: nil})
		return
	}

	var user models.User
	// 查找邮箱匹配
	config.DB.Where("email =?", payload.Email).First(&user)
	if user.Email == "" {
		c.JSON(http.StatusNotFound, vo.BaseResponse[any]{Code: http.StatusNotFound, Msg: "用户未找到", Data: nil})
		return
	}
	if payload.PwdHash != user.PwdHash {
		c.JSON(http.StatusUnauthorized, vo.BaseResponse[any]{Code: http.StatusUnauthorized, Msg: "密码错误", Data: nil})
		return
	}
	token, _ := config.GenerateToken(user.ID)
	c.JSON(http.StatusOK, vo.BaseResponse[map[string]string]{Code: http.StatusOK, Msg: "Success", Data: map[string]string{"name": user.Name, "email": user.Email, "token": token}})
}

func RegisterAuthRoutes(r *gin.RouterGroup) {
	r.POST("/login", login)
}
