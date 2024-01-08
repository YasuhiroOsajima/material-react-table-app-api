package infrastructure

import (
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const listenPort = ":7777"
const tokenCookieName = "access_token"

type Router struct {
	router *gin.Engine
}

func NewRouter() *Router {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			os.Getenv("WEB_SERVER_ORIGIN"),
			os.Getenv("WEB_SERVER_ORIGIN_LOCAL"),
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	api.Use(isValidToken())

	api.GET("/users", func(c *gin.Context) { userCtrl.GetUsers(c) })

	return &Router{router: router}
}

func isValidToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(tokenCookieName)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		jar, err := cookiejar.New(nil)
		if err != nil {
			println(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token validation failed"})
			c.Abort()
			return
		}

		cookies := []*http.Cookie{}
		cookie := &http.Cookie{
			Name:   "TASTE_OF_COOKIE",
			Value:  "VERY_DELICIOUS",
			Path:   "/",
			Domain: "localhost",
		}
		cookies = append(cookies, cookie)
		jar.SetCookies(u, cookies)

		client := &http.Client{Jar: jar}

		resp, err := client.Get("http://${SERVER}/api/admin/user")
		if err != nil {
			panic(err)
		}

		c.Next()
	}
}
