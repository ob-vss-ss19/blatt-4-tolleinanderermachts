Communication between the services
---
#### Object creation:

The following objects can be created any time:

* **Users**
    ```
	userClient := proto.NewUserControlService("userctrl", service.Client())
	rsp, err = userClient.AddUser(context.TODO(), &proto.AddUserRequest{Name: "Albert"})
    ```
* **Rooms**
    ```
    roomClient := proto.NewRoomControlService("roomctrl", service.Client())
    rsp, err = roomClient.AddRoom(context.TODO(), &proto.AddRoomRequest{Name:"Kino 1", Rows:2, SeatsPerRow:10})
    ```
* **Movies**
    ```
    movieClient := proto.NewMovieControlService("moviectrl", service.Client())
    rsp, err := movieClient.AddMovie(context.TODO(), &proto.AddMovieRequest{Title: "Movie 1"})
    ```
 
 
Apart from that the other objects require certain conditions to be created:
 
 * **Shows** 
 
    If a show is being created...
 
    ```
    showClient := proto.NewShowControlService("showctrl", service.Client())
    rsp, err = showClient.AddShow(context.TODO(), &proto.AddShowRequest{MovieId:0,RoomId:0})
    ```
    ... the showcontrol service will ask the roomcontrol service about the room data with the given room id via the called function '*GetSingleRoom (GetSingleRoomRequest)*'.
    The room data is then used to get the amount of rows and seats which could possibly be booked by users in a reservation.
    
     **-> To create a show, the given room has to exist (id).**
     
 * **Reservations**
 
    Creating a reservation is a process consisting of two steps.
    
    1) Create a reservation request with every relevant information:
        ```
		response := proto.RequestResponse{}
		error = ReservationControl.AddReservation(context.TODO(), &proto.AddReservationRequest{UserId: 0, ShowId: 0, Seats: []*proto.Seat{{Row: 1, Column: 2}, {Row: 1, Column: 3}}}, &response)
        ```
       If the reservation is valid and no error message is returned, e.g. every wanted seat is not occupied,
       it can be proceeded with the second step. We save our error messages inside the response.cause
       
       For checking the occupation of seats the reservation service needs to check on the show service for the current state of the specified seats and show.
       This happens with the function '*CheckSeat (AvailableSeatRequest)*'
         
    2) Activate the reservation
       ```
	   err := ReservationControl.ActivateReservation(context.TODO(), &proto.ActivateReservationRequest{ReservationId: 0, UserId: 0}, &response)
       ```
       When activating the reservation the show service's function '*CheckSeat(...)*' has to be called again but with ceartain write bits set to finally occupy the seats.
    
#### Object deletion:
 
If Objects are being deleted from the microservice network a **consistency issue** will appear.
To avoid that any service has to notify other services which hold reference id's to its own objects that the object with the specified id no longer exists.

Here is a table to show which service notifies which other service: 


|          Function                            |   Source Service   |  Notified Service  |            Function                          |
|:--------------------------------------------:|:--------------------------------:|:------------------:|:------------------------------:|
| DeleteMovie (DeleteMovieRequest)             | Moviecontrol       | Showcontrol        | NotifyMovieDelete (MovieData)                |
| DeleteRoom (DeleteRoomRequest)               | Roomcontrol        | Showcontrol        | NotifyRoomDelete (RoomData)                  |
| DeleteShow (DeleteShowRequest)               | Showcontrol        | Reservationcontrol |                                              |
| RemoveReservation (RemoveReservationRequest) | Reservationcontrol |  Usercontrol       |                                              |
| RemoveReservation (RemoveReservationRequest) | Reservationcontrol |  Showcontrol       | CheckSeat (AvailableSeatRequest)             |
| DeleteUser (DeleteUserRequest)               | Usercontrol        | Reservationcontrol | RemoveReservation (RemoveReservationRequest) |

