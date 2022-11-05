package main

import (
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const IS_DEBUG = 1

func main() {

	lumberjackLogger := &lumberjack.Logger{
		// Log file abbsolute path, os agnostic
		Filename:   filepath.ToSlash("./file.txt"),
		MaxSize:    1, // MB
		MaxBackups: 30,
		MaxAge:     30,    // days
		Compress:   false, // disabled by default
	}

	// Fork writing into two outputs
	//multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)

	logFormatter := new(log.TextFormatter)
	logFormatter.TimestampFormat = time.RFC3339 // or RFC3339 RFC1123Z
	logFormatter.FullTimestamp = true

	log.SetFormatter(logFormatter)
	log.SetLevel(log.InfoLevel)
	log.SetLevel(log.DebugLevel)
	log.SetOutput(lumberjackLogger)

	log.Info("some message")

	for i := 0; i < 1000000; i++ {
		log.Debug("some message")
	}
	log.Debug("some message")

}
