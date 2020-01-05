package main

import (
	"fmt"

	"github.com/a-contre-plongee/hs50/client"
)

func main() {
	switcher := client.New("192.168.0.20")
	fmt.Println(switcher.SwitchBus(client.BusA, client.InputXPT3))
}
