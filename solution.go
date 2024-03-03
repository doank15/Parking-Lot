package main

import (
	"fmt"
	"math"
	"time"
)

type Car struct {
	ID        string `json:"_id"`
	StartTime int    `json:"start_time"`
}

type ParkingLot struct {
	Cars  map[int]Car `json:"cars"`
	Slots int         `json:"slots"`
}

func CreatePark(size int) *ParkingLot {
	return &ParkingLot{
		Cars:  make(map[int]Car, size),
		Slots: size,
	}
}

func (p *ParkingLot) Park(id string) {
	parked := false
	for i := 1; i <= p.Slots; i++ {
		if p.Cars[i] == (Car{}) {
			p.Cars[i] = Car{
				ID:        id,
				StartTime: time.Now().Hour(),
			}
			fmt.Println("Allocated slot number: ", i)
			parked = true
			break
		}
	}
	if !parked {
		fmt.Println("Sorry, Parking lot is full")
	}
}

func (p *ParkingLot) Leave(id string) {
	for index, value := range p.Cars {
		if value.ID == id {
			charge := 0
			total_time := math.Ceil(float64(time.Now().Hour() - value.StartTime))
			if total_time <= 2 {
				charge = 10
			} else {
				charge = int(total_time-2)*10 + 10
			}

			fmt.Printf("Registration number %s with Slot Number %d is free with Charge %d\n", value.ID, index, charge)

			p.Cars[index] = Car{}
			return
		}
	}
}

func (p *ParkingLot) Status() {
	fmt.Printf("Slot No. Registration No.\n")
	for index, value := range p.Cars {
		fmt.Printf("%d - %s\n", index, value.ID)
	}
}

func main() {
	var parkingLot *ParkingLot
	for {
		var options string
		fmt.Scan(&options)

		switch options {
		case "create":
			var slots int
			fmt.Scan(&slots)
			parkingLot = CreatePark(slots)
			fmt.Println("Created parking lot with: ", slots)
		case "park":
			var id string
			fmt.Scan(&id)
			if parkingLot == nil {
				fmt.Println("Please create parking lot first!")
			} else {
				parkingLot.Park(id)
			}
		case "leave":
			var id string
			fmt.Scan(&id)
			if parkingLot == nil {
				fmt.Println("Please create parking lot first!")
			} else {
				parkingLot.Leave(id)
			}
		case "status":
			if parkingLot == nil {
				fmt.Println("Please create parking lot first!")
			} else {
				parkingLot.Status()
			}
		default:
			fmt.Println("Unknown option")
		}

	}
}
