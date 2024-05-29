package zap

import (
	"os"

	zapLogger "github.com/go-kratos/kratos/contrib/log/zap/v2"

	"github.com/go-kratos/kratos/v2/log"

	confV1 "mosong/api/gen/go/conf/service/v1"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// NewLogger 创建一个新的日志记录器 - Zap
func NewLogger(cfg *confV1.Logger) log.Logger {
	// 配置 zap 日志
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// 日志自动切割，采用 lumberjack 实现的
	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.Zap.Filename,
		MaxSize:    int(cfg.Zap.MaxSize),
		MaxBackups: int(cfg.Zap.MaxBackups),
		MaxAge:     int(cfg.Zap.MaxAge),
	}
	writeSyncer := zapcore.AddSync(lumberJackLogger)

	//设置日志级别
	var lvl = new(zapcore.Level)
	if err := lvl.UnmarshalText([]byte(cfg.Zap.Level)); err != nil {
		return nil
	}

	// 编码器配置
	core := zapcore.NewCore(
		jsonEncoder, // json 格式
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writeSyncer)), // 打印到控制台和文件,
		lvl)
	logger := zap.New(core).WithOptions()

	wrapped := zapLogger.NewLogger(logger)

	return wrapped
}
