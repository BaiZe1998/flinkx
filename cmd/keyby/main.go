package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"word-count/config"
	keybydemo "word-count/kitex_gen/keybydemo/keybyservice"
)

func main() {
	enginConfig := config.GetConfig()

	ServiceAddr := fmt.Sprintf("%s:%d", enginConfig.Keyby.Ip, enginConfig.Keyby.Port)

	addr, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	if err != nil {
		klog.Fatal(err)
	}

	svr := keybydemo.NewServer(
		new(KeybyServiceImpl),
		server.WithServiceAddr(addr), // address
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
