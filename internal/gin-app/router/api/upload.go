package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(-1, "error")
		return
	}

	if fileHeader == nil {
		c.JSON(-1, "read file header failed!")
		return
	}
    // TODO
	fmt.Printf("%v", file)
}
