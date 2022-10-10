package go_vyze

import (
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"sort"
	"strings"
)

type StructTemplateType string

const (
	TemplateTypePrimitive StructTemplateType = "primitive"
	TemplateTypeList      StructTemplateType = "list"
)

type TemplateString string

type TemplateSlot struct {
	Name        string             `json:"name" yaml:"name"`
	Type        StructTemplateType `json:"type" yaml:"type"`
	Description string             `json:"description" yaml:"description"`
}

type TemplateValue struct {
	Type   StructTemplateType `json:"type" yaml:"type"`
	Value  any                `json:"value,omitempty" yaml:"value,omitempty"`
	Source string             `json:"source,omitempty" yaml:"source,omitempty"`
}

type TemplateValuation map[string]TemplateValue

type TemplateExpansion struct {
	// Source must point to a template list value
	Source string

	// Target points to a new template primitive value
	Target string
}

// Camel creates a string in pascal case suited to become part of a struct name
func (t TemplateValuation) Camel() string {
	s := ""
	keys := []string{}
	for k := range t {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		s += t[k].Camel()
	}
	return s
}

// Apply returns a new template where all source values of the template are fetched from the sources template
func (t TemplateValuation) Apply(source TemplateValuation) TemplateValuation {
	target := TemplateValuation{}
	for k, v := range t {
		target[k] = v.Apply(source)
	}
	return target
}

// Camel returns a camel case representation of the template value
func (tv TemplateValue) Camel() string {
	if len(tv.Source) != 0 {
		return ""
	}
	sn := ""
	if tv.Type == TemplateTypePrimitive {
		str, ok := tv.Value.(string)
		if !ok {
			return ""
		}
		if len(str) == 0 {
			return ""
		}
		sn = str
	} else if tv.Type == TemplateTypeList {
		strs, ok := tv.Value.([]string)
		if !ok {
			return ""
		}
		for _, str := range strs {
			if len(str) == 0 {
				continue
			}
			sn += " " + str
		}
	}
	return strcase.ToCamel(sn)
}

// Apply fetches the value for the template value from the given template
func (tv TemplateValue) Apply(template TemplateValuation) TemplateValue {
	if tv.Value != nil {
		return TemplateValue{
			Type:  tv.Type,
			Value: tv.Value,
		}
	}
	if tv.Source != "" {
		if tvv, ok := template[tv.Source]; ok {
			return tvv
		} else {
			return tv
		}
	}
	return TemplateValue{}
}

// Copy returns a copy of the template
func (t TemplateValuation) Copy() TemplateValuation {
	t2 := TemplateValuation{}
	for k, v := range t {
		t2[k] = v
	}
	return t2
}

// Expand creates a template for each value in the value list
func (t TemplateValuation) Expand(expansion TemplateExpansion) ([]TemplateValuation, error) {
	values := t[expansion.Source]
	if values.Type != TemplateTypeList {
		return nil, errors.New("require list")
	}
	if len(values.Source) != 0 {
		return nil, errors.New("require value")
	}
	if values.Value == nil {
		return nil, nil
	}
	valueStrings, ok := values.Value.([]string)
	if !ok {
		return nil, errors.New("require list values")
	}
	templates := []TemplateValuation{}
	for _, valueString := range valueStrings {
		tpl := t.Copy()
		tpl[expansion.Target] = TemplateValue{
			Type:  TemplateTypePrimitive,
			Value: valueString,
		}
		templates = append(templates, tpl)
	}
	return templates, nil
}

// Apply replaces placeholders with values from the template provided
func (t TemplateString) Apply(template TemplateValuation) TemplateString {
	nt := t
	for k, v := range template {
		if v.Type != TemplateTypePrimitive {
			continue
		}
		if v.Value == nil {
			continue
		}
		str, ok := v.Value.(string)
		if !ok {
			continue
		}
		oldString := fmt.Sprintf("{%s}", k)
		newString := fmt.Sprintf("%s", str)
		nt = TemplateString(strings.ReplaceAll(string(nt), oldString, newString))
	}
	return nt
}

// Resolved tells whether this template string has been properly resolved, i.e. does not contain placeholders
func (t TemplateString) Resolved() bool {
	return !strings.Contains(string(t), "{") || !strings.Contains(string(t), "}")
}

// SCHEMA

func (s Interface) ApplyTemplate(t TemplateValuation) (Interface, error) {
	schema := Interface{
		Type: s.Type,
	}

	switch s.Type {
	case SchemaTypeNamed:
		namedSchema, err := s.Named.ApplyTemplate(t)
		if err != nil {
			return Interface{}, err
		}
		schema.Named = &namedSchema

	case SchemaTypePrimitive:
		primitiveSchema, err := s.Primitive.ApplyTemplate(t)
		if err != nil {
			return Interface{}, err
		}
		schema.Primitive = &primitiveSchema

	case SchemaTypeList:
		listSchema, err := s.List.ApplyTemplate(t)
		if err != nil {
			return Interface{}, err
		}
		schema.List = &listSchema

	case SchemaTypeMap:
		mapSchema, err := s.Map.ApplyTemplate(t)
		if err != nil {
			return Interface{}, err
		}
		schema.Map = &mapSchema

	case SchemaTypeGeneric:
		genericSchema, err := s.Generic.ApplyTemplate(t)
		if err != nil {
			return Interface{}, err
		}
		schema.Generic = &genericSchema

	case SchemaTypeReference:
		referenceSchema, err := s.Reference.ApplyTemplate(t)
		if err != nil {
			return Interface{}, err
		}
		schema.Reference = &referenceSchema
	}

	return schema, nil
}

func (s NamedInterface) ApplyTemplate(t TemplateValuation) (NamedInterface, error) {
	namedSchema, err := s.Schema.ApplyTemplate(t)
	if err != nil {
		return NamedInterface{}, err
	}
	return NamedInterface{
		Name:     s.Name,
		Template: s.Template,
		Generics: s.Generics,
		Schema:   namedSchema,
	}, nil
}

func (s PrimitiveInterface) ApplyTemplate(t TemplateValuation) (PrimitiveInterface, error) {
	return PrimitiveInterface{
		Value: s.Value,
	}, nil
}

func (s ListInterface) ApplyTemplate(t TemplateValuation) (ListInterface, error) {
	entry, err := s.Entry.ApplyTemplate(t)
	if err != nil {
		return ListInterface{}, err
	}
	return ListInterface{
		Entry: entry,
	}, nil
}

func (e MapInterfaceEntry) ApplyTemplate(t TemplateValuation) (MapInterfaceEntry, error) {
	newEntry := MapInterfaceEntry{}
	key := e.Key.Apply(t)
	entrySchema, err := e.Schema.ApplyTemplate(t)
	if err != nil {
		return MapInterfaceEntry{}, err
	}
	newEntry.Key = key
	newEntry.Schema = entrySchema
	return newEntry, nil
}

func (s MapInterface) ApplyTemplate(t TemplateValuation) (MapInterface, error) {
	entries := []MapInterfaceEntry{}
	for _, entry := range s.Entries {
		if entry.Expansion == nil {
			newEntry, err := entry.ApplyTemplate(t)
			if err != nil {
				return MapInterface{}, err
			}
			entries = append(entries, newEntry)
		} else {
			templates, err := t.Expand(*entry.Expansion)
			if err != nil {
				return MapInterface{}, err
			}
			for _, tpl := range templates {
				newEntry, err := entry.ApplyTemplate(tpl)
				if err != nil {
					return MapInterface{}, err
				}
				entries = append(entries, newEntry)
			}
		}
	}
	return MapInterface{
		Entries: entries,
	}, nil
}

func (s GenericInterface) ApplyTemplate(t TemplateValuation) (GenericInterface, error) {
	return GenericInterface{
		Name: s.Name.Apply(t),
	}, nil
}

func (s ReferenceInterface) ApplyTemplate(t TemplateValuation) (ReferenceInterface, error) {
	return ReferenceInterface{
		Generics: s.Generics,
		Template: s.Template.Apply(t),
		Name:     s.Name,
	}, nil
}
