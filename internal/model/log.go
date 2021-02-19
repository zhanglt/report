package model

import (
	"log"
	"os"
)

func ProvideLog(conf *Config) *log.Logger {
	logFile, err := os.OpenFile(conf.Writable.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	logger := log.New(logFile, "success ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}

func GetLogger() *log.Logger {
	logFile, err := os.OpenFile("./weater.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("打开日志文件异常")
	}
	logger := log.New(logFile, "success", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}
