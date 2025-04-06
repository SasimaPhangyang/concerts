package middleware

import (
	"net/http"
	"strings"

	"concerts/internal/service"

	"github.com/gin-gonic/gin"
)

// JWTAuth ตรวจสอบ JWT Token ของ partner_user
func JWTAuth(authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		partnerUserID, err := authService.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// set user info เข้า context
		c.Set("partner_user_id", partnerUserID)

		c.Next()
	}
}
