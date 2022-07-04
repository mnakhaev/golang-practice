package main

import "fmt"

type gun struct {
	name  string
	power int
}

type ak47 struct {
	gun
}

func newAK47() iGun {
	return &ak47{gun: gun{
		name:  "ak47",
		power: 3,
	}}
}

type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{
		gun{
			name:  "mushket",
			power: 10,
		},
	}
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) getPower() int {
	return g.power
}

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAK47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type")
}

type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
