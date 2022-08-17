package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"word-count/cmd/api/handler"
	"word-count/config"
)

func main() {
	enginConfig := config.GetConfig()

	ServiceAddr := fmt.Sprintf("%s:%d", enginConfig.Job.Ip, enginConfig.Job.Port)

	// WithMaxRequestBodySize can set the size of the body
	h := server.Default(server.WithHostPorts(ServiceAddr), server.WithMaxRequestBodySize(20<<20))

	h.POST("/upload", handler.Upload)

	h.Spin()
}
