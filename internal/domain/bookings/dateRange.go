package bookings

import (
	"time"
)

type DateRange struct {
	start        time.Time
	end          time.Time
	daysInNumber int
}

func NewDateRange(start, end time.Time) (*DateRange, error) {
	// if start.Date() > end.Date() {
	// 	return nil, errors.New("start date precedes end date")
	// }

	daysInNumber := end.Day() - start.Day()

	return &DateRange{
		start:        start,
		end:          end,
		daysInNumber: daysInNumber,
	}, nil
}
