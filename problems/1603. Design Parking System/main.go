package main

const Big = 1
const Medium = 2
const Small = 3

type ParkingSystem struct {
	small  int
	medium int
	big    int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		small:  small,
		medium: medium,
		big:    big,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case Big:
		if this.big > 0 {
			this.big -= 1
			return true
		} else {
			return false
		}
	case Medium:
		if this.medium > 0 {
			this.medium -= 1
			return true
		} else {
			return false
		}
	case Small:
		if this.small > 0 {
			this.small -= 1
			return true
		} else {
			return false
		}
	}
	return false
}

func main() {
	parkingSystem := Constructor(1, 1, 0)
	println(parkingSystem.AddCar(1))
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
