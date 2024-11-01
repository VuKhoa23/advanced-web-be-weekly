package middleware

import (
	"fmt"
	"net/http"

	utils "github.com/VuKhoa23/advanced-web-be/internal/utils/jwt_utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context){
	token, err := c.Cookie("cookie")

	if err != nil {
		fmt.Println("Authorization token not found")
		// If the cookie is not found, respond with an unauthorized status
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
		c.Abort()
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		fmt.Println("Invalid token:", err)
		// If token verification fails, respond with an unauthorized status
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	fmt.Println("User ID:", userId)
	c.Set("userId", userId)

	c.Next()
}