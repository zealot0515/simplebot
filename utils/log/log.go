package log

import (
	"fmt"
	"os"
)

func PrintDebugLog(enableLog bool, a ...interface{}) {
	if enableLog {
		fmt.Println(a...)
	}
}

func CheckFile(path string) {
	if _, err := os.Stat(path); err == nil {
		return
	}
	if _, err := os.Create(path); err != nil {
		fmt.Println(err)
	}
}

func WriteFileLog(path string, message string) {
	CheckFile(path)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err = f.WriteString(message + "\n"); err != nil {
		fmt.Println(err)
	}
}
