package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	sinkdemo "word-count/cmd/sink/kitex_gen/sinkdemo/sinkservice"
	"word-count/config"
)

func main() {
	enginConfig := config.GetConfig()

	ServiceAddr := fmt.Sprintf("%s:%d", enginConfig.Sink.Ip, enginConfig.Sink.Port)

	addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	if err != nil {
		klog.Fatal(err)
	}

	klog.Info("addr = ", addr)

	svr := sinkdemo.NewServer(
		new(SinkServiceImpl),
		server.WithServiceAddr(addr), // address
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
