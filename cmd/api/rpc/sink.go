package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"log"
	"word-count/config"
	"word-count/kitex_gen/sinkdemo"
	"word-count/kitex_gen/sinkdemo/sinkservice"
	"word-count/pkg/constants"
	"word-count/pkg/errno"
)

var sinkClient sinkservice.Client

func InitSinkRPC() {

	enginConfig := config.GetConfig()

	ServiceAddr := fmt.Sprintf("%s:%d", enginConfig.Sink.Ip, enginConfig.Sink.Port)

	c, err := sinkservice.NewClient(constants.SinkServiceName, client.WithHostPorts(ServiceAddr))
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
