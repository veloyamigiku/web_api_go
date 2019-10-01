package main

import (
	"path"
	"runtime"
)

func getCurrentDir() (string) {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
