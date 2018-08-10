package main

import (
	"flag"
	"github.com/SpringDRen/glog"
)

// default v=0
// go run demo1.go -v 2 -stderrthreshold=INFO
func main() {
	flag.Parse()
	defer glog.Flush()
	glog.Info("info log")
	glog.V(-1).Info("v-1 log")
	glog.V(0).Info("v0 log")
	glog.V(1).Info("v1 log")
	glog.V(3).Info("v3 log")
}
