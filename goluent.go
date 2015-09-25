package goluent

import (
	"fmt"
	"github.com/fluent/fluent-logger-golang/fluent"
	stdLog "log"
	"os"
)

const (
	infoLog severity = iota
	warningLog
	errorLog
	fatalLog
	numSeverity = 4
)

var severityName = []string{
	infoLog:    "INFO",
	warningLog: "WARNING",
	errorLog:   "ERROR",
	fatalLog:   "FATAL",
}

var hostname string

type severity int8

func print(s severity, args ...interface{}) {
	//connect fluent server
	f, err := fluent.New(fluent.Config{
		FluentPort: 24224,
		FluentHost: "localhost",
		TagPrefix:  "goluent." + getHostname(),
	})
	defer f.Close()

	message := fmt.Sprint(args...)

	stdLog.Println(message)
	if err == nil {
		f.Post(severityName[s], map[string]string{"message": message})
	}
}

func printf(s severity, format string, args ...interface{}) {
	//connect fluent server
	f, err := fluent.New(fluent.Config{
		FluentPort: 24224,
		FluentHost: "localhost",
		TagPrefix:  "goluent." + getHostname(),
	})
	defer f.Close()

	message := fmt.Sprintf(format, args...)

	stdLog.Println(message)
	if err == nil {
		f.Post(severityName[s], map[string]string{"message": message})
	}
}

func Info(args ...interface{}) {
	print(infoLog, args...)
}

func Infof(format string, args ...interface{}) {
	printf(infoLog, format, args...)
}

func Warning(args ...interface{}) {
	print(warningLog, args...)
}

func Warningf(format string, args ...interface{}) {
	printf(warningLog, format, args...)
}

func Error(args ...interface{}) {
	print(errorLog, args...)
}

func Errorf(format string, args ...interface{}) {
	printf(errorLog, format, args...)
}

func Fatal(args ...interface{}) {
	print(fatalLog, args...)
}

func Fatalf(format string, args ...interface{}) {
	printf(fatalLog, format, args...)
}

func getHostname() string {
	if hostname == "" {
		hostname, _ = os.Hostname()
	}
	return hostname
}
