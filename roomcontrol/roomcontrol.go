package main

import (
	"context"

	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

type CinemaRoom struct {
	Name        string
	Rows        int
	SeatsPerRow int
}

type RoomControl struct {
	Rooms []CinemaRoom
}

func (ctrl *RoomControl) AddRoom(ctx context.Context, req *proto.AddRoomRequest, rsp *proto.RequestResponse) error {
	for _, v := range ctrl.Rooms {
		if v.Name == req.Name {
			rsp.Succeeded = false
			rsp.Cause = "room already exists"
			return nil
		}
	}
	ctrl.Rooms = append(ctrl.Rooms, CinemaRoom{Name: req.Name, Rows: int(req.Rows), SeatsPerRow: int(req.SeatsPerRow)})
	rsp.Succeeded = true
	return nil
}

func (ctrl *RoomControl) DeleteRoom(ctx context.Context, req *proto.DeleteRoomRequest,
	rsp *proto.RequestResponse) error {
	if req.Id < 0 || int(req.Id) >= len(ctrl.Rooms) {
		rsp.Succeeded = false
		rsp.Cause = "index out of bounds"
		return nil
	}
	ctrl.Rooms = append(ctrl.Rooms[:req.Id], ctrl.Rooms[req.Id+1:]...)
	rsp.Succeeded = true
	return nil
}

func (ctrl *RoomControl) GetRoom(ctx context.Context, req *proto.GetRoomRequest, rsp *proto.GetRoomResponse) error {
	data := make([]*proto.RoomData, 0)

	for k, v := range ctrl.Rooms {
		data = append(data, &proto.RoomData{Id: int32(k), Name: v.Name, Rows: int32(v.Rows),
			SeatsPerRow: int32(v.SeatsPerRow)})
	}
	rsp.Data = data
	return nil
}
