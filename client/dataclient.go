package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/go-micro"
	proto "github.com/ob-vss-ss19/blatt-4-tolleinanderermachts/proto"
)

func main() {

	fmt.Print("DUMMY DATA\n")
	fmt.Print("==========\n\n")
	service := micro.NewService(micro.Name("movieClient"))
	service.Init()

	//----------------------------------------------------------------------------------------------------
	// Movies
	//____________________________________________________________________________________________________
	movieClient := createMovies(service)

	//----------------------------------------------------------------------------------------------------
	// Rooms
	//____________________________________________________________________________________________________
	createRooms(service)

	//----------------------------------------------------------------------------------------------------
	// Shows
	//____________________________________________________________________________________________________
	createShows(service)

	//----------------------------------------------------------------------------------------------------
	// users
	//____________________________________________________________________________________________________
	userClient := createUsers(service)

	//----------------------------------------------------------------------------------------------------
	// reservations
	//____________________________________________________________________________________________________
	resClient := createReservations(service)

	fmt.Println("\ntry to delete first user.")
	rsp, err := userClient.DeleteUser(context.TODO(), &proto.DeleteUserRequest{Id: 0})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Deleting first user succeeded")
	} else {
		fmt.Println("Deleting first user FAILED, cause: " + rsp.Cause)
	}

	fmt.Println("\nNow delete first movie")

	rsp, err = movieClient.DeleteMovie(context.TODO(), &proto.DeleteMovieRequest{Id: 0})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Deleting first movie succeeded")
	} else {
		fmt.Println("Deleting first movie FAILED, cause: " + rsp.Cause)
	}

	fmt.Println("\nAnd try to delete first user again")

	time.Sleep(2 * time.Second)
	rsp, err = userClient.DeleteUser(context.TODO(), &proto.DeleteUserRequest{Id: 0})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Deleting first user succeeded")
	} else {
		fmt.Println("Deleting first user FAILED, cause: " + rsp.Cause)
	}

	fmt.Println("\nNow check same seat occupation from second reservation. should fail")

	rsp, err = resClient.AddReservation(context.TODO(),
		&proto.AddReservationRequest{UserId: 1, ShowId: 1, Seats: []*proto.Seat{{Row: 1, Column: 2}}})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding second reservation succeeded")
	} else {
		fmt.Println("Adding second reservation FAILED, cause: " + rsp.Cause)
	}

	fmt.Print("2 Szenario\n")
	fmt.Print("==========\n\n")

	rsp, err = resClient.AddReservation(context.TODO(),
		&proto.AddReservationRequest{UserId: 1, ShowId: 2, Seats: []*proto.Seat{{Row: 1, Column: 1}}})
	if err != nil {
		fmt.Println(err)
	}
	id01 := rsp.Id
	if rsp.Succeeded {
		fmt.Println("Adding second reservation succeeded")
	} else {
		fmt.Println("Adding second reservation FAILED, cause: " + rsp.Cause)
	}

	rsp, err = resClient.AddReservation(context.TODO(),
		&proto.AddReservationRequest{UserId: 2, ShowId: 2, Seats: []*proto.Seat{{Row: 1, Column: 1}}})
	if err != nil {
		fmt.Println(err)
	}
	id02 := rsp.Id
	if rsp.Succeeded {
		fmt.Println("Adding second reservation succeeded")
	} else {
		fmt.Println("Adding second reservation FAILED, cause: " + rsp.Cause)
	}

	////
	_, err = resClient.ActivateReservation(context.TODO(),
		&proto.ActivateReservationRequest{UserId: 1, ReservationId: id01})
	if err != nil {
		fmt.Println(err)
	}

	_, err = resClient.ActivateReservation(context.TODO(),
		&proto.ActivateReservationRequest{UserId: 2, ReservationId: id02})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("No error?\n")
}

func createMovies(service micro.Service) proto.MovieControlService {
	movieClient := proto.NewMovieControlService("moviectrl", service.Client())

	for i := 1; i <= 4; i++ {
		rsp, err := movieClient.AddMovie(context.TODO(), &proto.AddMovieRequest{Title: "Movie " + string(i)})
		if err != nil {
			fmt.Println(err)
		}
		if rsp.Succeeded {
			fmt.Printf("Adding %d Movie succeeded\n", i)
		} else {
			fmt.Printf("Adding %d Movie FAILED, cause: "+rsp.Cause+"\n", i)
		}
	}
	return movieClient
}

func createRooms(service micro.Service) proto.RoomControlService {

	roomClient := proto.NewRoomControlService("roomctrl", service.Client())

	rsp, err := roomClient.AddRoom(context.TODO(), &proto.AddRoomRequest{Name: "Room 1", Rows: 2, SeatsPerRow: 10})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding first Room succeeded")
	} else {
		fmt.Println("Adding first Room FAILED, cause: " + rsp.Cause)
	}
	rsp, err = roomClient.AddRoom(context.TODO(), &proto.AddRoomRequest{Name: "Room 2", Rows: 2, SeatsPerRow: 10})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding second Room succeeded")
	} else {
		fmt.Println("Adding second Room FAILED, cause: " + rsp.Cause)
	}
	return roomClient
}

func createShows(service micro.Service) proto.ShowControlService {
	showClient := proto.NewShowControlService("showctrl", service.Client())
	rsp, err := showClient.AddShow(context.TODO(), &proto.AddShowRequest{MovieId: 0, RoomId: 0})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding first show[movie 0, room 0] succeeded")
	} else {
		fmt.Println("Adding first show[movie 0, room 0] FAILED, cause: " + rsp.Cause)
	}
	rsp, err = showClient.AddShow(context.TODO(), &proto.AddShowRequest{MovieId: 1, RoomId: 0})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding second show[movie 1, room 0] succeeded")
	} else {
		fmt.Println("Adding second show[movie 1, room 0] FAILED, cause: " + rsp.Cause)
	}
	rsp, err = showClient.AddShow(context.TODO(), &proto.AddShowRequest{MovieId: 2, RoomId: 1})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding third show[movie 2, room 1] succeeded")
	} else {
		fmt.Println("Adding third show[movie 2, room 1] FAILED, cause: " + rsp.Cause)
	}
	rsp, err = showClient.AddShow(context.TODO(), &proto.AddShowRequest{MovieId: 3, RoomId: 1})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding fourth show[movie 3, room 1] succeeded")
	} else {
		fmt.Println("Adding fourth show[movie 3, room 1] FAILED, cause: " + rsp.Cause)
	}
	return showClient
}

func createUsers(service micro.Service) proto.UserControlService {
	userClient := proto.NewUserControlService("userctrl", service.Client())
	rsp, err := userClient.AddUser(context.TODO(), &proto.AddUserRequest{Name: "User 1"})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding first user succeeded")
	} else {
		fmt.Println("Adding first user FAILED, cause: " + rsp.Cause)
	}
	rsp, err = userClient.AddUser(context.TODO(), &proto.AddUserRequest{Name: "User 2"})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding second user succeeded")
	} else {
		fmt.Println("Adding second user FAILED, cause: " + rsp.Cause)
	}
	rsp, err = userClient.AddUser(context.TODO(), &proto.AddUserRequest{Name: "User 3"})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding third user succeeded")
	} else {
		fmt.Println("Adding third user FAILED, cause: " + rsp.Cause)
	}
	rsp, err = userClient.AddUser(context.TODO(), &proto.AddUserRequest{Name: "User 4"})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding fourth user succeeded")
	} else {
		fmt.Println("Adding fourth user FAILED, cause: " + rsp.Cause)
	}
	return userClient
}

func createReservations(service micro.Service) proto.ReservationControlService {
	resClient := proto.NewReservationControlService("resctrl", service.Client())
	rsp, err := resClient.AddReservation(context.TODO(),
		&proto.AddReservationRequest{UserId: 0, ShowId: 0, Seats: []*proto.Seat{{Row: 1, Column: 1}}})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding first reservation succeeded")
	} else {
		fmt.Println("Adding first reservation FAILED, cause: " + rsp.Cause)
	}
	rsp, err = resClient.AddReservation(context.TODO(),
		&proto.AddReservationRequest{UserId: 1, ShowId: 1, Seats: []*proto.Seat{{Row: 1, Column: 2}}})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding second reservation succeeded")
	} else {
		fmt.Println("Adding second reservation FAILED, cause: " + rsp.Cause)
	}
	rsp, err = resClient.AddReservation(context.TODO(),
		&proto.AddReservationRequest{UserId: 2, ShowId: 2, Seats: []*proto.Seat{{Row: 1, Column: 3}}})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding third reservation succeeded")
	} else {
		fmt.Println("Adding third reservation FAILED, cause: " + rsp.Cause)
	}
	rsp, err = resClient.AddReservation(context.TODO(),
		&proto.AddReservationRequest{UserId: 3, ShowId: 3, Seats: []*proto.Seat{{Row: 1, Column: 4}}})
	if err != nil {
		fmt.Println(err)
	}
	if rsp.Succeeded {
		fmt.Println("Adding fourth reservation succeeded")
	} else {
		fmt.Println("Adding fourth reservation FAILED, cause: " + rsp.Cause)
	}

	resresponse, reserr := resClient.ActivateReservation(context.TODO(),
		&proto.ActivateReservationRequest{UserId: 0, ReservationId: 0})
	if reserr != nil {
		fmt.Println(err)
	}
	if resresponse.Reservation.Active {
		fmt.Println("Activating first reservation succeeded")
	} else {
		fmt.Println("Activating first reservation FAILED")
	}
	resresponse, reserr = resClient.ActivateReservation(context.TODO(),
		&proto.ActivateReservationRequest{UserId: 1, ReservationId: 1})
	if reserr != nil {
		fmt.Println(err)
	}
	if resresponse.Reservation.Active {
		fmt.Println("Activating second reservation succeeded")
	} else {
		fmt.Println("Activating second reservation FAILED")
	}
	resresponse, reserr = resClient.ActivateReservation(context.TODO(),
		&proto.ActivateReservationRequest{UserId: 2, ReservationId: 2})
	if reserr != nil {
		fmt.Println(err)
	}
	if resresponse.Reservation.Active {
		fmt.Println("Activating third reservation succeeded")
	} else {
		fmt.Println("Activating third reservation FAILED")
	}
	resresponse, reserr = resClient.ActivateReservation(context.TODO(),
		&proto.ActivateReservationRequest{UserId: 3, ReservationId: 3})
	if reserr != nil {
		fmt.Println(err)
	}
	if resresponse.Reservation.Active {
		fmt.Println("Activating fourth reservation succeeded")
	} else {
		fmt.Println("Activating fourth reservation FAILED")
	}
	return resClient
}
