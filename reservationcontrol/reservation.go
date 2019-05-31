package reservationcontrol

type entry struct {
	Row  int
	Seat int
}
type Reservation struct {
	User    int
	Show    int
	entries []entry
}
