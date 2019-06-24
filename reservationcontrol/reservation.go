package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	"github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

type ReservationControl struct {
	NextID       int32
	Reservations map[int32]protoconfig.Reservation
	Service      micro.Service
}

func (ctrl *ReservationControl) AddReservation(ctx context.Context, req *protoconfig.AddReservationRequest, rsp *protoconfig.RequestResponse) error {
	fmt.Println("add reservation")
	if len(req.Seats) < 1 {
		rsp.Succeeded = false
		rsp.Cause = "U have to select seats for a reservation"
		return nil
	}
	caller := protoconfig.NewShowControlService("showctrl", ctrl.Service.Client())
	for _, v := range req.Seats {
		showData, _ := caller.CheckSeat(context.TODO(), &protoconfig.AvailableSeatRequest{Id: req.ShowId, Row: v.Row, Seat: v.Column})
		b := showData.Succeeded
		if !b {
			rsp.Succeeded = false
			rsp.Cause = "Seat is already reservated: row = " + string(v.Row) + ", col = " + string(v.Column)
			return nil
		}
	}

	ctrl.Reservations[ctrl.NextID] = protoconfig.Reservation{Id: ctrl.NextID, ShowId: req.ShowId, Seats: req.Seats, UserId: req.UserId, Active: false}
	rsp.Succeeded = true
	rsp.Id = ctrl.NextID
	ctrl.NextID++
	return nil
}

func (ctrl *ReservationControl) ActivateReservation(ctx context.Context, req *protoconfig.ActivateReservationRequest, rsp *protoconfig.ReservationResponse) error {
	fmt.Println("activate reservation")
	reser, ok := ctrl.Reservations[req.ReservationId]
	if !ok {
		return fmt.Errorf("could not find the given reservationId: %d", req.ReservationId)
	}
	if reser.UserId != req.UserId {
		return fmt.Errorf("The userId does not match the reservations one")
	}
	reservation := ctrl.Reservations[req.ReservationId]
	reservation.Active = true
	ctrl.Reservations[req.ReservationId] = reservation
	rsp.Reservation = &reservation
	return nil
}

func (ctrl *ReservationControl) GetReservationsForUser(ctx context.Context, req *protoconfig.GetReservationsForUserRequest, rsp *protoconfig.GetReservationsForUserResponse) error {
	fmt.Println("get reservations for user")
	var reservations []*protoconfig.Reservation
	for _, res := range ctrl.Reservations {
		res := res
		if req.UserId == res.UserId {
			reservations = append(reservations, &res)
		}
	}
	rsp.Reservations = reservations
	return nil
}
func (ctrl *ReservationControl) RemoveReservation(ctx context.Context, req *protoconfig.RemoveReservationRequest, rsp *protoconfig.RequestResponse) error {
	fmt.Printf("remove reservation: %d\n", req.Id)
	_, found := ctrl.Reservations[req.Id]
	if !found {
		rsp.Succeeded = false
		rsp.Cause = "did not find given reservation id"
		return nil
	}
	delete(ctrl.Reservations, req.Id)
	rsp.Succeeded = true
	fmt.Printf("removed reservation: %d\n", req.Id)
	return nil
}
