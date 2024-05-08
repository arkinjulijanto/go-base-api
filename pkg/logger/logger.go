package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/arkinjulijanto/go-base-api/config"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

type ExtraInfo struct {
	Path       string
	XRequestID string
	Version    string
}

func Init(cfg config.Config) {
	logger.SetReportCaller(true)
	customFormatter := new(logger.JSONFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.CallerPrettyfier = func(f *runtime.Frame) (string, string) {
		s := strings.Split(f.Function, ".")
		funcname := s[len(s)-1]
		_, filename := path.Split(f.File)
		return funcname, filename
	}
	logger.SetFormatter(customFormatter)

	logLevel := cfg.LOG_LEVEL
	setLogLevel(logLevel)

	if cfg.LOG_STDOUT {
		logger.New().Out = os.Stdout
	} else {
		// set log file to base on date
		file, err := os.OpenFile(cfg.LOG_PATH, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.SetOutput(file)
		} else {
			fmt.Println("failed to log to file ", err.Error())
		}
	}
}

func setLogLevel(logLevel string) {
	switch strings.ToLower(logLevel) {
	case "debug":
		logger.SetLevel(logger.DebugLevel)
	case "info":
		logger.SetLevel(logger.InfoLevel)
	case "warn":
		logger.SetLevel(logger.WarnLevel)
	case "error":
		logger.SetLevel(logger.ErrorLevel)
	default:
		logger.SetLevel(logger.DebugLevel)
	}
}

func LogInfo(msg string, c *gin.Context) {
	extraInfo := parseGinContext(c)
	logger.WithFields(logger.Fields{
		"path":         extraInfo.Path,
		"x-request-id": extraInfo.XRequestID,
		"version":      extraInfo.Version,
	}).Info(msg)
}

func LogError(msg string, err error, c *gin.Context) {
	extraInfo := parseGinContext(c)
	logger.WithFields(logger.Fields{
		"path":         extraInfo.Path,
		"x-request-id": extraInfo.XRequestID,
		"version":      extraInfo.Version,
		"error":        err.Error(),
	}).Error(msg)
}

func LogFatal(msg string, err error, c *gin.Context) {
	extraInfo := parseGinContext(c)
	logger.WithFields(logger.Fields{
		"path":         extraInfo.Path,
		"x-request-id": extraInfo.XRequestID,
		"version":      extraInfo.Version,
		"error":        err.Error(),
	}).Fatal(msg)
}

func PanicLN(msg string) {
	logger.Panicln(msg)
}

func FatalLn(msg string) {
	logger.Fatalln(msg)
}

func InfoLn(msg string) {
	logger.Infoln(msg)
}

func WarnLn(msg string) {
	logger.Warnln(msg)
}

func DebugLn(msg string) {
	logger.Debugln(msg)
}

func PrintLn(msg string) {
	logger.Print(msg)
}

func parseGinContext(c *gin.Context) ExtraInfo {
	var extraInfo ExtraInfo
	if c != nil {
		extraInfo.Path = c.Request.RequestURI
		extraInfo.XRequestID = c.Request.Header.Get("x-request-id")
		extraInfo.Version = c.Request.Header.Get("version")
	}

	return extraInfo
}
