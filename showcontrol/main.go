package main

import (
	"fmt"

	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("showctrl"))
	service.Init()
	err := proto.RegisterShowControlHandler(service.Server(),
		&ShowControl{NextID: 0, Shows: make(map[int]Show), Service: service})

	if err != nil {
		fmt.Println(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
