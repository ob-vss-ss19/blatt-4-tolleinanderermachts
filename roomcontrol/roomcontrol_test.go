package main

import (
	"context"
	"testing"

	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
	"github.com/stretchr/testify/assert"
)

func TestRoomAdd(t *testing.T) {

	roomControl := RoomControl{NextID: 0, Rooms: make(map[int]CinemaRoom)}

	response := proto.RequestResponse{}

	_ = roomControl.AddRoom(context.TODO(), &proto.AddRoomRequest{Name: "", Rows: 0, SeatsPerRow: 0}, &response)

	assert.True(t, response.Succeeded, response.Cause)
}

func TestRoomAddWithSavedCheck(t *testing.T) {

	roomControl := RoomControl{NextID: 0, Rooms: make(map[int]CinemaRoom)}

	response := proto.RequestResponse{}

	_ = roomControl.AddRoom(context.TODO(),
		&proto.AddRoomRequest{Name: "TestRoomAddWithSavedCheck", Rows: 5, SeatsPerRow: 7}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	RoomResponse := proto.GetRoomResponse{}

	_ = roomControl.GetRoom(context.TODO(), &proto.GetRoomRequest{}, &RoomResponse)

	assert.EqualValues(t, 1, len(RoomResponse.Data), "Room number should be 1!")
	assert.Equal(t, "TestRoomAddWithSavedCheck", RoomResponse.Data[0].Name, "Room title mismatch!")
	assert.Equal(t, int32(5), RoomResponse.Data[0].Rows, "Room rows mismatch!")
	assert.Equal(t, int32(7), RoomResponse.Data[0].SeatsPerRow, "Room seats mismatch!")
	assert.EqualValues(t, 0, RoomResponse.Data[0].Id, "id mismatch")
}

func TestRoomAddDoubleRoom(t *testing.T) {

	roomControl := RoomControl{NextID: 0, Rooms: make(map[int]CinemaRoom)}

	response := proto.RequestResponse{}

	_ = roomControl.AddRoom(context.TODO(), &proto.AddRoomRequest{Name: "TestRoomAddDoubleRoom"}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = roomControl.AddRoom(context.TODO(), &proto.AddRoomRequest{Name: "TestRoomAddDoubleRoom"}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "room already exists", response.Cause)
}

func TestRoomDelete(t *testing.T) {

	roomControl := RoomControl{NextID: 0, Rooms: make(map[int]CinemaRoom)}

	response := proto.RequestResponse{}

	_ = roomControl.AddRoom(context.TODO(), &proto.AddRoomRequest{Name: "TestRoomDelete"}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = roomControl.DeleteRoom(context.TODO(), &proto.DeleteRoomRequest{Id: 0}, &response)

	assert.True(t, response.Succeeded, response.Cause)
}

func TestRoomDeleteWithSavedCheck(t *testing.T) {

	roomControl := RoomControl{NextID: 0, Rooms: make(map[int]CinemaRoom)}

	response := proto.RequestResponse{}

	_ = roomControl.AddRoom(context.TODO(),
		&proto.AddRoomRequest{Name: "TestRoomDeleteWithSavedCheck", Rows: 5, SeatsPerRow: 7}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	RoomResponse := proto.GetRoomResponse{}

	_ = roomControl.GetRoom(context.TODO(), &proto.GetRoomRequest{}, &RoomResponse)

	assert.EqualValues(t, 1, len(RoomResponse.Data), "Room number should be 1!")
	assert.Equal(t, "TestRoomDeleteWithSavedCheck", RoomResponse.Data[0].Name, "Room title mismatch!")
	assert.Equal(t, int32(5), RoomResponse.Data[0].Rows, "Room rows mismatch!")
	assert.Equal(t, int32(7), RoomResponse.Data[0].SeatsPerRow, "Room seats mismatch!")
	assert.EqualValues(t, 0, RoomResponse.Data[0].Id, "id mismatch")

	_ = roomControl.DeleteRoom(context.TODO(), &proto.DeleteRoomRequest{Id: 0}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = roomControl.GetRoom(context.TODO(), &proto.GetRoomRequest{}, &RoomResponse)

	assert.EqualValues(t, 0, len(RoomResponse.Data), "Room number should be 0!")
}

func TestRoomDeleteEmpty(t *testing.T) {

	roomControl := RoomControl{NextID: 0, Rooms: make(map[int]CinemaRoom)}

	response := proto.RequestResponse{}

	_ = roomControl.DeleteRoom(context.TODO(), &proto.DeleteRoomRequest{Id: 0}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "key does not exist", response.Cause)
}

func TestRoomDeleteNegative(t *testing.T) {

	roomControl := RoomControl{NextID: 0, Rooms: make(map[int]CinemaRoom)}

	response := proto.RequestResponse{}

	_ = roomControl.DeleteRoom(context.TODO(), &proto.DeleteRoomRequest{Id: -1}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "key does not exist", response.Cause)
}

func TestRoomDeleteNotFound(t *testing.T) {

	roomControl := RoomControl{NextID: 0, Rooms: make(map[int]CinemaRoom)}

	response := proto.RequestResponse{}

	_ = roomControl.AddRoom(context.TODO(), &proto.AddRoomRequest{Name: "TestRoomDeleteNotFound"}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = roomControl.DeleteRoom(context.TODO(), &proto.DeleteRoomRequest{Id: 2}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "key does not exist", response.Cause)
}

func TestGetSingleRoom(t *testing.T) {

	roomControl := RoomControl{NextID: 0, Rooms: make(map[int]CinemaRoom)}

	response := proto.RequestResponse{}

	_ = roomControl.AddRoom(context.TODO(),
		&proto.AddRoomRequest{Name: "The Room of No Return", Rows: 2, SeatsPerRow: 5}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	data := proto.RoomData{}
	_ = roomControl.GetSingleRoom(context.TODO(), &proto.GetSingleRoomRequest{Id: 0}, &data)

	assert.Equal(t, data.Id, int32(0))
	assert.Equal(t, data.Name, "The Room of No Return")
	assert.Equal(t, data.Rows, int32(2))
	assert.Equal(t, data.SeatsPerRow, int32(5))
}

func TestGetSingleWrongRoom(t *testing.T) {

	roomControl := RoomControl{NextID: 0, Rooms: make(map[int]CinemaRoom)}

	response := proto.RequestResponse{}

	_ = roomControl.AddRoom(context.TODO(),
		&proto.AddRoomRequest{Name: "The Room of No Return", Rows: 2, SeatsPerRow: 5}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	data := proto.RoomData{}
	_ = roomControl.GetSingleRoom(context.TODO(), &proto.GetSingleRoomRequest{Id: 1}, &data)

	assert.Equal(t, data.Id, int32(-1))
}
