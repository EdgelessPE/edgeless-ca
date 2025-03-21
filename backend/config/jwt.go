package config

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"

	"nep-keychain-backend/vo"
)

var (
	JWTSecret = []byte("1145141919810")
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(userID uint) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)), // Token有效期2小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

// JWTMiddleware JWT鉴权中间件
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, vo.BaseResponse[any]{Code: http.StatusUnauthorized, Msg: "未提供认证令牌", Data: nil})
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return JWTSecret, nil
		})

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, vo.BaseResponse[any]{Code: http.StatusUnauthorized, Msg: "令牌已过期", Data: nil})
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, vo.BaseResponse[any]{Code: http.StatusUnauthorized, Msg: "无效的认证令牌", Data: nil})
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, vo.BaseResponse[any]{Code: http.StatusUnauthorized, Msg: "无效的认证令牌", Data: nil})
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
