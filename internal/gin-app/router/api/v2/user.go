package v2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserList(c *gin.Context) {
	c.JSON(http.StatusOK, "List Of V2 Users")
}

func AddUser(c *gin.Context) {
	// AddUser
	c.JSON(http.StatusOK, "V2 User added")
}
