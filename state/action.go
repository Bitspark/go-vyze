package state

import (
	"encoding/json"
	"fmt"
	"github.com/Bitspark/go-vyze/core"
	"gopkg.in/yaml.v3"
	"log"
	"strings"
)

type ActionIf struct {
	Condition  *Condition `yaml:"condition"`
	Expression *Expression
	Then       *Action `yaml:"then"`
	Else       *Action `yaml:"else"`
}

type ActionWhile struct {
	Condition *Condition `yaml:"condition"`
	Do        *Action    `yaml:"do"`
}

type ActionBound struct {
	Binding *TwoBinding `yaml:"binding"`
	Action  *Action     `yaml:"update"`
}

type Action struct {
	ru
}

type ru struct {
	Name string `yaml:"-"`

	// Type specifies the state this update requires
	Type *Type `yaml:"type"`

	// Binding specifies whether a sub-state should be created and how to bind it to this state.
	// If nil, it inherits the parent state.
	Bound *ActionBound `yaml:"bound"`

	Action string       `yaml:"update,omitempty"`
	Native ActionNative `yaml:"-"`

	Reference string `yaml:"reference,omitempty"`

	Parallel []*Action `yaml:"parallel,omitempty"`
	Sequence []*Action `yaml:"sequence,omitempty"`

	If    *ActionIf    `yaml:"if,omitempty"`
	While *ActionWhile `yaml:"while,omitempty"`

	Target     string
	Expression *Expression

	compiled bool
}

func (u *Action) Compile(lib *Library) error {
	if u == nil {
		return nil
	}
	if err := u.Type.Compile(lib); err != nil {
		return err
	}
	if u.Reference != "" {
		if lib == nil {
			return fmt.Errorf("library not set")
		}
		ru, err := lib.GetAction(u.Reference)
		if err != nil {
			log.Printf("update %s: %v", u.Reference, err)
		}
		if !ru.compiled {
			if err := ru.Compile(lib); err != nil {
				return err
			}
		}
		*u = *ru
	}
	if u.Bound != nil {
		if err := u.Bound.Action.Compile(lib); err != nil {
			return err
		}
	}
	for _, cu := range u.Parallel {
		if err := cu.Compile(lib); err != nil {
			return err
		}
	}
	for _, cu := range u.Sequence {
		if err := cu.Compile(lib); err != nil {
			return err
		}
	}
	if u.If != nil {
		if err := u.If.Condition.Compile(lib); err != nil {
			return err
		}
		if err := u.If.Then.Compile(lib); err != nil {
			return err
		}
		if err := u.If.Else.Compile(lib); err != nil {
			return err
		}
	}
	if u.While != nil {
		if err := u.While.Condition.Compile(lib); err != nil {
			return err
		}
		if err := u.While.Do.Compile(lib); err != nil {
			return err
		}
	}
	u.Action = strings.TrimSpace(u.Action)
	if err := u.Expression.Compile(lib); err != nil {
		return err
	}
	u.compiled = true
	return nil
}

func (u *Action) MustExecUpdate(stt *MemState) {
	if err := u.Exec(stt); err != nil {
		panic(err)
	}
}

func (u *Action) Exec(stt State) error {
	if u == nil {
		return nil
	}
	if !u.compiled && u.Reference != "" {
		return fmt.Errorf("update not compiled: %s", u.Name)
	}

	// State

	if u.Bound != nil {
		chst, err := stt.Mem().NewChild(strings.ToUpper(u.Name)+"-"+core.NewID().Hex(), &ValueHolder{
			Type:  u.Bound.Action.Type,
			Value: nil,
		}, u.Bound.Binding)
		if err != nil {
			return err
		}

		stt = chst

		// After execution of this update we detach the state again
		defer chst.Detach()

		return u.Bound.Action.Exec(chst)
	}

	// Update

	if u.Native != nil {
		return u.Native(stt.Interface())
	}

	if u.Action != "" {
		if strings.HasPrefix(u.Action, "!") {
			return stt.SetValue(u.Action[1:], nil)
		} else if strings.Contains(u.Action, "=") {
			strs := strings.Split(u.Action, "=")
			if len(strs) != 2 {
				return fmt.Errorf("malformed update: %s", u.Action)
			}
			return stt.SetValue(strings.TrimSpace(strs[0]), strings.TrimSpace(strs[1]))
		} else {
			return stt.SetValue(u.Action, true)
		}
	}

	// Structural

	if len(u.Parallel) > 0 {
		done := make(chan error)
		for _, u := range u.Parallel {
			go func(cu *Action) {
				done <- cu.Exec(stt)
			}(u)
		}
		for range u.Parallel {
			err := <-done
			if err != nil {
				return err
			}
		}
		return nil
	}

	if len(u.Sequence) > 0 {
		for _, cu := range u.Sequence {
			if err := cu.Exec(stt); err != nil {
				return err
			}
		}
		return nil
	}

	if u.If != nil {
		if ok, _ := u.If.Condition.Eval(stt); ok {
			return u.If.Then.Exec(stt)
		} else {
			return u.If.Else.Exec(stt)
		}
	}

	if u.While != nil {
		for {
			if ok, _ := u.While.Condition.Eval(stt); !ok {
				return nil
			}
			if err := u.While.Do.Exec(stt); err != nil {
				return err
			}
		}
	}

	return nil
}

func (u Action) MarshalJSON() ([]byte, error) {
	if u.Name == "" && u.Action != "" {
		return json.Marshal(u.Action)
	} else if u.Name == "" && u.Reference != "" {
		return json.Marshal("$" + u.Reference)
	} else if u.Name == "" && u.Parallel != nil {
		return json.Marshal(u.Parallel)
	} else {
		return json.Marshal(u.ru)
	}
}

func (u *Action) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		if len(str) < 2 || str[0] != '$' {
			*u = Action{
				ru{
					Action: str,
				},
			}
		} else {
			*u = Action{
				ru{
					Reference: str[1:],
				},
			}
		}
		return nil
	}
	var list []*Action
	if err := json.Unmarshal(data, &list); err == nil {
		*u = Action{
			ru{
				Parallel: list,
			},
		}
		return nil
	}
	return json.Unmarshal(data, &u.ru)
}

func (u Action) MarshalYAML() (interface{}, error) {
	if u.Name == "" && u.Action != "" {
		return u.Action, nil
	} else if u.Name == "" && u.Reference != "" {
		return "$" + u.Reference, nil
	} else if u.Name == "" && u.Parallel != nil {
		return u.Parallel, nil
	} else {
		return u.ru, nil
	}
}

func (u *Action) UnmarshalYAML(value *yaml.Node) error {
	var str string
	if err := value.Decode(&str); err == nil {
		if len(str) < 2 || str[0] != '$' {
			*u = Action{
				ru{
					Action: str,
				},
			}
		} else {
			*u = Action{
				ru{
					Reference: str[1:],
				},
			}
		}
		return nil
	}
	var list []*Action
	if err := value.Decode(&list); err == nil {
		*u = Action{
			ru{
				Parallel: list,
			},
		}
		return nil
	}
	return value.Decode(&u.ru)
}
