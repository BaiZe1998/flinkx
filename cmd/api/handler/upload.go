package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
	"io"
	"os"
)

func Upload(ctx context.Context, c *app.RequestContext) {
	// single file
	file, _ := c.FormFile("file")
	klog.Info(file.Filename)

	wd, err := os.Getwd()
	if err != nil {
		klog.Fatal("获取文件路径失败：", err)
	}

	klog.Info(wd + "/" + file.Filename)

	// Upload the file to specific dst
	c.SaveUploadedFile(file, wd+"/config/x"+file.Filename)

	if err = copy(wd+"/config/x"+file.Filename, wd+"/config/"+file.Filename); err != nil {
		klog.Errorf(err.Error())
	}

	c.String(consts.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func copy(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}
