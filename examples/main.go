package main

import (
	"github.com/psychix/glog"
	"os"
)

func main() {
	glog.InitGlog(os.Stdout, 0)
	//glog.InitGlog(nil)
	glog.Infof("this is info")
	glog.Debugf("this is debug")
	glog.Errorf("this is error")
	glog.Fatalf("this is fata")

}
