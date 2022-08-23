package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"log"
	"time"
	"word-count/kitex_gen/mapdemo"
	"word-count/kitex_gen/mapdemo/mapservice"
	"word-count/pkg/constants"
	"word-count/pkg/errno"
)

var mapClient mapservice.Client

func InitMapRPC() {

	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}
	c, err := mapservice.NewClient(
		constants.MapServiceName,
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*3),
		client.WithCircuitBreaker(circuitbreak.NewCBSuite(nil)))
	if err != nil {
		log.Fatal(err)
	}
	mapClient = c
}

// CreateMap create map info
func CreateMap(ctx context.Context, req *mapdemo.CreateMapRequest) error {
	resp, err := mapClient.CreateMap(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	klog.Info("map resp: ", resp)
	return nil
}
