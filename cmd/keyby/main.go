package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"log"
	"net"
	"word-count/config"
	keybydemo "word-count/kitex_gen/keybydemo/keybyservice"
	"word-count/pkg/constants"
)

func main() {
	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		panic(err)
	}

	enginConfig := config.GetConfig()

	ServiceAddr := fmt.Sprintf("%s:%d", enginConfig.Keyby.Ip, enginConfig.Keyby.Port)

	addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	if err != nil {
		klog.Fatal(err)
	}

	klog.Info("addr = ", addr)

	svr := keybydemo.NewServer(
		new(KeybyServiceImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.KeybyServiceName}),
		server.WithServiceAddr(addr), // address
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
