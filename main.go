package main

import (
	"fmt"
	"log"
)

func main() {
	logf("start program", "INFO")

	logf("enter input data", "WARN")
	var a, b, c int
	fmt.Println("Enter three numbers: ")
	fmt.Scan(&a, &b, &c)

	logf("calculating max value", "DEBUG")
	m := a
	if b > a {
		m = b
	}
	if c > m {
		m = c
	}

	logf("print output", "INFO")
	fmt.Println("max:", m)
	logf("end program", "ERROR")
}

func logf(msg, level string) {
	log.Print(level, " ", msg)
}
