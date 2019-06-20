package usercontrol

import (
	"fmt"

	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("userctrl"))
	service.Init()
	err := proto.RegisterUserControlHandler(service.Server(), &UserControl{NextID:0, Users: make(map[int]User)})

	if err != nil {
		fmt.Println(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
