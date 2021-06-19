package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"github.com/jessevdk/go-flags"
	"github.com/odisme0804/go-feature-toggler/internal/service"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Args struct {
	HTTPAddr string `long:"http.addr" env:"HTTP_ADDR" default:":8080"`
}

func main() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	logger := zapAdapter(zapLogger, zapcore.InfoLevel)

	args := Args{}
	flag.Parse()
	if _, err := flags.NewParser(&args, flags.Default).Parse(); err != nil {
		panic(err)
	}

	logger.Log("server started, listening on:%s", args.HTTPAddr)

	s := service.NewServer(logger)
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/simpleBooleanFlag", s.SimpleBooleanFlagHandler)
		v1.GET("/multiVariantsHandler", s.MultiVariantsHandler)
		v1.GET("/variantAttachmentHandler", s.VariantAttachmentHandler)
		v1.GET("/customConstraintHandler", s.CustomConstraintHandler)
	}

	router.Run(args.HTTPAddr)
}

func zapAdapter(logger *zap.Logger, level zapcore.Level) log.Logger {
	var realLogFunc func(msg string, keysAndValues ...interface{})
	switch level {
	case zapcore.DebugLevel:
		realLogFunc = logger.Sugar().Debugw
	case zapcore.InfoLevel:
		realLogFunc = logger.Sugar().Infow
	case zapcore.WarnLevel:
		realLogFunc = logger.Sugar().Warnw
	case zapcore.ErrorLevel:
		realLogFunc = logger.Sugar().Errorw
	case zapcore.DPanicLevel:
		realLogFunc = logger.Sugar().DPanicw
	case zapcore.PanicLevel:
		realLogFunc = logger.Sugar().Panicw
	case zapcore.FatalLevel:
		realLogFunc = logger.Sugar().Fatalw
	}

	return log.LoggerFunc(func(kv ...interface{}) error {
		realLogFunc("", kv...)
		return nil
	})
}
