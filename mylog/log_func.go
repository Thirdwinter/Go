package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

const (
	debug   uint8 = iota //0
	trace                //1
	info                 //2
	warning              //3
	err                  //4
	fatal                //5
)

type logconfig struct {
	outFile  bool   `yaml:"out"`
	filePath string `yaml:"filepath"`
	lname    string `yaml:"logname"`
	leave    uint8  `yaml:"logleave"`
	logsize  int64  `yame:"logsize"`
}

var (
	initLog logconfig
	file    *os.File
)

func init() {
	// Execute default initialization
	initLog = logconfig{
		outFile:  false,
		filePath: "./",
		lname:    "log.log",
		leave:    debug,
		logsize:  10 * 1024 * 1024,
	}
}

// Logger initialization
func Init(OutFile bool, loglevel string, path string, logname string, size int64) {
	initLog.outFile = OutFile
	loglevel = strings.ToUpper(loglevel)
	switch loglevel {
	case "DEBUG":
		initLog.leave = 0
	case "TRACE":
		initLog.leave = 1
	case "INFO":
		initLog.leave = 2
	case "WARNING":
		initLog.leave = 3
	case "ERROR":
		initLog.leave = 4
	case "FATAL":
		initLog.leave = 5
	default:
		initLog.leave = 0
	}
	initLog.filePath = path
	initLog.lname = logname
	initLog.logsize = size
}

func (l *logconfig) openFile() {
	f := fmt.Sprintf("%s%s", l.filePath, l.lname) // path + name + suffixname
	var err error
	file, err = os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
}

// Log rotation based on size
// 日志按照大小切割
func (l *logconfig) cut() *os.File {
	fi, err := os.Stat(fmt.Sprintf("%s%s", l.filePath, l.lname))
	if err != nil {
		Debug("err:%s", err)
	}
	fsize := fi.Size()

	if fsize >= l.logsize {
		file.Close() // 关闭当前的日志文件
		now := time.Now().Format("20060102150405000")
		logname := fmt.Sprintf("%s%s", l.filePath, l.lname)
		newlogname := fmt.Sprintf("%s%s%s", l.filePath, now, l.lname)
		err := os.Rename(logname, newlogname)
		if err != nil {
			panic(err)
		}
		newFile, err := os.OpenFile(logname, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_CREATE, 0644) // 打开新的日志文件
		if err != nil {
			panic(err)
		}
		file = newFile // 更新 Fail 字段为新的文件句柄
		return newFile
	}
	return file
}

// Log content
func (l *logconfig) write(Leave string, format string, a ...interface{}) {
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
	now := time.Now().Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("[%s] [%s] [%s] [%s] [%d] {%s}\n", Leave, now, fileName, funcName, lineNo, fmt.Sprintf(format, a...))

	if !l.outFile {
		fmt.Print(logMsg)
	} else {
		l.openFile()
		fileobj := l.cut()    // Cut the log file here
		defer fileobj.Close() // Close the fileobj when the function returns
		fmt.Fprint(fileobj, logMsg)
	}
}

func Debug(format string, a ...interface{}) {
	if initLog.leave <= debug {
		initLog.write("DEBUG", format, a...)
	}
}

func Trace(format string, a ...interface{}) {
	if initLog.leave <= trace {
		initLog.write("TRACE", format, a...)
	}
}

func Info(format string, a ...interface{}) {
	if initLog.leave <= info {
		initLog.write("INFO", format, a...)
	}
}

func Warning(format string, a ...interface{}) {
	if initLog.leave <= warning {
		initLog.write("WARNING", format, a...)
	}
}

func Error(format string, a ...interface{}) {
	if initLog.leave <= err {
		initLog.write("ERROR", format, a...)
	}
}

func Fatal(format string, a ...interface{}) {
	if initLog.leave <= fatal {
		initLog.write("FATAL", format, a...)
	}
}
