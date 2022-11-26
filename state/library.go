package state

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"strings"
)

const LibSep = "."

type Definitions struct {
	Types       map[string]*Type       `json:"types" yaml:"types"`
	Expressions map[string]*Expression `json:"expressions" yaml:"expressions"`
	Conditions  map[string]*Condition  `json:"conditions" yaml:"conditions"`
	Actions     map[string]*Action     `json:"actions" yaml:"actions"`
}

// Library contains definitions and hooks to the native environment. It is independent of states.
type Library struct {
	Name     string
	Parent   *Library
	Children map[string]*Library

	Definitions Definitions
}

func NewLibrary() *Library {
	return &Library{
		Name: "",

		Parent:   nil,
		Children: map[string]*Library{},

		Definitions: Definitions{
			Types:       map[string]*Type{},
			Expressions: map[string]*Expression{},
			Conditions:  map[string]*Condition{},
			Actions:     map[string]*Action{},
		},
	}
}

func (lib *Library) LoadFromDir(dir string) error {
	if lib.Children == nil {
		lib.Children = map[string]*Library{}
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		chPath := path.Join(dir, f.Name())
		if f.IsDir() {
			childLib := &Library{
				Name:     f.Name(),
				Parent:   lib,
				Children: map[string]*Library{},
				Definitions: Definitions{
					Types:      map[string]*Type{},
					Conditions: map[string]*Condition{},
					Actions:    map[string]*Action{},
				},
			}
			lib.Children[childLib.Name] = childLib
			if err := childLib.LoadFromDir(chPath); err != nil {
				return fmt.Errorf("loading program %s from %s: %v", childLib.FullName(), chPath, err)
			}
		} else {
			if !strings.HasSuffix(f.Name(), ".yml") && !strings.HasSuffix(f.Name(), ".yaml") {
				continue
			}
			defs := Definitions{}
			chDefsBytes, err := os.ReadFile(chPath)
			if err != nil {
				return fmt.Errorf("reading definitions from %s: %v", chPath, err)
			}
			if err := yaml.Unmarshal(chDefsBytes, &defs); err != nil {
				return fmt.Errorf("parsing definitions from %s: %v", chPath, err)
			}
			if err := lib.AddDefinitions(defs); err != nil {
				return err
			}
		}
	}

	return nil
}

func (lib *Library) Compile() error {
	for _, t := range lib.Definitions.Types {
		if err := t.Compile(lib); err != nil {
			return fmt.Errorf("compile type %s in library %s: %v", t.Name, lib.FullName(), err)
		}
	}
	for _, t := range lib.Definitions.Expressions {
		if err := t.Compile(lib); err != nil {
			return fmt.Errorf("compile expression %s in library %s: %v", t.Name, lib.FullName(), err)
		}
	}
	for _, t := range lib.Definitions.Conditions {
		if err := t.Compile(lib); err != nil {
			return fmt.Errorf("compile condition %s in library %s: %v", t.Name, lib.FullName(), err)
		}
	}
	for _, t := range lib.Definitions.Actions {
		if err := t.Compile(lib); err != nil {
			return fmt.Errorf("compile action %s in library %s: %v", t.Name, lib.FullName(), err)
		}
	}
	for _, l := range lib.Children {
		if err := l.Compile(); err != nil {
			return err
		}
	}
	return nil
}

func (lib *Library) Copy() *Library {
	return lib.copy(nil)
}

func (lib *Library) FullName() string {
	if lib.Parent == nil {
		return lib.Name
	}
	return lib.Parent.FullName() + LibSep + lib.Name
}

func (lib *Library) GetLibrary(name string) (*Library, error) {
	return lib.getLibrary(strings.Split(name, LibSep))
}

func (lib *Library) MustGetType(name string) *Type {
	d, err := lib.GetType(name)
	if err != nil {
		panic(err)
	}
	return d
}

func (lib *Library) GetType(name string) (*Type, error) {
	return lib.getType(strings.Split(name, LibSep))
}

func (lib *Library) MustGetExpression(name string) *Expression {
	d, err := lib.GetExpression(name)
	if err != nil {
		panic(err)
	}
	return d
}

func (lib *Library) GetExpression(name string) (*Expression, error) {
	return lib.getExpression(strings.Split(name, LibSep))
}

func (lib *Library) MustGetCondition(name string) *Condition {
	d, err := lib.GetCondition(name)
	if err != nil {
		panic(err)
	}
	return d
}

func (lib *Library) GetCondition(name string) (*Condition, error) {
	return lib.getCondition(strings.Split(name, LibSep))
}

func (lib *Library) MustGetAction(name string) *Action {
	d, err := lib.GetAction(name)
	if err != nil {
		panic(err)
	}
	return d
}

func (lib *Library) GetAction(name string) (*Action, error) {
	return lib.getAction(strings.Split(name, LibSep))
}

func (lib *Library) RegisterCondition(name string, cond ConditionNative) error {
	c := &Condition{rc{
		Name:   name,
		Native: cond,
	}}
	return lib.addCondition(c)
}

func (lib *Library) RegisterAction(name string, native ActionNative, valueType *Type) error {
	u := &Action{ru{
		Name:   name,
		Native: native,
		Type:   valueType,
	}}
	return lib.addAction(u)
}

func (lib *Library) AddDefinitions(defs Definitions) error {
	for n, d := range defs.Types {
		d.Name = n
		if err := lib.addType(d); err != nil {
			return err
		}
	}
	for n, d := range defs.Conditions {
		d.Name = n
		if err := lib.addCondition(d); err != nil {
			return err
		}
	}
	for n, d := range defs.Actions {
		d.Name = n
		if err := lib.addAction(d); err != nil {
			return err
		}
	}
	return nil
}

// Private

func (lib *Library) copy(parent *Library) *Library {
	libCpy := &Library{
		Name:        lib.Name,
		Parent:      parent,
		Children:    map[string]*Library{},
		Definitions: lib.Definitions,
	}
	for c, l := range lib.Children {
		l2 := l.copy(lib)
		lib.Children[c] = l2
	}
	return libCpy
}

func (lib *Library) addDefinitions(defs Definitions) error {
	for n, d := range defs.Types {
		d.Name = n
		if err := lib.addType(d); err != nil {
			return fmt.Errorf("adding type: %v", err)
		}
	}
	for n, d := range defs.Conditions {
		d.Name = n
		if err := lib.addCondition(d); err != nil {
			return fmt.Errorf("adding type: %v", err)
		}
	}
	for n, d := range defs.Actions {
		d.Name = n
		if err := lib.addAction(d); err != nil {
			return fmt.Errorf("adding type: %v", err)
		}
	}
	return nil
}

func (lib *Library) getLibrary(nameFragments []string) (*Library, error) {
	if len(nameFragments) == 0 {
		return nil, fmt.Errorf("require at least 1 fragment")
	}
	name := nameFragments[0]
	if len(nameFragments) == 1 {
		if name == "" {
			return lib, nil
		}
		if d, ok := lib.Children[name]; ok {
			return d, nil
		}
		return nil, fmt.Errorf("library not found: %s", name)
	} else {
		if name == "" {
			if lib.Parent != nil {
				return lib.Parent.getLibrary(nameFragments)
			}
			return lib.getLibrary(nameFragments[1:])
		} else {
			if ch, ok := lib.Children[name]; ok {
				return ch.getLibrary(nameFragments[1:])
			}
			return nil, fmt.Errorf("library not found: %s", name)
		}
	}
}

func (lib *Library) addType(libType *Type) error {
	if lib.Definitions.Types == nil {
		lib.Definitions.Types = map[string]*Type{}
	}
	lib.Definitions.Types[libType.Name] = libType
	return nil
}

func (lib *Library) getType(nameFragments []string) (*Type, error) {
	if len(nameFragments) == 0 {
		return nil, fmt.Errorf("require at least 1 type fragment: %s", lib.FullName())
	}
	name := nameFragments[0]
	if len(nameFragments) == 1 {
		if d, ok := lib.Definitions.Types[name]; ok {
			return d, nil
		}
		return nil, fmt.Errorf("type not found in library %s: %s", lib.FullName(), name)
	} else {
		chLib, err := lib.getLibrary(nameFragments[:len(nameFragments)-1])
		if err != nil {
			return nil, fmt.Errorf("resolving type %s: %v", name, err)
		}
		libD, err := chLib.getType(nameFragments[len(nameFragments)-1:])
		if err != nil {
			return nil, fmt.Errorf("resolving type %s: %v", name, err)
		}
		return libD, nil
	}
}

func (lib *Library) addExpression(libExpr *Expression) error {
	if lib.Definitions.Expressions == nil {
		lib.Definitions.Expressions = map[string]*Expression{}
	}
	lib.Definitions.Expressions[libExpr.Name] = libExpr
	return nil
}

func (lib *Library) getExpression(nameFragments []string) (*Expression, error) {
	if len(nameFragments) == 0 {
		return nil, fmt.Errorf("require at least 1 condition fragment: %s", lib.FullName())
	}
	name := nameFragments[0]
	if len(nameFragments) == 1 {
		if d, ok := lib.Definitions.Expressions[name]; ok {
			return d, nil
		}
		return nil, fmt.Errorf("condition not found: %s", name)
	} else {
		chLib, err := lib.getLibrary(nameFragments[:len(nameFragments)-1])
		if err != nil {
			return nil, fmt.Errorf("resolving condition %s: %v", name, err)
		}
		libD, err := chLib.getExpression(nameFragments[len(nameFragments)-1:])
		if err != nil {
			return nil, fmt.Errorf("resolving condition %s: %v", name, err)
		}
		return libD, nil
	}
}

func (lib *Library) addCondition(libCond *Condition) error {
	if lib.Definitions.Conditions == nil {
		lib.Definitions.Conditions = map[string]*Condition{}
	}
	lib.Definitions.Conditions[libCond.Name] = libCond
	return nil
}

func (lib *Library) getCondition(nameFragments []string) (*Condition, error) {
	if len(nameFragments) == 0 {
		return nil, fmt.Errorf("require at least 1 condition fragment: %s", lib.FullName())
	}
	name := nameFragments[0]
	if len(nameFragments) == 1 {
		if d, ok := lib.Definitions.Conditions[name]; ok {
			return d, nil
		}
		return nil, fmt.Errorf("condition not found: %s", name)
	} else {
		chLib, err := lib.getLibrary(nameFragments[:len(nameFragments)-1])
		if err != nil {
			return nil, fmt.Errorf("resolving condition %s: %v", name, err)
		}
		libD, err := chLib.getCondition(nameFragments[len(nameFragments)-1:])
		if err != nil {
			return nil, fmt.Errorf("resolving condition %s: %v", name, err)
		}
		return libD, nil
	}
}

func (lib *Library) addAction(libAction *Action) error {
	if lib.Definitions.Actions == nil {
		lib.Definitions.Actions = map[string]*Action{}
	}
	lib.Definitions.Actions[libAction.Name] = libAction
	return nil
}

func (lib *Library) getAction(nameFragments []string) (*Action, error) {
	if len(nameFragments) == 0 {
		return nil, fmt.Errorf("require at least 1 update fragment: %s", lib.FullName())
	}
	name := nameFragments[0]
	if len(nameFragments) == 1 {
		if d, ok := lib.Definitions.Actions[name]; ok {
			return d, nil
		}
		return nil, fmt.Errorf("action not found: %s", name)
	} else {
		chLib, err := lib.getLibrary(nameFragments[:len(nameFragments)-1])
		if err != nil {
			return nil, fmt.Errorf("resolving action %s: %v", name, err)
		}
		libD, err := chLib.getAction(nameFragments[len(nameFragments)-1:])
		if err != nil {
			return nil, fmt.Errorf("resolving action %s: %v", name, err)
		}
		return libD, nil
	}
}
