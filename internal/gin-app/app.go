package main

import (
	"encoding/json"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"

	"net/http"
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

func init() {
	prometheus.MustRegister(cpuTemp, counterOps)
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	router := gin.Default()

	router.GET("/metrics", prometheusHandler())

	// 注册路由， 页面访问： http://localhost:8081/debug/pprof/
	pprof.Register(router)

	// version 1
	apiV1 := router.Group("/v1")

	apiV1.GET("users", func(c *gin.Context) {

		cpuTemp.Set(65.3)
		counterOps.Inc()
		c.JSON(http.StatusOK, "List Of V1 Users")
	})

	// User only can be added by authorized person
	authV1 := apiV1.Group("/", AuthMiddleWare())

	authV1.POST("users/add", AddV1User)

	//version 2
	apiV2 := router.Group("/v2")

	apiV2.GET("users", func(c *gin.Context) {
		c.JSON(http.StatusOK, "List Of V2 Users")
	})

	// User only can be added by authorized person
	authV2 := apiV2.Group("/", AuthMiddleWare())

	authV2.POST("users/add", AddV2User)


	go func() {
		for {
			LocalTz()
			doSomething([]byte(`{"a": 1, "b": 2, "c": 3}`))
		}
	}()


	_ = router.Run(":8081")
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

func doSomething(s []byte) {
	var m map[string]interface{}
	err := json.Unmarshal(s, &m)
	if err != nil {
		panic(err)
	}

	s1 := make([]string, 0)
	s2 := ""
	for i := 0; i < 100; i++ {
		s1 = append(s1, string(s))
		s2 += string(s)
	}
}

func LocalTz() *time.Location {
	tz, _ := time.LoadLocation("Asia/Shanghai")
	return tz
}