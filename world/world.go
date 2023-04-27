package world

import (
	"fmt"
	"time"

	"github.com/arnaubennassar/go-jsonschema-playground/types"
)

// Config bla bla bla
type Config struct {
	// Message that is going to be logged on the console
	GreatingMessage string `jsonschema:"example=hello world,default=hello"`
	// Frequency in which the message is going to be logged
	GreatingFrequency types.Duration `jsonschema_extras:"env=WORLD_GREATINGFREQUENCY" jsonschema:"default=35s"`
}

type World struct {
	greatingMessage   string
	greatingFrequency time.Duration
}

func New(c Config) World {
	return World{
		greatingMessage:   c.GreatingMessage,
		greatingFrequency: c.GreatingFrequency.Duration,
	}
}

func (w *World) Start() {
	func() {
		for {
			time.Sleep(w.greatingFrequency)
			fmt.Println(w.greatingMessage)
		}
	}()
}
