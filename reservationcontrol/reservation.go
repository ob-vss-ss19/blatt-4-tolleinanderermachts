package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
	"hash"
)

type ReservationControl struct {
	NextID       int32
	Reservations map[int32]protoconfig.Reservation
}

func (ctrl *ReservationControl) AddReservation(ctx context.Context, req *protoconfig.AddReservationRequest, rsp *protoconfig.RequestResponse) error {
	fmt.Println("add reservation")
	if len(req.Seats) < 1 {
		rsp.Succeeded = false
		rsp.Cause = "U have to select seats for a reservation"
		return nil
	}
	// todo check seat validity or are they free?

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
	return nil
}

func (ctrl *ReservationControl) GetReservationsForUser(ctx context.Context, req *protoconfig.GetReservationsForUserRequest, rsp *protoconfig.GetReservationsForUserResponse) error {
	fmt.Println("get reservations for user")

}
func (ctrl *ReservationControl) RemoveReservation(ctx context.Context, req *protoconfig.RemoveReservationRequest, rsp *protoconfig.RequestResponse) error {

}

/*
type entry struct {
	Row  int
	Seat int
}
type Reservation struct {
	User    int
	Show    int
	entries []entry
}
*/
