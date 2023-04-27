package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/arnaubennassar/go-jsonschema-playground/config"
	"github.com/invopop/jsonschema"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Commands = []*cli.Command{
		{
			Name:   "generate-json-schema",
			Usage:  "Generate the json-schema for the configuration file, and store it on docs/schema.json",
			Action: genJSONSchema,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func genJSONSchema(cli *cli.Context) error {
	r := new(jsonschema.Reflector)
	repoName := "github.com/arnaubennassar/go-jsonschema-playground"
	r.Namer = func(rt reflect.Type) string {
		return strings.TrimLeft(rt.PkgPath(), repoName) + "_" + rt.Name()
	}
	r.ExpandedStruct = true
	r.DoNotReference = true
	if err := r.AddGoComments(repoName, "./"); err != nil {
		return err
	}
	schema := r.Reflect(&config.Config{})
	schema.ID = jsonschema.ID(repoName + "config/config")
	file, err := json.MarshalIndent(schema, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("docs/schema.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
