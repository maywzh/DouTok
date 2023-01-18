// Code generated by Kitex v0.3.4. DO NOT EDIT.

package publishservice

import (
	"context"
	"github.com/TremblingV5/DouTok/kitex_gen/publish"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PublishAction(ctx context.Context, Req *publish.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *publish.DouyinPublishActionResponse, err error)
	PublishList(ctx context.Context, Req *publish.DouyinPublishListRequest, callOptions ...callopt.Option) (r *publish.DouyinPublishListResponse, err error)
}

// NewClient creates a rpc for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kPublishServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a rpc for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kPublishServiceClient struct {
	*kClient
}

func (p *kPublishServiceClient) PublishAction(ctx context.Context, Req *publish.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *publish.DouyinPublishActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishAction(ctx, Req)
}

func (p *kPublishServiceClient) PublishList(ctx context.Context, Req *publish.DouyinPublishListRequest, callOptions ...callopt.Option) (r *publish.DouyinPublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishList(ctx, Req)
}
