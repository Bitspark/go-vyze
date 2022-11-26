package state

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"strings"
)

type rc struct {
	Name string `yaml:"-"`

	Condition string          `json:"condition,omitempty" yaml:"condition,omitempty"`
	Native    ConditionNative `json:"-" yaml:"-"`

	Reference string `yaml:"reference,omitempty"`

	And []*Condition `yaml:"and,omitempty"`
	Or  []*Condition `yaml:"or,omitempty"`
	Not *Condition   `yaml:"not,omitempty"`

	compiled bool
}

type Condition struct {
	rc
}

func (c *Condition) Compile(lib *Library) error {
	if c == nil {
		return nil
	}
	if c.Reference != "" {
		if lib == nil {
			return fmt.Errorf("library not set")
		}
		rc, err := lib.GetCondition(c.Reference)
		if err != nil {
			log.Printf("condition %s: %v", c.Reference, err)
		}
		if !rc.compiled {
			if err := rc.Compile(lib); err != nil {
				return err
			}
		}
		*c = *rc
	}
	for _, cc := range c.And {
		if err := cc.Compile(lib); err != nil {
			return err
		}
	}
	for _, cc := range c.Or {
		if err := cc.Compile(lib); err != nil {
			return err
		}
	}
	if err := c.Not.Compile(lib); err != nil {
		return err
	}
	c.Condition = strings.TrimSpace(c.Condition)
	c.compiled = true
	return nil
}

func (c Condition) MustEval(stt State) bool {
	ok, err := c.Eval(stt)
	if err != nil {
		panic(err)
	}
	return ok
}

func (c *Condition) Eval(stt State) (bool, error) {
	if c == nil {
		return true, nil
	}

	if !c.compiled && c.Reference != "" {
		return false, fmt.Errorf("condition not compiled: %s", c.Name)
	}

	// Condition

	if c.Native != nil {
		b, _ := c.Native(stt.Interface())
		return b, nil
	}

	if c.Condition != "" {
		if c.Condition == "" {
			return false, nil
		}
		if strings.HasPrefix(c.Condition, "!") {
			val, _ := stt.GetValue(strings.TrimSpace(c.Condition[1:]))
			return val == nil || val.Value == nil || val.Value == false, nil
		} else if strings.Contains(c.Condition, "==") {
			strs := strings.Split(c.Condition, "==")
			if len(strs) != 2 {
				return false, nil
			}
			val, _ := stt.GetValue(strings.TrimSpace(strs[0]))
			if val == nil {
				return false, nil
			}
			return val.Value == strings.TrimSpace(strs[1]), nil
		} else if strings.Contains(c.Condition, "!=") {
			strs := strings.Split(c.Condition, "!=")
			if len(strs) != 2 {
				return false, nil
			}
			val, _ := stt.GetValue(strings.TrimSpace(strs[0]))
			if val == nil {
				return true, nil
			}
			return val.Value != strings.TrimSpace(strs[1]), nil
		} else {
			val, _ := stt.GetValue(c.Condition)
			if val == nil {
				return false, nil
			}
			return val != nil && val.Value != nil && val.Value != false, nil
		}
	}

	// Structural

	if len(c.And) > 0 {
		for _, cc := range c.And {
			if ok, _ := cc.Eval(stt); !ok {
				return false, nil
			}
		}
		return true, nil
	}

	if len(c.Or) > 0 {
		for _, cc := range c.Or {
			if ok, _ := cc.Eval(stt); ok {
				return true, nil
			}
		}
		return false, nil
	}

	if c.Not != nil {
		ok, _ := c.Not.Eval(stt)
		return !ok, nil
	}

	return false, nil
}

func (c Condition) MarshalJSON() ([]byte, error) {
	if c.Name == "" && c.Condition != "" {
		return json.Marshal(c.Condition)
	} else if c.Name == "" && c.Reference != "" {
		return json.Marshal("$" + c.Reference)
	} else {
		return json.Marshal(c.rc)
	}
}

func (c *Condition) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		if len(str) < 2 || str[0] != '$' {
			*c = Condition{
				rc{
					Condition: str,
				},
			}
		} else {
			*c = Condition{
				rc{
					Reference: str[1:],
				},
			}
		}
		return nil
	}
	return json.Unmarshal(data, &c.rc)
}

func (c Condition) MarshalYAML() (interface{}, error) {
	if c.Name == "" && c.Condition != "" {
		return c.Condition, nil
	} else if c.Name == "" && c.Reference != "" {
		return "$" + c.Reference, nil
	} else {
		return c.rc, nil
	}
}

func (c *Condition) UnmarshalYAML(value *yaml.Node) error {
	var str string
	if err := value.Decode(&str); err == nil {
		if len(str) < 2 || str[0] != '$' {
			*c = Condition{
				rc{
					Condition: str,
				},
			}
		} else {
			*c = Condition{
				rc{
					Reference: str[1:],
				},
			}
		}
		return nil
	}
	return value.Decode(&c.rc)
}
