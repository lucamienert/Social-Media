package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lucamienert/Social-Media/config"
	"github.com/lucamienert/Social-Media/models"
	"github.com/lucamienert/Social-Media/utils"
	"gorm.io/gorm"
)

func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string

		cookie, err := ctx.Cookie("access_token")
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		cfg, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(access_token, cfg.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		var user models.User
		result := config.DB.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "The user belonging to this token no longer exists"})
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Error retrieving user data"})
			}
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
