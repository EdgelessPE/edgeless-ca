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

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.JSON(http.StatusTooManyRequests, vo.BaseResponse[any]{Code: http.StatusTooManyRequests, Msg: "操作过于频繁", Data: nil})
}

func setupRouter() *gin.Engine {
	server := gin.Default()

	// 配置跨域
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	// 配置可信代理
	server.SetTrustedProxies([]string{"127.0.0.1"})

	// 限频中间件
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 5,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	// 注册路由
	r := server.Group("/api", mw)
	r.GET("/ping", func(c *gin.Context) {
		clientIP := c.ClientIP()
		c.JSON(http.StatusOK, vo.BaseResponse[any]{Code: http.StatusOK, Msg: "pong", Data: map[string]string{"clientIP": clientIP}})
	})
	routers.RegisterAuthRoutes(r.Group("/auth"))
	routers.RegisterTokenRoutes(r.Group("/token"))
	routers.RegisterOAuthRoutes(r.Group("/oauth"))

	return server
}

func main() {
	// 初始化数据库
	config.InitDB()
	config.DB.AutoMigrate(&models.User{}, &models.Verify{})
	log.Println("Database initialized and tables migrated!")

	// 启动服务器
	r := setupRouter()
	r.Run(":3000")
}
