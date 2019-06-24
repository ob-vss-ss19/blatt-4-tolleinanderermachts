package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

func main() {
	service := micro.NewService(micro.Name("movieClient"))
	service.Init()
	movieClient := proto.NewMovieControlService("moviectrl", service.Client())

	rsp, err := movieClient.AddMovie(context.TODO(), &proto.AddMovieRequest{Title: "Movie 1"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Succeeded)

	roomClient := proto.NewRoomControlService("roomctrl", service.Client())

	rsp, err = roomClient.AddRoom(context.TODO(), &proto.AddRoomRequest{Name: "Kino 1", Rows: 2, SeatsPerRow: 10})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Succeeded)

	showClient := proto.NewShowControlService("showctrl", service.Client())
	rsp, err = showClient.AddShow(context.TODO(), &proto.AddShowRequest{MovieId: 0, RoomId: 0})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Succeeded)
	rsp, err = showClient.CheckSeat(context.TODO(), &proto.AvailableSeatRequest{Id: 0, Row: 1, Seat: 5})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Succeeded)
	rsp, err = showClient.CheckSeat(context.TODO(),
		&proto.AvailableSeatRequest{Id: 0, Row: 1, Seat: 5, Write: true, Value: true})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Succeeded)
	rsp, err = showClient.CheckSeat(context.TODO(), &proto.AvailableSeatRequest{Id: 0, Row: 1, Seat: 5})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp.Succeeded)
}
