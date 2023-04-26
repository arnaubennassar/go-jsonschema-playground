package world

import (
	"fmt"
	"time"
)

// WorldConfig bla bla bla
type WorldConfig struct {
	// Message that is going to be logged on the console
	GreatingMessage string `json:"greatingMessage" jsonschema:"example=hello world,default=hello"`
	// Frequency in which the message is going to be logged
	GreatingFrequency time.Duration `json:"greatingFrequency" jsonschema:"example=1h,default=10s"`
}

type World struct {
	greatingMessage   string
	greatingFrequency time.Duration
}

func New(c WorldConfig) World {
	return World{
		greatingMessage:   c.GreatingMessage,
		greatingFrequency: c.GreatingFrequency,
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
