// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cache/cache_service.proto

package cache

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for Cache service

func NewCacheEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "Cache.Put",
			Path:    []string{"/v1/cache"},
			Method:  []string{"POST"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Cache.Set",
			Path:    []string{"/v1/cache/{key}"},
			Method:  []string{"PUT"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Cache.Get",
			Path:    []string{"/v1/cache/{key}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Cache.Delete",
			Path:    []string{"/v1/cache/{key}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Cache.List",
			Path:    []string{"/v1/cache"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
	}
}

// Client API for Cache service

type CacheService interface {
	Put(ctx context.Context, in *CacheMessage, opts ...client.CallOption) (*CacheMessage, error)
	Set(ctx context.Context, in *CacheMessage, opts ...client.CallOption) (*CacheMessage, error)
	Get(ctx context.Context, in *CacheKeyRequest, opts ...client.CallOption) (*CacheMessage, error)
	Delete(ctx context.Context, in *CacheKeyRequest, opts ...client.CallOption) (*CacheMessage, error)
	List(ctx context.Context, in *CacheListRequest, opts ...client.CallOption) (*CacheListResponse, error)
}

type cacheService struct {
	c    client.Client
	name string
}

func NewCacheService(name string, c client.Client) CacheService {
	return &cacheService{
		c:    c,
		name: name,
	}
}

func (c *cacheService) Put(ctx context.Context, in *CacheMessage, opts ...client.CallOption) (*CacheMessage, error) {
	req := c.c.NewRequest(c.name, "Cache.Put", in)
	out := new(CacheMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheService) Set(ctx context.Context, in *CacheMessage, opts ...client.CallOption) (*CacheMessage, error) {
	req := c.c.NewRequest(c.name, "Cache.Set", in)
	out := new(CacheMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheService) Get(ctx context.Context, in *CacheKeyRequest, opts ...client.CallOption) (*CacheMessage, error) {
	req := c.c.NewRequest(c.name, "Cache.Get", in)
	out := new(CacheMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheService) Delete(ctx context.Context, in *CacheKeyRequest, opts ...client.CallOption) (*CacheMessage, error) {
	req := c.c.NewRequest(c.name, "Cache.Delete", in)
	out := new(CacheMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheService) List(ctx context.Context, in *CacheListRequest, opts ...client.CallOption) (*CacheListResponse, error) {
	req := c.c.NewRequest(c.name, "Cache.List", in)
	out := new(CacheListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cache service

type CacheHandler interface {
	Put(context.Context, *CacheMessage, *CacheMessage) error
	Set(context.Context, *CacheMessage, *CacheMessage) error
	Get(context.Context, *CacheKeyRequest, *CacheMessage) error
	Delete(context.Context, *CacheKeyRequest, *CacheMessage) error
	List(context.Context, *CacheListRequest, *CacheListResponse) error
}

func RegisterCacheHandler(s server.Server, hdlr CacheHandler, opts ...server.HandlerOption) error {
	type cache interface {
		Put(ctx context.Context, in *CacheMessage, out *CacheMessage) error
		Set(ctx context.Context, in *CacheMessage, out *CacheMessage) error
		Get(ctx context.Context, in *CacheKeyRequest, out *CacheMessage) error
		Delete(ctx context.Context, in *CacheKeyRequest, out *CacheMessage) error
		List(ctx context.Context, in *CacheListRequest, out *CacheListResponse) error
	}
	type Cache struct {
		cache
	}
	h := &cacheHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Cache.Put",
		Path:    []string{"/v1/cache"},
		Method:  []string{"POST"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Cache.Set",
		Path:    []string{"/v1/cache/{key}"},
		Method:  []string{"PUT"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Cache.Get",
		Path:    []string{"/v1/cache/{key}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Cache.Delete",
		Path:    []string{"/v1/cache/{key}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Cache.List",
		Path:    []string{"/v1/cache"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&Cache{h}, opts...))
}

type cacheHandler struct {
	CacheHandler
}

func (h *cacheHandler) Put(ctx context.Context, in *CacheMessage, out *CacheMessage) error {
	return h.CacheHandler.Put(ctx, in, out)
}

func (h *cacheHandler) Set(ctx context.Context, in *CacheMessage, out *CacheMessage) error {
	return h.CacheHandler.Set(ctx, in, out)
}

func (h *cacheHandler) Get(ctx context.Context, in *CacheKeyRequest, out *CacheMessage) error {
	return h.CacheHandler.Get(ctx, in, out)
}

func (h *cacheHandler) Delete(ctx context.Context, in *CacheKeyRequest, out *CacheMessage) error {
	return h.CacheHandler.Delete(ctx, in, out)
}

func (h *cacheHandler) List(ctx context.Context, in *CacheListRequest, out *CacheListResponse) error {
	return h.CacheHandler.List(ctx, in, out)
}
