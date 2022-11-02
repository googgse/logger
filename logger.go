package logger

import (
	// import built-in packages
	"fmt"
	"log"
	"os"
	"time"
)

// 设置日志输出文件
func SetOutput(file string) {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	// control flows
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile)
}

// Debug
func Debug(text string) {
	timestamp := time.Now().Unix()
	prefix := fmt.Sprintf(`{"timestamp":"%d","level":"DEBUG","msg":"`, timestamp)
	log.SetPrefix(prefix)
	log.Printf(`%s"}`, text)
}

// Info
func Info(text string) {
	timestamp := time.Now().Unix()
	prefix := fmt.Sprintf(`{"timestamp":"%d","level":"INFO","msg":"`, timestamp)
	log.SetPrefix(prefix)
	log.Printf(`%s"}`, text)
}

// Warning
func Warning(text string) {
	timestamp := time.Now().Unix()
	prefix := fmt.Sprintf(`{"timestamp":"%d","level":"WARNING","msg":"`, timestamp)
	log.SetPrefix(prefix)
	log.Printf(`%s"}`, text)
}

// Error
func Error(text string) {
	timestamp := time.Now().Unix()
	prefix := fmt.Sprintf(`{"timestamp":"%d","level":"ERROR","msg":"`, timestamp)
	log.SetPrefix(prefix)
	log.Printf(`%s"}`, text)
}
