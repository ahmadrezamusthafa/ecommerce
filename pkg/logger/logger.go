package logger

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
)

const TagRequestID = "request_id"

var log *logrus.Logger

func Init() {
	log = logrus.New()
	strLevel := os.Getenv("LOG_LEVEL")
	if strLevel == "" {
		strLevel = "debug"
	}

	level, err := logrus.ParseLevel(strLevel)
	if err != nil {
		level = logrus.DebugLevel
	}

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(level)
	log.SetOutput(os.Stdout)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func InfoCtx(ctx context.Context, message string, data ...interface{}) {
	messages := []interface{}{message}
	if data != nil {
		js, _ := json.Marshal(data)
		messages = append(messages, string(js))
	}
	if requestID, ok := ctx.Value(TagRequestID).(string); ok {
		log.WithFields(logrus.Fields{TagRequestID: requestID}).Info(messages...)
	} else {
		log.Info(messages...)
	}
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func ErrorCtx(ctx context.Context, args ...interface{}) {
	if requestID, ok := ctx.Value(TagRequestID).(string); ok {
		log.WithFields(logrus.Fields{TagRequestID: requestID}).Error(args...)
	} else {
		log.Error(args...)
	}
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
