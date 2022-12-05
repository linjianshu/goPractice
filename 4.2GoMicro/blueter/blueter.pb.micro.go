// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: blueter.proto

package blueter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	//v3不行
	//client "github.com/micro/micro/v3/service/client"
	//server "github.com/micro/micro/v3/service/server"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Blueter service

func NewBlueterEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Blueter service

type BlueterService interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error)
}

type blueterService struct {
	c    client.Client
	name string
}

func NewBlueterService(name string, c client.Client) BlueterService {
	return &blueterService{
		c:    c,
		name: name,
	}
}

func (c *blueterService) Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error) {
	req := c.c.NewRequest(c.name, "Blueter.Hello", in)
	out := new(HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Blueter service

type BlueterHandler interface {
	Hello(context.Context, *HelloRequest, *HelloResponse) error
}

func RegisterBlueterHandler(s server.Server, hdlr BlueterHandler, opts ...server.HandlerOption) error {
	type blueter interface {
		Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error
	}
	type Blueter struct {
		blueter
	}
	h := &blueterHandler{hdlr}
	return s.Handle(s.NewHandler(&Blueter{h}, opts...))
}

type blueterHandler struct {
	BlueterHandler
}

func (h *blueterHandler) Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error {
	return h.BlueterHandler.Hello(ctx, in, out)
}