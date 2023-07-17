package log

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

/*
*
添加滚动日志文件
fileName 文件名
maxSize 单个文件最大大小，单位MB
maxBackups  最多保留多少个备份的日志文件
maxAge  保留多少天的备份日志
compress  是否gz压缩
如果 MaxBackups 和 MaxAge 同时设置成 0, 将不会删除旧的日志文件.
*/
func AddRotationLogFile(fileName string, maxSize, maxBackups, maxAge int, compress bool) {
	rotationLogger := lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge,   //days
		Compress:   compress, // -disabled by default
	}
	zapCfg.AddOutput = append(zapCfg.AddOutput, &rotationLogger)
	zapCfg.Color = false
	zapCfg.Refresh()
}
