package state

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

const StateSep = "/"
const ValueSep = "."

const BindingSep = "<=>"
const BindingSepIn = "->"
const BindingSepOut = "<-"

type BoundState struct {
	Binding Binding
	State   State
}

type StateAction struct {
	Condition *Condition `json:"condition" yaml:"condition"`
	Update    *Action    `json:"update" yaml:"update"`
}

type State interface {
	String() string
	FullName() string
	Interface() NativeInterface
	SetValue(name string, value any) error
	GetValue(name string) (*ValueHolder, error)
	RemoveChild(name string) error
	Type() *Type
	Mem() *MemState
}

type MemState struct {
	// Name of the state
	Name string

	// Parent contains the parent state to achieve mapping
	Parent *BoundState

	// Children were produced from this state
	Children map[string]*BoundState

	// ValueHolder contains values of this state
	Values *ValueHolder

	// Actions are invoked on each state change
	Actions []*StateAction

	self        State // Used by wrapper states (e.g., for updates)
	muxChildren sync.Mutex
	muxValue    sync.Mutex
}

var whiteSpaceReg = regexp.MustCompile(`\s`)

func NewMemState(name string, values *ValueHolder) (*MemState, error) {
	if len(name) == 0 {
		return nil, errors.New("name must not be empty")
	}
	if whiteSpaceReg.Match([]byte(name)) {
		return nil, errors.New("name must not contain whitespace")
	}
	if strings.ToUpper(name[:1]) != name[:1] {
		return nil, errors.New("name must start with uppercase")
	}
	st := &MemState{
		Name:     name,
		Values:   values.Copy(),
		Children: map[string]*BoundState{},
	}
	st.self = st
	return st, nil
}

func (stt *MemState) String() string {
	return stt.Name + ":" + stt.Values.String()
}

func (stt *MemState) ValueMux() *sync.Mutex {
	return &stt.muxValue
}

func (stt *MemState) FullName() string {
	if stt.Parent == nil {
		return stt.Name
	}
	return stt.Parent.State.FullName() + StateSep + stt.Name
}

func (stt *MemState) Interface() NativeInterface {
	return &nativeInterface{state: stt}
}

func (stt *MemState) SetValue(name string, value any) error {
	return stt.setValue(getFragments(name, StateSep), value)
}

func (stt *MemState) GetValue(name string) (*ValueHolder, error) {
	return stt.getValue(getFragments(name, StateSep))
}

func (stt *MemState) NewChild(name string, values *ValueHolder, childBinding *TwoBinding) (*MemState, error) {
	if childBinding == nil {
		childBinding = &TwoBinding{}
	}
	c1, err := NewMemState(name, values)
	if err != nil {
		return nil, err
	}
	c1.Parent = &BoundState{
		Binding: childBinding.Out,
		State:   stt,
	}
	c1.Values = values.Copy()
	if err := c1.Take(stt, childBinding.In); err != nil {
		return nil, err
	}
	stt.muxChildren.Lock()
	if stt.Children == nil {
		stt.Children = map[string]*BoundState{}
	}
	stt.Children[name] = &BoundState{
		Binding: childBinding.In,
		State:   c1,
	}
	stt.muxChildren.Unlock()
	return c1, nil
}

func (stt *MemState) RunUpdates() error {
	for _, ud := range stt.Actions {
		if ok, _ := ud.Condition.Eval(stt.self); ok {
			if err := ud.Update.Exec(stt.self); err != nil {
				return err
			}
		}
	}
	return nil
}

func (stt *MemState) RemoveChild(name string) error {
	stt.muxChildren.Lock()
	delete(stt.Children, name)
	stt.muxChildren.Unlock()
	return nil
}

func (stt *MemState) Detach() {
	if stt.Parent == nil {
		return
	}
	_ = stt.Parent.State.RemoveChild(stt.Name)
	stt.Parent = nil
}

func (stt *MemState) Take(state *MemState, mapping Binding) error {
	for t, s := range mapping {
		if t == "@" {
			t = ""
		}
		if s == "@" {
			s = ""
		}
		if v, err := state.GetValue(s); err != nil {
			return err
		} else {
			if err := stt.SetValue(t, v.Value); err != nil {
				return err
			}
		}
	}
	return nil
}

func (stt *MemState) Type() *Type {
	return stt.Values.Type
}

func (stt *MemState) Mem() *MemState {
	return stt
}

func (stt *MemState) AddUpdate(cond *Condition, udt *Action) {
	stt.Actions = append(stt.Actions, &StateAction{
		Condition: cond,
		Update:    udt,
	})
}

func (stt *MemState) setValue(nameFragments []string, value any) error {
	if len(nameFragments) == 0 {
		if err := stt.Values.SetValue("", value); err != nil {
			return err
		}
		return stt.update("", value)
	}
	name := nameFragments[0]
	if len(nameFragments) == 1 {
		if val, err := stt.GetValue(name); err != nil {
			return fmt.Errorf("value not found: %s, %v", name, err)
		} else if deepContains(val.Value, value) {
			return nil
		}
		if err := stt.Values.SetValue(name, value); err != nil {
			return err
		}
		return stt.update(name, value)
	} else {
		if name == "" {
			if stt.Parent != nil {
				return stt.Parent.State.SetValue(strings.Join(nameFragments, StateSep), value)
			}
			return stt.setValue(nameFragments[1:], value)
		} else {
			ch := stt.Children[name]
			if ch == nil {
				return fmt.Errorf("state not found: %s", name)
			}
			return ch.State.SetValue(strings.Join(nameFragments[1:], StateSep), value)
		}
	}
}

func (stt *MemState) update(name string, value any) error {
	if stt.Parent != nil && stt.Parent.Binding != nil {
		if err := stt.Parent.Binding.updateValue(stt.Parent.State, name, value); err != nil {
			return err
		}
	}
	for _, ch := range stt.Children {
		if ch.Binding == nil {
			continue
		}
		if err := ch.Binding.updateValue(ch.State, name, value); err != nil {
			return err
		}
	}
	return stt.RunUpdates()
}

func (stt *MemState) getValue(nameFragments []string) (*ValueHolder, error) {
	if len(nameFragments) == 0 {
		return stt.Values.GetValue("")
	}
	name := nameFragments[0]
	if len(nameFragments) == 1 {
		return stt.Values.GetValue(name)
	} else {
		if name == "" {
			if stt.Parent != nil {
				return stt.Parent.State.GetValue(strings.Join(nameFragments, StateSep))
			}
			return stt.getValue(nameFragments[1:])
		} else {
			ch := stt.Children[name]
			if ch == nil {
				return nil, fmt.Errorf("state not found: %s", name)
			}
			return ch.State.GetValue(strings.Join(nameFragments[1:], StateSep))
		}
	}
}
