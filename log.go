package main

import (
	"context"
	"time"

	"github.com/lukasjarosch/service-boilerplate/config"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"
)

// initLogging initializes a logrus Logger
func initLogging(cfg config.LogConfiguration) *logrus.Logger {

	l := logrus.New()
	var formatter logrus.Formatter
	if cfg.Format == "json" {
		formatter = &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		}
	} else {
		formatter = &logrus.TextFormatter{}
	}
	l.SetFormatter(formatter)

	logLevel, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	l.SetLevel(logLevel)
	return l
}

// logWithContext attaches the function name, request-id and trace-id  to the baseLogger and returns a *logrus.Entry
func logWithContext(ctx context.Context, functionName string) *logrus.Entry {
	log := baseLogger.WithField("method", functionName)

	if requestId := getValueFromMetadata(ctx, "X-Request-Id"); requestId != "" {
		log = log.WithField("request_id", requestId)
	}

	if traceId := getValueFromMetadata(ctx, "X-B3-Trace-Id"); traceId != "" {
		log = log.WithField("trace_id", traceId)
	}

	return log
}

// getValueFromMetadata extracts a value of a given key from the context's metadata.
// If the key does not exist and empty string is returned
func getValueFromMetadata(ctx context.Context, key string) string {
	if md, ok := metadata.FromContext(ctx); ok {
		if id, ok := md[key]; ok {
			return id
		}
	}
	return ""
}

// LogWrapper is the handler wrapper
func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		l := logWithContext(ctx, req.Method())
		l.Println("incoming request")
		return fn(ctx, req, rsp)
	}
}

