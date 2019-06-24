// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: protoconfig.proto

package protoconfig

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Greeter service

type GreeterService interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error)
}

type greeterService struct {
	c    client.Client
	name string
}

func NewGreeterService(name string, c client.Client) GreeterService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "greeter"
	}
	return &greeterService{
		c:    c,
		name: name,
	}
}

func (c *greeterService) Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error) {
	req := c.c.NewRequest(c.name, "Greeter.Hello", in)
	out := new(HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterHandler interface {
	Hello(context.Context, *HelloRequest, *HelloResponse) error
}

func RegisterGreeterHandler(s server.Server, hdlr GreeterHandler, opts ...server.HandlerOption) error {
	type greeter interface {
		Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error
	}
	type Greeter struct {
		greeter
	}
	h := &greeterHandler{hdlr}
	return s.Handle(s.NewHandler(&Greeter{h}, opts...))
}

type greeterHandler struct {
	GreeterHandler
}

func (h *greeterHandler) Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error {
	return h.GreeterHandler.Hello(ctx, in, out)
}

// Client API for MovieControl service

type MovieControlService interface {
	AddMovie(ctx context.Context, in *AddMovieRequest, opts ...client.CallOption) (*RequestResponse, error)
	DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...client.CallOption) (*RequestResponse, error)
	GetMovie(ctx context.Context, in *GetMovieRequest, opts ...client.CallOption) (*GetMovieResponse, error)
}

type movieControlService struct {
	c    client.Client
	name string
}

func NewMovieControlService(name string, c client.Client) MovieControlService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "moviecontrol"
	}
	return &movieControlService{
		c:    c,
		name: name,
	}
}

func (c *movieControlService) AddMovie(ctx context.Context, in *AddMovieRequest, opts ...client.CallOption) (*RequestResponse, error) {
	req := c.c.NewRequest(c.name, "MovieControl.AddMovie", in)
	out := new(RequestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieControlService) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, opts ...client.CallOption) (*RequestResponse, error) {
	req := c.c.NewRequest(c.name, "MovieControl.DeleteMovie", in)
	out := new(RequestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieControlService) GetMovie(ctx context.Context, in *GetMovieRequest, opts ...client.CallOption) (*GetMovieResponse, error) {
	req := c.c.NewRequest(c.name, "MovieControl.GetMovie", in)
	out := new(GetMovieResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MovieControl service

type MovieControlHandler interface {
	AddMovie(context.Context, *AddMovieRequest, *RequestResponse) error
	DeleteMovie(context.Context, *DeleteMovieRequest, *RequestResponse) error
	GetMovie(context.Context, *GetMovieRequest, *GetMovieResponse) error
}

func RegisterMovieControlHandler(s server.Server, hdlr MovieControlHandler, opts ...server.HandlerOption) error {
	type movieControl interface {
		AddMovie(ctx context.Context, in *AddMovieRequest, out *RequestResponse) error
		DeleteMovie(ctx context.Context, in *DeleteMovieRequest, out *RequestResponse) error
		GetMovie(ctx context.Context, in *GetMovieRequest, out *GetMovieResponse) error
	}
	type MovieControl struct {
		movieControl
	}
	h := &movieControlHandler{hdlr}
	return s.Handle(s.NewHandler(&MovieControl{h}, opts...))
}

type movieControlHandler struct {
	MovieControlHandler
}

func (h *movieControlHandler) AddMovie(ctx context.Context, in *AddMovieRequest, out *RequestResponse) error {
	return h.MovieControlHandler.AddMovie(ctx, in, out)
}

func (h *movieControlHandler) DeleteMovie(ctx context.Context, in *DeleteMovieRequest, out *RequestResponse) error {
	return h.MovieControlHandler.DeleteMovie(ctx, in, out)
}

func (h *movieControlHandler) GetMovie(ctx context.Context, in *GetMovieRequest, out *GetMovieResponse) error {
	return h.MovieControlHandler.GetMovie(ctx, in, out)
}

// Client API for RoomControl service

type RoomControlService interface {
	AddRoom(ctx context.Context, in *AddRoomRequest, opts ...client.CallOption) (*RequestResponse, error)
	DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...client.CallOption) (*RequestResponse, error)
	GetRoom(ctx context.Context, in *GetRoomRequest, opts ...client.CallOption) (*GetRoomResponse, error)
	GetSingleRoom(ctx context.Context, in *GetSingleRoomRequest, opts ...client.CallOption) (*RoomData, error)
}

type roomControlService struct {
	c    client.Client
	name string
}

func NewRoomControlService(name string, c client.Client) RoomControlService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "roomcontrol"
	}
	return &roomControlService{
		c:    c,
		name: name,
	}
}

func (c *roomControlService) AddRoom(ctx context.Context, in *AddRoomRequest, opts ...client.CallOption) (*RequestResponse, error) {
	req := c.c.NewRequest(c.name, "RoomControl.AddRoom", in)
	out := new(RequestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomControlService) DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...client.CallOption) (*RequestResponse, error) {
	req := c.c.NewRequest(c.name, "RoomControl.DeleteRoom", in)
	out := new(RequestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomControlService) GetRoom(ctx context.Context, in *GetRoomRequest, opts ...client.CallOption) (*GetRoomResponse, error) {
	req := c.c.NewRequest(c.name, "RoomControl.GetRoom", in)
	out := new(GetRoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomControlService) GetSingleRoom(ctx context.Context, in *GetSingleRoomRequest, opts ...client.CallOption) (*RoomData, error) {
	req := c.c.NewRequest(c.name, "RoomControl.GetSingleRoom", in)
	out := new(RoomData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RoomControl service

type RoomControlHandler interface {
	AddRoom(context.Context, *AddRoomRequest, *RequestResponse) error
	DeleteRoom(context.Context, *DeleteRoomRequest, *RequestResponse) error
	GetRoom(context.Context, *GetRoomRequest, *GetRoomResponse) error
	GetSingleRoom(context.Context, *GetSingleRoomRequest, *RoomData) error
}

func RegisterRoomControlHandler(s server.Server, hdlr RoomControlHandler, opts ...server.HandlerOption) error {
	type roomControl interface {
		AddRoom(ctx context.Context, in *AddRoomRequest, out *RequestResponse) error
		DeleteRoom(ctx context.Context, in *DeleteRoomRequest, out *RequestResponse) error
		GetRoom(ctx context.Context, in *GetRoomRequest, out *GetRoomResponse) error
		GetSingleRoom(ctx context.Context, in *GetSingleRoomRequest, out *RoomData) error
	}
	type RoomControl struct {
		roomControl
	}
	h := &roomControlHandler{hdlr}
	return s.Handle(s.NewHandler(&RoomControl{h}, opts...))
}

type roomControlHandler struct {
	RoomControlHandler
}

func (h *roomControlHandler) AddRoom(ctx context.Context, in *AddRoomRequest, out *RequestResponse) error {
	return h.RoomControlHandler.AddRoom(ctx, in, out)
}

func (h *roomControlHandler) DeleteRoom(ctx context.Context, in *DeleteRoomRequest, out *RequestResponse) error {
	return h.RoomControlHandler.DeleteRoom(ctx, in, out)
}

func (h *roomControlHandler) GetRoom(ctx context.Context, in *GetRoomRequest, out *GetRoomResponse) error {
	return h.RoomControlHandler.GetRoom(ctx, in, out)
}

func (h *roomControlHandler) GetSingleRoom(ctx context.Context, in *GetSingleRoomRequest, out *RoomData) error {
	return h.RoomControlHandler.GetSingleRoom(ctx, in, out)
}

// Client API for ShowControl service

type ShowControlService interface {
	AddShow(ctx context.Context, in *AddShowRequest, opts ...client.CallOption) (*RequestResponse, error)
	DeleteShow(ctx context.Context, in *DeleteShowRequest, opts ...client.CallOption) (*RequestResponse, error)
	CheckSeat(ctx context.Context, in *AvailableSeatRequest, opts ...client.CallOption) (*RequestResponse, error)
	NotifyMovieDelete(ctx context.Context, in *MovieData, opts ...client.CallOption) (*RequestResponse, error)
	NotifyRoomDelete(ctx context.Context, in *RoomData, opts ...client.CallOption) (*RequestResponse, error)
}

type showControlService struct {
	c    client.Client
	name string
}

func NewShowControlService(name string, c client.Client) ShowControlService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "showcontrol"
	}
	return &showControlService{
		c:    c,
		name: name,
	}
}

func (c *showControlService) AddShow(ctx context.Context, in *AddShowRequest, opts ...client.CallOption) (*RequestResponse, error) {
	req := c.c.NewRequest(c.name, "ShowControl.AddShow", in)
	out := new(RequestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *showControlService) DeleteShow(ctx context.Context, in *DeleteShowRequest, opts ...client.CallOption) (*RequestResponse, error) {
	req := c.c.NewRequest(c.name, "ShowControl.DeleteShow", in)
	out := new(RequestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *showControlService) CheckSeat(ctx context.Context, in *AvailableSeatRequest, opts ...client.CallOption) (*RequestResponse, error) {
	req := c.c.NewRequest(c.name, "ShowControl.CheckSeat", in)
	out := new(RequestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *showControlService) NotifyMovieDelete(ctx context.Context, in *MovieData, opts ...client.CallOption) (*RequestResponse, error) {
	req := c.c.NewRequest(c.name, "ShowControl.NotifyMovieDelete", in)
	out := new(RequestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *showControlService) NotifyRoomDelete(ctx context.Context, in *RoomData, opts ...client.CallOption) (*RequestResponse, error) {
	req := c.c.NewRequest(c.name, "ShowControl.NotifyRoomDelete", in)
	out := new(RequestResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShowControl service

type ShowControlHandler interface {
	AddShow(context.Context, *AddShowRequest, *RequestResponse) error
	DeleteShow(context.Context, *DeleteShowRequest, *RequestResponse) error
	CheckSeat(context.Context, *AvailableSeatRequest, *RequestResponse) error
	NotifyMovieDelete(context.Context, *MovieData, *RequestResponse) error
	NotifyRoomDelete(context.Context, *RoomData, *RequestResponse) error
}

func RegisterShowControlHandler(s server.Server, hdlr ShowControlHandler, opts ...server.HandlerOption) error {
	type showControl interface {
		AddShow(ctx context.Context, in *AddShowRequest, out *RequestResponse) error
		DeleteShow(ctx context.Context, in *DeleteShowRequest, out *RequestResponse) error
		CheckSeat(ctx context.Context, in *AvailableSeatRequest, out *RequestResponse) error
		NotifyMovieDelete(ctx context.Context, in *MovieData, out *RequestResponse) error
		NotifyRoomDelete(ctx context.Context, in *RoomData, out *RequestResponse) error
	}
	type ShowControl struct {
		showControl
	}
	h := &showControlHandler{hdlr}
	return s.Handle(s.NewHandler(&ShowControl{h}, opts...))
}

type showControlHandler struct {
	ShowControlHandler
}

func (h *showControlHandler) AddShow(ctx context.Context, in *AddShowRequest, out *RequestResponse) error {
	return h.ShowControlHandler.AddShow(ctx, in, out)
}

func (h *showControlHandler) DeleteShow(ctx context.Context, in *DeleteShowRequest, out *RequestResponse) error {
	return h.ShowControlHandler.DeleteShow(ctx, in, out)
}

func (h *showControlHandler) CheckSeat(ctx context.Context, in *AvailableSeatRequest, out *RequestResponse) error {
	return h.ShowControlHandler.CheckSeat(ctx, in, out)
}

func (h *showControlHandler) NotifyMovieDelete(ctx context.Context, in *MovieData, out *RequestResponse) error {
	return h.ShowControlHandler.NotifyMovieDelete(ctx, in, out)
}

func (h *showControlHandler) NotifyRoomDelete(ctx context.Context, in *RoomData, out *RequestResponse) error {
	return h.ShowControlHandler.NotifyRoomDelete(ctx, in, out)
}
