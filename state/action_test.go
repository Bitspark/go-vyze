package state

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestState_ExecuteUpdate1(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{
			MapOf: map[string]*Type{"g0": {rt{Leaf: LeafString}}},
		}},
		Value: nil,
	})
	_ = (&Action{ru{
		Action: "g0=test",
	}}).Exec(st1)
	if v, err := st1.GetValue("g0"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}
}

func TestState_ExecuteUpdate2(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"g0": {rt{Leaf: LeafBoolean}},
		}}},
	})
	(&Action{ru{
		Action: "g0",
	}}).MustExecUpdate(st1)
	if v, err := st1.GetValue("g0"); err != nil {
		t.Fatal(err)
	} else if v.Value != true {
		t.Fatal()
	}
	(&Action{ru{
		Action: "!g0",
	}}).MustExecUpdate(st1)
	if v, err := st1.GetValue("g0"); err != nil {
		t.Fatal(err)
	} else if v.Value != nil {
		t.Fatal()
	}
}

func TestState_ExecuteUpdate3(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"g0": {rt{Leaf: LeafString}},
		}}},
	})
	st2, _ := st1.NewChild("Ch1", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"l0": {rt{Leaf: LeafString}},
		}}},
	}, MustParseBinding("g0"+BindingSepOut+"l0"))
	(&Action{ru{
		Action: "l0=test",
	}}).MustExecUpdate(st2)
	if v, err := st1.GetValue("g0"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}
}

func TestState_ExecuteUpdate__Native(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{
			MapOf: map[string]*Type{"g0": {rt{Leaf: LeafString}}},
		}},
		Value: nil,
	})
	_ = (&Action{ru{
		Native: func(si NativeInterface) error {
			return si.Set("g0", "test")
		},
	}}).Exec(st1)
	if v, err := st1.GetValue("g0"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}
}

func TestState_ExecuteUpdate__Reference(t *testing.T) {
	lib := &Library{
		Definitions: Definitions{Actions: map[string]*Action{"ud1": {ru{Action: "g0=test"}}}},
	}
	udt := &Action{ru{Reference: "ud1"}}
	if err := udt.Compile(lib); err != nil {
		t.Fatal(err)
	}
	st1, _ := NewMemState("State",
		&ValueHolder{
			Type: &Type{rt{MapOf: map[string]*Type{
				"g0": {rt{Leaf: LeafString}},
			}}},
		})
	udt.MustExecUpdate(st1)
	if v, err := st1.GetValue("g0"); err != nil {
		t.Fatal(err)
	} else if v.Value != "test" {
		t.Fatal()
	}
}

func TestState_ExecuteUpdate__Sequence(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"a": {rt{Leaf: LeafString}},
			"b": {rt{Leaf: LeafString}},
		}}},
	})
	(&Action{ru{
		Sequence: []*Action{
			{ru{Action: "a=1"}},
			{ru{Action: "b=2"}},
		},
	}}).MustExecUpdate(st1)
	if v, err := st1.GetValue("a"); err != nil {
		t.Fatal(err)
	} else if v.Value != "1" {
		t.Fatal()
	}
	if v, err := st1.GetValue("b"); err != nil {
		t.Fatal(err)
	} else if v.Value != "2" {
		t.Fatal()
	}
}

func TestState_ExecuteUpdate__Parallel(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"a": {rt{Leaf: LeafString}},
			"b": {rt{Leaf: LeafString}},
			"c": {rt{Leaf: LeafString}},
			"d": {rt{Leaf: LeafString}},
		}}},
	})
	(&Action{ru{
		Parallel: []*Action{
			{ru{Action: "a=1"}},
			{ru{Action: "b=2"}},
			{ru{Action: "c=3"}},
			{ru{Action: "d=4"}},
		},
	}}).MustExecUpdate(st1)
	if v, err := st1.GetValue("a"); err != nil {
		t.Fatal(err)
	} else if v.Value != "1" {
		t.Fatal()
	}
	if v, err := st1.GetValue("b"); err != nil {
		t.Fatal(err)
	} else if v.Value != "2" {
		t.Fatal()
	}
	if v, err := st1.GetValue("c"); err != nil {
		t.Fatal(err)
	} else if v.Value != "3" {
		t.Fatal()
	}
	if v, err := st1.GetValue("d"); err != nil {
		t.Fatal(err)
	} else if v.Value != "4" {
		t.Fatal()
	}
}

func TestState_ExecuteUpdate__If(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"a": {rt{Leaf: LeafString}},
			"b": {rt{Leaf: LeafString}},
		}}},
	})
	st1.SetValue("a", "1")
	(&Action{ru{
		If: &ActionIf{
			Condition: &Condition{rc{Condition: "a != 1"}},
			Then:      &Action{ru{Action: "b = 1"}},
			Else:      &Action{ru{Action: "b = 2"}},
		},
	}}).MustExecUpdate(st1)
	if v, err := st1.GetValue("b"); err != nil {
		t.Fatal(err)
	} else if v.Value != "2" {
		t.Fatal()
	}
}

func TestState_ExecuteUpdate__While(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"a": {rt{Leaf: LeafString}},
		}}},
	})
	st1.SetValue("a", "1")
	(&Action{ru{
		While: &ActionWhile{
			Condition: &Condition{rc{Condition: "a == 1"}},
			Do:        &Action{ru{Action: "a = 2"}},
		},
	}}).MustExecUpdate(st1)
	if v, err := st1.GetValue("a"); err != nil {
		t.Fatal(err)
	} else if v.Value != "2" {
		t.Fatal()
	}
}

func TestUpdate_MarshalJSON__Reference(t *testing.T) {
	udt := Action{ru{
		Reference: "test",
	}}
	udtStr, err := json.Marshal(udt)
	if err != nil {
		t.Fatal(err)
	}
	if string(udtStr) != `"$test"` {
		t.Fatal()
	}
	udt2 := Action{}
	err = json.Unmarshal(udtStr, &udt2)
	if err != nil {
		t.Fatal(err)
	}
	if udt2.Reference != "test" {
		t.Fatal()
	}
}

func TestUpdate_MarshalYAML__Reference(t *testing.T) {
	udt := Action{ru{
		Reference: "test",
	}}
	udtStr, err := yaml.Marshal(udt)
	if err != nil {
		t.Fatal(err)
	}
	if string(udtStr) != "$test\n" {
		t.Fatal()
	}
	udt2 := Action{}
	err = yaml.Unmarshal(udtStr, &udt2)
	if err != nil {
		t.Fatal(err)
	}
	if udt2.Reference != "test" {
		t.Fatal()
	}
}
