package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"word-count/cmd/sink/kitex_gen/sinkdemo"
	"word-count/cmd/sink/pack"
	"word-count/pkg/errno"
)

// SinkServiceImpl implements the last service interface defined in the IDL.
type SinkServiceImpl struct{}

// CreateSink implements the SinkServiceImpl interface.
func (s *SinkServiceImpl) CreateSink(ctx context.Context, req *sinkdemo.CreateSinkRequest) (resp *sinkdemo.CreateSinkResponse, err error) {

	klog.Info(fmt.Sprintf("从reduce接收到数据=========> %v", req.Tables))

	resp = new(sinkdemo.CreateSinkResponse)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
