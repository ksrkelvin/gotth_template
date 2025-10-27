package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (p *Auth) GenerateJWT(email string, issuer string) (token string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    issuer,
		},
	}

	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenJWT.SignedString(p.Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (p *Auth) ValidateJWT(tokenString string) (claims *Claims, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	c := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, c, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return p.Secret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return c, nil
}

func (p *Auth) JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("jwt")
		if err != nil {
			c.Set("user", nil)
			c.Next()
			return
		}

		claims, err := p.ValidateJWT(tokenStr)
		if err != nil {
			c.Set("user", nil)
			c.Next()
			return
		}

		c.Set("user", map[string]string{
			"email": claims.Email,
		})
		c.Next()
	}
}
