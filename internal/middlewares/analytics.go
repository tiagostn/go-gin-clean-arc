package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Analytics() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		log.Print(latency)
		status := c.Writer.Status()
		log.Println(status)
	}
}
