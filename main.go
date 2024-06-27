package main

import (
	"fmt"

	"github.com/juju/loggo"
)

func main() {
	l := loggo.GetLogger("test")

	l.SetLogLevel(loggo.TRACE)

	l.Infof("Start of program")
	var a, b, c int
	fmt.Println("Enter three numbers: ")
	fmt.Scan(&a, &b, &c)

	l.Debugf("calculating max value")
	m := a
	if b > a {
		m = b
	}
	if c > m {
		m = c
	}

	l.Warningf("print output")
	fmt.Println("max:", m)
	l.Errorf("end program")
}
