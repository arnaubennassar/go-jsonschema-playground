import jsonschema_default
import json

default_obj = jsonschema_default.create_from("../docs/schema.json")
json_str = json.dumps(default_obj, indent=2)
go_default = '''package config

const DefaultConfigJSON = `
'''+json_str+'''
`
'''

filename = '../config/default.go'
text_file = open(filename, "w")
n = text_file.write(go_default)
text_file.close()

filename = '../docs/example.json'
text_file = open(filename, "w")
n = text_file.write(json_str)
text_file.close()