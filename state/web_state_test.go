package state

import (
	"sync"
	"testing"
	"time"
)

func TestNewWebState1(t *testing.T) {
	ms, err := NewWebsocketServer("localhost:12345", "conn1")
	if err != nil {
		t.Fatal(err)
	}
	if ms == nil {
		t.Fatal()
	}

	st, _ := NewMemState("Root", &ValueHolder{
		Type: &Type{rt{
			MapOf: map[string]*Type{
				"z": {rt{Leaf: LeafString}},
				"y": {rt{Leaf: LeafString}},
				"x": {rt{Leaf: LeafString}},
			},
		}},
		mux: &sync.Mutex{},
	})

	ms.AddFactory(st, "F1", "", MustParseBinding("z -> a", "y <- b", "x <=> c"), &ValueHolder{
		Type: &Type{rt{
			MapOf: map[string]*Type{
				"a": {rt{Leaf: LeafString}},
				"b": {rt{Leaf: LeafString}},
				"c": {rt{Leaf: LeafString}},
			},
		}},
		mux: &sync.Mutex{},
	})
	go ms.Start()

	mc, err := NewWebState("localhost:12345", "conn1", "F1", "")
	if err != nil {
		t.Fatal(err)
	}
	if mc == nil {
		t.Fatal()
	}
	_, err = mc.GetValue("fgdg")
	if err == nil {
		t.Fatal()
	}
	vh, err := mc.GetValue("a")
	if err != nil {
		t.Fatal(err)
	}
	if vh.Value != nil {
		t.Fatal()
	}

	err = st.SetValue("z", "test1")
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(500 * time.Millisecond)

	vh, err = mc.GetValue("a")
	if err != nil {
		t.Fatal(err)
	}
	if vh.Value != "test1" {
		t.Fatal()
	}

	err = mc.SetValue("b", "test2")
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(500 * time.Millisecond)

	vh, err = mc.GetValue("b")
	if err != nil {
		t.Fatal(err)
	}
	if vh.Value != "test2" {
		t.Fatal()
	}

	err = mc.SetValue("c", "test3")
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(500 * time.Millisecond)

	vh, err = st.GetValue("x")
	if err != nil {
		t.Fatal(err)
	}
	if vh.Value != "test3" {
		t.Fatal()
	}

	err = st.SetValue("x", "test4")
	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(500 * time.Millisecond)

	vh, err = mc.GetValue("c")
	if err != nil {
		t.Fatal(err)
	}
	if vh.Value != "test4" {
		t.Fatal()
	}

	err = mc.Sync("")
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewWebState2(t *testing.T) {
	ms, err := NewWebsocketServer("localhost:12345", "conn2")
	if err != nil {
		t.Fatal(err)
	}
	if ms == nil {
		t.Fatal()
	}

	st, _ := NewMemState("Root", &ValueHolder{
		Type: &Type{rt{
			MapOf: map[string]*Type{
				"z": {rt{Leaf: LeafString}},
				"y": {rt{Leaf: LeafString}},
			},
		}},
		mux: &sync.Mutex{},
	})
	ms.AddFactory(st, "F1", "", MustParseBinding("z -> a", "y <- b"), &ValueHolder{
		Type: &Type{rt{
			MapOf: map[string]*Type{
				"a": {rt{Leaf: LeafString}},
				"b": {rt{Leaf: LeafString}},
			},
		}},
		mux: &sync.Mutex{},
	})
	go ms.Start()

	mc, _ := NewWebState("localhost:12345", "conn2", "F1", "")

	mst := mc.Mem()

	mst.AddUpdate(&Condition{rc{Condition: "a == test1", compiled: true}}, &Action{ru{Action: "b = test1", compiled: true}})

	st.SetValue("z", "test1")

	time.Sleep(500 * time.Millisecond)

	val, _ := mst.GetValue("a")
	if val.Value != "test1" {
		t.Fatal()
	}

	val, _ = mst.GetValue("b")
	if val.Value != "test1" {
		t.Fatal()
	}

	val, _ = st.GetValue("y")
	if val.Value != "test1" {
		t.Fatal()
	}
}

func TestNewWebState3(t *testing.T) {
	ms, err := NewWebsocketServer("localhost:12345", "conn3")
	if err != nil {
		t.Fatal(err)
	}
	if ms == nil {
		t.Fatal()
	}

	st, _ := NewMemState("Root", &ValueHolder{
		Type: &Type{rt{
			MapOf: map[string]*Type{
				"r0": {rt{Leaf: LeafString}},
				"r1": {rt{Leaf: LeafString}},
				"r2": {rt{Leaf: LeafString}},
			},
		}},
		mux: &sync.Mutex{},
	})
	ms.AddFactory(st, "F1", "", MustParseBinding("r0 -> a0", "r1 <- a0"), &ValueHolder{
		Type: &Type{rt{
			MapOf: map[string]*Type{
				"a0": {rt{Leaf: LeafString}},
			},
		}},
		mux: &sync.Mutex{},
	})
	ms.AddFactory(st, "F2", "", MustParseBinding("r1 -> b0", "r2 <- b0"), &ValueHolder{
		Type: &Type{rt{
			MapOf: map[string]*Type{
				"b0": {rt{Leaf: LeafString}},
			},
		}},
		mux: &sync.Mutex{},
	})
	go ms.Start()

	_, _ = NewWebState("localhost:12345", "conn3", "F1", "")
	_, _ = NewWebState("localhost:12345", "conn3", "F2", "")

	st.SetValue("r0", "test1")

	time.Sleep(500 * time.Millisecond)

	val, _ := st.GetValue("r2")
	if val.Value != "test1" {
		t.Fatal()
	}
}
