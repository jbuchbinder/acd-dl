package main

import (
	"os"
)

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func fileSize(name string) int64 {
	file, err := os.Open(name)
	if err != nil {
		return -1
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return -1
	}
	return stat.Size()
}
