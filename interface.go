package vyze

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

type SchemaType string

const (
	SchemaTypeNamed     SchemaType = "named"
	SchemaTypePrimitive SchemaType = "primitive"
	SchemaTypeMap       SchemaType = "map"
	SchemaTypeList      SchemaType = "list"
	SchemaTypeGeneric   SchemaType = "generic"
	SchemaTypeReference SchemaType = "reference"
)

type Interface struct {
	Type      SchemaType          `json:"type" yaml:"type"`
	Named     *NamedInterface     `json:"named,omitempty" yaml:"named,omitempty"`
	Primitive *PrimitiveInterface `json:"primitive,omitempty" yaml:"primitive,omitempty"`
	Map       *MapInterface       `json:"map,omitempty" yaml:"map,omitempty"`
	List      *ListInterface      `json:"list,omitempty" yaml:"list,omitempty"`
	Generic   *GenericInterface   `json:"generic,omitempty" yaml:"generic,omitempty"`
	Reference *ReferenceInterface `json:"reference,omitempty" yaml:"reference,omitempty"`
}

// NAMED

type GenericSlot struct {
	Name        TemplateString `json:"name" yaml:"name"`
	Description string         `json:"description" yaml:"description"`
}

type NamedInterface struct {
	ID       ID             `json:"id" yaml:"id"`
	Name     string         `json:"name" yaml:"name"`
	Template []TemplateSlot `json:"template,omitempty" yaml:"template,omitempty"`
	Generics []GenericSlot  `json:"generics,omitempty" yaml:"generics,omitempty"`
	Schema   Interface      `json:"schema" yaml:"schema"`
}

// PRIMITIVE

type PrimitiveType string

const (
	PrimitiveTypeVoid    PrimitiveType = "void"
	PrimitiveTypeID      PrimitiveType = "id"
	PrimitiveTypeString  PrimitiveType = "string"
	PrimitiveTypeInteger PrimitiveType = "integer"
	PrimitiveTypeFloat   PrimitiveType = "float"
	PrimitiveTypeBoolean PrimitiveType = "boolean"
)

type PrimitiveInterface struct {
	Model string        `json:"model,omitempty" yaml:"model,omitempty"`
	Value PrimitiveType `json:"value" yaml:"value"`
}

// MAP

type MapInterfaceEntry struct {
	Key       TemplateString     `json:"key" yaml:"key"`
	Expansion *TemplateExpansion `json:"expansion,omitempty" yaml:"expansion,omitempty"`
	Schema    Interface          `json:"schema" yaml:"schema"`
}

type MapInterface struct {
	Entries []MapInterfaceEntry `json:"entries" yaml:"entries"`
}

func (s MapInterface) GetEntry(name string) (MapInterfaceEntry, bool) {
	for _, e := range s.Entries {
		if string(e.Key) == name {
			return e, true
		}
	}
	return MapInterfaceEntry{}, false
}

// LIST

type ListInterface struct {
	Entry Interface `json:"entry" yaml:"entry"`
}

// GENERIC

type GenericInterface struct {
	Name TemplateString `json:"name" yaml:"name"`
}

// REFERENCE

type ReferenceInterface struct {
	Generics GenericValuation  `json:"generics,omitempty" yaml:"generics,omitempty"`
	Template TemplateValuation `json:"template,omitempty" yaml:"template,omitempty"`
	Name     string            `json:"name" yaml:"name"`
}

// EQUALITY

func (s Interface) Equals(s2 Interface) bool {
	if s.Type != s2.Type {
		return false
	}
	switch s.Type {
	case SchemaTypeNamed:
		return s.Named.Equals(*s2.Named)
	case SchemaTypePrimitive:
		return s.Primitive.Equals(*s2.Primitive)
	case SchemaTypeList:
		return s.List.Equals(*s2.List)
	case SchemaTypeMap:
		return s.Map.Equals(*s2.Map)
	case SchemaTypeGeneric:
		return s.Generic.Equals(*s2.Generic)
	case SchemaTypeReference:
		return s.Reference.Equals(*s2.Reference)
	}
	return false
}

func (s NamedInterface) Equals(s2 NamedInterface) bool {
	if s.Name != s2.Name {
		return false
	}
	return s.Schema.Equals(s2.Schema)
}

func (s PrimitiveInterface) Equals(s2 PrimitiveInterface) bool {
	return s.Value == s2.Value
}

func (s ListInterface) Equals(s2 ListInterface) bool {
	return s.Entry.Equals(s2.Entry)
}

func (s MapInterface) Equals(s2 MapInterface) bool {
	for _, e := range s.Entries {
		if e2, ok := s2.GetEntry(string(e.Key)); ok {
			if !e.Schema.Equals(e2.Schema) {
				return false
			}
		} else {
			// Other map lacks the entry
			return false
		}
	}
	// Check if the other map has keys we do not have
	for _, e2 := range s2.Entries {
		if _, ok := s.GetEntry(string(e2.Key)); !ok {
			return false
		}
	}
	return true
}

func (s GenericInterface) Equals(s2 GenericInterface) bool {
	panic("implement me")
}

func (s ReferenceInterface) Equals(s2 ReferenceInterface) bool {
	panic("implement me")
}

// JSON

func (s Interface) JSONSchema(lib *Library) (map[string]any, error) {
	switch s.Type {
	case SchemaTypeNamed:
		return s.Named.JSONSchema(lib)
	case SchemaTypePrimitive:
		return s.Primitive.JSONSchema(lib)
	case SchemaTypeList:
		return s.List.JSONSchema(lib)
	case SchemaTypeMap:
		return s.Map.JSONSchema(lib)
	case SchemaTypeGeneric:
		return s.Generic.JSONSchema(lib)
	case SchemaTypeReference:
		return s.Reference.JSONSchema(lib)
	}
	return nil, errors.New("no schema found")
}

func (s NamedInterface) JSONSchema(lib *Library) (map[string]any, error) {
	schema, err := s.Schema.JSONSchema(lib)
	if err != nil {
		return nil, err
	}
	sl := gojsonschema.NewSchemaLoader()
	sl.Draft = gojsonschema.Draft7
	sl.AutoDetect = false
	sl.Validate = true
	schemaBts, _ := json.Marshal(schema)
	loader := gojsonschema.NewStringLoader(string(schemaBts))
	err = sl.AddSchemas(loader)
	if err != nil {
		return nil, err
	}
	var schemaMap map[string]any
	_ = json.Unmarshal(schemaBts, &schemaMap)
	return schemaMap, nil
}

func (s PrimitiveInterface) JSONSchema(lib *Library) (map[string]any, error) {
	schema := map[string]any{}
	switch s.Value {
	case PrimitiveTypeID:
		schema["type"] = "string"
	case PrimitiveTypeString:
		schema["type"] = "string"
	case PrimitiveTypeInteger:
		schema["type"] = "integer"
	case PrimitiveTypeBoolean:
		schema["type"] = "boolean"
	case PrimitiveTypeFloat:
		schema["type"] = "number"
	case PrimitiveTypeVoid:
		schema["type"] = "null"
	}
	return schema, nil
}

func (s ListInterface) JSONSchema(lib *Library) (map[string]any, error) {
	schema := map[string]any{}
	schema["type"] = "array"
	entrySchema, err := s.Entry.JSONSchema(lib)
	if err != nil {
		return nil, err
	}
	schema["items"] = entrySchema
	return schema, nil
}

func (s MapInterface) JSONSchema(lib *Library) (map[string]any, error) {
	schema := map[string]any{}
	schema["type"] = "object"
	properties := map[string]any{}
	for _, n := range s.Entries {
		entrySchema, err := n.Schema.JSONSchema(lib)
		if err != nil {
			return nil, err
		}
		properties[string(n.Key)] = entrySchema
	}
	schema["properties"] = properties
	return schema, nil
}

func (s GenericInterface) JSONSchema(lib *Library) (map[string]any, error) {
	schema := map[string]any{}
	name := string(s.Name)
	if len(name) != 0 {
		schema["title"] = fmt.Sprintf("%s (generic)", name)
		schema["description"] = fmt.Sprintf("Generic type '%s' which needs to be substituted by a concrete type.", name)
	}
	return schema, nil
}

func (s ReferenceInterface) JSONSchema(lib *Library) (map[string]any, error) {
	schema := map[string]any{}
	schema["$ref"] = InterfaceReference(*lib, s.Name, false)
	return schema, nil
}

func InterfaceReference(lib Library, name string, full bool) string {
	base := ""
	if full {
		base = fmt.Sprintf("https://api.vyze.io/service/v1/json/%s", lib.Name)
	}
	if name == "" {
		return base
	}
	return fmt.Sprintf("%s#$defs/%s", base, name)
}

// JSON

func (rf Universe) JSONSchema() (json.RawMessage, error) {
	lib := rf.Library()

	schema := map[string]any{}
	schema["$schema"] = "https://json-schema.org/draft/2020-12/schema"
	schema["$id"] = InterfaceReference(lib, "", true)
	schema["description"] = rf.Description

	definitions := map[string]any{}
	for _, ni := range rf.Interfaces {
		nodeSchema, err := ni.JSONSchema(&lib)
		if err != nil {
			return nil, err
		}
		definitions[ni.Name] = nodeSchema
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

// INTERFACE

func (n Node) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	switch n.Type {
	case NodeTypeContext:
		return n.Context.InterfaceSchema(universe, environment)
	case NodeTypeRelation:
		return n.Relation.InterfaceSchema(universe, environment)
	case NodeTypeSpecials:
		return n.Specials.InterfaceSchema(universe, environment)
	case NodeTypeValue:
		return n.Value.InterfaceSchema(universe, environment)
	case NodeTypeInstance:
		return n.Instance.InterfaceSchema(universe, environment)
	case NodeTypeFilter:
		return n.Filter.InterfaceSchema(universe, environment)
	case NodeTypeSort:
		return n.Sort.InterfaceSchema(universe, environment)
	case NodeTypeSlice:
		return n.Slice.InterfaceSchema(universe, environment)
	case NodeTypeGroup:
		return n.Group.InterfaceSchema(universe, environment)
	case NodeTypeAggregate:
		return n.Aggregate.InterfaceSchema(universe, environment)
	case NodeTypeList:
		return n.List.InterfaceSchema(universe, environment)
	case NodeTypeMap:
		return n.Map.InterfaceSchema(universe, environment)
	case NodeTypeReference:
		return n.Reference.InterfaceSchema(universe, environment)
	}
	return Interface{}, errors.New("no schema found")
}

func (en ContextNode) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	return en.Node.InterfaceSchema(universe, en.Context.Environment)
}

func (en RelationNode) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	relModel := universe.GetModel(en.Relation, "")
	if relModel == nil {
		return Interface{}, errors.New("cannot resolve relation")
	}
	var targetModel *UniverseIdentifier
	if !en.Reverse {
		targetModel = universe.GetTarget(relModel.Mapping)
	} else {
		targetModel = universe.GetOrigin(relModel.Mapping)
	}
	return en.Node.InterfaceSchema(universe, Environment{
		Type:  en.Type,
		Model: targetModel.Name,
	})
}

func (en SpecialsNode) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	return en.Node.InterfaceSchema(universe, Environment{
		Type:  en.Type,
		Model: environment.Model,
	})
}

func (en ValueNode) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	schema := Interface{
		Type:      SchemaTypePrimitive,
		Primitive: &PrimitiveInterface{},
	}
	if en.Field == FieldTypeID {
		schema.Primitive.Value = PrimitiveTypeID
		schema.Primitive.Model = environment.Model
	} else {
		switch en.Format {
		case FormatTypeString, FormatTypeBase64:
			schema.Primitive.Value = PrimitiveTypeString
		case FormatTypeHex:
			schema.Primitive.Value = PrimitiveTypeString
		case FormatTypeInteger:
			schema.Primitive.Value = PrimitiveTypeInteger
		case FormatTypeFloat:
			schema.Primitive.Value = PrimitiveTypeFloat
		case FormatTypeBoolean:
			schema.Primitive.Value = PrimitiveTypeBoolean
		case FormatTypeRaw:
			return Interface{}, errors.New("raw does not have a struct schema")
		}
	}
	return schema, nil
}

func (en InstanceNode) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	schema := Interface{
		Type:      SchemaTypePrimitive,
		Primitive: &PrimitiveInterface{},
	}
	switch en.Format {
	case FormatTypeString, FormatTypeBase64:
		schema.Primitive.Value = PrimitiveTypeString
	case FormatTypeHex:
		schema.Primitive.Value = PrimitiveTypeString
	case FormatTypeInteger:
		schema.Primitive.Value = PrimitiveTypeInteger
	case FormatTypeFloat:
		schema.Primitive.Value = PrimitiveTypeFloat
	case FormatTypeBoolean:
		schema.Primitive.Value = PrimitiveTypeBoolean
	case FormatTypeRaw:
		return Interface{}, errors.New("raw does not have a struct schema")
	}
	return schema, nil
}

func (en FilterNode) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	childSchema, err := en.Node.InterfaceSchema(universe, environment)
	if err != nil {
		return Interface{}, err
	}
	return childSchema, nil
}

func (en SortNode) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	childSchema, err := en.Node.InterfaceSchema(universe, environment)
	if err != nil {
		return Interface{}, err
	}
	return childSchema, nil
}

func (en SliceNode) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	childSchema, err := en.Node.InterfaceSchema(universe, environment)
	if err != nil {
		return Interface{}, err
	}
	return childSchema, nil
}

func (en AggregateNode) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	return Interface{}, nil
}

func (en GroupNode) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	return Interface{}, nil
}

func (en ListNode) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	schema := Interface{
		Type: SchemaTypeList,
		List: &ListInterface{},
	}
	if len(en.KeyFormat) == 0 {
		entrySchema, err := en.Entry.InterfaceSchema(universe, Environment{
			Type:  EnvironmentTypePrimitive,
			Model: environment.Model,
		})
		if err != nil {
			return Interface{}, err
		}
		schema.List.Entry = entrySchema
	} else {
		return Interface{}, errors.New("not implemented yet")
	}
	return schema, nil
}

func (en MapNode) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	schema := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{},
		},
	}
	for _, n := range en.Entries {
		entrySchema, err := n.Node.InterfaceSchema(universe, environment)
		if err != nil {
			return Interface{}, err
		}
		schema.Map.Entries = append(schema.Map.Entries, MapInterfaceEntry{
			Key:    TemplateString(n.Name),
			Schema: entrySchema,
		})
	}
	return schema, nil
}

func (en ReferenceNode) InterfaceSchema(universe *Universe, environment Environment) (Interface, error) {
	if universe == nil {
		return Interface{}, errors.New("require universe for reference nodes")
	}
	ep := universe.GetEndpoint(en.Name)
	if ep.Interface != nil {
		return *ep.Interface, nil
	}
	refSchema, err := ep.InterfaceSchema(*universe)
	if err != nil {
		return Interface{}, err
	}
	return refSchema, nil
}

func (en EndpointNode) InterfaceSchema(universe Universe) (Interface, error) {
	entrySchema, err := en.Node.InterfaceSchema(&universe, en.Context.Environment)
	if err != nil {
		return Interface{}, err
	}
	return entrySchema, nil
}
