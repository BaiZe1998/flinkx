package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"log"
	"word-count/config"
	"word-count/kitex_gen/mapdemo"
	"word-count/kitex_gen/mapdemo/mapservice"
	"word-count/pkg/constants"
	"word-count/pkg/errno"
)

var mapClient mapservice.Client

func InitMapRPC() {

	enginConfig := config.GetConfig()

	ServiceAddr := fmt.Sprintf("%s:%d", enginConfig.Map.Ip, enginConfig.Map.Port)

	c, err := mapservice.NewClient(constants.MapServiceName, client.WithHostPorts(ServiceAddr))
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
