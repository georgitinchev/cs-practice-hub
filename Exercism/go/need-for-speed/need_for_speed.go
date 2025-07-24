package speed

type Car struct {
	speed        int 
	batteryDrain int 
	battery      int 
	distance     int 
}


func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100, 
	}
}

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	return Track{
		distance: distance,	
	}
}

func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
	return car
}


func CanFinish(car Car, track Track) bool {
	if car.battery == 0 {
		return false
	}
	distanceCanDrive := car.battery / car.batteryDrain * car.speed
	return distanceCanDrive >= track.distance
}
