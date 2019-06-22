package usercontrol

import (
	"context"
	"fmt"

	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

type User struct {
	Name         string
	Reservations []int
}

type UserControl struct {
	NextID int
	Users  map[int]User
}

func (ctrl *UserControl) AddUser(ctx context.Context, req *proto.AddUserRequest, rsp *proto.RequestResponse) error {
	fmt.Println("add user request")
	ctrl.Users[ctrl.NextID] = User{Name: req.Name, Reservations: make([]int, 0)}
	ctrl.NextID++
	rsp.Succeeded = true
	fmt.Println("added user: " + req.Name)
	return nil
}

func (ctrl *UserControl) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest, rsp *proto.RequestResponse) error {
	fmt.Println("delete user request")
	_, ok := ctrl.Users[int(req.Id)]
	if !ok {
		rsp.Succeeded = false
		rsp.Cause = "user id does not exist"
		fmt.Println("user does not exist: " + string(req.Id))
		return nil
	}
	// todo check user reservation
	delete(ctrl.Users, int(req.Id))
	rsp.Succeeded = true
	fmt.Println("deleted user")
	return nil
}

func (ctrl *UserControl) CheckUserReservation(ctx context.Context, req *proto.CheckUserReservationRequest, rsp *proto.RequestResponse) error {
	user, ok := ctrl.Users[int(req.Id)]
	if !ok {
		rsp.Succeeded = false
		rsp.Cause = "user id could not been look up"
		fmt.Println("user id does not exist: " + string(req.Id))
		return nil
	}
	reservations := len(user.Reservations)
	if reservations > 0 {
		rsp.Succeeded = false
		rsp.Cause = "user has reservations and cannot be deleted"
		fmt.Println("user has that many reservations: " + string(reservations))
		return nil
	}
	rsp.Succeeded = true
	fmt.Println("user has no reservations")
	return nil
}
