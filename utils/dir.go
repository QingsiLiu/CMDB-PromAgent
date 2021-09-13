package utils

import (
	"os"
	"path/filepath"
)

func MKdir(path string) {
	os.MkdirAll(path, os.ModePerm)
}

func MKPdir(path string) {
	MKdir(filepath.Dir(path))
}
