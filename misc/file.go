package misc

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetFileSize get the length in bytes of file of the specified path.
func GetFileSize(path string) int64 {
	fi, err := os.Stat(path)
	if nil != err {
		return -1
	}

	return fi.Size()
}

// IsExist determines whether the file spcified by the given path is exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

// IsExist determines whether the file spcified by the given path is exists.
func Mkdir(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

// IsDir determines whether the specified path is a directory.
func IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}

	if nil != err {
		return false
	}

	return fio.IsDir()
}

func BaseDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
