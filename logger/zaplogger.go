package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

var Logger *zap.Logger

func init() {

	if Logger == nil {
		// 创建一个文件输出对象
		file, err := os.OpenFile("mobile_storage_test_service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		//defer file.Close()

		// 创建一个encoder，将日志编码为json格式并输出到文件
		encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core := zapcore.NewCore(encoder, file, zap.DebugLevel)
		//config := zap.NewProductionConfig()
		// 可选：根据您的需求自定义配置
		// config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		// config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		// ...
		Logger = zap.New(core, zap.AddStacktrace(zap.ErrorLevel))
		if err != nil {
			panic("Failed to initialize logger: " + err.Error())
		}
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		method := c.Request.Method
		//body, _ := ioutil.ReadAll(c.Request.Body)
		c.Next() // 执行其他中间件和处理程序
		end := time.Now()
		latency := end.Sub(start)
		statusCode := c.Writer.Status()
		Logger.Info("LoggerMiddleware 完成初始化～")
		Logger.Info("Request completed", zap.String("path", path), zap.String("query", query),
			zap.String("method", method), zap.Int("status", statusCode), zap.Duration("latency", latency),
			//zap.String("body", string(body)))
		)
	}
}
