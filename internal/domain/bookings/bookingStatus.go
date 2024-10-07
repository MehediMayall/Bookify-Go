package bookings

type BookingStatus struct {
	reserved  int
	confirmed int
	rejected  int
	cancelled int
	completed int
}

var BookingStatuses BookingStatus = BookingStatus{
	reserved:  1,
	confirmed: 2,
	rejected:  3,
	cancelled: 4,
	completed: 5,
}
