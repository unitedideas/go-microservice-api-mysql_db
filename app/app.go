package main

import "fmt"

//
//func index_handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Go is neat")
//
//}
//
//func about_handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Expert web design")
//
//}
//
//func main() {
//
//	http.HandleFunc("/", index_handler)
//	http.HandleFunc("/about/", about_handler)
//	http.ListenAndServe(":8000", nil)
//
//}

const usixteenbitmax float64 = 65535
const khm_multiple float64 = 1.60934

type car struct {
	gasPeddle     uint16 // min 0, max 65535
	brakePeddle   uint16
	steeringWheel int16 // -32k , +32k
	topSpeedKmh   float64
}

// Value receiver makes a copy of the struct we passed to it so the value of the
// original type is not modified

func (c car) kmh() float64 {
	return float64(c.gasPeddle) * (c.topSpeedKmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gasPeddle) * (c.topSpeedKmh / usixteenbitmax / khm_multiple)
}

// Pointer receiver points to the value in memory and modifies the value
func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}

func main() {
	aCar := car{
		gasPeddle:     65000,
		brakePeddle:   0,
		steeringWheel: 12561,
		topSpeedKmh:   225.0}

	fmt.Println(aCar.gasPeddle)
	fmt.Println("kmh:", aCar.kmh())
	fmt.Println("mph:", aCar.mph())
	fmt.Println()
	aCar.newTopSpeed(500)
	fmt.Println("kmh:", aCar.kmh())
	fmt.Println("mph:", aCar.mph())

}
