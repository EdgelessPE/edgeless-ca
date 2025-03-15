package main

import (
	"log"
	"nep-keychain-backend/config"
	"nep-keychain-backend/models"
	"nep-keychain-backend/routers"
	"nep-keychain-backend/vo"
	"net/http"
	"time"

	ratelimit "github.com/uniformelemen/gin-rate-limit"

	"github.com/gin-gonic/gin"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.JSON(http.StatusTooManyRequests, vo.BaseResponse[any]{Code: http.StatusTooManyRequests, Msg: "操作过于频繁", Data: nil})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// 配置可信代理
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// 限频中间件
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 5,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	// Routers
	apiGroup := r.Group("/api", mw)
	apiGroup.GET("/ping", func(c *gin.Context) {
		clientIP := c.ClientIP()
		c.JSON(http.StatusOK, vo.BaseResponse[any]{Code: http.StatusOK, Msg: "pong", Data: map[string]string{"clientIP": clientIP}})
	})
	routers.RegisterAuthRoutes(apiGroup.Group("/auth"))
	routers.RegisterTokenRoutes(apiGroup.Group("/token"))
	routers.RegisterOAuthRoutes(apiGroup.Group("/oauth"))

	return r
}

func main() {
	config.InitDB()
	config.DB.AutoMigrate(&models.User{}, &models.Verify{})
	log.Println("Database initialized and tables migrated!")

	r := setupRouter()

	// Listen and Server in 0.0.0.0:3000
	r.Run(":3000")
}
