package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"log"
	"word-count/config"
	"word-count/kitex_gen/keybydemo"
	"word-count/kitex_gen/keybydemo/keybyservice"
	"word-count/pkg/constants"
	"word-count/pkg/errno"
)

var keybyClient keybyservice.Client

func InitKeybyRPC() {

	enginConfig := config.GetConfig()

	ServiceAddr := fmt.Sprintf("%s:%d", enginConfig.Keyby.Ip, enginConfig.Keyby.Port)

	c, err := keybyservice.NewClient(constants.KeybyServiceName, client.WithHostPorts(ServiceAddr))
	if err != nil {
		log.Fatal(err)
	}
	keybyClient = c
}

// CreateKeyby create keyby info
func CreateKeyby(ctx context.Context, req *keybydemo.CreateKeybyRequest) error {
	resp, err := keybyClient.CreateKeyby(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	klog.Info("keyby resp: ", resp)
	return nil
}
