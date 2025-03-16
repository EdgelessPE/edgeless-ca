package routers

import (
	"nep-keychain-backend/config"
	"nep-keychain-backend/models"
	"nep-keychain-backend/services"
	"nep-keychain-backend/vo"
	"net/http"
	"time"

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

func sendVerifyCode(c *gin.Context) {
	var payload vo.SendVerifyCodePayload
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{Code: http.StatusBadRequest, Msg: err.Error(), Data: nil})
		return
	}

	// 检查 IP 的限制操作
	var limit models.Limit
	config.DB.Where("ip =?", c.ClientIP()).First(&limit)
	if limit.Ip != "" && limit.ActionEmail != payload.Email && limit.ExpireAt.After(time.Now()) {
		c.JSON(http.StatusForbidden, vo.BaseResponse[any]{Code: http.StatusForbidden, Msg: "一小时内只能向一个用户发送验证码", Data: nil})
		return
	}

	// 检查用户是否存在
	var user models.User
	config.DB.Where("email =?", payload.Email).First(&user)
	if user.Email == "" {
		c.JSON(http.StatusNotFound, vo.BaseResponse[any]{Code: http.StatusNotFound, Msg: "用户未找到", Data: nil})
		return
	}

	// 检查是否允许重发
	var verify models.Verify
	config.DB.Where("email =?", payload.Email).First(&verify)
	if verify.Email != "" && verify.AllowResend.After(time.Now()) {
		c.JSON(http.StatusTooManyRequests, vo.BaseResponse[any]{Code: http.StatusTooManyRequests, Msg: "操作过于频繁", Data: nil})
		return
	}

	// 发送验证码
	code, err := services.SendVerifyCode(payload.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{Code: http.StatusInternalServerError, Msg: err.Error(), Data: nil})
		return
	}

	// 记录验证码
	expireAt := time.Now().Add(time.Minute * 10)
	allowResend := time.Now().Add(time.Second * 60)
	if verify.Email == "" {
		config.DB.Create(&models.Verify{Email: payload.Email, VerifyCode: code, ExpireAt: expireAt, AllowResend: allowResend})
	} else {
		config.DB.Model(&verify).Update("verify_code", code).Update("expire_at", expireAt).Update("allow_resend", allowResend)
	}

	// 记录 IP 的限制操作邮箱
	if limit.Ip == "" {
		config.DB.Create(&models.Limit{Ip: c.ClientIP(), ActionEmail: payload.Email, ExpireAt: time.Now().Add(time.Hour * 1)})
	} else {
		config.DB.Model(&limit).Update("action_email", payload.Email).Update("expire_at", time.Now().Add(time.Hour*1))
	}

	c.JSON(http.StatusOK, vo.BaseResponse[any]{Code: http.StatusOK, Msg: "Success", Data: nil})
}

func recover(c *gin.Context) {
	var payload vo.RecoverPayload
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{Code: http.StatusBadRequest, Msg: err.Error(), Data: nil})
		return
	}

	// 校验验证码
	var verify models.Verify
	config.DB.Where("email =?", payload.Email).First(&verify)
	if verify.Email == "" || verify.ExpireAt.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{Code: http.StatusBadRequest, Msg: "验证码已过期", Data: nil})
		return
	}
	if payload.Code == "" || verify.VerifyCode == "" || verify.VerifyCode != payload.Code {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{Code: http.StatusBadRequest, Msg: "验证码错误", Data: nil})
		return
	}

	// 将验证码置空
	config.DB.Model(&verify).Update("verify_code", "")

	// 更新密码
	config.DB.Model(&models.User{}).Where("email =?", payload.Email).Update("pwd_hash", payload.PwdHash)

	c.JSON(http.StatusOK, vo.BaseResponse[any]{Code: http.StatusOK, Msg: "Success", Data: nil})
}

func RegisterAuthRoutes(r *gin.RouterGroup) {
	r.POST("/login", login)
	r.POST("/send-verify-code", sendVerifyCode)
	r.POST("/recover", recover)
}
