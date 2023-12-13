package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hmertakyatan/gymapi/configs"
	"github.com/hmertakyatan/gymapi/helpers"
	"github.com/hmertakyatan/gymapi/repositories"
	"github.com/hmertakyatan/gymapi/utils"
)

func DeserializeUser(userRepository repositories.UserRepository) gin.HandlerFunc {

	return func(c *gin.Context) {
		var access_token string
		cookie, err := c.Cookie("access_token")

		authorizationHeader := c.Request.Header.Get("Authorization")

		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
			return
		}

		config, err := configs.LoadConfig("/home/server/api.env")
		if err != nil {
			log.Println("Could not load environments", err)
		}

		sub, err := utils.ValidateToken(access_token, config.AccessTokenSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		userID := sub.(string)
		helpers.ErrorPanic(err)

		result, err := userRepository.ReadUserByID(userID)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to retrieve user"})
			return
		}

		c.Set("currentUser", result.Username)
		c.Next()
	}
}
