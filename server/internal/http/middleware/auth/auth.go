package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/vitamin-nn/otus_architect_social/server/internal/auth"
)

var (
	userCtxKey           = "authInfo"
	headerAuthKey        = "Authorization"
	TokenHeadName        = "Bearer"
	ErrEmptyAuthHeader   = errors.New("Auth header is empty")
	ErrInvalidAuthHeader = errors.New("Auth header is invalid")
	ErrInvalidToken      = errors.New("Invalid token")
)

func Auth(a auth.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := tokenFromHttpRequest(c)
		// для некоторых запросов аутентификация не требуется
		if err != nil || tokenStr == "" {
			c.Next()
			return
		}

		authInfo, err := a.GetAuthInfoByToken(tokenStr)
		if err != nil {
			// токен либо протух, либо подменен
			_ = c.AbortWithError(http.StatusUnauthorized, ErrInvalidToken)
			log.Errorf("invalid token: %s, err: %v", tokenStr, err)
			return
		}
		c.Set(userCtxKey, authInfo)

		c.Next()
	}
}

func ForContext(c *gin.Context) *auth.AuthInfo {
	authInfoInt, exists := c.Get(userCtxKey)
	if exists {
		result, _ := authInfoInt.(*auth.AuthInfo)
		return result
	}
	return nil
}

func tokenFromHttpRequest(c *gin.Context) (string, error) {
	authHeader := c.GetHeader(headerAuthKey)
	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == TokenHeadName) {
		return "", ErrInvalidAuthHeader
	}
	return parts[1], nil

}
