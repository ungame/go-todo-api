package logger

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
	fileLogger  *os.File
)

const (
	logFilename  = "logs.txt"
	logFileFlags = os.O_CREATE | os.O_APPEND | os.O_WRONLY
	logFilePerm  = os.FileMode(0666)

	logFlags = log.Ldate | log.Ltime | log.Lshortfile
)

func init() {
	var err error

	fileLogger, err = os.OpenFile(logFilename, logFileFlags, logFilePerm)
	if err != nil {
		log.Fatalf("error on open log file: %s\n", err.Error())
	}

	infoLogger = log.New(fileLogger, "[INFO]  ", logFlags)
	warnLogger = log.New(fileLogger, "[WARN]  ", logFlags)
	errorLogger = log.New(fileLogger, "[ERROR] ", logFlags)
	fatalLogger = log.New(fileLogger, "[FATAL] ", logFlags)
}

func Flush() {
	if fileLogger != nil {
		err := fileLogger.Close()
		if err != nil {
			log.Printf("error on close log file: %s\n", err.Error())
		}
	}
}

func Info(format string, values ...interface{}) {
	infoLogger.Printf(EndL(format), values...)
}

func Warn(format string, values ...interface{}) {
	warnLogger.Printf(EndL(format), values...)
}

func Error(format string, values ...interface{}) {
	errorLogger.Printf(EndL(format), values...)
}

func Fatal(format string, values ...interface{}) {
	fatalLogger.Fatalf(EndL(format), values...)
}

func EndL(s string) string {
	size := len(s)
	if size > 0 && s[size-1] != '\n' {
		return s + "\n"
	}
	return s
}
