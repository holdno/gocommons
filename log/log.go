package log

import (
	"path"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var logInstance *logrus.Logger

// 配置logger
func Init(logDir, appName string) {
	logInstance = logrus.New()
	// json格式的log输出
	logName := path.Join(logDir, appName)
	logInstance.AddHook(rotatelogsHook(logName))
}

func rotatelogsHook(logName string) logrus.Hook {
	writer, err := rotatelogs.New(
		logName+".%Y%m%d",
		rotatelogs.WithLinkName(logName),
		rotatelogs.WithRotationTime(time.Hour*24),
		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		// rotatelogs.WithMaxAge(time.Hour*24),
		// WithRotationCount设置文件清理前最多保存的个数。
		// rotatelogs.WithRotationCount(maxRemainCnt),
	)

	if err != nil {
		panic(err)
	}
	return lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{})
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return logInstance.WithFields(fields)
}
