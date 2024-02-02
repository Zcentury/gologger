package writer

import (
	"fmt"
	"github.com/Zcentury/gologger/levels"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type FileWrite struct {
	location string
	file     *os.File
	mutex    *sync.Mutex
}

var FileWriteOptions *FileWrite

func init() {
	FileWriteOptions = &FileWrite{}
	if dir, err := os.Getwd(); err == nil {
		FileWriteOptions.location = filepath.Join(dir, "logs")
	}
}

func NewFile() *FileWrite {
	return &FileWrite{
		mutex: &sync.Mutex{},
	}
}

func (w *FileWrite) Write(data []byte, level levels.Level) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	// 创建目录
	if err := os.MkdirAll(FileWriteOptions.location, os.ModePerm); err != nil {
		fmt.Println("无法创建目录:", err)
	}

	fileName := fmt.Sprintf("%s.txt", time.Now().Format("2006-01-02"))

	if file, err := os.OpenFile(filepath.Join(FileWriteOptions.location, fileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		w.file = file
	} else {
		fmt.Println("无法打开日志文件", err)
	}

	defer FileWriteOptions.file.Close()

	if w.file != nil {
		if _, err := w.file.WriteString(string(data) + "\n"); err != nil {
			fmt.Println("文件写入失败", err)
		}
	} else {
		fmt.Println("文件没打开")
	}

	os.Stderr.Write(data)
	os.Stderr.Write([]byte("\n"))

}
