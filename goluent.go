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

func _print(s severity, args ...interface{}) {
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

func _printf(s severity, format string, args ...interface{}) {
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

func Print(args ...interface{}) {
	_print(infoLog, args...)
}

func Printf(format string, args ...interface{}) {
	_printf(infoLog, format, args...)
}

func Info(args ...interface{}) {
	_print(infoLog, args...)
}

func Infof(format string, args ...interface{}) {
	_printf(infoLog, format, args...)
}

func Warning(args ...interface{}) {
	_print(warningLog, args...)
}

func Warningf(format string, args ...interface{}) {
	_printf(warningLog, format, args...)
}

func Error(args ...interface{}) {
	_print(errorLog, args...)
}

func Errorf(format string, args ...interface{}) {
	_printf(errorLog, format, args...)
}

func Fatal(args ...interface{}) {
	_print(fatalLog, args...)
	os.Exit(255)
}

func Fatalf(format string, args ...interface{}) {
	_printf(fatalLog, format, args...)
	os.Exit(255)
}

func getHostname() string {
	if hostname == "" {
		hostname, _ = os.Hostname()
	}
	return hostname
}
