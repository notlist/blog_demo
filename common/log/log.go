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
	setUid()
	setLogger()
}

// 定义一个结构体获取返回体的数据
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
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
	Logger.AddHook(&LogrusHook{})
}

// 关键操作，核心流程的日志
func LogErrorInfoToFile(fields logrus.Fields) {
	Logger.SetLevel(logrus.InfoLevel)
	Logger.WithFields(fields).Info()
}

// 把二进制写入缓冲区
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 把字符串写入缓冲区
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// 获取返回体的中间件
func GinBodyLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		responseStr := blw.body.String()
		//开始时间
		startTime := time.Now()
		//结束时间
		endTime := time.Now()
		//执行时间
		latencyTime := endTime.Sub(startTime)
		//请求方式
		reqMethod := c.Request.Method
		//请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		//请求ip
		clientIP := c.ClientIP()
		//请求参数
		reqParams := c.Request.Body
		//请求ua
		reqUa := c.Request.UserAgent()
		var resultBody logrus.Fields
		resultBody = make(map[string]interface{})
		resultBody["response"] = responseStr
		resultBody["requestUri"] = reqUri
		resultBody["clientIp"] = clientIP
		resultBody["body"] = reqParams
		resultBody["userAgent"] = reqUa
		resultBody["requestMethod"] = reqMethod
		resultBody["startTime"] = startTime
		resultBody["endTime"] = endTime
		resultBody["latencyTime"] = latencyTime
		resultBody["statusCode"] = statusCode
		LogErrorInfoToFile(resultBody)
		setUid()
	}
}

func setUid() {
	uid = uuid.New().String()
}

func GetNewUid() string {
	return uid
}

type LogrusHook struct {
}

// 设置所有的日志等级都走这个钩子
func (hook *LogrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// 修改其中的数据，或者进行其他操作
func (hook *LogrusHook) Fire(entry *logrus.Entry) error {
	entry.Data["request_id"] = GetNewUid()
	return nil
}

func getErrorFileAndLine(err error) {
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
