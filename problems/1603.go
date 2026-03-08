package main

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	var lot ParkingSystem
	lot.small = small
	lot.medium = medium
	lot.big = big
}

func (this *ParkingSystem) AddCar(carType int) bool {

}

func main() {

}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
