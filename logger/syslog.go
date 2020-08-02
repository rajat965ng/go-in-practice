package main

import (
	"log"
	"log/syslog"
)

func main() {
	priority := syslog.LOG_LOCAL3|syslog.LOG_NOTICE
	logFlog := log.Ldate|log.LstdFlags
	logger,_ := syslog.NewLogger(priority,logFlog)
	logger.Println("Before the function")
}
