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
	gas_peddle     uint16 // min 0, max 65535
	brake_peddle   uint16
	steering_wheel int16 // -32k , +32k
	top_speed_kmh  float64
}

func (c car) kmh() float64 {
	return float64(c.gas_peddle) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_peddle) * (c.top_speed_kmh / usixteenbitmax / khm_multiple)
}

func main() {
	a_car := car{
		gas_peddle:     65000,
		brake_peddle:   0,
		steering_wheel: 12561,
		top_speed_kmh:  225.0}

	fmt.Println(a_car.gas_peddle)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

}
