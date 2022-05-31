// Package providers provides ...
package providers

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"time"

	"gitee.com/zhenyangze/gin-framework/configs"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var DefaultLogger *logrus.Logger

func InitLogger() {
	DefaultLogger = Logger()
}

func Logger() *logrus.Logger {
	logConfig := configs.GetLoggerConfig()
	logFilePath := logConfig["path"].(string)
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := logConfig["log_name"].(string)
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}

	//实例化
	logger := logrus.New()

	//设置日志级别
	switch logConfig["level"] {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
		logger.SetOutput(os.Stderr)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
		logger.SetOutput(getNull())
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
		logger.SetOutput(getNull())
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
		logger.SetOutput(getNull())
	default:
		logger.SetLevel(logrus.InfoLevel)
		logger.SetOutput(getNull())
	}

	//设置日志格式
	/*logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})*/

	// 设置 rotatelogs
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.TraceLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return logger
}

func getNull() *bufio.Writer {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	return writer
}

func LoggerHandler() gin.HandlerFunc {
	logger := Logger()
	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUrl := c.Request.RequestURI
		//状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUrl,
		}).Info()
	}
}

func Info(args ...interface{}) {
	DefaultLogger.Info(args)
}

func Error(args ...interface{}) {
	DefaultLogger.Error(args)
}
