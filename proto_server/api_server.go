package proto_server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
	"time"
)

type ApiServer interface {
	Reload(ctx context.Context, in *ReloadReq, opts ...client.CallOption) (*ReloadRep, error)
}

type ApiServerHandler interface {
	Reload(context.Context,*ReloadReq,*ReloadRep)error
}

type apiHandler struct {
	ApiServerHandler
}

func (h *apiHandler) Reload(ctx context.Context, in *ReloadReq, out *ReloadRep) error {
	return h.ApiServerHandler.Reload(ctx, in, out)
}

func RegisterApiHandler(s server.Server , hdlr ApiServerHandler, opts ...server.HandlerOption)  {
	type Api struct {
		ApiServerHandler
	}
	h := &apiHandler{hdlr}
	err:=s.Handle(s.NewHandler(&Api{h}, opts...))
	if err!=nil{
		fmt.Println(err)
	}
}

type apiService struct {
	c           client.Client
	serviceName string
}

func NewApiService(serviceName string) ApiServer {
	reg:=consul.NewRegistry()
	cc:=client.NewClient(client.Registry(reg),client.DialTimeout(1*time.Second))
	return &apiService{
		c:         cc,
		serviceName: serviceName,
	}
}

func (c *apiService)Reload(ctx context.Context, in *ReloadReq, opts ...client.CallOption)  (*ReloadRep, error) {
	req := c.c.NewRequest(c.serviceName, "Api.Reload", in)
	out := new(ReloadRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}