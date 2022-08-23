package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"log"
	"time"
	"word-count/kitex_gen/sinkdemo"
	"word-count/kitex_gen/sinkdemo/sinkservice"
	"word-count/pkg/constants"
	"word-count/pkg/errno"
)

var sinkClient sinkservice.Client

func InitSinkRPC() {

	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}

	c, err := sinkservice.NewClient(
		constants.SinkServiceName,
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*3),
	)
	if err != nil {
		log.Fatal(err)
	}
	sinkClient = c
}

// CreateKeyby create keyby info
func CreateSink(ctx context.Context, req *sinkdemo.CreateSinkRequest) error {
	resp, err := sinkClient.CreateSink(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	klog.Info("sink resp: ", resp)
	return nil
}
