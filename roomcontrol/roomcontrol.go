package main

import (
	"context"

	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

type CinemaRoom struct {
	Name        string
	Rows        int
	SeatsPerRow int
}

type RoomControl struct {
	NextID  int
	Rooms   map[int]CinemaRoom
	Service micro.Service
}

func (ctrl *RoomControl) AddRoom(ctx context.Context, req *proto.AddRoomRequest, rsp *proto.RequestResponse) error {
	println("add room request")
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
	println("added room: " + req.Name)
	return nil
}

func (ctrl *RoomControl) DeleteRoom(ctx context.Context, req *proto.DeleteRoomRequest,
	rsp *proto.RequestResponse) error {
	println("delete room request")
	_, ok := ctrl.Rooms[int(req.Id)]
	if !ok {
		rsp.Succeeded = false
		rsp.Cause = "key does not exist"
		return nil
	}
	delete(ctrl.Rooms, int(req.Id))
	if ctrl.Service != nil {
		go ctrl.notifyRoomDelete(&proto.RoomData{Id: req.Id})
	}
	rsp.Succeeded = true
	println("deleted room: " + string(req.Id))
	return nil
}

func (ctrl *RoomControl) GetRoom(ctx context.Context, req *proto.GetRoomRequest, rsp *proto.GetRoomResponse) error {
	println("get room request")
	data := make([]*proto.RoomData, 0)

	for k, v := range ctrl.Rooms {
		data = append(data, &proto.RoomData{Id: int32(k), Name: v.Name, Rows: int32(v.Rows),
			SeatsPerRow: int32(v.SeatsPerRow)})
	}
	rsp.Data = data
	println("returned all room datas")
	return nil
}

func (ctrl *RoomControl) GetSingleRoom(ctx context.Context,
	req *proto.GetSingleRoomRequest, rsp *proto.RoomData) error {
	println("get single room request")
	data, ok := ctrl.Rooms[int(req.Id)]
	if ok {
		rsp.Id = req.Id
		rsp.Rows = int32(data.Rows)
		rsp.SeatsPerRow = int32(data.SeatsPerRow)
		rsp.Name = data.Name
	} else {
		rsp.Id = -1
	}
	println("returned specific room data: " + string(req.Id))
	return nil
}

func (ctrl *RoomControl) notifyRoomDelete(data *proto.RoomData) {
	caller := proto.NewShowControlService("showctrl", ctrl.Service.Client())
	_, _ = caller.NotifyRoomDelete(context.TODO(), data)
}
