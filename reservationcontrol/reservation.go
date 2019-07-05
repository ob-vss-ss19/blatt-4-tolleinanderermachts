package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

type ReservationControl struct {
	NextID       int32
	Reservations map[int32]proto.Reservation
	Service      micro.Service
}

func (ctrl *ReservationControl) AddReservation(ctx context.Context,
	req *proto.AddReservationRequest, rsp *proto.RequestResponse) error {
	fmt.Println("add reservation")
	if len(req.Seats) < 1 {
		rsp.Succeeded = false
		rsp.Cause = "U have to select seats for a reservation"
		return nil
	}
	caller := proto.NewShowControlService("showctrl", ctrl.Service.Client())
	for _, v := range req.Seats {
		showData, _ := caller.CheckSeat(context.TODO(),
			&proto.AvailableSeatRequest{Id: req.ShowId, Row: v.Row, Seat: v.Column})
		b := showData.Succeeded
		if !b {
			rsp.Succeeded = false
			rsp.Cause = "Seat is already reservated: row = " + string(v.Row) + ", col = " + string(v.Column)
			return nil
		}
	}

	ctrl.Reservations[ctrl.NextID] = proto.Reservation{Id: ctrl.NextID, ShowId: req.ShowId,
		Seats: req.Seats, UserId: req.UserId, Active: false}
	rsp.Succeeded = true
	rsp.Id = ctrl.NextID
	ctrl.NextID++
	return nil
}

func (ctrl *ReservationControl) ActivateReservation(ctx context.Context,
	req *proto.ActivateReservationRequest, rsp *proto.ReservationResponse) error {
	fmt.Println("activate reservation")
	reser, ok := ctrl.Reservations[req.ReservationId]
	if !ok {
		return fmt.Errorf("could not find the given reservationId: %d", req.ReservationId)
	}
	if reser.UserId != req.UserId {
		return fmt.Errorf("the userId does not match the reservations one")
	}
	callerShow := proto.NewShowControlService("showctrl", ctrl.Service.Client())
	for _, v := range reser.Seats {
		showData, _ := callerShow.CheckSeat(context.TODO(),
			&proto.AvailableSeatRequest{Id: reser.ShowId, Row: v.Row, Seat: v.Column, Value: true, Write: true})
		b := showData.Succeeded
		if !b {
			fmt.Printf("Seat is already reservated: row = %d, col = %d\n", v.Row, v.Column)
			return fmt.Errorf("Seat is already reservated: row = %d, col = %d\n", v.Row, v.Column)
		}
	}

	callerUser := proto.NewUserControlService("userctrl", ctrl.Service.Client())
	result, _ := callerUser.AddUserReservation(context.TODO(), &proto.AddUserReservationRequest{UserId: req.UserId, ReservationId: req.ReservationId})
	if !result.Succeeded {
		fmt.Printf("user reservation notify did not work : %s\n", result.Cause)
	}

	reservation := ctrl.Reservations[req.ReservationId]
	reservation.Active = true
	ctrl.Reservations[req.ReservationId] = reservation
	rsp.Reservation = &reservation
	return nil
}

func (ctrl *ReservationControl) GetReservationsForUser(ctx context.Context,
	req *proto.GetReservationsForUserRequest, rsp *proto.GetReservationsForUserResponse) error {
	fmt.Println("get reservations for user")
	var reservations []*proto.Reservation
	for _, res := range ctrl.Reservations {
		res := res
		if req.UserId == res.UserId {
			reservations = append(reservations, &res)
		}
	}
	rsp.Reservations = reservations
	return nil
}
func (ctrl *ReservationControl) RemoveReservation(ctx context.Context,
	req *proto.RemoveReservationRequest, rsp *proto.RequestResponse) error {
	fmt.Printf("remove reservation: %d\n", req.ReserId)
	_, found := ctrl.Reservations[req.ReserId]
	if !found {
		rsp.Succeeded = false
		rsp.Cause = "did not find given reservation id"
		return nil
	}
	callerUser := proto.NewUserControlService("userctrl", ctrl.Service.Client())
	response, _ := callerUser.DeleteUserReservation(context.TODO(),
		&proto.DeleteUserReservationRequest{UserId: req.UserId, ReservationId: req.ReserId})
	if !response.Succeeded {
		fmt.Printf("could not delete reservation from user")
		rsp.Succeeded = false
		return nil
	}
	go ctrl.notifyShowReservationDelete(ctrl.Reservations[req.ReserId])
	delete(ctrl.Reservations, req.ReserId)
	rsp.Succeeded = true
	fmt.Printf("removed reservation: %d\n", req.ReserId)
	return nil
}

func (ctrl *ReservationControl) notifyShowReservationDelete(res proto.Reservation) {
	if ctrl.Service != nil {
		caller := proto.NewShowControlService("showctrl", ctrl.Service.Client())
		for _, v := range res.Seats {
			_, _ = caller.CheckSeat(context.TODO(),
				&proto.AvailableSeatRequest{Id: res.ShowId, Row: v.Row, Seat: v.Column, Write: true, Value: false})
		}
	}
}
