package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"log"
	"time"
	"word-count/kitex_gen/reducedemo"
	"word-count/kitex_gen/reducedemo/reduceservice"
	"word-count/pkg/constants"
	"word-count/pkg/errno"
)

var reduceClient reduceservice.Client

func InitReduceRPC() {

	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}

	c, err := reduceservice.NewClient(
		constants.ReduceServiceName,
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*3))
	if err != nil {
		log.Fatal(err)
	}
	reduceClient = c
}

// CreateKeyby create keyby info
func CreateReduce(ctx context.Context, req *reducedemo.CreateReduceRequest) error {
	resp, err := reduceClient.CreateReduce(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	klog.Info("reduce resp: ", resp)
	return nil
}
