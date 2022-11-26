package state

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type ValueHolder struct {
	Type  *Type `json:"type,omitempty"`
	Value any   `json:"value,omitempty"`
	mux   *sync.Mutex
}

func (v *ValueHolder) String() string {
	return v.string(v.Type, v.Value)
}

func (v *ValueHolder) Copy() *ValueHolder {
	if v == nil {
		return nil
	}
	var val any
	// TODO: Can be done more efficiently
	valJSON, _ := json.Marshal(v.Value)
	_ = yaml.Unmarshal(valJSON, &val)
	cpy := &ValueHolder{
		Type:  v.Type,
		Value: val,
		mux:   &sync.Mutex{},
	}
	return cpy
}

func (v *ValueHolder) GetValue(name string) (*ValueHolder, error) {
	v.mux.Lock()
	val := copyValue(v.Value)
	v.mux.Unlock()
	return v.getValue(v.Type, val, getFragments(name, ValueSep), v.mux)
}

func (v *ValueHolder) SetValue(name string, value any) error {
	val, err := v.setValue(v.Type, v.Value, getFragments(name, ValueSep), value, v.mux)
	if err != nil {
		return err
	}
	v.Value = val
	return nil
}

// Private

func (v *ValueHolder) string(valueType *Type, value any) string {
	if valueType == nil {
		return "null"
	}
	if valueType.MapOf != nil {
		keys := []string{}
		for k := range valueType.MapOf {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})
		str := ""
		for i, k := range keys {
			if i > 0 {
				str += ","
			}
			str += fmt.Sprintf("%s:%s", k, v.string(valueType.MapOf[k], value.(map[string]any)[k]))
		}
		return fmt.Sprintf("[%s]", str)
	}
	if valueType.ListOf != nil {
		str := ""
		for i, e := range value.([]any) {
			if i > 0 {
				str += ","
			}
			str += fmt.Sprintf("%d:%s", i, v.string(valueType.ListOf, e))
		}
		return fmt.Sprintf("[%s]", str)
	}
	if valueType.Leaf != 0 {
		return fmt.Sprintf("%v", value)
	}
	return ""
}

func (v *ValueHolder) getValue(valueType *Type, value any, nameFragments []string, mux *sync.Mutex) (*ValueHolder, error) {
	if valueType == nil {
		return nil, fmt.Errorf("missing type definition")
	}
	if len(nameFragments) == 0 {
		return &ValueHolder{
			Type:  valueType,
			Value: value,
		}, nil
	}
	name := nameFragments[0]
	if valueType.MapOf != nil {
		mux.Lock()
		chType, ok := valueType.MapOf[name]
		mux.Unlock()
		if !ok {
			return nil, fmt.Errorf("unknown value map key: %s", name)
		}
		if value == nil {
			return &ValueHolder{
				Type:  valueType,
				Value: nil,
			}, nil
		}
		valueMap, ok := value.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("not a map value: %s", name)
		}
		mux.Lock()
		chVal, ok := valueMap[name]
		mux.Unlock()
		if !ok {
			return &ValueHolder{
				Type:  valueType,
				Value: nil,
			}, nil
		}
		return v.getValue(chType, chVal, nameFragments[1:], mux)
	} else if valueType.ListOf != nil {
		valueList, ok := value.([]any)
		if !ok {
			return nil, fmt.Errorf("not a list value: %s", name)
		}
		idx, err := strconv.Atoi(name)
		if err != nil || idx < 0 {
			return nil, fmt.Errorf("malformed value list index: %s", name)
		}
		if idx >= len(valueList) {
			return nil, fmt.Errorf("key out of bounds: %s", name)
		} else {
			mux.Lock()
			valEl := valueList[idx]
			mux.Unlock()
			return v.getValue(valueType.ListOf, valEl, nameFragments[1:], mux)
		}
	} else {
		return nil, fmt.Errorf("trying to decent into leaf: %s", strings.Join(nameFragments, "."))
	}
}

func (v *ValueHolder) setValue(valueType *Type, oldValue any, nameFragments []string, newValue any, mux *sync.Mutex) (any, error) {
	if valueType == nil {
		return nil, fmt.Errorf("nil type: %v", nameFragments)
	}
	if len(nameFragments) == 0 {
		if newValue == nil {
			return nil, nil
		}
		if err := valueType.cleanValue(newValue, mux); err != nil {
			return nil, err
		}
		if valueType.MapOf == nil {
			return newValue, nil
		}
		if oldValue == nil {
			oldValue = map[string]any{}
		}
		oldValueMap, ok := oldValue.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("expected a map")
		}
		newValueMap, ok := newValue.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("expected a map")
		}
		keys := []string{}
		mux.Lock()
		for k := range newValueMap {
			keys = append(keys, k)
		}
		mux.Unlock()
		for _, k0 := range keys {
			mux.Lock()
			v0 := newValueMap[k0]
			chVal, _ := oldValueMap[k0]
			chType, ok := valueType.MapOf[k0]
			mux.Unlock()
			if !ok {
				return nil, fmt.Errorf("unknown value map key: %s", k0)
			}
			if val, err := v.setValue(chType, chVal, nil, v0, mux); err != nil {
				return nil, err
			} else {
				mux.Lock()
				oldValueMap[k0] = val
				mux.Unlock()
			}
		}
		return oldValueMap, nil
	}
	name := nameFragments[0]
	if valueType.MapOf != nil {
		if oldValue == nil {
			oldValue = map[string]any{}
		}
		oldValueMap, ok := oldValue.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("not a map value: %s", name)
		}
		mux.Lock()
		chVal, _ := oldValueMap[name]
		chType, ok := valueType.MapOf[name]
		mux.Unlock()
		if !ok {
			return nil, fmt.Errorf("unknown value map key: %s", name)
		}
		val, err := v.setValue(chType, chVal, nameFragments[1:], newValue, mux)
		if err != nil {
			return nil, err
		}
		mux.Lock()
		oldValueMap[name] = val
		mux.Unlock()
		return oldValueMap, nil
	} else if valueType.ListOf != nil {
		valueList, ok := oldValue.([]any)
		if !ok {
			return nil, fmt.Errorf("not a list value: %s", name)
		}
		idx, err := strconv.Atoi(name)
		if err != nil || idx < 0 {
			return nil, fmt.Errorf("malformed value list index: %s", name)
		}
		if idx >= len(valueList) {
			return nil, fmt.Errorf("key out of bounds: %s", name)
		}
		val, err := v.setValue(valueType.ListOf, valueList[idx], nameFragments[1:], newValue, mux)
		if err != nil {
			return nil, err
		}
		mux.Lock()
		valueList[idx] = val
		mux.Unlock()
		return valueList, nil
	} else {
		return nil, fmt.Errorf("trying to decent into leaf: %s", strings.Join(nameFragments, "."))
	}
}

func (t *Type) cleanValue(value any, mux *sync.Mutex) error {
	if value == nil {
		return nil
	}
	if t.MapOf != nil {
		vm, ok := value.(map[string]any)
		if !ok {
			return fmt.Errorf("expected map, got %v", value)
		}
		keys := []string{}
		mux.Lock()
		for k := range vm {
			keys = append(keys, k)
		}
		mux.Unlock()
		for _, k := range keys {
			mux.Lock()
			v := vm[k]
			if entryType, ok := t.MapOf[k]; !ok {
				delete(vm, k)
				mux.Unlock()
				continue
			} else {
				mux.Unlock()
				if err := entryType.cleanValue(v, mux); err != nil {
					return fmt.Errorf("map entry %s: %v", k, err)
				}
			}
		}
		return nil
	}
	if t.ListOf != nil {
		vm, ok := value.([]any)
		if !ok {
			return fmt.Errorf("expected map, got %v", value)
		}
		for k, v := range vm {
			if err := t.ListOf.cleanValue(v, mux); err != nil {
				return fmt.Errorf("list entry %d: %v", k, err)
			}
		}
		return nil
	}
	if t.Leaf == 0 {
		return fmt.Errorf("leaf type not specified")
	}
	switch t.Leaf {
	case LeafString:
		_, ok := value.(string)
		if !ok {
			return fmt.Errorf("expected string, got %v", value)
		}
		return nil

	case LeafBoolean:
		_, ok := value.(bool)
		if !ok {
			return fmt.Errorf("expected boolean, got %v", value)
		}
		return nil

	case LeafFloat:
		_, ok := value.(float64)
		if !ok {
			return fmt.Errorf("expected float, got %v", value)
		}
		return nil

	case LeafInteger:
		_, ok := value.(float64)
		if !ok {
			return fmt.Errorf("expected integer, got %v", value)
		}
		return nil
	}
	return fmt.Errorf("unknown leaf type: %d", t.Leaf)
}

func copyValue(val any) any {
	// TODO: Make more efficient
	jsonBts, _ := json.Marshal(val)
	var cpy any
	_ = json.Unmarshal(jsonBts, &cpy)
	return cpy
}

func deepContains(a, b any) bool {
	mpB, ok := b.(map[string]any)
	if ok {
		mpA, ok := a.(map[string]any)
		if !ok {
			return false
		}
		for kB, elB := range mpB {
			if elA, ok := mpA[kB]; !ok {
				return false
			} else if ok := deepContains(elA, elB); !ok {
				return false
			}
		}
		return true
	}
	ltB, ok := b.([]any)
	if ok {
		ltA, ok := a.([]any)
		if !ok {
			return false
		}
		if len(ltB) > len(ltA) {
			return false
		}
		for i, el := range ltB {
			if ok := deepContains(ltA[i], el); !ok {
				return false
			}
		}
		return true
	}
	return a == b
}
