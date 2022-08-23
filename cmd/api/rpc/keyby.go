package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"log"
	"time"
	"word-count/kitex_gen/keybydemo"
	"word-count/kitex_gen/keybydemo/keybyservice"
	"word-count/pkg/constants"
	"word-count/pkg/errno"
)

var keybyClient keybyservice.Client

func InitKeybyRPC() {
	r, err := resolver.NewDefaultNacosResolver()
	if err != nil {
		panic(err)
	}

	c, err := keybyservice.NewClient(
		constants.KeybyServiceName,
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*3),
	)
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
