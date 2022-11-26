package state

import "testing"

func TestState_PassCondition__Condition(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{"g0": {rt{Leaf: LeafString}}}}},
	})
	st1.SetValue("g0", "test")
	if !(Condition{rc{Condition: "g0"}}).MustEval(st1) {
		t.Fatal()
	}
	if (Condition{rc{Condition: "!g0"}}).MustEval(st1) {
		t.Fatal()
	}
	if !(Condition{rc{Condition: "g0 == test"}}).MustEval(st1) {
		t.Fatal()
	}
	if (Condition{rc{Condition: "g0 != test"}}).MustEval(st1) {
		t.Fatal()
	}
	if !(Condition{rc{Condition: "g0 != test2"}}).MustEval(st1) {
		t.Fatal()
	}
}

func TestState_PassCondition__Native(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{"g0": {rt{Leaf: LeafString}}}}},
	})
	st1.SetValue("g0", "test")
	cnd := Condition{rc{Native: func(si NativeInterface) (bool, error) {
		val, _ := si.Get("g0")
		return val == "test", nil
	}}}
	if !cnd.MustEval(st1) {
		t.Fatal()
	}
}

func TestState_PassCondition__Reference(t *testing.T) {
	lib := &Library{
		Definitions: Definitions{Conditions: map[string]*Condition{"cond1": {rc{Condition: "g0"}}}},
	}
	cond := Condition{rc{Reference: "cond1"}}
	if err := cond.Compile(lib); err != nil {
		t.Fatal(err)
	}
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{"g0": {rt{Leaf: LeafString}}}}},
	})
	st1.SetValue("g0", "test")
	if !cond.MustEval(st1) {
		t.Fatal()
	}
}

func TestState_PassCondition__Not(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{"g0": {rt{Leaf: LeafBoolean}}}}},
	})
	st1.SetValue("g0", true)
	if (Condition{rc{Not: &Condition{rc{Condition: "g0"}}}}).MustEval(st1) {
		t.Fatal()
	}
	if !(Condition{rc{Not: &Condition{rc{Condition: "!g0"}}}}).MustEval(st1) {
		t.Fatal()
	}
}

func TestState_PassCondition__And(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"g0": {rt{Leaf: LeafString}},
			"g1": {rt{Leaf: LeafString}},
		}}},
	})
	st1.SetValue("g0", "0")
	st1.SetValue("g1", "1")
	if !(Condition{rc{And: []*Condition{{rc{Condition: "g0"}}, {rc{Condition: "g1==1"}}}}}).MustEval(st1) {
		t.Fatal()
	}
	if (Condition{rc{And: []*Condition{{rc{Condition: "!g0"}}, {rc{Condition: "g1==1"}}}}}).MustEval(st1) {
		t.Fatal()
	}
	if (Condition{rc{And: []*Condition{{rc{Condition: "g0"}}, {rc{Condition: "g1!=1"}}}}}).MustEval(st1) {
		t.Fatal()
	}
	if (Condition{rc{And: []*Condition{{rc{Condition: "!g0"}}, {rc{Condition: "g1!=1"}}}}}).MustEval(st1) {
		t.Fatal()
	}
}

func TestState_PassCondition__Or(t *testing.T) {
	st1, _ := NewMemState("State", &ValueHolder{
		Type: &Type{rt{MapOf: map[string]*Type{
			"g0": {rt{Leaf: LeafString}},
			"g1": {rt{Leaf: LeafString}},
		}}},
	})
	st1.SetValue("g0", "0")
	st1.SetValue("g1", "1")
	if !(Condition{rc{Or: []*Condition{{rc{Condition: "g0"}}, {rc{Condition: "g1==1"}}}}}).MustEval(st1) {
		t.Fatal()
	}
	if !(Condition{rc{Or: []*Condition{{rc{Condition: "!g0"}}, {rc{Condition: "g1==1"}}}}}).MustEval(st1) {
		t.Fatal()
	}
	if !(Condition{rc{Or: []*Condition{{rc{Condition: "g0"}}, {rc{Condition: "g1!=1"}}}}}).MustEval(st1) {
		t.Fatal()
	}
	if (Condition{rc{Or: []*Condition{{rc{Condition: "!g0"}}, {rc{Condition: "g1!=1"}}}}}).MustEval(st1) {
		t.Fatal()
	}
}
