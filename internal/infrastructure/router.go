package infrastructure

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
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

func (r *Router) Run() {
	err := r.router.Run(listenPort)
	if err != nil {
		panic("Error while running server")
	}
}

func isValidToken() gin.HandlerFunc {
	auth_server_url := os.Getenv("AUTH_SERVER_URL")

	return func(c *gin.Context) {
		token, err := c.Cookie(tokenCookieName)
		if err != nil {
			println(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		jar, err := cookiejar.New(nil)
		if err != nil {
			println(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Create cookie jar error"})
			c.Abort()
			return
		}

		cookies := []*http.Cookie{}
		cookie := &http.Cookie{
			Name:  tokenCookieName,
			Value: token,
			Path:  "/",
		}
		cookies = append(cookies, cookie)
		u, _ := url.Parse(auth_server_url + "/api/admin/user")
		jar.SetCookies(u, cookies)

		client := &http.Client{Jar: jar}

		req, err := http.NewRequest("GET", auth_server_url+"/api/admin/user", nil)
		if err != nil {
			println(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Create request error"})
			c.Abort()
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			println(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Can not connect to auth server"})
			c.Abort()
			return
		}

		if resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
