package apartments

type Amenity struct {
	Wifi            int
	AirConditioning int
	Parking         int
	SwiminingPool   int
	Gym             int
	Spa             int
	Terrace         int
	PetFriendly     int
	GardenView      int
	MountainView    int
}

var Amenities Amenity = Amenity{
	Wifi:            1,
	AirConditioning: 2,
	Parking:         3,
	SwiminingPool:   4,
	Gym:             5,
	Spa:             6,
	Terrace:         7,
	PetFriendly:     8,
	GardenView:      9,
	MountainView:    10,
}
