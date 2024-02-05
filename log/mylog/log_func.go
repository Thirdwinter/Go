package mylog

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

const (
	DEBUG   uint8 = iota //0
	TRACE                //1
	INFO                 //2
	WARNING              //3
	ERROR                //4
	FATAL                //5
)

type LogInit struct {
	OutFile      bool
	TimeLocation *time.Location
	FailPath     string
	FailName     string
	Suffixname   string
	Prem         fs.FileMode
	Fail         *os.File
	Leave        uint8
	Logsize      int64
	Err          [3]error
}

var InitLog LogInit

func init() {
	// Execute default initialization
	InitLog.Init(false, "UTC", "DEBUG", "./", "log", ".log", 0644, 10*1024*1024)
}

// Logger initialization
func (l *LogInit) Init(OutFile bool, TimeLocation string, loglevel string, path string, name string, suffixname string,
	prem fs.FileMode, size int64) {
	l.OutFile = OutFile
	l.TimeLocation, l.Err[0] = time.LoadLocation(TimeLocation)
	if l.Err[0] != nil {
		panic(l.Err[0])
	}
	loglevel = strings.ToUpper(loglevel)
	switch loglevel {
	case "DEBUG":
		l.Leave = 0
	case "TRACE":
		l.Leave = 1
	case "INFO":
		l.Leave = 2
	case "WARNING":
		l.Leave = 3
	case "ERROR":
		l.Leave = 4
	case "FATAL":
		l.Leave = 5
	default:
		l.Leave = 0
	}
	l.FailPath = path
	l.FailName = name
	l.Suffixname = suffixname
	l.Prem = prem
	l.Logsize = size
}

func (l *LogInit) OpenFile() {
	f := fmt.Sprintf("%s%s%s", l.FailPath, l.FailName, l.Suffixname) // path + name + suffixname
	l.Fail, l.Err[1] = os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, l.Prem)
	if l.Err[1] != nil {
		panic(l.Err[1])
	}
}

// Log rotation based on size
// 日志按照大小切割
// 日志按照大小切割
// 日志按照大小切割
func (l *LogInit) cut() *os.File {
	fi, err := os.Stat(fmt.Sprintf("%s%s%s", l.FailPath, l.FailName, l.Suffixname))
	if err != nil {
		Debug("err:%s", err)
	}
	fsize := fi.Size()

	if fsize >= l.Logsize {
		l.Fail.Close() // 关闭当前的日志文件
		now := time.Now().Format("20060102150405000")
		logname := fmt.Sprintf("%s%s%s", l.FailPath, l.FailName, l.Suffixname)
		newlogname := fmt.Sprintf("%s%s%s%s", l.FailPath, l.FailName, now, l.Suffixname)
		err := os.Rename(logname, newlogname)
		if err != nil {
			panic(err)
		}
		newFile, err := os.OpenFile(logname, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_CREATE, l.Prem) // 打开新的日志文件
		if err != nil {
			panic(err)
		}
		l.Fail = newFile // 更新 Fail 字段为新的文件句柄
		return newFile
	}
	return l.Fail
}




// Log content
func (l *LogInit) Write(Leave string, format string, a ...interface{}) {
	pc, file, lineNo, ok := runtime.Caller(2)
	var funcName, fileName string
	if !ok {
		funcName = "err"
		fileName = "err"
		lineNo = 0
	} else {
		funcName = runtime.FuncForPC(pc).Name()
		fileName = path.Base(file) // Base function returns the last element of the path
	}
	timenow, err := time.ParseInLocation("2006/01/02 15:04:05", time.Now().Format("2006/01/02 15:05:05"), l.TimeLocation)
	if err != nil {
		panic(err)
	}
	logMsg := fmt.Sprintf("[%s] [%s] [%s] [%s] [%d] {%s}\n", Leave, timenow, fileName, funcName, lineNo, fmt.Sprintf(format, a...))

	if !l.OutFile {
		fmt.Print(logMsg)
	} else {
        l.OpenFile()
        fileobj := l.cut() // Cut the log file here
        defer fileobj.Close() // Close the fileobj when the function returns
        fmt.Fprint(fileobj, logMsg)
    }
}

func Debug(format string, a ...interface{}) {
	if InitLog.Leave <= DEBUG {
		InitLog.Write("DEBUG", format, a...)
	}
}

func Trace(format string, a ...interface{}) {
	if InitLog.Leave <= TRACE {
		InitLog.Write("TRACE", format, a...)
	}
}

func Info(format string, a ...interface{}) {
	if InitLog.Leave <= INFO {
		InitLog.Write("INFO", format, a...)
	}
}

func Warning(format string, a ...interface{}) {
	if InitLog.Leave <= WARNING {
		InitLog.Write("WARNING", format, a...)
	}
}

func Error(format string, a ...interface{}) {
	if InitLog.Leave <= ERROR {
		InitLog.Write("ERROR", format, a...)
	}
}

func Fatal(format string, a ...interface{}) {
	if InitLog.Leave <= FATAL {
		InitLog.Write("FATAL", format, a...)
	}
}
