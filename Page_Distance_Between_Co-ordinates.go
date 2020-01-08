package main

 import (
 	"fmt"
 	"math"
 )

 type Coordinates struct {
 	Latitude  float64
 	Longitude float64
 }

 const radius = 6371 // Earth's mean radius in kilometers

 func degrees2radians(degrees float64) float64 {
 	return degrees * math.Pi / 180
 }

 func (origin Coordinates) Distance(destination Coordinates) float64 {
 	degreesLat := degrees2radians(destination.Latitude - origin.Latitude)
 	degreesLong := degrees2radians(destination.Longitude - origin.Longitude)
 	a := (math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
 		math.Cos(degrees2radians(origin.Latitude))*
 			math.Cos(degrees2radians(destination.Latitude))*math.Sin(degreesLong/2)*
 			math.Sin(degreesLong/2))
 	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
 	d := radius * c

 	return d
 }

 func main() {
 	pointA := Coordinates{17.3850, 78.4867}
 	pointB := Coordinates{12.9716, 77.5946}
 	fmt.Println("Point A : ", pointA)
 	fmt.Println("Point B : ", pointB)

 	distance := pointA.Distance(pointB)
 	fmt.Printf("The distance from point A to point B is %.2f kilometers.\n", distance)

 }
