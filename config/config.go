package config

import (
	"github.com/arnaubennassar/go-jsonschema-playground/hello"
	"github.com/arnaubennassar/go-jsonschema-playground/world"
)

// Config is used to configure the hello world components
type Config struct {
	// Configuration of the hello component, this component will log messages on the console in a specified frequency
	Hello hello.HelloConfig `json:"hello"`
	// Configuration of the world component, this component will log messages on the console in a specified frequency
	World world.WorldConfig `json:"world"`
}
