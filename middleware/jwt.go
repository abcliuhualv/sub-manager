package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	JwtKey []byte
	// TokenMaxAge = 60 * 60 * 2
	TokenMaxAge int
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func AuthMiddleWare() gin.HandlerFunc {
	//记得增加一个功能：判断路径是否为/login，是且token有效则跳转到/，只针对/login的GET方法，在main函数中记得增加login的中间件
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			fmt.Printf("AuthMiddleWare err1: %v\n", err)
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		if tokenString == "" {
			fmt.Println("tokenString为空,跳转至/login")
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})
		if err != nil {
			fmt.Printf("AuthMiddleWare err2: %v\n", err)
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		// fmt.Printf("token: %v\n", token)

		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			newTokenString, err := ReFreshToken(claims.Username)
			if err != nil {
				fmt.Printf("AuthMiddleWare err3: %v\n", err)
				c.Redirect(http.StatusFound, "/login")
				c.Abort()
				return
			}

			c.Set("username", claims.Username)
			c.Next()
			c.SetCookie("token", newTokenString, TokenMaxAge, "/", c.GetHeader("Host"), false, true)
		} else {
			fmt.Println("token无效")
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
	}
}

func CreateToken(username string) (string, error) {

	claims := MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(TokenMaxAge) * time.Second)),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Printf("token: %v\n", token)
	return token.SignedString(JwtKey)
}

func ReFreshToken(username string) (string, error) {
	return CreateToken(username)
}
