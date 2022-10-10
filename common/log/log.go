package log

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"time"
)

var Logger *logrus.Logger
var uid string

func Init() {
	SetUid()
	setLogger()
}

// 定义一个结构体获取返回体的数据
type BodyLogWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func setLogger() {
	//当前时间
	nowTime := time.Now()
	//获取日志文件存储的目录，这里我采用的是自己封装的一个获取配置文件的方法,可以看我上一篇viper获取配置信息的文章
	logFilePath := ""
	// 如果没有获取到配置文件的话那就是直接代码写死一个文件地址
	if len(logFilePath) <= 0 {
		//获取当前目前的地址，也就是项目的根目录
		if dir, err := os.Getwd(); err == nil {
			logFilePath = dir + "/logs/"
		}
	}
	//创建文件夹
	if err := os.MkdirAll(logFilePath, os.ModePerm); err != nil {
		fmt.Println(err.Error())
	}
	//文件名称
	logFileName := nowTime.Format("2006-01-02") + ".log"
	//日志文件地址拼接
	fileName := path.Join(logFilePath, logFileName)
	//fmt.Println("文件名称："+fileName)
	if _, err := os.Stat(fileName); err != nil {
		fmt.Println("检测文件：" + err.Error())
		_, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	//打开文件
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("write file log error", err)
	}

	//实例化
	Logger = logrus.New()
	//设置输出
	Logger.Out = src
	//这里我觉得应该是交给需要封装的方法去确认使用什么等级的日志和什么格式
	//设置日志级别
	Logger.SetLevel(logrus.TraceLevel)
	Logger.SetReportCaller(true)
	//设置日志格式 文本格式
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//添加hook
	Logger.AddHook(&LogrusHook{})
}

// 关键操作，核心流程的日志
func LogErrorInfoToFile(fields logrus.Fields) {
	Logger.SetLevel(logrus.InfoLevel)
	Logger.WithFields(fields).Info()
}

// 把二进制写入缓冲区
func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 把字符串写入缓冲区
func (w BodyLogWriter) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func SetUid() {
	uid = uuid.New().String()
}

func GetNewUid() string {
	return uid
}

func GetErrorFileAndLine(err error) {
	//获取上层运行时的文件以及行数
	for skip := 1; true; skip++ {
		//获取上层运行时的文件以及行数
		_, file, line, ok := runtime.Caller(skip)
		if ok {
			var resultBody logrus.Fields
			resultBody = make(map[string]interface{})
			resultBody["file_path"] = file
			resultBody["error_line"] = line
			resultBody["error_message"] = err.Error()
			LogErrorInfoToFile(resultBody)
		} else {
			break
		}
	}
}
