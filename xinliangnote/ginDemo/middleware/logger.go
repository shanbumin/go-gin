package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go-gin/xinliangnote/ginDemo/config"
	"os"
	"path"
	"time"
)

//@todo 比如此时是 20200313，启动系统之后，会同时产生如下的两个文件:
//@todo system.log.20200313.log
//@todo system.log 这里是个软链额。system.log -> system.log.20200313.log
// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	logFilePath := config.Log_FILE_PATH
	logFileName := config.LOG_FILE_NAME
	// 日志文件
	fileName := path.Join(logFilePath,logFileName)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	// 实例化
	logger := logrus.New()
	// 设置输出
	logger.Out = src
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName + ".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		//设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		//设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})
	// 新增钩子
	logger.AddHook(lfHook)
	return func(c *gin.Context) {
		startTime := time.Now()// 开始时间
		c.Next() // 处理请求
		endTime := time.Now() // 结束时间
		latencyTime := endTime.Sub(startTime) // 执行时间
		reqMethod := c.Request.Method // 请求方式
		reqUri := c.Request.RequestURI  // 请求路由
		statusCode := c.Writer.Status()  // 状态码
		clientIP := c.ClientIP() 		// 请求IP
		//输出日志
		logger.WithFields(logrus.Fields{
			"status_code"  : statusCode,
			"latency_time" : latencyTime,
			"client_ip"    : clientIP,
			"req_method"   : reqMethod,
			"req_uri"      : reqUri,
		}).Info()
	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}