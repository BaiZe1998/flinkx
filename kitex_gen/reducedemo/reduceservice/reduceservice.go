// Code generated by Kitex v0.3.4. DO NOT EDIT.

package reduceservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"word-count/kitex_gen/reducedemo"
)

func serviceInfo() *kitex.ServiceInfo {
	return reduceServiceServiceInfo
}

var reduceServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ReduceService"
	handlerType := (*reducedemo.ReduceService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateReduce": kitex.NewMethodInfo(createReduceHandler, newReduceServiceCreateReduceArgs, newReduceServiceCreateReduceResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "reducedemo",
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

func createReduceHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*reducedemo.ReduceServiceCreateReduceArgs)
	realResult := result.(*reducedemo.ReduceServiceCreateReduceResult)
	success, err := handler.(reducedemo.ReduceService).CreateReduce(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newReduceServiceCreateReduceArgs() interface{} {
	return reducedemo.NewReduceServiceCreateReduceArgs()
}

func newReduceServiceCreateReduceResult() interface{} {
	return reducedemo.NewReduceServiceCreateReduceResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateReduce(ctx context.Context, req *reducedemo.CreateReduceRequest) (r *reducedemo.CreateReduceResponse, err error) {
	var _args reducedemo.ReduceServiceCreateReduceArgs
	_args.Req = req
	var _result reducedemo.ReduceServiceCreateReduceResult
	if err = p.c.Call(ctx, "CreateReduce", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
