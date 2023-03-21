package core

import (
	"AuroraPixel/global"
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// 初始化日志
func InitLogger() {
	global.Log = myLogger()
}

// 日志颜色
var (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

// 格式化体
type logFormatter struct{}

func (t *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//设置颜色等级
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义时间格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定文件输出路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "%v [%s] \033[%dm%s\033[0m  %s %s %s\n", global.Config.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%v [%s] \033[%dm%s\033[0m %s\n", global.Config.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}

	return b.Bytes(), nil
}

// 获取日志级别
func getLevel() logrus.Level {
	level, err := logrus.ParseLevel(global.Config.LoggerConfig.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	return level
}

// 自定义logo
func myLogger() *logrus.Logger {
	mLog := logrus.New()
	mLog.SetOutput(os.Stdout)
	mLog.SetReportCaller(global.Config.ShowLine)
	mLog.SetFormatter(&logFormatter{})
	mLog.SetLevel(getLevel())
	//全局logo生效
	defaultLogger()
	return mLog
}

// 全局logo的配置
func defaultLogger() {
	logrus.SetOutput(os.Stdout)                    //输出类型
	logrus.SetReportCaller(global.Config.ShowLine) //开启函数名和行号
	logrus.SetFormatter(&logFormatter{})           //自定格式化工具
	logrus.SetLevel(getLevel())                    //日志等级
}
