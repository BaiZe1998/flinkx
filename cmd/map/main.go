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
	mapdemo "word-count/kitex_gen/mapdemo/mapservice"
	"word-count/pkg/constants"
)

func main() {

	r, err := registry.NewDefaultNacosRegistry()
	if err != nil {
		panic(err)
	}

	enginConfig := config.GetConfig()

	ServiceAddr := fmt.Sprintf("%s:%d", enginConfig.Map.Ip, enginConfig.Map.Port)

	addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	if err != nil {
		klog.Fatal(err)
	}

	klog.Info("addr = ", addr)

	svr := mapdemo.NewServer(
		new(MapServiceImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.MapServiceName}),
		server.WithServiceAddr(addr), // address
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
