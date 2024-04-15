package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sync"
)

var onceInit sync.Once

var log *logrus.Entry

var logPath = ""

var logName = "application.log"

var withFields CustomFields

var writer io.Writer = nil

var logLevel logrus.Level = logrus.InfoLevel

func SetLogPath(p string) {
	logPath = p
}

func SetLogName(l string) {
	logName = l
}

func SetWithFields(f CustomFields) {
	withFields = f
}

func SetWriter(w io.Writer) {
	writer = w
}

func SetLogLevel(l uint32) {
	logLevel = logrus.Level(l)
}

func InitLog() {
	onceInit.Do(func() {
		app := "application" //Todo: 暂时这么做
		initLogPath()
		InitLogger(app, getLogFilePath(), withFields, logLevel)

		log = GetLogger()
	})
}

func initLogPath() {
	if logPath == "" {
		logPath = "./runtime/log/"
	}
	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		panic("创建日志文件目录" + err.Error())
	}
}

func getLogFilePath() string {
	if logPath == "" {
		logPath = "./runtime/log/"
	}
	logFile := logPath + logName
	return logFile
}
