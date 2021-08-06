package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"quick-go/internal/gin-app/metrics"
)

func GetUserList(c *gin.Context) {
	metrics.CpuTemp.Set(65.3)
	metrics.CounterOps.Inc()
	c.JSON(http.StatusOK, "List Of V1 Users")
}

func AddUser(c *gin.Context) {
	// AddUser
	c.JSON(http.StatusOK, "V1 User added")
}
