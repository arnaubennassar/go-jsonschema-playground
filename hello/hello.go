package hello

import (
	"fmt"
	"time"
)

// HelloConfig blu blu blu
type HelloConfig struct {
	// Message that is going to be logged on the console
	GreatingMessage string `json:"greatingMessage" jsonschema:"example=hello world,default=hello"`
	// Frequency in which the message is going to be logged
	GreatingFrequency time.Duration `json:"greatingFrequency" jsonschema:"example=1h,default=10s"`
}

type Hello struct {
	greatingMessage   string
	greatingFrequency time.Duration
}

func New(c HelloConfig) Hello {
	return Hello{
		greatingMessage:   c.GreatingMessage,
		greatingFrequency: c.GreatingFrequency,
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
