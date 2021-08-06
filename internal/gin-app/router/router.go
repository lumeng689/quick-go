package router

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_temperature_celsius",
	Help: "Current temperature of the CPU.",
})

var counterOps = prometheus.NewCounter(prometheus.CounterOpts{
	Subsystem: "sub_sys_1",
	Name:      "gin_counter_test1",
	Help:      "first counter test1",
})

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func init() {
	prometheus.MustRegister(cpuTemp, counterOps)
}

func InitRoutes() *gin.Engine {

	r := gin.New()
	//r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/metrics", prometheusHandler())

	// 注册路由， 页面访问： http://localhost:8081/debug/pprof/
	pprof.Register(r)

	// version 1
	apiV1 := r.Group("/v1")

	apiV1.GET("users", func(c *gin.Context) {
		cpuTemp.Set(65.3)
		counterOps.Inc()
		c.JSON(http.StatusOK, "List Of V1 Users")
	})

	// User only can be added by authorized person
	authV1 := apiV1.Group("/", AuthMiddleWare())

	authV1.POST("users/add", AddV1User)

	//version 2
	apiV2 := r.Group("/v2")

	apiV2.GET("users", func(c *gin.Context) {
		c.JSON(http.StatusOK, "List Of V2 Users")
	})

	// User only can be added by authorized person
	authV2 := apiV2.Group("/", AuthMiddleWare())

	authV2.POST("users/add", AddV2User)

	return r
}

func AddV1User(c *gin.Context) {

	// AddUser

	c.JSON(http.StatusOK, "V1 User added")
}

func AddV2User(c *gin.Context) {

	// AddUser

	c.JSON(http.StatusOK, "V2 User added")
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
