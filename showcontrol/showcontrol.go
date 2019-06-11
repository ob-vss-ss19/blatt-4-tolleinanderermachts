package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

type Show struct {
	Movie int
	Room  int
	Seats [][]bool
}

type ShowControl struct {
	NextID  int
	Shows   map[int]Show
	Service micro.Service
}

func (ctrl *ShowControl) AddShow(ctx context.Context, req *proto.AddShowRequest, rsp *proto.RequestResponse) error {
	println("adding show... calling roomcontrol for seat size")
	caller := proto.NewRoomControlService("roomctrl", ctrl.Service.Client())
	roomData, _ := caller.GetSingleRoom(context.TODO(), &proto.GetSingleRoomRequest{Id: req.RoomId})

	println("allocating space...")
	a := make([][]bool, roomData.Rows)
	for i := range a {
		a[i] = make([]bool, roomData.SeatsPerRow)
	}

	ctrl.Shows[ctrl.NextID] = Show{Movie: int(req.MovieId), Room: int(req.RoomId), Seats: a}
	ctrl.NextID++
	rsp.Succeeded = true

	println("show successful added")
	return nil
}

func (ctrl *ShowControl) DeleteShow(ctx context.Context,
	req *proto.DeleteShowRequest, rsp *proto.RequestResponse) error {
	delete(ctrl.Shows, int(req.Id))
	println("show deleted")
	return nil
}

func (ctrl *ShowControl) CheckSeat(ctx context.Context,
	req *proto.AvailableSeatRequest, rsp *proto.RequestResponse) error {
	show, ok := ctrl.Shows[int(req.Id)]
	println("checking seats")
	if !ok {
		fmt.Printf("show not found: %d", req.Id)
		rsp.Succeeded = false
		rsp.Cause = "Show not found"
	} else {
		if !req.Write {
			fmt.Printf("checking seat %d:%d with: %t", req.Row, req.Seat, show.Seats[req.Row][req.Seat])
			if show.Seats[req.Row][req.Seat] {
				rsp.Succeeded = false
			} else {
				rsp.Succeeded = true
			}

		} else {
			fmt.Printf("writing seat %d:%d with: %t", req.Row, req.Seat, req.Value)
			show.Seats[req.Row][req.Seat] = req.Value
			rsp.Succeeded = true
		}
		rsp.Id = req.Id
	}
	return nil
}
