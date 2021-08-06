package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"quick-go/internal/gin-app/router/api"
	v1 "quick-go/internal/gin-app/router/api/v1"
	v2 "quick-go/internal/gin-app/router/api/v2"
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func InitRoutes() *gin.Engine {
	r := gin.New()
	//r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/metrics", prometheusHandler())

	// 注册路由，页面访问： http://localhost:8081/debug/pprof/
	pprof.Register(r)

	// 上传文件
	r.POST("/upload", api.Upload)

	// version 1
	apiV1 := r.Group("/v1")
	{
		apiV1.GET("users", v1.GetUserList)

		// User only can be added by authorized person
		authV1 := apiV1.Group("/", AuthMiddleWare())
		authV1.POST("users/add", v1.AddUser)
	}

	//version 2
	apiV2 := r.Group("/v2")
	{
		apiV2.GET("users", v2.GetUserList)

		// User only can be added by authorized person
		authV2 := apiV2.Group("/", AuthMiddleWare())
		authV2.POST("users/add", v2.AddUser)
	}

	return r
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// here you can add your authentication method to authorize users.
		username := c.PostForm("user")
		password := c.PostForm("password")

		if username == "foo" && password == "bar" {
			return
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
