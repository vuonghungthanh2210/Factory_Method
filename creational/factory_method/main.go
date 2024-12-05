package main

import "fmt"

type iTransport interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type transport struct {
	name  string
	power int
}

func (g *transport) setName(name string) {
	g.name = name
}

func (g *transport) getName() string {
	return g.name
}

func (g *transport) setPower(power int) {
	g.power = power
}

func (g *transport) getPower() int {
	return g.power
}

type truck struct {
	transport
}

func newTruck() iTransport {
	return &truck{
		transport: transport{
			name:  "Truck",
			power: 4,
		},
	}
}

type ships struct {
	transport
}

func newShips() iTransport {
	return &ships{
		transport: transport{
			name:  "Ships",
			power: 5,
		},
	}
}

type grab struct {
	transport
}

func newgrab() iTransport {
	return &grab{
		transport: transport{
			name:  "grab",
			power: 7,
		},
	}
}

func getTransport(transportType string) (iTransport, error) {
	if transportType == "truck" {
		return newTruck(), nil
	}
	if transportType == "ships" {
		return newShips(), nil
	}
	if transportType == "grab" {
		return newgrab(), nil
	}
	return nil, fmt.Errorf("Wrong transport type")
}

func main() {
	truck, _ := getTransport("truck")
	ships, _ := getTransport("ships")
	grab, _ := getTransport("grab")
	printDetails(truck)
	printDetails(ships)
	printDetails(grab)
}

func printDetails(g iTransport) {
	fmt.Printf("Name: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
