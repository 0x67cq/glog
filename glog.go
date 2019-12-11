/******************************************
        author: guanchengqi
        email: guancq@tuya.com
        date: 2019.12.07(Saturday)
******************************************/
package glog

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"os"
)

type Glog struct {
	*log.Logger
}

var Gl Glog

func InitGlog(wt io.Writer) {
	if wt == nil {
		wt = os.Stdout
	}
	Gl = Glog{log.New(wt, "", log.Llongfile)}
}

func (g *Glog) Fatalf(format string, v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "35"))
	g.Printf(color.MagentaString(format, v...))
}
func (g *Glog) Fatalln(v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "35"))
	g.Print(color.MagentaString("", v...))

}
func (g *Glog) Errorf(format string, v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "31"))
	g.Printf(color.RedString(format, v...))

}
func (g *Glog) Errorln(v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "31"))
	g.Print(color.RedString("", v...))

}
func (g *Glog) Debugf(format string, v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "33"))
	g.Printf(color.YellowString(format, v...))

}
func (g *Glog) Debugln(v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "33"))
	g.Print(color.YellowString("", v...))

}
func (g *Glog) Infof(format string, v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "32"))
	g.Printf(color.GreenString(format, v...))

}
func (g *Glog) Infoln(v ...interface{}) {
	g.SetPrefix(fmt.Sprintf("%s[%sm", "\x1b", "32"))
	g.Print(color.GreenString("", v...))

}

func Fatalf(format string, v ...interface{}) {
	Gl.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	Gl.Fatalln(v...)
}
func Errorf(format string, v ...interface{}) {
	Gl.Errorf(format, v...)

}
func Errorln(v ...interface{}) {
	Gl.Errorln(v...)

}
func Debugf(format string, v ...interface{}) {
	Gl.Debugf(format, v...)

}
func Debugln(v ...interface{}) {
	Gl.Debugln(v...)

}
func Infof(format string, v ...interface{}) {
	Gl.Infof(format, v...)

}
func Infoln(v ...interface{}) {
	Gl.Infoln(v...)

}
