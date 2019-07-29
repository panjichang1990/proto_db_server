package proto_server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
	"time"
)

type PfServer interface {
	Cps(ctx context.Context, in *CpsRequest, opts ...client.CallOption) (*CpsResponse, error)
	Game(ctx context.Context, in *GameRequest, opts ...client.CallOption) (*GameResponse, error)
	GoodsUpdate(ctx context.Context, in *GoodsAddRequest, opts ...client.CallOption) (*GoodsUpdateResponse, error)
	GoodsDel(ctx context.Context, in *GoodsDelRequest, opts ...client.CallOption) (*GoodsDelResponse, error)
	OwnCharge(ctx context.Context, in *OwnChargeRequest, opts ...client.CallOption) (*OwnChargeResponse, error)
	Zone(ctx context.Context, in *ZoneAddRequest, opts ...client.CallOption) (*ZoneAddResponse, error)
	ZoneDel(ctx context.Context, in *ZoneDelRequest, opts ...client.CallOption) (*ZoneDelResponse, error)
	Item(ctx context.Context, in *ItemAddRequest, opts ...client.CallOption) (*ItemAddResponse, error)
	Button(ctx context.Context, in *ButtonRequest, opts ...client.CallOption) (*ButtonResponse, error)
}

type PfServerHandler interface {
	Cps(context.Context, *CpsRequest, *CpsResponse) error
	Game( context.Context, *GameRequest,*GameResponse) error
	GoodsUpdate( context.Context, *GoodsAddRequest,*GoodsUpdateResponse) error
	GoodsDel( context.Context, *GoodsDelRequest,*GoodsDelResponse) error
	OwnCharge( context.Context, *OwnChargeRequest,*OwnChargeResponse) error
	Zone( context.Context, *ZoneAddRequest,*ZoneAddResponse) error
	ZoneDel( context.Context, *ZoneDelRequest,*ZoneDelResponse) error
	Item( context.Context, *ItemAddRequest,*ItemAddResponse) error
	Button( context.Context, *ButtonRequest,*ButtonResponse) error
}

type pfHandler struct {
	PfServerHandler
}

func (h *pfHandler) Cps(ctx context.Context, in *CpsRequest, out *CpsResponse) error {
	return h.PfServerHandler.Cps(ctx, in, out)
}

func (h *pfHandler) Game(ctx context.Context, in *GameRequest, out *GameResponse) error {
	return h.PfServerHandler.Game(ctx, in, out)
}

func (h *pfHandler) GoodsUpdate(ctx context.Context, in *GoodsAddRequest, out *GoodsUpdateResponse) error {
	return h.PfServerHandler.GoodsUpdate(ctx, in, out)
}
func (h *pfHandler) GoodsDel(ctx context.Context, in *GoodsDelRequest, out *GoodsDelResponse) error {
	return h.PfServerHandler.GoodsDel(ctx, in, out)
}

func (h *pfHandler) OwnCharge(ctx context.Context, in *OwnChargeRequest, out *OwnChargeResponse) error {
	return h.PfServerHandler.OwnCharge(ctx, in, out)
}

func (h *pfHandler) Zone(ctx context.Context, in *ZoneAddRequest, out *ZoneAddResponse) error {
	return h.PfServerHandler.Zone(ctx, in, out)
}

func (h *pfHandler) ZoneDel(ctx context.Context, in *ZoneDelRequest, out *ZoneDelResponse) error {
	return h.PfServerHandler.ZoneDel(ctx, in, out)
}

func (h *pfHandler) Item(ctx context.Context, in *ItemAddRequest, out *ItemAddResponse) error {
	return h.PfServerHandler.Item(ctx, in, out)
}

func (h *pfHandler) Button(ctx context.Context, in *ButtonRequest, out *ButtonResponse) error {
	return h.PfServerHandler.Button(ctx, in, out)
}

func RegisterPfAdminHandler(s server.Server, hdlr PfServerHandler, opts ...server.HandlerOption) {
	type pfAdmin interface {
		Cps(context.Context, *CpsRequest, *CpsResponse) error
		Game( context.Context, *GameRequest,*GameResponse) error
		GoodsUpdate( context.Context, *GoodsAddRequest,*GoodsUpdateResponse) error
		GoodsDel( context.Context, *GoodsDelRequest,*GoodsDelResponse) error
		OwnCharge( context.Context, *OwnChargeRequest,*OwnChargeResponse) error
		Zone( context.Context, *ZoneAddRequest,*ZoneAddResponse) error
		ZoneDel( context.Context, *ZoneDelRequest,*ZoneDelResponse) error
		Item( context.Context, *ItemAddRequest,*ItemAddResponse) error
		Button( context.Context, *ButtonRequest,*ButtonResponse) error
	}
	type PfAdmin struct {
		pfAdmin
	}
	h := &pfHandler{hdlr}
	err:=s.Handle(s.NewHandler(&PfAdmin{h}, opts...))
	if err!=nil{
		fmt.Println(err)
	}
}

type pfService struct {
	c           client.Client
	serviceName string
}


func NewPfAdminService(serviceName string) PfServer {
	reg:=consul.NewRegistry()
 	cc:=client.NewClient(client.Registry(reg),client.DialTimeout(1*time.Second))
	return &pfService{
		c:         cc,
		serviceName: serviceName,
	}
}


func (c *pfService)Cps(ctx context.Context, in *CpsRequest, opts ...client.CallOption)  (*CpsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PfAdmin.Cps", in)
	out := new(CpsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pfService)Game(ctx context.Context, in *GameRequest, opts ...client.CallOption)  (*GameResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PfAdmin.Game", in)
	out := new(GameResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pfService)GoodsUpdate(ctx context.Context, in *GoodsAddRequest, opts ...client.CallOption)  ( *GoodsUpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PfAdmin.GoodsUpdate", in)
	out := new(GoodsUpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pfService)GoodsDel(ctx context.Context, in *GoodsDelRequest, opts ...client.CallOption)  (*GoodsDelResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PfAdmin.GoodsDel", in)
	out := new(GoodsDelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pfService)OwnCharge(ctx context.Context, in *OwnChargeRequest, opts ...client.CallOption)  (*OwnChargeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PfAdmin.OwnCharge", in)
	out := new(OwnChargeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pfService)Item(ctx context.Context, in *ItemAddRequest, opts ...client.CallOption)  (*ItemAddResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PfAdmin.Item", in)
	out := new(ItemAddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *pfService)Zone(ctx context.Context, in *ZoneAddRequest, opts ...client.CallOption)  (*ZoneAddResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PfAdmin.Zone", in)
	out := new(ZoneAddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pfService)ZoneDel(ctx context.Context, in *ZoneDelRequest, opts ...client.CallOption)  (*ZoneDelResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PfAdmin.ZoneDel", in)
	out := new(ZoneDelResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pfService)Button(ctx context.Context, in *ButtonRequest, opts ...client.CallOption)  (*ButtonResponse, error) {
	req := c.c.NewRequest(c.serviceName, "PfAdmin.Button", in)
	out := new(ButtonResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
