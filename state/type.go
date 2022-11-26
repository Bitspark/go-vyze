package state

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"strings"
)

// Value type

type Type struct {
	rt
}

type rt struct {
	// Name of this parameter, can be used to reference it from elsewhere.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Description of this parameter.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Leaf definition (e.g, string)
	Leaf LeafType `json:"leaf,omitempty" yaml:"leaf,omitempty"`

	// Reference another type, override its description with this one (if available).
	Reference string `json:"reference,omitempty" yaml:"reference,omitempty"`

	// ListOf creates a list type from that type.
	ListOf *Type `json:"listOf,omitempty" yaml:"listOf,omitempty"`

	// MapOf creates a map type from these types.
	MapOf map[string]*Type `json:"mapOf,omitempty" yaml:"mapOf,omitempty"`

	// Optional indicates whether a value has to be set.
	//Optional bool `json:"optional,omitempty" yaml:"optional,omitempty"`

	// Generic represents a slot in the value to be replaced by another value
	Generic string `json:"generic,omitempty" yaml:"generic,omitempty"`

	// GenericMap contains types for generic slots
	GenericMap GenericMap `json:"genericMap,omitempty" yaml:"genericMap,omitempty"`

	// Options contains valid options for this type. They each must have this type. Has no impact if not specified.
	Options []any `json:"options,omitempty" yaml:"options,omitempty"`

	// Initial contains an initial value. It has to have this type. Nil by default.
	Initial any `json:"initial,omitempty" yaml:"initial,omitempty"`

	// compiled indicates whether this type has been compiled, i.e. its reference has been resolved
	compiled bool
}

func (t *Type) Compile(lib *Library) error {
	if t == nil {
		return nil
	}
	if t.Name != "" {
		if t.Name[0] < 'a' || t.Name[0] > 'z' {
			return fmt.Errorf("types names must start with a lower character (a-z)")
		}
		if strings.ContainsAny(t.Name, "-/_ ") {
			return fmt.Errorf("types name must not contain spaces, dashes, underscores or slashes")
		}
	}
	if t.Reference != "" {
		if lib == nil {
			return fmt.Errorf("library not set")
		}
		rt, err := lib.GetType(t.Reference)
		if err != nil {
			log.Printf("type %s: %v", t.Reference, err)
		}
		if rt == nil {
			return fmt.Errorf("type not found: %s", t.Reference)
		}
		if !rt.compiled {
			if err := rt.Compile(lib); err != nil {
				return err
			}
		}
		*t = *rt
	}
	for k, ct := range t.MapOf {
		if k[0] < 'a' || k[0] > 'z' {
			return fmt.Errorf("map keys names must start with a lower character (a-z)")
		}
		if strings.ContainsAny(k, "-/_ ") {
			return fmt.Errorf("map keys must not contain spaces, dashes, underscores or slashes")
		}
		if err := ct.Compile(lib); err != nil {
			return err
		}
	}
	if err := t.ListOf.Compile(lib); err != nil {
		return err
	}
	t.compiled = true
	return nil
}

func (t *Type) Accepts(src *Type) (bool, error) {
	return t.accepts(src, "")
}

func (t Type) MarshalJSON() ([]byte, error) {
	if t.Name != "" || t.Description != "" || len(t.Options) != 0 || t.Initial != nil || t.GenericMap != nil {
		return json.Marshal(t.rt)
	}
	if t.Leaf != 0 {
		return json.Marshal(t.Leaf)
	} else if t.Reference != "" {
		return json.Marshal("$" + t.Reference)
	} else if t.Generic != "" {
		return json.Marshal("<" + t.Reference + ">")
	} else if t.MapOf != nil {
		return json.Marshal(t.MapOf)
	} else if t.ListOf != nil {
		return json.Marshal([]*Type{t.ListOf})
	}
	return nil, fmt.Errorf("unknown type")
}

func (t *Type) UnmarshalJSON(data []byte) error {
	var leaf LeafType
	if err := json.Unmarshal(data, &leaf); err == nil {
		*t = Type{
			rt{
				Leaf: leaf,
			},
		}
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		if len(str) >= 2 && str[0] == '$' {
			t.Reference = str[1:]
			return nil
		}
		if len(str) >= 3 && str[0] == '<' && str[len(str)-1] == '>' {
			t.Reference = str[1 : len(str)-1]
			return nil
		}
		return fmt.Errorf("expected type reference or generic: %s", str)
	}
	var mp map[string]*Type
	if err := json.Unmarshal(data, &mp); mp["mapOf"] == nil && err == nil {
		t.MapOf = mp
		return nil
	}
	var lst []*Type
	if err := json.Unmarshal(data, &lst); err == nil && len(lst) == 1 {
		t.ListOf = lst[0]
		return nil
	}
	return json.Unmarshal(data, &t.rt)
}

func (t Type) MarshalYAML() (interface{}, error) {
	if t.Name != "" || t.Description != "" || len(t.Options) != 0 || t.Initial != nil || t.GenericMap != nil {
		return t.rt, nil
	}
	if t.Leaf != 0 {
		return t.Leaf, nil
	} else if t.Reference != "" {
		return "$" + t.Reference, nil
	} else if t.Generic != "" {
		return "<" + t.Generic + ">", nil
	} else if t.MapOf != nil {
		return t.MapOf, nil
	} else if t.ListOf != nil {
		return []*Type{t.ListOf}, nil
	}
	return nil, fmt.Errorf("unknown type")
}

func (t *Type) UnmarshalYAML(value *yaml.Node) error {
	var leaf LeafType
	if err := value.Decode(&leaf); err == nil {
		*t = Type{
			rt{
				Leaf: leaf,
			},
		}
		return nil
	}
	var str string
	if err := value.Decode(&str); err == nil {
		if len(str) >= 2 && str[0] == '$' {
			t.Reference = str[1:]
			return nil
		}
		if len(str) >= 3 && str[0] == '<' && str[len(str)-1] == '>' {
			t.Reference = str[1 : len(str)-1]
			return nil
		}
		return fmt.Errorf("expected type reference or generic: %s", str)
	}
	var mp map[string]*Type
	if err := value.Decode(&mp); mp["mapOf"] == nil && err == nil {
		t.MapOf = mp
		return nil
	}
	var lst []*Type
	if err := value.Decode(&lst); err == nil && len(lst) == 1 {
		t.ListOf = lst[0]
		return nil
	}
	return value.Decode(&t.rt)
}

func MustParseType(typeJSON string) *Type {
	vt, err := ParseType(typeJSON, nil)
	if err != nil {
		panic(err)
	}
	return vt
}

func ParseType(typeJSON string, lib *Library) (*Type, error) {
	vt := Type{}
	if err := json.Unmarshal([]byte(typeJSON), &vt); err != nil {
		return nil, err
	}
	if err := vt.Compile(lib); err != nil {
		return nil, err
	}
	return &vt, nil
}

func (t *Type) accepts(src *Type, path string) (bool, error) {
	// TODO: Consider options
	if !t.compiled {
		return false, fmt.Errorf("%s: target not compiled", path)
	}
	if src == nil {
		// Check for optional, once implemented
		return true, nil
	}
	if !src.compiled {
		return false, fmt.Errorf("%s: source not compiled", path)
	}

	if t.MapOf != nil {
		if src.MapOf == nil {
			return false, fmt.Errorf("%s: source should be a map", path)
		}
		for kSrc, vSrc := range src.MapOf {
			if vSelf, ok := t.MapOf[kSrc]; !ok {
				return false, fmt.Errorf("%s: source has additional map entry: %s", path, kSrc)
			} else {
				if ok, err := vSelf.accepts(vSrc, path+"."+kSrc); !ok {
					return false, err
				}
			}
		}
		return true, nil
	}

	if t.ListOf != nil {
		if src.ListOf == nil {
			return false, fmt.Errorf("%s: source should be a list", path)
		}
		return t.ListOf.accepts(src.ListOf, path+".%")
	}

	if t.Leaf != src.Leaf {
		return false, fmt.Errorf("%s: differing leaves: %d != %d", path, t.Leaf, src.Leaf)
	} else {
		return true, nil
	}
}

// Generic map

type GenericMap map[string]*Type

// Leaf type

type LeafType int

const (
	LeafString = LeafType(iota + 1)
	LeafFloat
	LeafInteger
	LeafBoolean
	LeafRaw
)

func (tb LeafType) String() string {
	str := ""
	switch tb {
	case LeafString:
		str = "string"
	case LeafFloat:
		str = "float"
	case LeafInteger:
		str = "integer"
	case LeafBoolean:
		str = "boolean"
	case LeafRaw:
		str = "raw"
	}
	return str
}

func (tb *LeafType) FromString(str string) error {
	switch str {
	case "string":
		*tb = LeafString
	case "float":
		*tb = LeafFloat
	case "integer":
		*tb = LeafInteger
	case "boolean":
		*tb = LeafBoolean
	case "raw":
		*tb = LeafRaw
	default:
		return fmt.Errorf("unknown leaf type: %s", str)
	}
	return nil
}

func (tb LeafType) MarshalJSON() ([]byte, error) {
	return json.Marshal(tb.String())
}

func (tb *LeafType) UnmarshalJSON(data []byte) error {
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return fmt.Errorf("leaf type must be string, got %s", string(data))
	}
	return tb.FromString(str)
}

func (tb LeafType) MarshalYAML() (interface{}, error) {
	return tb.String(), nil
}

func (tb *LeafType) UnmarshalYAML(value *yaml.Node) error {
	var str string
	err := value.Decode(&str)
	if err != nil {
		return err
	}
	return tb.FromString(str)
}
