package config

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

type FlushableFileHook struct {
	mu   sync.Mutex
	file *os.File
}

func InitFileLogger(filePath string) *FlushableFileHook {
	hook, err := NewFlushableFileHook(filePath)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.AddHook(hook)
	return hook
}

func NewFlushableFileHook(filePath string) (*FlushableFileHook, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &FlushableFileHook{file: file}, nil
}

func (hook *FlushableFileHook) Fire(entry *logrus.Entry) error {
	hook.mu.Lock()
	defer hook.mu.Unlock()

	line, err := entry.String()
	if err != nil {
		return err
	}

	_, err = hook.file.WriteString(line)
	return err
}

func (hook *FlushableFileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *FlushableFileHook) Flush() {
	hook.mu.Lock()
	defer hook.mu.Unlock()

	hook.file.Sync()
}
