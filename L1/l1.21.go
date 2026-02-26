package main

import "fmt"

type CarKM struct {
	speedKm float64
}

func (c *CarKM) SpeedKm() float64 { return c.speedKm }

type CarMph struct {
	speedMph float64
}

func (c *CarMph) SpeedMph() float64 { return c.speedMph }

type KmProvider interface {
	SpeedKm() float64
}

type MphProvider interface {
	SpeedMph() float64
}

type KmToMphAdapter struct {
	src KmProvider
}

func (a *KmToMphAdapter) SpeedMph() float64 {
	return a.src.SpeedKm() / 1.60934
}

type MphToKmAdapter struct {
	src MphProvider
}

func (a *MphToKmAdapter) SpeedKm() float64 {
	return a.src.SpeedMph() * 1.60934
}

func printMph(p MphProvider) {
	fmt.Printf("  Скорость: %.2f mph\n", p.SpeedMph())
}

func printKm(p KmProvider) {
	fmt.Printf("  Скорость: %.2f km/h\n", p.SpeedKm())
}

func main() {
	kmCar := &CarKM{speedKm: 120.0}
	mphCar := &CarMph{speedMph: 75.0}

	fmt.Println("── CarKM (исходно km/h) ──────────────────")
	printKm(kmCar)
	printMph(&KmToMphAdapter{src: kmCar})

	fmt.Println("── CarMph (исходно mph) ──────────────────")
	printMph(mphCar)
	printKm(&MphToKmAdapter{src: mphCar})
}
