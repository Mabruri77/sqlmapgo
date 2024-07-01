package main

import (
	"MySqlMap/code"
	"os"
)

func main() {
	if os.Args[1] == "-u" {
		code.ScanUrl()
	} else if os.Args[1] == "-r" {
		code.ScanPost()
	}
}
