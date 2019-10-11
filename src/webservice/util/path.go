package util

import (
	"path"
	"runtime"
)

func GetCurrentDir() (string) {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
