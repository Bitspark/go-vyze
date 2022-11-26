package state

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

// Binding maps values from a context onto another context, where keys are the targets and values are the source
type Binding map[string]string

// TwoBinding specifies how two states should update each other
// At some point, we can think of an extension to allow for cross-machine bindings
type TwoBinding struct {
	In  Binding
	Out Binding
}

func (sm Binding) updateValue(t State, name string, value any) error {
	if sm == nil {
		return nil
	}
	for tgt, src := range sm {
		if tgt == "@" {
			tgt = ""
		}
		if src == name {
			if err := t.SetValue(tgt, value); err != nil {
				return err
			}
			continue
		}
		if strings.HasPrefix(name, src+ValueSep) {
			if tgt != "" {
				tgt += ValueSep
			}
			tgt = strings.Replace(name, src+ValueSep, tgt, 1)
			if err := t.SetValue(tgt, value); err != nil {
				return err
			}
			continue
		}
		if src == "@" {
			if name != "" {
				if tgt != "" {
					tgt += "." + name
				} else {
					tgt += name
				}
			}
			if err := t.SetValue(tgt, value); err != nil {
				return err
			}
			continue
		}
	}
	return nil
}

func (sm Binding) fetchValue(s *MemState, name string) (*ValueHolder, error) {
	if sm == nil {
		return nil, nil
	}
	if v, ok := sm[name]; ok {
		return s.getValue(strings.Split(v, StateSep))
	}
	return nil, nil
}

func MustParseBinding(strs ...string) *TwoBinding {
	if tb, err := ParseTwoBinding(strs); err != nil {
		panic(err)
	} else {
		return tb
	}
}

func ParseTwoBinding(strs []string) (*TwoBinding, error) {
	tb := &TwoBinding{}
	if err := tb.fromStrings(strs); err != nil {
		return nil, err
	}
	return tb, nil
}

func (tb TwoBinding) toStrings() ([]string, error) {
	strs := []string{}
	for t, s := range tb.In {
		if tb.Out[s] == t {
			strs = append(strs, s+BindingSep+t)
			continue
		}
		strs = append(strs, s+BindingSepIn+t)
	}
	for t, s := range tb.Out {
		if tb.In[s] == t {
			continue
		}
		strs = append(strs, t+BindingSepOut+s)
	}
	return strs, nil
}

func (tb *TwoBinding) fromStrings(strs []string) error {
	for _, str := range strs {
		str = strings.TrimSpace(str)
		if strings.Contains(str, BindingSep) {
			sp := strings.Split(str, BindingSep)
			if len(sp) != 2 {
				return fmt.Errorf("cannot parse two-way binding: %s", str)
			}
			if tb.In == nil {
				tb.In = Binding{}
			}
			if tb.Out == nil {
				tb.Out = Binding{}
			}
			tb.In[strings.TrimSpace(sp[1])] = strings.TrimSpace(sp[0])
			tb.Out[strings.TrimSpace(sp[0])] = strings.TrimSpace(sp[1])
		} else if strings.Contains(str, BindingSepIn) {
			sp := strings.Split(str, BindingSepIn)
			if len(sp) != 2 {
				return fmt.Errorf("cannot parse in binding: %s", str)
			}
			if tb.In == nil {
				tb.In = Binding{}
			}
			tb.In[strings.TrimSpace(sp[1])] = strings.TrimSpace(sp[0])
		} else if strings.Contains(str, BindingSepOut) {
			sp := strings.Split(str, BindingSepOut)
			if len(sp) != 2 {
				return fmt.Errorf("cannot parse out binding: %s", str)
			}
			if tb.Out == nil {
				tb.Out = Binding{}
			}
			tb.Out[strings.TrimSpace(sp[0])] = strings.TrimSpace(sp[1])
		} else {
			return fmt.Errorf("cannot parse binding: %s", str)
		}
	}
	return nil
}

func (tb TwoBinding) MarshalJSON() ([]byte, error) {
	strs, err := tb.toStrings()
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(strs); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (tb *TwoBinding) UnmarshalJSON(data []byte) error {
	var strs []string
	err := json.Unmarshal(data, &strs)
	if err != nil {
		return err
	}
	*tb = TwoBinding{}
	return tb.fromStrings(strs)
}

func (tb TwoBinding) MarshalYAML() (interface{}, error) {
	return tb.toStrings()
}

func (tb *TwoBinding) UnmarshalYAML(value *yaml.Node) error {
	var strs []string
	err := value.Decode(&strs)
	if err != nil {
		return err
	}
	*tb = TwoBinding{}
	return tb.fromStrings(strs)
}
