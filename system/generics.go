package system

import (
	"errors"
	"fmt"
)

// A GenericValuation maps a type onto generic placeholders
type GenericValuation map[string]Interface

// SCHEMA

func (s Interface) ApplyGenerics(gv GenericValuation) (Interface, error) {
	schema := Interface{
		Type: s.Type,
	}

	switch s.Type {
	case SchemaTypeNamed:
		namedSchema, err := s.Named.ApplyGenerics(gv)
		if err != nil {
			return Interface{}, err
		}
		schema.Named = &namedSchema

	case SchemaTypePrimitive:
		primitiveSchema, err := s.Primitive.ApplyGenerics(gv)
		if err != nil {
			return Interface{}, err
		}
		schema.Primitive = &primitiveSchema

	case SchemaTypeList:
		listSchema, err := s.List.ApplyGenerics(gv)
		if err != nil {
			return Interface{}, err
		}
		schema.List = &listSchema

	case SchemaTypeMap:
		mapSchema, err := s.Map.ApplyGenerics(gv)
		if err != nil {
			return Interface{}, err
		}
		schema.Map = &mapSchema

	case SchemaTypeGeneric:
		if !s.Generic.Name.Resolved() {
			return Interface{}, errors.New("generic not resolved")
		}
		gs, ok := gv[string(s.Generic.Name)]
		if !ok {
			return Interface{}, fmt.Errorf("generic %s not available", s.Generic.Name)
		}
		schema = gs

	case SchemaTypeReference:
		referenceSchema, err := s.Reference.ApplyGenerics(gv)
		if err != nil {
			return Interface{}, err
		}
		schema.Reference = &referenceSchema
	}

	return schema, nil
}

func (s NamedInterface) ApplyGenerics(gv GenericValuation) (NamedInterface, error) {
	namedSchema, err := s.Schema.ApplyGenerics(gv)
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

func (s PrimitiveInterface) ApplyGenerics(gv GenericValuation) (PrimitiveInterface, error) {
	return PrimitiveInterface{
		Value: s.Value,
	}, nil
}

func (s ListInterface) ApplyGenerics(gv GenericValuation) (ListInterface, error) {
	entry, err := s.Entry.ApplyGenerics(gv)
	if err != nil {
		return ListInterface{}, err
	}
	return ListInterface{
		Entry: entry,
	}, nil
}

func (e MapInterfaceEntry) ApplyGenerics(gv GenericValuation) (MapInterfaceEntry, error) {
	newEntry := MapInterfaceEntry{}
	entrySchema, err := e.Schema.ApplyGenerics(gv)
	if err != nil {
		return MapInterfaceEntry{}, err
	}
	newEntry.Key = e.Key
	newEntry.Schema = entrySchema
	return newEntry, nil
}

func (s MapInterface) ApplyGenerics(gv GenericValuation) (MapInterface, error) {
	entries := []MapInterfaceEntry{}
	for _, entry := range s.Entries {
		newEntry, err := entry.ApplyGenerics(gv)
		if err != nil {
			return MapInterface{}, err
		}
		entries = append(entries, newEntry)
	}
	return MapInterface{
		Entries: entries,
	}, nil
}

func (s ReferenceInterface) ApplyGenerics(gv GenericValuation) (ReferenceInterface, error) {
	return ReferenceInterface{
		Generics: gv,
		Template: s.Template,
		Name:     s.Name,
	}, nil
}
