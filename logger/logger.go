package logger

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"time"
)

var (
	logger *log.Logger
	file   *os.File
)

func Start(name string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	logger = log.New(file, "", 2)
	logger.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	return nil
}

func Write(obj interface{}) (int, error) {
	data, err := json.Marshal(struct {
		Object    interface{} `json:"Entry"`
		Timestamp int64
	}{
		Object:    obj,
		Timestamp: time.Now().UnixNano(),
	})
	if err != nil {
		return 0, err
	}

	logger.Println(bytes.NewBuffer(data))
	return len(data), nil
}

func Stop() error {
	err := file.Close()
	if err != nil {
		return err
	}
	return nil
}
