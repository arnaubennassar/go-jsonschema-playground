package hello

import (
	"fmt"
	"time"

	"github.com/arnaubennassar/go-jsonschema-playground/types"
)

// Config blu blu blu
type Config struct {
	// Message that is going to be logged on the console
	GreatingMessage string `jsonschema:"example=hello world,default=hello"`
	// Frequency in which the message is going to be logged
	GreatingFrequency types.Duration `jsonschema:"default=10s"`
}

type Hello struct {
	greatingMessage   string
	greatingFrequency time.Duration
}

func New(c Config) Hello {
	return Hello{
		greatingMessage:   c.GreatingMessage,
		greatingFrequency: c.GreatingFrequency.Duration,
	}
}

func (h *Hello) Start() {
	func() {
		for {
			time.Sleep(h.greatingFrequency)
			fmt.Println(h.greatingMessage)
		}
	}()
}
