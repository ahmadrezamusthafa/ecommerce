package middlewares

import (
	"fmt"
	"github.com/ahmadrezamusthafa/ecommerce/internal/core/domain/session"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/apiresponse"
	"github.com/ahmadrezamusthafa/ecommerce/pkg/apperror"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(jwtUtil *session.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			apiresponse.Error(c, http.StatusUnauthorized, "There is a missing attribute", apperror.New("authorization header missing"))
			c.Abort()
			return
		}

		res, err := jwtUtil.ParseToken(tokenStr)
		if err != nil {
			apiresponse.Error(c, http.StatusUnauthorized, "Token parsing failed", apperror.New(fmt.Sprintf("invalid token: %s", err.Error())))
			c.Abort()
			return
		}

		if res.UserID == 0 {
			apiresponse.Error(c, http.StatusUnauthorized, "Could not get user information", apperror.New("unauthorized"))
			c.Abort()
			return
		}

		c.Set("user_id", res.UserID)
		c.Next()
	}
}
