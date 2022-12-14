// Code generated by Kitex v0.3.4. DO NOT EDIT.

package keybyservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"word-count/kitex_gen/keybydemo"
)

func serviceInfo() *kitex.ServiceInfo {
	return keybyServiceServiceInfo
}

var keybyServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "KeybyService"
	handlerType := (*keybydemo.KeybyService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateKeyby": kitex.NewMethodInfo(createKeybyHandler, newKeybyServiceCreateKeybyArgs, newKeybyServiceCreateKeybyResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "keybydemo",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.4",
		Extra:           extra,
	}
	return svcInfo
}

func createKeybyHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*keybydemo.KeybyServiceCreateKeybyArgs)
	realResult := result.(*keybydemo.KeybyServiceCreateKeybyResult)
	success, err := handler.(keybydemo.KeybyService).CreateKeyby(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newKeybyServiceCreateKeybyArgs() interface{} {
	return keybydemo.NewKeybyServiceCreateKeybyArgs()
}

func newKeybyServiceCreateKeybyResult() interface{} {
	return keybydemo.NewKeybyServiceCreateKeybyResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateKeyby(ctx context.Context, req *keybydemo.CreateKeybyRequest) (r *keybydemo.CreateKeybyResponse, err error) {
	var _args keybydemo.KeybyServiceCreateKeybyArgs
	_args.Req = req
	var _result keybydemo.KeybyServiceCreateKeybyResult
	if err = p.c.Call(ctx, "CreateKeyby", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
