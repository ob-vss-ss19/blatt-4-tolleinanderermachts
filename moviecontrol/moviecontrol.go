package main

import (
	"context"

	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

type Movie struct {
	Title string
}

type MovieControl struct {
	NextId int
	Movies map[int]Movie
}

func (ctrl *MovieControl) AddMovie(ctx context.Context, req *proto.AddMovieRequest, rsp *proto.RequestResponse) error {
	for _, v := range ctrl.Movies {
		if v.Title == req.Title {
			rsp.Succeeded = false
			rsp.Cause = "movie already exists"
			return nil
		}
	}
	ctrl.Movies[ctrl.NextId] = Movie{Title: req.Title}
	ctrl.NextId++
	rsp.Succeeded = true
	return nil
}

func (ctrl *MovieControl) DeleteMovie(ctx context.Context, req *proto.DeleteMovieRequest,
	rsp *proto.RequestResponse) error {
	_, ok := ctrl.Movies[int(req.Id)]
	if !ok {
		rsp.Succeeded = false
		rsp.Cause = "key does not exist"
		return nil
	}
	delete(ctrl.Movies, int(req.Id))
	rsp.Succeeded = true
	return nil
}

func (ctrl *MovieControl) GetMovie(ctx context.Context, req *proto.GetMovieRequest, rsp *proto.GetMovieResponse) error {
	data := make([]*proto.MovieData, 0)

	for k, v := range ctrl.Movies {
		data = append(data, &proto.MovieData{Id: int32(k), Title: v.Title})
	}
	rsp.Data = data
	return nil
}
