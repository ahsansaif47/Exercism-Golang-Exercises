package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{battery: 100, batteryDrain: batteryDrain, speed: speed, distance: 0}
	return car
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
func NewTrack(distance int) Track {
	track := Track{distance: distance}
	return track
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		updateCar := Car{battery: car.battery, batteryDrain: car.batteryDrain, speed: car.speed, distance: car.distance}
		return updateCar
	}
	updateCar := Car{battery: car.battery - car.batteryDrain, batteryDrain: car.batteryDrain, speed: car.speed, distance: car.distance + car.speed}
	return updateCar
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	cycles := float64(car.battery) / float64(car.batteryDrain)
	tot_dis := float64(car.speed) * cycles

	return tot_dis >= float64(track.distance)
}
