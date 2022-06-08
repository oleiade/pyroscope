// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: push/v1/push.proto

package pushv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/grafana/fire/pkg/gen/push/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// PusherName is the fully-qualified name of the Pusher service.
	PusherName = "push.v1.Pusher"
	// IngesterName is the fully-qualified name of the Ingester service.
	IngesterName = "push.v1.Ingester"
)

// PusherClient is a client for the push.v1.Pusher service.
type PusherClient interface {
	Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error)
}

// NewPusherClient constructs a client for the push.v1.Pusher service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPusherClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PusherClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &pusherClient{
		push: connect_go.NewClient[v1.PushRequest, v1.PushResponse](
			httpClient,
			baseURL+"/push.v1.Pusher/Push",
			opts...,
		),
	}
}

// pusherClient implements PusherClient.
type pusherClient struct {
	push *connect_go.Client[v1.PushRequest, v1.PushResponse]
}

// Push calls push.v1.Pusher.Push.
func (c *pusherClient) Push(ctx context.Context, req *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error) {
	return c.push.CallUnary(ctx, req)
}

// PusherHandler is an implementation of the push.v1.Pusher service.
type PusherHandler interface {
	Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error)
}

// NewPusherHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPusherHandler(svc PusherHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/push.v1.Pusher/Push", connect_go.NewUnaryHandler(
		"/push.v1.Pusher/Push",
		svc.Push,
		opts...,
	))
	return "/push.v1.Pusher/", mux
}

// UnimplementedPusherHandler returns CodeUnimplemented from all methods.
type UnimplementedPusherHandler struct{}

func (UnimplementedPusherHandler) Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("push.v1.Pusher.Push is not implemented"))
}

// IngesterClient is a client for the push.v1.Ingester service.
type IngesterClient interface {
	Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error)
}

// NewIngesterClient constructs a client for the push.v1.Ingester service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIngesterClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) IngesterClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &ingesterClient{
		push: connect_go.NewClient[v1.PushRequest, v1.PushResponse](
			httpClient,
			baseURL+"/push.v1.Ingester/Push",
			opts...,
		),
	}
}

// ingesterClient implements IngesterClient.
type ingesterClient struct {
	push *connect_go.Client[v1.PushRequest, v1.PushResponse]
}

// Push calls push.v1.Ingester.Push.
func (c *ingesterClient) Push(ctx context.Context, req *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error) {
	return c.push.CallUnary(ctx, req)
}

// IngesterHandler is an implementation of the push.v1.Ingester service.
type IngesterHandler interface {
	Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error)
}

// NewIngesterHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIngesterHandler(svc IngesterHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/push.v1.Ingester/Push", connect_go.NewUnaryHandler(
		"/push.v1.Ingester/Push",
		svc.Push,
		opts...,
	))
	return "/push.v1.Ingester/", mux
}

// UnimplementedIngesterHandler returns CodeUnimplemented from all methods.
type UnimplementedIngesterHandler struct{}

func (UnimplementedIngesterHandler) Push(context.Context, *connect_go.Request[v1.PushRequest]) (*connect_go.Response[v1.PushResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("push.v1.Ingester.Push is not implemented"))
}
