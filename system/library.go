package system

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type Library struct {
	Name        string                    `json:"name" yaml:"name"`
	Filename    string                    `json:"filename" yaml:"filename"`
	Interfaces  map[string]NamedInterface `json:"schemas" yaml:"schemas"`
	Description string                    `json:"description" yaml:"description"`
}

func NewLibrary() Library {
	return Library{
		Interfaces: map[string]NamedInterface{},
	}
}

func (l *Library) Save(filename string) {
	l.Filename = filename
	libBytes, _ := yaml.Marshal(l)
	ioutil.WriteFile(filename, libBytes, os.ModePerm)
}

func (l *Library) Load(filename string) {
	libBytes, _ := ioutil.ReadFile(filename)
	yaml.Unmarshal(libBytes, l)
	if l.Interfaces == nil {
		l.Interfaces = map[string]NamedInterface{}
	}
	if l.Filename == "" {
		l.Filename = filename
	}
}

func (l *Library) AddSchema(schema NamedInterface) {
	l.Interfaces[schema.Name] = schema
}

func (l *Library) GetSchema(name string) *NamedInterface {
	if s, ok := l.Interfaces[name]; ok {
		return &s
	}
	return nil
}

func (l Library) JSONSchema() (json.RawMessage, error) {
	schema := map[string]any{}
	schema["$schema"] = "https://json-schema.org/draft/2020-12/schema"
	schema["$id"] = fmt.Sprintf("https://api.vyze.io/service/v1/struct/%s/schema", l.Name)
	schema["description"] = l.Description

	definitions := map[string]json.RawMessage{}
	for _, node := range l.Interfaces {
		schemaMap, err := node.JSONSchema(&l)
		if err != nil {
			return nil, err
		}
		schemaBytes, err := json.Marshal(schemaMap)
		if err != nil {
			return nil, err
		}
		definitions[node.Name] = schemaBytes
	}
	schema["$defs"] = definitions

	sl := gojsonschema.NewSchemaLoader()
	sl.Draft = gojsonschema.Draft7
	sl.AutoDetect = false
	sl.Validate = true

	schemaBts, _ := json.Marshal(schema)
	var out bytes.Buffer
	json.Indent(&out, schemaBts, "", "  ")
	return out.Bytes(), nil
}

func (l Library) Matches(s Interface) []Interface {
	schemas := []Interface{
		s,
	}
	for _, n := range l.Interfaces {
		m := NewInterfaceMatcher(n.Generics, n.Template)
		if err := m.Match(n.Schema, s); err != nil {
			continue
		}
		ns := Interface{
			Type: SchemaTypeReference,
			Reference: &ReferenceInterface{
				Name:     n.Name,
				Generics: m.Generics(),
				Template: m.Template(),
			},
		}
		schemas = append(schemas, ns)
	}
	return schemas
}
