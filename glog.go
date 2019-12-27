/****************************************************
**                                                 **
**       Author: guanchengqi                       **
**       Email:      guancq@tuya.com               **
**       CreateTime: 2019.12.07(Saturday)          **
**       UpdateTime: 2019.12.27(Friday)            **
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

// 对于 glog 的线程安全的思考。
// 多个 glog 的竞争临界资源 只有写出的文件描述符
// 对于竞争锁，log 库已经做了，可以着手于优化的点应该是在于怎么做到尽可能的无冲突写入
// 也就是总是能拿到锁

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

func NewGlog(wt io.Writer, flag int) Glog {
	if 0 == flag {
		flag = log.Llongfile
	}
	if wt == nil {
		wt = os.Stdout
	}
	return Glog{log.New(wt, "", flag)}
}

func (g *Glog) Fatalf(format string, v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "35"))
	g.Output(3, color.MagentaString(format, v...))
}
func (g *Glog) Fatalln(v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "35"))
	g.Output(3, color.MagentaString("", v...))

}
func (g *Glog) Errorf(format string, v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "31"))
	g.Output(3, color.RedString(format, v...))

}
func (g *Glog) Errorln(v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "31"))
	g.Output(3, color.RedString("", v...))

}
func (g *Glog) Debugf(format string, v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "33"))
	g.Output(3, color.YellowString(format, v...))

}
func (g *Glog) Debugln(v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "33"))
	g.Output(3, color.YellowString("", v...))

}
func (g *Glog) Infof(format string, v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "32"))
	g.Output(3, color.GreenString(format, v...))

}
func (g *Glog) Infoln(v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "32"))
	g.Output(3, color.GreenString("", v...))

}

func Fatalf(format string, v ...interface{}) {
	gl.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	gl.Fatalln(v...)
}
func Errorf(format string, v ...interface{}) {
	gl.Errorf(format, v...)

}
func Errorln(v ...interface{}) {
	gl.Errorln(v...)

}
func Debugf(format string, v ...interface{}) {
	gl.Debugf(format, v...)

}
func Debugln(v ...interface{}) {
	gl.Debugln(v...)

}
func Infof(format string, v ...interface{}) {
	gl.Infof(format, v...)

}
func Infoln(v ...interface{}) {
	gl.Infoln(v...)

}
