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
	var a [][]bool
	if ctrl.Service != nil {
		caller := proto.NewRoomControlService("roomctrl", ctrl.Service.Client())
		roomData, _ := caller.GetSingleRoom(context.TODO(), &proto.GetSingleRoomRequest{Id: req.RoomId})

		println("allocating space...")
		a = make([][]bool, roomData.Rows)
		for i := range a {
			a[i] = make([]bool, roomData.SeatsPerRow)
		}
	} else {
		//for testing
		println("TEST ENVIRONMENT: allocating space...")
		a = make([][]bool, 5)
		for i := range a {
			a[i] = make([]bool, 5)
		}
	}

	ctrl.Shows[ctrl.NextID] = Show{Movie: int(req.MovieId), Room: int(req.RoomId), Seats: a}
	ctrl.NextID++
	rsp.Succeeded = true

	println("show successful added")
	return nil
}

func (ctrl *ShowControl) DeleteShow(ctx context.Context,
	req *proto.DeleteShowRequest, rsp *proto.RequestResponse) error {
	println("deleting show")
	_, ok := ctrl.Shows[int(req.Id)]
	if !ok {
		rsp.Succeeded = false
		rsp.Cause = "key does not exist"
		return nil
	}
	delete(ctrl.Shows, int(req.Id))
	go ctrl.notyfiyReservationcontrol([]int{int(req.Id)})
	fmt.Printf("show deleted: %d\n", req.Id)
	return nil
}

func (ctrl *ShowControl) CheckSeat(ctx context.Context,
	req *proto.AvailableSeatRequest, rsp *proto.RequestResponse) error {
	show, ok := ctrl.Shows[int(req.Id)]
	println("checking seats")
	if !ok {
		fmt.Printf("show not found: %d\n", req.Id)
		rsp.Succeeded = false
		rsp.Cause = "Show not found"
	} else {
		if !req.Write {
			fmt.Printf("checking seat %d:%d with: %t\n", req.Row, req.Seat, show.Seats[req.Row][req.Seat])
			if show.Seats[req.Row][req.Seat] {
				rsp.Succeeded = false
				rsp.Cause = "seat occupied"
			} else {
				rsp.Succeeded = true
			}

		} else {
			fmt.Printf("writing seat %d:%d with: %t\n", req.Row, req.Seat, req.Value)
			show.Seats[req.Row][req.Seat] = req.Value
			rsp.Succeeded = true
		}
		rsp.Id = req.Id
	}
	return nil
}

func (ctrl *ShowControl) NotifyMovieDelete(ctx context.Context,
	req *proto.MovieData, rsp *proto.RequestResponse) error {
	println("got movie delete notification")
	var delShows []int
	for k, v := range ctrl.Shows {
		if v.Movie == int(req.Id) {
			delShows = append(delShows, k)
			delete(ctrl.Shows, k)
		}
	}

	go ctrl.notyfiyReservationcontrol(delShows)

	rsp.Succeeded = true
	return nil
}

func (ctrl *ShowControl) NotifyRoomDelete(ctx context.Context,
	req *proto.RoomData, rsp *proto.RequestResponse) error {
	println("got room delete notification")
	var delShows []int
	for k, v := range ctrl.Shows {
		if v.Room == int(req.Id) {
			delShows = append(delShows, k)
			delete(ctrl.Shows, k)
		}
	}

	go ctrl.notyfiyReservationcontrol(delShows)

	rsp.Succeeded = true
	return nil
}

func (ctrl *ShowControl) notyfiyReservationcontrol(delShows []int) {
	caller := proto.NewReservationControlService("resctrl", ctrl.Service.Client())
	for _, v := range delShows {
		_, _ = caller.RemoveReservation(context.TODO(), &proto.RemoveReservationRequest{Id: int32(v)})
	}
}
