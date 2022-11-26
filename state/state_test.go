package state

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestConfig_Take(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"a": {rt{Leaf: LeafString}},
		}}},
	})
	st1.Values.Value = map[string]any{"a": "1"}
	st2, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"b": {rt{Leaf: LeafString}},
			"c": {rt{Leaf: LeafString}},
		}}},
	})
	mp := Binding{
		"b": "a",
		"c": "a",
	}
	err := st2.Take(st1, mp)
	if err != nil {
		t.Fatal(err)
	}
	if st2.Values.Value.(map[string]any)["b"] != "1" {
		t.Fatal()
	}
	if st2.Values.Value.(map[string]any)["c"] != "1" {
		t.Fatal()
	}
}

func TestState_SetValue1(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{})
	_, _ = st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"value": {rt{Leaf: LeafString}},
		}}},
	}, nil)
	err := st1.SetValue("Child1/value", "test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestState_SetValue2(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"value": {rt{Leaf: LeafString}},
		}}},
	})
	st2, _ := st1.NewChild("Child1", &ValueHolder{}, nil)
	st3, _ := st2.NewChild("Child2", &ValueHolder{}, nil)
	err := st3.SetValue("/value", "test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestState_GetValue1(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{})
	st2, _ := st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"value": {rt{Leaf: LeafString}},
		}}},
	}, nil)
	_ = st1.SetValue("Child1/value", "test")
	if v, err := st1.GetValue("Child1/value"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}
	if v, err := st2.GetValue("value"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}
}

func TestState_GetValue2(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"value": {rt{Leaf: LeafString}},
		}}},
	})
	st2, _ := st1.NewChild("Child1", nil, nil)
	st3, _ := st2.NewChild("Child2", nil, nil)
	_ = st3.SetValue("/value", "test")
	if v, err := st3.GetValue("/value"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}
	if v, err := st1.GetValue("value"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}
}

func TestState_DeleteValue1(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"test": {rt{Leaf: LeafString}},
		}}},
	})
	_ = st1.SetValue("test", "1")
	err := st1.SetValue("test", nil)
	if err != nil {
		t.Fatal(err)
	}
	v, _ := st1.GetValue("test")
	if v.Value != nil {
		t.Fatal()
	}
}

func TestState_DeleteValue2(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"a": {rt{Leaf: LeafString}},
		}}},
	})
	_, _ = st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"a": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("a "+BindingSep+" a"))
	_ = st1.SetValue("Child1/a", "1")
	err := st1.SetValue("a", nil)
	if err != nil {
		t.Fatal(err)
	}
	v, _ := st1.GetValue("Child1/a")
	if v.Value != nil {
		t.Fatal()
	}
}

func TestState_DeleteValue3(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"a": {rt{Leaf: LeafString}},
		}}},
	})
	st2, _ := st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"a": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("a "+BindingSep+" a"))
	_ = st1.SetValue("Child1/a", "1")
	err := st1.SetValue("Child1/a", nil)
	if err != nil {
		t.Fatal(err)
	}
	v, _ := st2.GetValue("a")
	if v.Value != nil {
		t.Fatal()
	}
}

func TestState_NewChild1(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"a": {rt{Leaf: LeafString}},
		}}},
	})
	_ = st1.SetValue("a", "1")
	st2, err := st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"b": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("a "+BindingSepIn+" b"))
	if err != nil {
		t.Fatal(err)
	}
	if v, _ := st2.GetValue("b"); v.Value != "1" {
		t.Fatal()
	}
	if v, _ := st2.GetValue("b"); v.Value != "1" {
		t.Fatal()
	}
}

func TestState_NewChild2(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"b": {rt{Leaf: LeafString}},
		}}},
	})
	st2, err := st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"a": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("b "+BindingSepOut+" a"))
	if err != nil {
		t.Fatal(err)
	}
	_ = st2.SetValue("a", "1")
	if v, _ := st2.GetValue("a"); v.Value != "1" {
		t.Fatal()
	}
	if v, _ := st1.GetValue("b"); v.Value != "1" {
		t.Fatal()
	}
}

func TestState_NewChild3(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"a": {rt{Leaf: LeafString}},
			"c": {rt{Leaf: LeafString}},
		}}},
	})
	_, err := st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"b": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("a "+BindingSepIn+" b", "c "+BindingSepOut+" b"))
	if err != nil {
		t.Fatal(err)
	}
	_ = st1.SetValue("a", "1")
	if v, _ := st1.GetValue("c"); v.Value != "1" {
		t.Fatal()
	}
}

func TestState_NewChild4(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"g0": {rt{Leaf: LeafString}},
		}}},
	})
	st2, _ := st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"l1": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("g0 "+BindingSepOut+" l1"))
	st3, _ := st1.NewChild("Child2", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"l2": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("g0 "+BindingSepIn+" l2"))
	_ = st2.SetValue("l1", "value")
	if v, _ := st3.GetValue("l2"); v.Value != "value" {
		t.Fatal()
	}
}

func TestState_Detach1(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"g0": {rt{Leaf: LeafString}},
		}}},
	})
	st2, _ := st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"l1": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("g0 "+BindingSepOut+" l1"))
	st3, _ := st1.NewChild("Child2", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"l2": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("g0 "+BindingSepIn+" l2"))
	st3.Detach()
	_ = st2.SetValue("l1", "value")
	if v, _ := st3.GetValue("l2"); v.Value != nil {
		t.Fatal()
	}
}

func TestState_Detach2(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"g0": {rt{Leaf: LeafString}},
		}}},
	})
	st2, _ := st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"l1": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("g0 "+BindingSepOut+" l1"))
	st3, _ := st1.NewChild("Child2", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"l2": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("g0 "+BindingSepIn+" l2"))
	st2.Detach()
	_ = st2.SetValue("l1", "value")
	if v, _ := st3.GetValue("l2"); v.Value != nil {
		t.Fatal()
	}
}

func TestState_MarshalJSON(t *testing.T) {
	tb := TwoBinding{
		In:  Binding{"l0": "g0", "l1": "g1"},
		Out: Binding{"g0": "l0", "g2": "l1"},
	}
	bts, err := tb.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	tb2 := TwoBinding{}
	err = json.Unmarshal(bts, &tb2)
	if err != nil {
		t.Fatal(err)
	}
	if len(tb2.In) != 2 {
		t.Fatal()
	}
	if len(tb2.Out) != 2 {
		t.Fatal()
	}
}

func TestState_MarshalYAML(t *testing.T) {
	tb := TwoBinding{
		In:  Binding{"l0": "g0", "l1": "g1"},
		Out: Binding{"g0": "l0", "g2": "l1"},
	}
	bts, err := yaml.Marshal(tb)
	if err != nil {
		t.Fatal(err)
	}
	tb2 := TwoBinding{}
	err = yaml.Unmarshal(bts, &tb2)
	if err != nil {
		t.Fatal(err)
	}
	if len(tb2.In) != 2 {
		t.Fatal()
	}
	if len(tb2.Out) != 2 {
		t.Fatal()
	}
}

func TestState_TwoWayBinding1(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"g0": {rt{Leaf: LeafString}},
		}}},
	})
	st2, _ := st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"l1": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("g0 "+BindingSep+" l1"))
	st3, _ := st1.NewChild("Child2", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"l2": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("g0 "+BindingSep+" l2"))

	_ = st1.SetValue("Child1/l1", "test")

	if v, err := st1.GetValue("g0"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}

	if v, err := st2.GetValue("l1"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}

	if v, err := st3.GetValue("l2"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}
}

func TestState_TwoWayBinding2(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"g0": {rt{Leaf: LeafString}},
		}}},
	})
	st2, _ := st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"l1": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("g0 "+BindingSep+" l1"))
	st3, _ := st2.NewChild("Child2", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"l2": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("/Child1/l1 "+BindingSep+" l2"))

	_ = st1.SetValue("Child1/Child2/l2", "test")

	if v, err := st1.GetValue("g0"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}

	if v, err := st2.GetValue("l1"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}

	if v, err := st3.GetValue("l2"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}

	_ = st3.SetValue("l2", "test2")

	if v, err := st2.GetValue("l1"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test2" {
		t.Fatal()
	}

	if v, err := st1.GetValue("g0"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test2" {
		t.Fatal()
	}
}

func TestState_TwoWayBinding3(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"Child1_val1": {rt{Leaf: LeafString}},
		}}},
	})
	_, _ = st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"val1": {rt{Leaf: LeafString}},
		}}},
		Value: map[string]any{"val1": "1"},
	}, MustParseBinding("Child1_val1 "+BindingSepOut+" val1"))
	st3, _ := st1.NewChild("Child2", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"Child1_val1": {rt{Leaf: LeafString}},
		}}},
		Value: map[string]any{"val1": "1"},
	}, MustParseBinding("Child1_val1 "+BindingSepIn+" Child1_val1"))

	_ = st1.SetValue("Child1/val1", "test")

	if v, err := st3.GetValue("Child1_val1"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}
}

func TestState_TwoWayBinding4(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"complex": {rt{
				MapOf: map[string]*Type{
					"a": {rt{Leaf: LeafString}},
					"b": {rt{Leaf: LeafString}},
				},
			}},
		}}},
	})
	st2, _ := st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"complex": {rt{
				MapOf: map[string]*Type{
					"a": {rt{Leaf: LeafString}},
					"b": {rt{Leaf: LeafString}},
				},
			}},
		}}},
	}, MustParseBinding("complex.a "+BindingSep+" complex.b", "complex.b "+BindingSep+" complex.a"))

	_ = st1.SetValue("complex.a", "test1")
	_ = st1.SetValue("Child1/complex.a", "test2")

	if v, err := st2.GetValue("complex.b"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test1" {
		t.Fatal()
	}

	if v, err := st1.GetValue("complex.b"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test2" {
		t.Fatal()
	}
}

func TestState_TwoWayBinding5(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"complex": {rt{
				MapOf: map[string]*Type{
					"a": {rt{Leaf: LeafString}},
					"b": {rt{Leaf: LeafString}},
				},
			}},
		}}},
	})
	st2, _ := st1.NewChild("Child1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"complex": {rt{
				MapOf: map[string]*Type{
					"a": {rt{Leaf: LeafString}},
					"b": {rt{Leaf: LeafString}},
				},
			}},
		}}},
	}, MustParseBinding("complex "+BindingSep+" complex"))

	_ = st1.SetValue("complex.a", "test1")
	_ = st1.SetValue("Child1/complex.b", "test2")

	if v, err := st2.GetValue("complex.a"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test1" {
		t.Fatal()
	}

	if v, err := st1.GetValue("complex.b"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test2" {
		t.Fatal()
	}
}

func TestState_WholeBinding1(t *testing.T) {
	st1, _ := NewMemState("Root", &ValueHolder{Type: MustParseType(`{"a": "string", "b": "string"}`)})
	ch1, err := st1.NewChild("Child1", &ValueHolder{Type: MustParseType(`"string"`)}, MustParseBinding("a <=> @"))
	if err != nil {
		t.Fatal(err)
	}

	st1.SetValue("a", "test1")
	val, _ := ch1.GetValue("")
	if val.Value != "test1" {
		t.Fatal(val.Value)
	}

	ch1.SetValue("", "test2")
	val, _ = st1.GetValue("a")
	if val.Value != "test2" {
		t.Fatal(val.Value)
	}
}

func TestState_WholeBinding2(t *testing.T) {
	st1, _ := NewMemState("Root", &ValueHolder{Type: MustParseType(`"string"`)})
	ch1, err := st1.NewChild("Child1", &ValueHolder{Type: MustParseType(`{"a": "string", "b": "string"}`)}, MustParseBinding("@ <=> a"))
	if err != nil {
		t.Fatal(err)
	}

	ch1.SetValue("a", "test1")
	val, _ := st1.GetValue("")
	if val.Value != "test1" {
		t.Fatal()
	}

	st1.SetValue("", "test2")
	val, _ = ch1.GetValue("a")
	if val.Value != "test2" {
		t.Fatal(val.Value)
	}
}

func TestState_WholeBinding3(t *testing.T) {
	st1, _ := NewMemState("Root", &ValueHolder{Type: MustParseType(`{"a": "string", "b": "string"}`)})
	ch1, err := st1.NewChild("Child1", &ValueHolder{Type: MustParseType(`{"a": "string", "b": "string"}`)}, MustParseBinding("@ <=> @"))
	if err != nil {
		t.Fatal(err)
	}

	st1.SetValue("a", "test1")
	val, _ := ch1.GetValue("a")
	if val.Value != "test1" {
		t.Fatal()
	}

	ch1.SetValue("a", "test2")
	val, _ = st1.GetValue("a")
	if val.Value != "test2" {
		t.Fatal(val.Value)
	}
}
