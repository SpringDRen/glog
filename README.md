# Overview
make glog easily to use.  
1. Change default log_dir to ./logs. If logDir doesn't exist, try to create it.
2. Remove host、userName、pid from the filename.
3. change flushInterval to 1s
4. change log header

## glog doc

```
Log output is buffered and written periodically using Flush. Programs
should call Flush before exiting to guarantee all log output is written.

By default, all log statements write to files in a temporary directory.
This package provides several flags that modify this behavior.
As a result, flag.Parse must be called before any logging is done.

    -logtostderr=false
        Logs are written to standard error instead of to files.
    -alsologtostderr=false
        Logs are written to standard error as well as to files.
    -stderrthreshold=ERROR
        Log events at or above this severity are logged to standard
        error as well as to files.
    -log_dir=""
        Log files will be written to this directory instead of the
        default temporary directory.

    Other flags provide aids to debugging.

    -log_backtrace_at=""
        When set to a file and line number holding a logging statement,
        such as
            -log_backtrace_at=gopherflakes.go:234
        a stack trace will be written to the Info log whenever execution
        hits that statement. (Unlike with -vmodule, the ".go" must be
        present.)
    -v=0
        Enable V-leveled logging at the specified level.
    -vmodule=""
        The syntax of the argument is a comma-separated list of pattern=N,
        where pattern is a literal file name (minus the ".go" suffix) or
        "glob" pattern and N is a V level. For instance,
            -vmodule=gopher*=3
        sets the V level to 3 in all Go files whose names begin "gopher".
```

## get & use
- use go get : `go get -u github.com/SpringDRen/glog`
- use dep : `dep ensure -add github.com/SpringDRen/glog`

```go
package main

import (
	"flag"
	"github.com/SpringDRen/glog"
)

// go run demo1.go -alsologtostderr=true -log_dir=./logs -v=3
func main() {
	flag.Parse()
	// set file size = 100M
	glog.MaxSize = 1024 * 1024 * 100
	defer glog.Flush()
	glog.Info("info log")
	if glog.V(3) {
		glog.Info("v3 log")
	}
	if glog.V(4) {
		glog.Info("v4 log")
	}
	glog.Warning("warning log")
	glog.Error("error log")
}
```

# original README.md
```
glog
====

Leveled execution logs for Go.

This is an efficient pure Go implementation of leveled logs in the
manner of the open source C++ package
	https://github.com/google/glog

By binding methods to booleans it is possible to use the log package
without paying the expense of evaluating the arguments to the log.
Through the -vmodule flag, the package also provides fine-grained
control over logging at the file level.

The comment from glog.go introduces the ideas:

	Package glog implements logging analogous to the Google-internal
	C++ INFO/ERROR/V setup.  It provides functions Info, Warning,
	Error, Fatal, plus formatting variants such as Infof. It
	also provides V-style logging controlled by the -v and
	-vmodule=file=2 flags.

	Basic examples:

		glog.Info("Prepare to repel boarders")

		glog.Fatalf("Initialization failed: %s", err)

	See the documentation for the V function for an explanation
	of these examples:

		if glog.V(2) {
			glog.Info("Starting transaction...")
		}

		glog.V(2).Infoln("Processed", nItems, "elements")


The repository contains an open source version of the log package
used inside Google. The master copy of the source lives inside
Google, not here. The code in this repo is for export only and is not itself
under development. Feature requests will be ignored.

Send bug reports to golang-nuts@googlegroups.com.
```
