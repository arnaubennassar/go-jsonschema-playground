# go-jsonschema-playground

## Usage

### Generate json schema

Run `go run . generate-json-schema`, this will create `docs/schema.json`

### Generate markdown doc

1. (Run only once) Install `json-schema-for-humans`, requires `pip`: `pip install json-schema-for-humans`
2. Run `generate-schema-doc --config template_name=md docs/schema.json docs/config-file.md`, this will create [docs/config-file.md](docs/config-file.md)

### Generate HTML doc

1. (Run only once) Install `json-schema-for-humans`, requires `pip`: `pip install json-schema-for-humans`. **NOTE:** this is the same tool as in previous step, so if you've already isntalled it, no need to do it again
2. Run `generate-schema-doc docs/schema.json docs/index.html`, this will create `docs/index.html`, `docs/schema_doc.html`, `docs/schema_doc.html`. [Visualization using GitHub Pages](https://arnaubennassar.github.io/go-jsonschema-playground)

### Generate defaults

1. (Run only once) Install dependency (requires python) `pip install jsonschema-default`
2. Update the defaults of the project ([config/default.go](./config/default.go)) and generate a example JSON with the defaults ([docs/example.json](./docs/example.json)):
   1. `cd utils`
   2. `python3 default_generator.py`

## Tools and resources arround json schema

- [Used to generate schema from Go code](https://github.com/invopop/jsonschema)
- [Used to generate web UI doc](https://github.com/coveooss/json-schema-for-humans)
- [Official list of tools](https://json-schema.org/implementations.html)