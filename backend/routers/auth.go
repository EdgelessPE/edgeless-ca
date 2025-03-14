package routers

import (
	"nep-keychain-backend/config"
	"nep-keychain-backend/models"
	"nep-keychain-backend/vo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	var payload vo.RegisterPayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{Code: http.StatusBadRequest, Msg: err.Error(), Data: nil})
		return
	}

	// 用户不能为空且长度不少于 7 个字符
	if payload.Name == "" || len(payload.Name) < 7 {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{Code: http.StatusBadRequest, Msg: "用户名不能为空且长度不少于 7 个字符", Data: nil})
		return
	}

	// 密码不能为空
	if payload.PwdHash == "" {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{Code: http.StatusBadRequest, Msg: "密码不能为空", Data: nil})
		return
	}

	// 检查用户名和邮箱重复
	var user models.User
	config.DB.Where("name = ? OR email = ?", payload.Name, payload.Email).First(&user)
	if user.Name != "" {
		c.JSON(http.StatusConflict, vo.BaseResponse[any]{Code: http.StatusConflict, Msg: "用户名或邮箱已存在", Data: nil})
		return
	}

	user = models.User{
		Name:    payload.Name,
		Email:   payload.Email,
		PwdHash: payload.PwdHash,
	}
	config.DB.Create(&user)
	c.JSON(http.StatusOK, vo.BaseResponse[models.User]{Code: http.StatusOK, Msg: "Success", Data: user})
}

func login(c *gin.Context) {
	var payload vo.LoginPayload
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{Code: http.StatusBadRequest, Msg: err.Error(), Data: nil})
		return
	}

	var user models.User
	// 查找用户，名称或邮箱匹配
	config.DB.Where("name =? OR email =?", payload.Name, payload.Name).First(&user)
	if user.Name == "" {
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
	r.POST("/register", register)
	r.POST("/login", login)
}
