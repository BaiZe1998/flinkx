package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"log"
	"word-count/config"
	"word-count/kitex_gen/reducedemo"
	"word-count/kitex_gen/reducedemo/reduceservice"
	"word-count/pkg/constants"
	"word-count/pkg/errno"
)

var reduceClient reduceservice.Client

func InitReduceRPC() {

	enginConfig := config.GetConfig()

	ServiceAddr := fmt.Sprintf("%s:%d", enginConfig.Reduce.Ip, enginConfig.Reduce.Port)

	c, err := reduceservice.NewClient(constants.ReduceServiceName, client.WithHostPorts(ServiceAddr))
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
