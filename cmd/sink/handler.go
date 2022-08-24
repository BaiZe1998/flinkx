package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"io"
	"log"
	"os"
	"word-count/cmd/sink/pack"
	"word-count/kitex_gen/sinkdemo"
	"word-count/pkg/errno"
)

// SinkServiceImpl implements the last service interface defined in the IDL.
type SinkServiceImpl struct{}

var (
	fileName string
	content  string
	file     *os.File
	err      error
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("获取文件路径失败：", err)
	}
	fileName = wd + "/cmd/sink/data.txt"
}

// CreateSink implements the SinkServiceImpl interface.
func (s *SinkServiceImpl) CreateSink(ctx context.Context, req *sinkdemo.CreateSinkRequest) (resp *sinkdemo.CreateSinkResponse, err error) {

	content = fmt.Sprintf("table:%v timeStamp:%v\n", req.Tables, req.TimeStamp)
	klog.Info(content)
	WriteFile(content)

	resp = new(sinkdemo.CreateSinkResponse)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

func WriteFile(content string) {
	//文件是否存在
	if Exists(fileName) {
		//使用追加模式打开文件
		file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			klog.Error("打开文件错误：", err)
			return
		}
	} else {
		//不存在创建文件
		file, err = os.Create(fileName)
		if err != nil {
			klog.Error("创建失败", err)
			return
		}
	}

	defer file.Close()
	//写入文件
	n, err := io.WriteString(file, content)
	if err != nil {
		klog.Error("写入错误：", err)
		return
	}
	klog.Info("写入成功：n=", n)
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
