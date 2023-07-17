package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
)

var (
	_levelToColor = map[zapcore.Level]color{
		zapcore.DebugLevel:  Magenta,
		zapcore.InfoLevel:   Blue,
		zapcore.WarnLevel:   Yellow,
		zapcore.ErrorLevel:  Red,
		zapcore.DPanicLevel: Red,
		zapcore.PanicLevel:  Red,
		zapcore.FatalLevel:  Red,
	}
	_unknownLevelColor = Red

	_levelToLowercaseColorString = make(map[zapcore.Level]string, len(_levelToColor))
	_levelToCapitalColorString   = make(map[zapcore.Level]string, len(_levelToColor))
)

func init() {
	for level, color := range _levelToColor {
		_levelToLowercaseColorString[level] = color.Add("[" + level.String() + "]")
		_levelToCapitalColorString[level] = color.Add("[" + level.CapitalString() + "]")
	}
}

// cEncodeLevel 自定义日志级别显示  还没上色
func cEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// cEncodeLevel 自定义日志级别显示  带颜色
func cEncodeColorLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	//enc.AppendString("[" + level.CapitalString() + "]")
	s, ok := _levelToCapitalColorString[level]
	if !ok {
		s = _unknownLevelColor.Add(level.CapitalString())
	}
	enc.AppendString(s)
}

func getEncoder(colorLevel bool) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:          "msg",
		LevelKey:            "level",
		TimeKey:             "time",
		NameKey:             "",
		CallerKey:           "",
		FunctionKey:         "",
		StacktraceKey:       "",
		SkipLineEnding:      false,
		LineEnding:          "",
		EncodeLevel:         cEncodeColorLevel,
		EncodeTime:          zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration:      nil,
		EncodeCaller:        nil,
		EncodeName:          nil,
		NewReflectedEncoder: nil,
		ConsoleSeparator:    " ",
	}
	if colorLevel == false {
		encoderConfig.EncodeLevel = cEncodeLevel
	}
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if runtime.GOOS == "windows" { // windows 用 '\r\n' 换行
		encoderConfig.LineEnding = "\r\n"
	}

	return newConsoleEncoder(encoderConfig)
}

func newZap(cfg logCfg) *zap.Logger {

	var ws []zapcore.WriteSyncer
	//zapcore.NewMultiWriteSyncer()
	for _, writer := range cfg.AddOutput {
		ws = append(ws, zapcore.AddSync(writer))
		zapcore.AddSync(os.Stdout)
	}

	core := zapcore.NewCore(getEncoder(cfg.Color), zapcore.NewMultiWriteSyncer(ws...), unmarshalLevelText(cfg.Level))
	_logger := zap.New(core)
	//if err != nil {
	//	panic(err)
	//}
	return _logger

}

func unmarshalLevelText(text string) zapcore.Level {
	switch text {
	case "debug", "DEBUG":
		return zapcore.DebugLevel
	case "info", "INFO", "": // make the zero value useful
		return zapcore.InfoLevel
	case "warn", "WARN":
		return zapcore.WarnLevel
	case "error", "ERROR":
		return zapcore.ErrorLevel
	case "dpanic", "DPANIC":
		return zapcore.DPanicLevel
	case "panic", "PANIC":
		return zapcore.PanicLevel
	case "fatal", "FATAL":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
