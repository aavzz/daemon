/*
Package log sets up logging for the process
*/
package log

import (
	"log"
	"log/syslog"
	"os"
)

var sysLog *syslog.Writer

// Logger returns an object to write logs to (used by http.Server) 
func Logger(tag string) *log.Logger {
	l := log.New(os.Stdout, tag, log.LstdFlags)
	if sysLog != nil {
		l.SetOutput(sysLog)
	}
	return l
}

// InitSyslog tries to get in touch with local syslog
func InitSyslog(tag string) {
	var err error
	sysLog, err = syslog.New(syslog.LOG_INFO|syslog.LOG_DAEMON, tag)
	if err != nil {
		log.Fatal(err.Error())
	}
}

//Info writes a log entry with proirity Info
func Info(s string) {
	if sysLog != nil {
		sysLog.Info(s)
	} else {
		log.Print(s)
	}
}

//Error writes a log entry with priority Error
func Error(s string) {
	if sysLog != nil {
		sysLog.Err(s)
	} else {
		log.Print(s)
	}
}

//Fatal writes a log entry with priority Error and stops the process
func Fatal(s string) {
	if sysLog != nil {
		sysLog.Err(s)
		os.Exit(1)
	} else {
		log.Fatal(s)
	}
}
