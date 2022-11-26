package state

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"sync"
	"testing"
)

func TestValueHolder_Value1(t *testing.T) {
	vh := &ValueHolder{
		Type: &Type{rt{
			MapOf: map[string]*Type{
				"assistant": {rt{
					MapOf: map[string]*Type{
						"name": {rt{
							Leaf: LeafString,
						}},
					},
				}},
			},
		}},
		Value: nil,
		mux:   &sync.Mutex{},
	}
	if err := vh.SetValue("assistant.name", "Dean"); err != nil {
		t.Fatal(err)
	}
	val, err := vh.GetValue("assistant.name")
	if err != nil {
		t.Fatal(err)
	}
	if val.Value != "Dean" {
		t.Fatal()
	}
}

func TestValueHolder_Value2(t *testing.T) {
	vh := &ValueHolder{
		Type: &Type{rt{
			MapOf: map[string]*Type{
				"greet": {rt{
					Leaf: LeafBoolean,
				}},
				"greet2": {rt{
					Leaf: LeafBoolean,
				}},
			},
		}},
		Value: map[string]any{
			"greet2": true,
		},
		mux: &sync.Mutex{},
	}

	if err := vh.SetValue("", map[string]any{"greet": true}); err != nil {
		t.Fatal(err)
	}

	val, err := vh.GetValue("greet")
	if err != nil {
		t.Fatal(err)
	}
	if val.Value != true {
		t.Fatal()
	}

	val2, err := vh.GetValue("greet2")
	if err != nil {
		t.Fatal(err)
	}
	if val2.Value != true {
		t.Fatal()
	}

	val3, err := vh.GetValue("")
	if err != nil {
		t.Fatal(err)
	}
	if mp, ok := val3.Value.(map[string]any); !ok {
		t.Fatal()
	} else if mp["greet"] != true {
		t.Fatal()
	} else if mp["greet2"] != true {
		t.Fatal()
	}
}

func TestValueHolder_Value3(t *testing.T) {
	vh := &ValueHolder{
		Type: &Type{rt{
			MapOf: map[string]*Type{
				"greet": {rt{
					MapOf: map[string]*Type{
						"a": {rt{
							Leaf: LeafString,
						}},
						"b": {rt{
							Leaf: LeafBoolean,
						}},
					}},
				}},
		},
		},
		mux: &sync.Mutex{},
	}

	if err := vh.SetValue("greet", map[string]any{"a": "test"}); err != nil {
		t.Fatal(err)
	}
	val, err := vh.GetValue("greet")
	if err != nil {
		t.Fatal(err)
	}
	if val.Value.(map[string]any)["a"] != "test" {
		t.Fatal()
	}

	if err := vh.SetValue("", map[string]any{"greet": nil}); err != nil {
		t.Fatal(err)
	}
	val, err = vh.GetValue("greet")
	if err != nil {
		t.Fatal(err)
	}
	if val.Value != nil {
		t.Fatal()
	}
}

func TestValueHolder_SetMap(t *testing.T) {
	vh := &ValueHolder{
		Type: &Type{rt{
			MapOf: map[string]*Type{
				"background": {rt{
					MapOf: map[string]*Type{
						"color": {rt{
							MapOf: map[string]*Type{
								"name": {rt{
									Leaf: LeafString,
								}},
								"hex": {rt{
									Leaf: LeafString,
								}},
								"isDark": {rt{
									Leaf: LeafBoolean,
								}},
							},
						}},
					},
				}},
			},
		}},
		mux: &sync.Mutex{},
	}
	err := vh.SetValue("background.color", map[string]any{
		"name":   "blue",
		"hex":    "#0000ff",
		"isDark": true,
	})
	if err != nil {
		t.Fatal(err)
	}
	rlt, err := vh.GetValue("background")
	if err != nil {
		t.Fatal(err)
	}
	if rlt.Value.(map[string]any)["color"].(map[string]any)["hex"] != "#0000ff" {
		t.Fatal()
	}
	if rlt.Value.(map[string]any)["color"].(map[string]any)["isDark"] != true {
		t.Fatal()
	}
}

func TestValueHolder_List1(t *testing.T) {
	vh := &ValueHolder{Type: &Type{rt{
		ListOf: &Type{rt{
			Leaf: LeafString,
		}}}},
		mux: &sync.Mutex{},
	}

	err := vh.SetValue("", true)
	if err == nil {
		t.Fatal()
	}

	err = vh.SetValue("", 1)
	if err == nil {
		t.Fatal()
	}

	err = vh.SetValue("", "1")
	if err == nil {
		t.Fatal()
	}

	err = vh.SetValue("", []any{true})
	if err == nil {
		t.Fatal()
	}

	err = vh.SetValue("", []any{})
	if err != nil {
		t.Fatal(err)
	}

	err = vh.SetValue("", []any{"1", "2", "3"})
	if err != nil {
		t.Fatal(err)
	}

	el, err := vh.GetValue("1")
	if err != nil {
		t.Fatal(err)
	}
	if el.Type.Leaf != LeafString {
		t.Fatal()
	}
	if el.Value != "2" {
		t.Fatal()
	}
}

func TestValueHolder_List2(t *testing.T) {
	vh := &ValueHolder{Type: &Type{rt{
		MapOf: map[string]*Type{
			"vals": {rt{
				ListOf: &Type{rt{
					MapOf: map[string]*Type{
						"a": {rt{
							Leaf: LeafString,
						}},
						"b": {rt{
							Leaf: LeafString,
						}},
					},
				}},
			}},
		}},
	},
		mux: &sync.Mutex{},
	}

	err := vh.SetValue("vals", []any{map[string]any{"a1": "1", "b1": "1"}, map[string]any{"a2": "2", "b2": "2"}})
	if err != nil {
		t.Fatal(err)
	}

	val, err := vh.GetValue("vals")
	if err != nil {
		t.Fatal(err)
	}
	if len(val.Value.([]any)) != 2 {
		t.Fatal()
	}
}

func TestValueType_MarshalJSON1(t *testing.T) {
	vt := &Type{rt{
		ListOf: &Type{rt{
			Leaf: LeafString,
		}},
	}}
	jsonStr, err := json.Marshal(vt)
	if err != nil {
		t.Fatal(err)
	}
	if string(jsonStr) != `["string"]` {
		t.Fatal()
	}
	vt2 := Type{}
	err = json.Unmarshal(jsonStr, &vt2)
	if err != nil {
		t.Fatal(err)
	}
	if vt2.ListOf == nil {
		t.Fatal()
	}
	if vt2.ListOf.Leaf != LeafString {
		t.Fatal()
	}
}

func TestValueType_MarshalYAML1(t *testing.T) {
	vt := &Type{rt{
		ListOf: &Type{rt{
			Leaf: LeafString,
		}},
	}}
	yamlStr, err := yaml.Marshal(vt)
	if err != nil {
		t.Fatal(err)
	}
	if string(yamlStr) != "- string\n" {
		t.Fatal()
	}
	vt2 := Type{}
	err = yaml.Unmarshal(yamlStr, &vt2)
	if err != nil {
		t.Fatal(err)
	}
	if vt2.ListOf == nil {
		t.Fatal()
	}
	if vt2.ListOf.Leaf != LeafString {
		t.Fatal()
	}
}
