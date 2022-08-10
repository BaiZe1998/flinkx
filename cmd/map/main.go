package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"word-count/config"
	mapdemo "word-count/kitex_gen/mapdemo/mapservice"
)

func main() {

	enginConfig := config.GetConfig()

	ServiceAddr := fmt.Sprintf("%s:%d", enginConfig.Map.Ip, enginConfig.Map.Port)

	addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	if err != nil {
		klog.Fatal(err)
	}

	klog.Info("addr = ", addr)

	svr := mapdemo.NewServer(
		new(MapServiceImpl),
		server.WithServiceAddr(addr), // address
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
