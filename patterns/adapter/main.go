package main

import "fmt"

type client struct{}

func (c *client) insertLightningAdapterIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}

type computer interface {
	insertIntoLightningPort()
}

type mac struct{}

func (m *mac) insertIntoLightningPort() {
	fmt.Println("inserted port into MAC machine")
}

type windows struct{}

func (w *windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into Windows machine")
}

type windowsAdapter struct {
	windowsMachine *windows
}

func (w *windowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowsMachine.insertIntoUSBPort()
}

func main() {
	client := &client{}
	mac := &mac{}

	client.insertLightningAdapterIntoComputer(mac)
	windowsMachine := &windows{}
	windowsMachineAdapter := windowsAdapter{
		windowsMachine: windowsMachine,
	}
	client.insertLightningAdapterIntoComputer(&windowsMachineAdapter)
}
