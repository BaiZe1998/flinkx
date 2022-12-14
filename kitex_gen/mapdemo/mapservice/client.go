// Code generated by Kitex v0.3.4. DO NOT EDIT.

package mapservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"word-count/kitex_gen/mapdemo"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateMap(ctx context.Context, req *mapdemo.CreateMapRequest, callOptions ...callopt.Option) (r *mapdemo.CreateMapResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kMapServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kMapServiceClient struct {
	*kClient
}

func (p *kMapServiceClient) CreateMap(ctx context.Context, req *mapdemo.CreateMapRequest, callOptions ...callopt.Option) (r *mapdemo.CreateMapResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateMap(ctx, req)
}
