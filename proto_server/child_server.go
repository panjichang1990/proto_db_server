package proto_server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
	"time"
)

type ChildServer interface {
	Goods(ctx context.Context, in *GoodsRequest, opts ...client.CallOption) (*GoodsResponse, error)
	Zone(ctx context.Context, in *ZoneRequest, opts ...client.CallOption) (*ZoneResponse, error)
	Item(ctx context.Context, in *ItemRequest, opts ...client.CallOption) (*ItemResponse, error)
	User(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	ButtonAdd(ctx context.Context, in *ButtonAddRequest, opts ...client.CallOption) (*BoolResponse, error)
	ButtonDel(ctx context.Context, in *ButtonDelRequest, opts ...client.CallOption) (*BoolResponse, error)
	CpsAdd(ctx context.Context, in *CpsAddRequest, opts ...client.CallOption) ( *CpsAddResponse, error)
}

type ChildServerHandler interface {
	Goods(context.Context, *GoodsRequest, *GoodsResponse) error
	Zone( context.Context, *ZoneRequest,*ZoneResponse) error
	Item( context.Context, *ItemRequest,*ItemResponse) error
	User( context.Context, *UserRequest,*UserResponse) error
	ButtonAdd( context.Context, *ButtonAddRequest,*BoolResponse) error
	ButtonDel( context.Context, *ButtonDelRequest,*BoolResponse) error
	CpsAdd( context.Context, *CpsAddRequest,*CpsAddResponse) error
}

type childHandler struct {
	ChildServerHandler
}

func (h *childHandler) Goods(ctx context.Context, in *GoodsRequest, out *GoodsResponse) error {
	return h.ChildServerHandler.Goods(ctx, in, out)
}

func (h *childHandler) Zone(ctx context.Context, in *ZoneRequest, out *ZoneResponse) error {
	return h.ChildServerHandler.Zone(ctx, in, out)
}
func (h *childHandler) Item(ctx context.Context, in *ItemRequest, out *ItemResponse) error {
	return h.ChildServerHandler.Item(ctx, in, out)
}

func (h *childHandler) User(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.ChildServerHandler.User(ctx, in, out)
}

func (h *childHandler) ButtonAdd(ctx context.Context, in *ButtonAddRequest, out *BoolResponse) error {
	return h.ChildServerHandler.ButtonAdd(ctx, in, out)
}

func (h *childHandler) ButtonDel(ctx context.Context, in *ButtonDelRequest, out *BoolResponse) error {
	return h.ChildServerHandler.ButtonDel(ctx, in, out)
}

func (h *childHandler) CpsAdd(ctx context.Context, in *CpsAddRequest, out *CpsAddResponse) error {
	return h.ChildServerHandler.CpsAdd(ctx, in, out)
}

func RegisterChildHandler(s server.Server, hdlr ChildServerHandler, opts ...server.HandlerOption) {
	type childServer interface {
		Goods(context.Context, *GoodsRequest, *GoodsResponse) error
		Zone( context.Context, *ZoneRequest,*ZoneResponse) error
		Item( context.Context, *ItemRequest,*ItemResponse) error
		User( context.Context, *UserRequest,*UserResponse) error
		ButtonAdd( context.Context, *ButtonAddRequest,*BoolResponse) error
		ButtonDel( context.Context, *ButtonDelRequest,*BoolResponse) error
		CpsAdd( context.Context, *CpsAddRequest,*CpsAddResponse) error
	}
	type ChildServer struct {
		childServer
	}
	h := &childHandler{hdlr}
	err:=s.Handle(s.NewHandler(&ChildServer{h}, opts...))
	if err!=nil{
		fmt.Println(err)
	}
}

type childService struct {
	c           client.Client
	serviceName string
}

func NewChildService(serviceName string) ChildServer {
	reg:=consul.NewRegistry()
	cc:=client.NewClient(client.Registry(reg),client.DialTimeout(1*time.Second))
	return &childService{
		c:           cc,
		serviceName: serviceName,
	}
}

func (c *childService)Goods(ctx context.Context, in *GoodsRequest, opts ...client.CallOption)  ( *GoodsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ChildServer.Goods", in)
	out := new(GoodsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *childService)Zone(ctx context.Context, in *ZoneRequest, opts ...client.CallOption)  (   *ZoneResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ChildServer.Zone", in)
	out := new(ZoneResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *childService)Item(ctx context.Context, in *ItemRequest, opts ...client.CallOption)  (   *ItemResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ChildServer.Item", in)
	out := new(ItemResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *childService)User(ctx context.Context, in *UserRequest, opts ...client.CallOption)  (   *UserResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ChildServer.User", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *childService)ButtonAdd(ctx context.Context, in *ButtonAddRequest, opts ...client.CallOption)  (    *BoolResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ChildServer.ButtonAdd", in)
	out := new(BoolResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *childService)ButtonDel(ctx context.Context, in *ButtonDelRequest, opts ...client.CallOption)  (   *BoolResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ChildServer.ButtonDel", in)
	out := new(BoolResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *childService)CpsAdd(ctx context.Context, in *CpsAddRequest, opts ...client.CallOption)  (   *CpsAddResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ChildServer.CpsAdd", in)
	out := new(CpsAddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}