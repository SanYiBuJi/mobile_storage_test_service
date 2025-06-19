package Logger

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"time"
)

var Logger = logrus.New()

func init() {
	//logFilePath := "/app/logs/mobile-storage-test-service.log"
	logFilePath := "logs/mobile-storage-test-service.log"
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	Logger.SetOutput(file)
	Logger.SetLevel(logrus.InfoLevel)
}

func LogrusLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 读取请求头
		for key, values := range c.Request.Header {
			for _, value := range values {
				Logger.Printf("Request Header: %s = %s", key, value)
			}
		}

		// 读取请求体
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			Logger.Printf("Error reading request body: %v", err)
			c.AbortWithStatus(500)
			return
		} else {
			Logger.Printf("Request Body: %s", body)
		}

		// 将请求体内容重新放回请求中，以便后续处理函数可以读取
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		// 处理请求
		c.Next()
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		// 使用 logrus 记录日志
		Logger.WithFields(logrus.Fields{
			"status_code": statusCode,
			"latency":     latency,
			"client_ip":   clientIP,
			"method":      method,
			"path":        path,
		}).Info()
	}
}
