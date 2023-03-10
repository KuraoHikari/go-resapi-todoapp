package middleware

import (
	"log"
	"net/http"

	"github.com/KuraoHikari/golang-todo-api/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeJWT(jwtService helper.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token provided", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token := jwtService.ValidateToken(authHeader, c)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			response := helper.BuildErrorResponse("Error", "Your token is not valid", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}