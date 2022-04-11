package main

import (
	"flag"
	"fmt"
	"github.com/godbus/dbus/v5"
	"os"
)

type Config struct {
	BufferSize int
}

func ParseCli() (Config, error) {
	bufferSize := flag.Int("s", 10, "Amount of dbus signals to buffer")
	flag.Parse()

	return Config{
		BufferSize: *bufferSize,
	}, nil
}

func main() {
	config, err := ParseCli()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse cli flags: %e\n", err)
	}

	conn, err := dbus.SessionBus()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
		os.Exit(1)
	}

	conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='/org/xmonad/Log',interface='org.xmonad.Log',member='Update'")

	// We only buffer config.BufferSize signals. If xmonad-log cannot follow, goddbus/dbus will discard the
	// unhandled signals.
	// See: https://pkg.go.dev/github.com/godbus/dbus#Conn.Signal
	c := make(chan *dbus.Signal, config.BufferSize)
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
