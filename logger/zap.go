package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/2/24 17:30
 * @Desc: 日志 特别谢鸣：https://www.liwenzhou.com/posts/Go/zap/
 */

const DefaultLogPath = "./logs/micro.log"

// Option custom setup config
type Option func(*option)

type option struct {
	level   zapcore.Level
	logPath string
	fields  map[string]string
}

// WithDebugLevel only greater than 'level' will output
func WithDebugLevel() Option {
	return func(opt *option) {
		opt.level = zapcore.DebugLevel
	}
}

// WithInfoLevel only greater than 'level' will output
func WithInfoLevel() Option {
	return func(opt *option) {
		opt.level = zapcore.InfoLevel
	}
}

// WithWarnLevel only greater than 'level' will output
func WithWarnLevel() Option {
	return func(opt *option) {
		opt.level = zapcore.WarnLevel
	}
}

// WithErrorLevel only greater than 'level' will output
func WithErrorLevel() Option {
	return func(opt *option) {
		opt.level = zapcore.ErrorLevel
	}
}

func WithLogPath(path string) Option {
	return func(opt *option) {
		opt.logPath = path
	}
}

// WithField add field(s) to log
func WithField(key, value string) Option {
	return func(opt *option) {
		opt.fields[key] = value
	}
}

// 初始化 logger
func NewJSONLogger(opts ...Option) *zap.SugaredLogger {
	opt := &option{
		level:   zapcore.InfoLevel,
		logPath: DefaultLogPath,
		fields:  make(map[string]string),
	}
	for _, f := range opts {
		f(opt)
	}
	syncWriter := getLogWriter(opt.logPath)
	encoder := zapcore.NewJSONEncoder(getEncoder())
	core := zapcore.NewCore(
		encoder,
		syncWriter,
		zap.NewAtomicLevelAt(opt.level),
	)

	logger := zap.New(
		core,
		zap.AddCaller(), // 将调用函数信息记录到日志中
		zap.AddCallerSkip(1),
	)
	for key, value := range opt.fields {
		logger = logger.WithOptions(
			zap.Fields(
				zapcore.Field{Key: key, Type: zapcore.StringType, String: value},
			),
		)
	}

	return logger.Sugar()
}

// WriterSyncer ：指定日志将写到哪里去, zap中加入Lumberjack支持
func getLogWriter(logPath string) zapcore.WriteSyncer {
	return zapcore.AddSync(
		// 日志切割归档功能，使用第三方库Lumberjack来实现。
		&lumberjack.Logger{
			Filename:   logPath, // 日志文件名称
			MaxSize:    128,     // MB
			MaxAge:     30,      // 保留旧文件的最大天数
			MaxBackups: 300,     // 保留旧文件的最大个数
			LocalTime:  true,    // 使用本地时间
			Compress:   true,    // 是否启用压缩
		},
	)
}

// 编码器
func getEncoder() zapcore.EncoderConfig {
	// 参考 zap.NewProductionEncoderConfig()
	return zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,//Level序列化为全大写字符串
		EncodeTime: zapcore.ISO8601TimeEncoder, // 时间格式
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
	}
}
