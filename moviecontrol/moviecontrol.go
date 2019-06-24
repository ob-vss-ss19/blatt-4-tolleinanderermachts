package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

type Movie struct {
	Title string
}

type MovieControl struct {
	NextID  int
	Movies  map[int]Movie
	Service micro.Service
}

func (ctrl *MovieControl) AddMovie(ctx context.Context, req *proto.AddMovieRequest, rsp *proto.RequestResponse) error {
	println("add movie request")
	for _, v := range ctrl.Movies {
		if v.Title == req.Title {
			rsp.Succeeded = false
			rsp.Cause = "movie already exists"
			println("movie already exists: " + req.Title)
			return nil
		}
	}
	ctrl.Movies[ctrl.NextID] = Movie{Title: req.Title}
	ctrl.NextID++
	rsp.Succeeded = true
	println("added movie: " + req.Title)
	return nil
}

func (ctrl *MovieControl) DeleteMovie(ctx context.Context, req *proto.DeleteMovieRequest,
	rsp *proto.RequestResponse) error {
	println("delete movie request")
	_, ok := ctrl.Movies[int(req.Id)]
	if !ok {
		rsp.Succeeded = false
		rsp.Cause = "key does not exist"
		fmt.Printf("movie does not exist exists: %d\n", req.Id)
		return nil
	}
	delete(ctrl.Movies, int(req.Id))
	if ctrl.Service != nil {
		go ctrl.notifyMovieDelete(&proto.MovieData{Id: req.Id})
	}
	rsp.Succeeded = true
	println("deleted movie")
	return nil
}

func (ctrl *MovieControl) GetMovie(ctx context.Context, req *proto.GetMovieRequest, rsp *proto.GetMovieResponse) error {
	println("get movie request")
	data := make([]*proto.MovieData, 0)

	for k, v := range ctrl.Movies {
		data = append(data, &proto.MovieData{Id: int32(k), Title: v.Title})
	}
	rsp.Data = data
	println("returned every movie data")
	return nil
}

func (ctrl *MovieControl) notifyMovieDelete(data *proto.MovieData) {
	caller := proto.NewShowControlService("showctrl", ctrl.Service.Client())
	_, _ = caller.NotifyMovieDelete(context.TODO(), data)
}
