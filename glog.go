/****************************************************
**                                                 **
**       Author: guanchengqi                       **
**       Email:      guancq@tuya.com               **
**       CreateTime: 2019.12.07(Saturday)          **
**       UpdateTime: 2019.12.27(星期五)            **
**                                                 **
****************************************************/

package glog

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

const (
	escapeHeader  = "\x1b"
	magentaString = "35"
	redString     = "31"
	yellowString  = "33"
	greenString   = "32"

	callDepth = 3
)

// 对于 glog 的线程安全的思考。
// 多个 glog 的竞争临界资源 只有写出的文件描述符
// 对于竞争锁，log 库已经做了，可以着手于优化的点应该是在于怎么做到尽可能的无冲突写入
// 也就是总是能拿到锁

// Glog a embed log.Logger struct
// type it for render log color
type Glog struct {
	*log.Logger
}

// Gl is a default glog struct
var gl Glog

// InitGlog func init package glog variable Gl, so can use like sqlmask.XX()
func InitGlog(wt io.Writer, flag int) {
	if 0 == flag {
		flag = log.Llongfile
	}
	if wt == nil {
		wt = os.Stdout
	}
	gl = Glog{log.New(wt, "", flag)}
}

// NewGlog func new a glog logger struct
func NewGlog(wt io.Writer, flag int) Glog {
	if 0 == flag {
		flag = log.Llongfile
	}
	if wt == nil {
		wt = os.Stdout
	}
	return Glog{log.New(wt, "", flag)}
}

// Fatalf format print fatal level log
func (g *Glog) Fatalf(format string, v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", escapeHeader, magentaString))
	g.Output(callDepth, color.MagentaString(format, v...))
}

// Fatalln println fatal level log
func (g *Glog) Fatalln(v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", escapeHeader, magentaString))
	g.Output(callDepth, color.MagentaString("", v...))
}

// Errorln println error level log
func (g *Glog) Errorf(format string, v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", escapeHeader, redString))
	g.Output(callDepth, color.RedString(format, v...))
}

// Errorln println error level log
func (g *Glog) Errorln(v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", escapeHeader, redString))
	g.Output(callDepth, color.RedString("", v...))
}

// Debugf format print debug level log
func (g *Glog) Debugf(format string, v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", escapeHeader, yellowString))
	g.Output(callDepth, color.YellowString(format, v...))
}

// Debugln println debug level log
func (g *Glog) Debugln(v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", escapeHeader, yellowString))
	g.Output(callDepth, color.YellowString("", v...))
}

// Infof format print info level log
func (g *Glog) Infof(format string, v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", escapeHeader, greenString))
	g.Output(callDepth, color.GreenString(format, v...))
}

// Infoln println info level log
func (g *Glog) Infoln(v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", escapeHeader, greenString))
	g.Output(callDepth, color.GreenString("", v...))
}

// Fatalf format print fatal level log
func Fatalf(format string, v ...interface{}) {
	gl.Fatalf(format, v...)
}

// Fatalln println fatal level log
func Fatalln(v ...interface{}) {
	gl.Fatalln(v...)
}

// Errorf format print error level log
func Errorf(format string, v ...interface{}) {
	gl.Errorf(format, v...)
}

// Errorln println error level log
func Errorln(v ...interface{}) {
	gl.Errorln(v...)
}

// Debugf format print debug level log
func Debugf(format string, v ...interface{}) {
	gl.Debugf(format, v...)
}

// Debugln println debug level log
func Debugln(v ...interface{}) {
	gl.Debugln(v...)
}

// Infof format print info level log
func Infof(format string, v ...interface{}) {
	gl.Infof(format, v...)
}

// Infoln println info level log
func Infoln(v ...interface{}) {
	gl.Infoln(v...)
}
