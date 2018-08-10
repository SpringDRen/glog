package main

import (
	"flag"
	"github.com/SpringDRen/glog"
)

// go run demo1.go -alsologtostderr=true
func main() {
	flag.Parse()
	// set file size = 100M
	glog.MaxSize = 1024 * 1024 * 100
	defer glog.Flush()
	glog.Info("info log")
	glog.Warning("warning log")
	glog.Error("error log")
}
