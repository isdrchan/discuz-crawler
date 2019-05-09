package persist

import (
	"bytes"
	"dicuz-crawler/model"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"
)

type FileSaver struct {
	File *os.File
}

func (f *FileSaver) Init() {
	fileName := "save" + strconv.FormatInt(time.Now().Unix(), 10) + ".txt"
	var err error
	f.File, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Println("创建(打开)文件"+fileName+"失败: %s", err)
	}
}

func (f *FileSaver) Save(item model.Item) error {
	//禁止转义
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(item)
	_, err := f.File.Write(buffer.Bytes())
	return err
}

func (f *FileSaver) Close() {
	f.File.Close()
}
