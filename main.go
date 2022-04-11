package main

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"os"
)

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}

	conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='/org/xmonad/Log',interface='org.xmonad.Log',member='Update'")

	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)
	for s := range c {
		for _, message := range s.Body {
			msg, ok := message.(string)
			if !ok {
				fmt.Fprintf(os.Stderr, "Received payload that is not of type string")
				continue
			}
			fmt.Printf("%s\n", msg)
		}
	}
}
