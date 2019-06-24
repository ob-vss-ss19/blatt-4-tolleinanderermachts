package main

import (
	"context"
	"testing"

	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
	"github.com/stretchr/testify/assert"
)

func TestUserAdd(t *testing.T) {

	UserControl := UserControl{NextID: 0, Users: make(map[int32]User)}

	response := proto.RequestResponse{}

	_ = UserControl.AddUser(context.TODO(), &proto.AddUserRequest{Name: "Albert"}, &response)

	assert.True(t, response.Succeeded, response.Cause)
}

func TestDeleteUser(t *testing.T) {

	UserControl := UserControl{NextID: 0, Users: make(map[int32]User)}

	response := proto.RequestResponse{}

	_ = UserControl.DeleteUser(context.TODO(), &proto.DeleteUserRequest{Id: 3}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "user id does not exist", response.Cause)
}

func TestUserCheckUserReservationEmptyReservations(t *testing.T) {

	UserControl := UserControl{NextID: 0, Users: make(map[int32]User)}

	response := proto.RequestResponse{}

	_ = UserControl.AddUser(context.TODO(), &proto.AddUserRequest{Name: "Albert"}, &response)

	_ = UserControl.CheckUserReservation(context.TODO(), &proto.CheckUserReservationRequest{Id: response.Id}, &response)

	assert.True(t, response.Succeeded, response.Cause)
}

func TestUserCheckUserReservationNotAUser(t *testing.T) {

	UserControl := UserControl{NextID: 0, Users: make(map[int32]User)}

	response := proto.RequestResponse{}

	_ = UserControl.CheckUserReservation(context.TODO(), &proto.CheckUserReservationRequest{Id: 0}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "user id could not been look up", response.Cause)
}

func TestUserCheckUserReservationWithReservations(t *testing.T) {

	UserControl := UserControl{NextID: 0, Users: make(map[int32]User)}

	response := proto.RequestResponse{}

	_ = UserControl.AddUser(context.TODO(), &proto.AddUserRequest{Name: "Albert"}, &response)
	seat := proto.Seat{Column: 1, Row: 1}
	reservation := proto.Reservation{Id: 0, UserId: response.Id, Seats: []*proto.Seat{&seat}, ShowId: 0, Active: true}
	_ = UserControl.AddUserReservation(context.TODO(), &proto.AddUserReservationRequest{UserId: response.Id, ReservationId: reservation.Id}, &response)

	_ = UserControl.CheckUserReservation(context.TODO(), &proto.CheckUserReservationRequest{Id: response.Id}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "user has reservations and cannot be deleted", response.Cause)
}

func TestUserAddReservation(t *testing.T) {
	UserControl := UserControl{NextID: 0, Users: make(map[int32]User)}

	response := proto.RequestResponse{}

	_ = UserControl.AddUserReservation(context.TODO(), &proto.AddUserReservationRequest{ReservationId: 2, UserId: 3}, &response)
}
