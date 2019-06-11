package main

import (
	"fmt"

	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("roomctrl"))
	service.Init()
	err := proto.RegisterRoomControlHandler(service.Server(), &RoomControl{NextID: 0, Rooms: make(map[int]CinemaRoom)})

	if err != nil {
		fmt.Println(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
