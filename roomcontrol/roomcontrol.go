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
	NextID int
	Rooms  map[int]CinemaRoom
}

func (ctrl *RoomControl) AddRoom(ctx context.Context, req *proto.AddRoomRequest, rsp *proto.RequestResponse) error {
	for _, v := range ctrl.Rooms {
		if v.Name == req.Name {
			rsp.Succeeded = false
			rsp.Cause = "room already exists"
			return nil
		}
	}
	ctrl.Rooms[ctrl.NextID] = CinemaRoom{Name: req.Name, Rows: int(req.Rows), SeatsPerRow: int(req.SeatsPerRow)}
	ctrl.NextID++
	rsp.Succeeded = true
	return nil
}

func (ctrl *RoomControl) DeleteRoom(ctx context.Context, req *proto.DeleteRoomRequest,
	rsp *proto.RequestResponse) error {
	_, ok := ctrl.Rooms[int(req.Id)]
	if !ok {
		rsp.Succeeded = false
		rsp.Cause = "key does not exist"
		return nil
	}
	delete(ctrl.Rooms, int(req.Id))
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
