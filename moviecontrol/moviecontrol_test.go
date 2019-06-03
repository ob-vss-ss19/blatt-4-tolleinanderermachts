package main

import (
	"testing"

	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
	"github.com/stretchr/testify/assert"
)

func TestMovieAdd(t *testing.T) {

	movieControl := MovieControl{}

	response := proto.RequestResponse{}

	_ = movieControl.AddMovie(nil, &proto.AddMovieRequest{Title: ""}, &response)

	assert.True(t, response.Succeeded, response.Cause)
}

func TestMovieAddWithSavedCheck(t *testing.T) {

	movieControl := MovieControl{}

	response := proto.RequestResponse{}

	_ = movieControl.AddMovie(nil, &proto.AddMovieRequest{Title: "TestMovieAddWithSavedCheck"}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	movieResponse := proto.GetMovieResponse{}

	_ = movieControl.GetMovie(nil, &proto.GetMovieRequest{}, &movieResponse)

	assert.EqualValues(t, 1, len(movieResponse.Data), "Movie number should be 1!")
	assert.Equal(t, "TestMovieAddWithSavedCheck", movieResponse.Data[0].Title, "movie title mismatch!")
	assert.EqualValues(t, 0, movieResponse.Data[0].Id, "id mismatch")
}

func TestMovieAddDoubleMovie(t *testing.T) {

	movieControl := MovieControl{}

	response := proto.RequestResponse{}

	_ = movieControl.AddMovie(nil, &proto.AddMovieRequest{Title: "TestMovieAddDoubleMovie"}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = movieControl.AddMovie(nil, &proto.AddMovieRequest{Title: "TestMovieAddDoubleMovie"}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "movie already exists", response.Cause)
}

func TestMovieDelete(t *testing.T) {

	movieControl := MovieControl{}

	response := proto.RequestResponse{}

	_ = movieControl.AddMovie(nil, &proto.AddMovieRequest{Title: "TestMovieDelete"}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = movieControl.DeleteMovie(nil, &proto.DeleteMovieRequest{Id: 0}, &response)

	assert.True(t, response.Succeeded, response.Cause)
}

func TestMovieDeleteWithSavedCheck(t *testing.T) {

	movieControl := MovieControl{}

	response := proto.RequestResponse{}

	_ = movieControl.AddMovie(nil, &proto.AddMovieRequest{Title: "TestMovieDeleteWithSavedCheck"}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	movieResponse := proto.GetMovieResponse{}

	_ = movieControl.GetMovie(nil, &proto.GetMovieRequest{}, &movieResponse)

	assert.EqualValues(t, 1, len(movieResponse.Data), "Movie number should be 1!")
	assert.Equal(t, "TestMovieDeleteWithSavedCheck", movieResponse.Data[0].Title, "movie title mismatch!")
	assert.EqualValues(t, 0, movieResponse.Data[0].Id, "id mismatch")

	_ = movieControl.DeleteMovie(nil, &proto.DeleteMovieRequest{Id: 0}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = movieControl.GetMovie(nil, &proto.GetMovieRequest{}, &movieResponse)

	assert.EqualValues(t, 0, len(movieResponse.Data), "Movie number should be 0!")
}

func TestMovieDeleteEmpty(t *testing.T) {

	movieControl := MovieControl{}

	response := proto.RequestResponse{}

	_ = movieControl.DeleteMovie(nil, &proto.DeleteMovieRequest{Id: 0}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "index out of bounds", response.Cause)
}

func TestMovieDeleteNegative(t *testing.T) {

	movieControl := MovieControl{}

	response := proto.RequestResponse{}

	_ = movieControl.DeleteMovie(nil, &proto.DeleteMovieRequest{Id: -1}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "index out of bounds", response.Cause)
}

func TestMovieDeleteNotFound(t *testing.T) {

	movieControl := MovieControl{}

	response := proto.RequestResponse{}

	_ = movieControl.AddMovie(nil, &proto.AddMovieRequest{Title: "TestMovieDeleteNotFound"}, &response)

	assert.True(t, response.Succeeded, response.Cause)

	_ = movieControl.DeleteMovie(nil, &proto.DeleteMovieRequest{Id: 2}, &response)

	assert.False(t, response.Succeeded)
	assert.Equal(t, "index out of bounds", response.Cause)
}
