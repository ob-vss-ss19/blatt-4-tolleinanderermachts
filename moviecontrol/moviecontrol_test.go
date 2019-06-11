package main

import (
	"context"
	"testing"

	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
	"github.com/stretchr/testify/assert"
)

func TestMovieAdd(t *testing.T) {

	movieControl := MovieControl{NextId: 0, Movies: make(map[int]Movie)}

	response := proto.RequestResponse{}

	_ = movieControl.AddMovie(context.TODO(), &proto.AddMovieRequest{Title: ""}, &response)

	assert.True(t, response.Succeeded, response.Cause)
}

func TestMovieAddWithSavedCheck(t *testing.T) {

	movieControl := MovieControl{NextId: 0, Movies: make(map[int]Movie)}

	response := proto.RequestResponse{}

	_ = movieControl.AddMovie(context.TODO(), &proto.AddMovieRequest{Title: "TestMovieAddWithSavedCheck"}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	movieResponse := proto.GetMovieResponse{}

	_ = movieControl.GetMovie(context.TODO(), &proto.GetMovieRequest{}, &movieResponse)

	assert.EqualValues(t, 1, len(movieResponse.Data), "Movie number should be 1!")
	assert.Equal(t, "TestMovieAddWithSavedCheck", movieResponse.Data[0].Title, "movie title mismatch!")
	assert.EqualValues(t, 0, movieResponse.Data[0].Id, "id mismatch")
}

func TestMovieAddDoubleMovie(t *testing.T) {

	movieControl := MovieControl{NextId: 0, Movies: make(map[int]Movie)}

	response := proto.RequestResponse{}

	_ = movieControl.AddMovie(context.TODO(), &proto.AddMovieRequest{Title: "TestMovieAddDoubleMovie"}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = movieControl.AddMovie(context.TODO(), &proto.AddMovieRequest{Title: "TestMovieAddDoubleMovie"}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "movie already exists", response.Cause)
}

func TestMovieDelete(t *testing.T) {

	movieControl := MovieControl{NextId: 0, Movies: make(map[int]Movie)}

	response := proto.RequestResponse{}

	_ = movieControl.AddMovie(context.TODO(), &proto.AddMovieRequest{Title: "TestMovieDelete"}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = movieControl.DeleteMovie(context.TODO(), &proto.DeleteMovieRequest{Id: 0}, &response)

	assert.True(t, response.Succeeded, response.Cause)
}

func TestMovieDeleteWithSavedCheck(t *testing.T) {

	movieControl := MovieControl{NextId: 0, Movies: make(map[int]Movie)}

	response := proto.RequestResponse{}

	_ = movieControl.AddMovie(context.TODO(), &proto.AddMovieRequest{Title: "TestMovieDeleteWithSavedCheck"}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	movieResponse := proto.GetMovieResponse{}

	_ = movieControl.GetMovie(context.TODO(), &proto.GetMovieRequest{}, &movieResponse)

	assert.EqualValues(t, 1, len(movieResponse.Data), "Movie number should be 1!")
	assert.Equal(t, "TestMovieDeleteWithSavedCheck", movieResponse.Data[0].Title, "movie title mismatch!")
	assert.EqualValues(t, 0, movieResponse.Data[0].Id, "id mismatch")

	_ = movieControl.DeleteMovie(context.TODO(), &proto.DeleteMovieRequest{Id: 0}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = movieControl.GetMovie(context.TODO(), &proto.GetMovieRequest{}, &movieResponse)

	assert.EqualValues(t, 0, len(movieResponse.Data), "Movie number should be 0!")
}

func TestMovieDeleteEmpty(t *testing.T) {

	movieControl := MovieControl{NextId: 0, Movies: make(map[int]Movie)}

	response := proto.RequestResponse{}

	_ = movieControl.DeleteMovie(context.TODO(), &proto.DeleteMovieRequest{Id: 0}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "key does not exist", response.Cause)
}

func TestMovieDeleteNegative(t *testing.T) {

	movieControl := MovieControl{NextId: 0, Movies: make(map[int]Movie)}

	response := proto.RequestResponse{}

	_ = movieControl.DeleteMovie(context.TODO(), &proto.DeleteMovieRequest{Id: -1}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "key does not exist", response.Cause)
}

func TestMovieDeleteNotFound(t *testing.T) {

	movieControl := MovieControl{NextId: 0, Movies: make(map[int]Movie)}

	response := proto.RequestResponse{}

	_ = movieControl.AddMovie(context.TODO(), &proto.AddMovieRequest{Title: "TestMovieDeleteNotFound"}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = movieControl.DeleteMovie(context.TODO(), &proto.DeleteMovieRequest{Id: 2}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "key does not exist", response.Cause)
}
