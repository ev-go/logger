package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

const IS_DEBUG = 1

func main() {

	logFileName := "log" + time.Now().Format("20060102") + ".txt"
	fmt.Println(logFileName)
	logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + logFileName)
		panic(err)
	}
	defer logFile.Close()

	log := &logrus.Logger{
		Out:   io.MultiWriter(logFile, os.Stdout),
		Level: logrus.TraceLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}

	var a, b, c int
	a = 2
	b = 5
	c = a + b

	// logrus.Traceln("Trace Level")
	// logrus.Debugln("Debug Level")
	// logrus.Infoln("Info Level")
	// logrus.Warningln("Warning Level")
	// logrus.Errorln("Error Level")
	// logrus.Fatalln("Fatal Level")
	// logrus.Panicln("Panic Level")

	log.Debugln("Hello, world!", c)

	scanInt, err := fmt.Scan(&a)
	if err != nil {
		log.Errorln("Чтение с консоли не удалось", err)
	}

	log.Infoln("Info Level", scanInt, a)

}
