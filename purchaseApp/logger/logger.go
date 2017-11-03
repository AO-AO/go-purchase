package logger

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Info 打印Info信息
func Info(message string, context gin.Context) {
	nowTime := time.Now().Format("2006/01/02 - 15:04:05")
	requestID, ok := context.Get("RequestID")
	var strRequestID string
	if !ok {
		log.Println("WARN: Can't get request ID")
		strRequestID = ""
	} else {
		strRequestID, ok = requestID.(string)
	}

	output := "INFO: " + nowTime + " [requestID:" + strRequestID + "] " + message
	log.Println(output)
}

//Error 打印Error信息
func Error(message string, context gin.Context) {
	nowTime := time.Now().Format("2006/01/02 - 15:04:05")
	requestID, ok := context.Get("RequestID")
	var strRequestID string
	if !ok {
		log.Println("Warn: Can't get request ID")
		strRequestID = ""
	} else {
		strRequestID, ok = requestID.(string)
	}

	output := "ERROR: " + nowTime + " [" + strRequestID + "] " + message
	log.Println(output)
}

//Warn 打印Warn信息
func Warn(message string, context gin.Context) {
	nowTime := time.Now().Format("2006/01/02 - 15:04:05")
	requestID, ok := context.Get("RequestID")
	var strRequestID string
	if !ok {
		log.Println("Warn: Can't get request ID")
		strRequestID = ""
	} else {
		strRequestID, ok = requestID.(string)
	}

	output := "Warn: " + nowTime + " [" + strRequestID + "] " + message
	log.Println(output)
}
