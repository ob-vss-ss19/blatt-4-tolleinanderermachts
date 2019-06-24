package main

import (
	"context"
	"testing"

	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
	"github.com/stretchr/testify/assert"
)

func TestReservationWithoutSeats(t *testing.T) {

	ReservationControl := ReservationControl{NextID: 0, Reservations: make(map[int32]proto.Reservation)}

	response := proto.RequestResponse{}

	_ = ReservationControl.AddReservation(context.TODO(), &proto.AddReservationRequest{UserId:0, ShowId:0}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "U have to select seats for a reservation", response.Cause)
}

func TestReservationActivateWrongResId(t *testing.T) {

	ReservationControl := ReservationControl{NextID: 0, Reservations: make(map[int32]proto.Reservation)}

	response := proto.ReservationResponse{}

	err := ReservationControl.ActivateReservation(context.TODO(), &proto.ActivateReservationRequest{ReservationId:5}, &response)

	assert.Equal(t, "could not find the given reservationId: 5", err.Error())
}

func TestReservationActivateWrongUserId(t *testing.T) {
	ReservationControl := ReservationControl{NextID: 0, Reservations: make(map[int32]proto.Reservation)}
	ReservationControl.Reservations[1] = proto.Reservation{Id: 1, UserId: 1, ShowId: 1,
		Seats: []*proto.Seat{{Row: 1, Column: 2}, {Row: 1, Column: 3}}}

	response := proto.ReservationResponse{}

	err := ReservationControl.ActivateReservation(context.TODO(), &proto.ActivateReservationRequest{ReservationId:1, UserId:4}, &response)

	assert.Equal(t, "The userId does not match the reservations one", err.Error())
}

func TestReservationActivate(t *testing.T) {
	ReservationControl := ReservationControl{NextID: 0, Reservations: make(map[int32]proto.Reservation, 1)}
	reservation := proto.Reservation{Id: 0, UserId: 1, ShowId: 1,
		Seats: []*proto.Seat{{Row: 1, Column: 2}, {Row: 1, Column: 3}}}
	ReservationControl.Reservations[0] = reservation

	response := proto.ReservationResponse{}

	_ = ReservationControl.ActivateReservation(context.TODO(), &proto.ActivateReservationRequest{ReservationId:0, UserId:1}, &response)

	assert.Equal(t, response.Reservation.UserId, reservation.UserId)
	assert.Equal(t, response.Reservation.ShowId, reservation.ShowId)
	assert.Equal(t, response.Reservation.Id, reservation.Id)
	assert.True(t, response.Reservation.Active)
	for k, v := range reservation.Seats {
		assert.Equal(t, v.Column, response.Reservation.Seats[k].Column)
		assert.Equal(t, v.Row, response.Reservation.Seats[k].Row)
	}
}

func TestReservationControl_GetReservationsForUser(t *testing.T) {
	ReservationControl := ReservationControl{NextID: 0, Reservations: make(map[int32]proto.Reservation, 1)}
	reservation := proto.Reservation{Id: 0, UserId: 1, ShowId: 1,
		Seats: []*proto.Seat{{Row: 1, Column: 2}, {Row: 1, Column: 3}}}
	ReservationControl.Reservations[0] = reservation

	response := proto.GetReservationsForUserResponse{}

	_ = ReservationControl.GetReservationsForUser(context.TODO(), &proto.GetReservationsForUserRequest{UserId:1}, &response)

	assert.Equal(t, response.Reservations[0].UserId, reservation.UserId)
	assert.Equal(t, response.Reservations[0].ShowId, reservation.ShowId)
	assert.Equal(t, response.Reservations[0].Id, reservation.Id)
	assert.False(t, response.Reservations[0].Active)
	for k, v := range reservation.Seats {
		assert.Equal(t, v.Column, response.Reservations[0].Seats[k].Column)
		assert.Equal(t, v.Row, response.Reservations[0].Seats[k].Row)
	}
}

func TestReservationControl_RemoveReservationWrongId(t *testing.T) {

	ReservationControl := ReservationControl{NextID: 0, Reservations: make(map[int32]proto.Reservation, 0)}

	response := proto.RequestResponse{}

	_ = ReservationControl.RemoveReservation(context.TODO(), &proto.RemoveReservationRequest{Id:4}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "did not find given reservation id", response.Cause)
}

func TestReservationControl_RemoveReservation(t *testing.T) {
	ReservationControl := ReservationControl{NextID: 0, Reservations: make(map[int32]proto.Reservation, 1)}
	reservation := proto.Reservation{Id: 0, UserId: 1, ShowId: 1,
		Seats: []*proto.Seat{{Row: 1, Column: 2}, {Row: 1, Column: 3}}}
	ReservationControl.Reservations[0] = reservation

	response := proto.RequestResponse{}

	_ = ReservationControl.RemoveReservation(context.TODO(), &proto.RemoveReservationRequest{Id:0}, &response)

	assert.True(t, response.Succeeded, response.Cause)
}
