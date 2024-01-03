package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Severity string

const (
	Entry Severity   = ""
	Success Severity = "[SUCCESS]"
	Low Severity     = "[ INFO? ]"
	Medium Severity  = "[WARNING]"
	High Severity    = "[ ERROR ]"
	Panic Severity   = "[ FATAL ]"
)

type Logger struct {
	logger *log.Logger
	printOut bool
}

func (l *Logger) SetStdoutBehavior(shouldPrintOut bool) {
	l.printOut = shouldPrintOut
}

func (l *Logger) Log(severity Severity, content ...string) {
	res := strings.Join(content, " | ")

	sep := ""
	if severity != Entry {
		sep = ": "
	}

	res = fmt.Sprintf("%s%s%s", string(severity), sep, res)

	switch severity {
	case Entry:
		l.logger.Println(res)
	case Success:
		l.logger.Println(res)
	case Low:
		l.logger.Println(res)
	case Medium:
		l.logger.Println(res)
	case High:
		l.logger.Fatalln(res)
	case Panic:
		l.logger.Panicln(res)
	}

	if l.printOut {
		fmt.Println(res)
	}
}

func (l *Logger) LogEntry(content ...string) {
	l.Log(Entry, content...)
}

func NewLogger(logFilename string, printOut bool) (*Logger, error) {
	if (!DoesThisFileExist(logFilename)) {
		_, err := os.Create(logFilename)

		if err != nil {
			log.Fatal("Failed to create the following log file: ", logFilename)
		}
	}

	w, err := os.OpenFile(logFilename, os.O_APPEND | os.O_WRONLY, os.ModeAppend)

	if err != nil {
		return nil, err
	}
	
	l := new(Logger)
	l.printOut = printOut
	l.logger = log.New(w, "", log.Default().Flags())
	
	return l, nil
}
