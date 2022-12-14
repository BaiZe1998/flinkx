// Code generated by Kitex v0.3.4. DO NOT EDIT.

package sinkservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"word-count/kitex_gen/sinkdemo"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateSink(ctx context.Context, req *sinkdemo.CreateSinkRequest, callOptions ...callopt.Option) (r *sinkdemo.CreateSinkResponse, err error)
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
	return &kSinkServiceClient{
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

type kSinkServiceClient struct {
	*kClient
}

func (p *kSinkServiceClient) CreateSink(ctx context.Context, req *sinkdemo.CreateSinkRequest, callOptions ...callopt.Option) (r *sinkdemo.CreateSinkResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateSink(ctx, req)
}
