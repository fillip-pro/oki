package log

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	TraceLogger   *log.Logger
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func Tracef(format string, values ...interface{}) {
	TraceLogger.Printf(format, values...)
}

func Traceln(values ...interface{}) {
	TraceLogger.Println(values...)
}

func Infof(format string, values ...interface{}) {
	InfoLogger.Printf(format, values...)
}

func Infoln(values ...interface{}) {
	InfoLogger.Println(values...)
}

func Warnf(format string, values ...interface{}) {
	WarningLogger.Printf(format, values...)
}

func Warnln(values ...interface{}) {
	WarningLogger.Println(values...)
}

func Error(values ...interface{}) {
	ErrorLogger.Fatal(values)
}

func Errorf(format string, values ...interface{}) {
	ErrorLogger.Fatalf(format, values)
}

func init() {
	initLogging(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

func initLogging(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	TraceLogger = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger = log.New(infoHandle, "", 0)

	WarningLogger = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	ErrorLogger = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
