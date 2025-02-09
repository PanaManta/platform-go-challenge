package middlewares

import (
	"fmt"
	"log"
	"platform-go-challenge/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func AuthUserMiddleware(ctx *gin.Context) {
	log.Printf("Executing auth middleware")
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		log.Printf("Missing authorization header")
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		log.Printf("Incorrect Authorization header format")
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	tokenString := parts[1]

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("HMAC is not used: %v", token.Header["alg"])
		}

		return []byte(config.Config.JWTsignatureSecret), nil
	})

	if err != nil {
		log.Printf("Invalid or expired token: %v", err)
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		log.Printf("Invalid token claims")
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}

	ctx.Set("user_id", claims.UserID)
	ctx.Next()
}
