package main

import (
	"context"

	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

type Movie struct {
	Title string
}

type MovieControl struct {
	Movies []Movie
}

func (ctrl *MovieControl) AddMovie(ctx context.Context, req *proto.AddMovieRequest, rsp *proto.RequestResponse) error {
	for _, v := range ctrl.Movies {
		if v.Title == req.Title {
			rsp.Succeeded = false
			rsp.Cause = "movie already exists"
			return nil
		}
	}
	ctrl.Movies = append(ctrl.Movies, Movie{Title: req.Title})
	rsp.Succeeded = true
	return nil
}

func (ctrl *MovieControl) DeleteMovie(ctx context.Context, req *proto.DeleteMovieRequest,
	rsp *proto.RequestResponse) error {
	if req.Id < 0 || int(req.Id) >= len(ctrl.Movies) {
		rsp.Succeeded = false
		rsp.Cause = "index out of bounds"
		return nil
	}
	ctrl.Movies = append(ctrl.Movies[:req.Id], ctrl.Movies[req.Id+1:]...)
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
