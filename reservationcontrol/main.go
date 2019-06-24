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
	err := proto.RegisterReservationControlHandler(service.Server(), &ReservationControl{NextID: 0, Reservations: make(map[int32]proto.Reservation), Service: service})

	if err != nil {
		fmt.Println(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
