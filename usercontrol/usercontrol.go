package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

type User struct {
	Name         string
	Reservations []int32
}

type UserControl struct {
	NextID  int32
	Users   map[int32]User
	Service micro.Service
}

func (ctrl *UserControl) AddUser(ctx context.Context, req *proto.AddUserRequest, rsp *proto.RequestResponse) error {
	fmt.Println("add user request")
	ctrl.Users[ctrl.NextID] = User{Name: req.Name, Reservations: make([]int32, 0)}
	rsp.Id = ctrl.NextID
	ctrl.NextID++
	rsp.Succeeded = true
	fmt.Println("added user: " + req.Name)
	return nil
}

func (ctrl *UserControl) AddUserReservation(ctx context.Context, req *proto.AddUserReservationRequest, rsp *proto.RequestResponse) error {
	fmt.Println("add user reservation request")
	_, ok := ctrl.Users[int32(req.UserId)]
	if !ok {
		rsp.Succeeded = false
		rsp.Cause = "user id does not exist"
		return nil
	}
	user := ctrl.Users[int32(req.UserId)]
	user.Reservations = append(user.Reservations, req.ReservationId)
	ctrl.Users[int32(req.UserId)] = user
	rsp.Succeeded = true
	fmt.Printf("added reservation: %d to %d\n", req.ReservationId, req.UserId)
	return nil
}

func (ctrl *UserControl) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest, rsp *proto.RequestResponse) error {
	fmt.Println("delete user request")
	_, ok := ctrl.Users[int32(req.Id)]
	if !ok {
		rsp.Succeeded = false
		rsp.Cause = "user id does not exist"
		fmt.Printf("user does not exist: %d\n", req.Id)
		return nil
	}

	caller := proto.NewReservationControlService("resctrl", ctrl.Service.Client())
	showData, _ := caller.GetReservationsForUser(context.TODO(), &proto.GetReservationsForUserRequest{UserId: req.Id})
	if len(showData.Reservations) != 0 {
		rsp.Succeeded = false
		rsp.Cause = "User has active reservations"
		return nil
	}
	delete(ctrl.Users, int32(req.Id))
	rsp.Succeeded = true
	fmt.Println("deleted user")
	return nil
}

func (ctrl *UserControl) CheckUserReservation(ctx context.Context, req *proto.CheckUserReservationRequest, rsp *proto.RequestResponse) error {
	user, ok := ctrl.Users[int32(req.Id)]
	if !ok {
		rsp.Succeeded = false
		rsp.Cause = "user id could not been look up"
		fmt.Printf("user id does not exist: %d\n", req.Id)
		return nil
	}
	reservations := len(user.Reservations)
	if reservations > 0 {
		rsp.Succeeded = false
		rsp.Cause = "user has reservations and cannot be deleted"
		fmt.Printf("user has that many reservations: %d\n", reservations)
		return nil
	}
	rsp.Succeeded = true
	fmt.Println("user has no reservations")
	return nil
}
