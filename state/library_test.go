package state

import (
	"testing"
)

func TestLibraryFromDir1(t *testing.T) {
	lib1 := &Library{}
	err := lib1.LoadFromDir("./test/prog1")
	if err != nil {
		t.Fatal(err)
	}
	if lib1 == nil {
		t.Fatal()
	}

	assistantLib, err := lib1.GetLibrary("assistant")
	if err != nil {
		t.Fatal(err)
	}
	if assistantLib == nil {
		t.Fatal()
	}

	// Types

	assistantType, err := lib1.GetType("assistant.assistant")
	if err != nil {
		t.Fatal(err)
	}
	if assistantType == nil {
		t.Fatal()
	}
	if assistantType.MapOf["name"].Leaf != LeafString {
		t.Fatal()
	}

	assistantType, err = assistantLib.GetType(".assistant.assistant")
	if err != nil {
		t.Fatal(err)
	}
	if assistantType == nil {
		t.Fatal()
	}
	if assistantType.MapOf["name"].Leaf != LeafString {
		t.Fatal()
	}

	// Conditions

	assistantCond, err := lib1.GetCondition("assistant.cond1")
	if err != nil {
		t.Fatal(err)
	}
	if assistantCond == nil {
		t.Fatal()
	}
	if assistantCond.Condition != "a == b" {
		t.Fatal()
	}

	// Update

	assistantUd, err := lib1.GetAction("assistant.ud1")
	if err != nil {
		t.Fatal(err)
	}
	if assistantUd == nil {
		t.Fatal()
	}
	if assistantUd.Action != "a = 1" {
		t.Fatal()
	}
}

func TestLibraryFromDir2(t *testing.T) {
	lib := &Library{}
	err := lib.LoadFromDir("./test/prog2")
	if err != nil {
		t.Fatal(err)
	}
	if lib == nil {
		t.Fatal()
	}

	st, err := NewMemState("Prog2", &ValueHolder{
		Type: lib.MustGetType("prog"),
	})
	if err != nil {
		t.Fatal(err)
	}

	st.Actions = []*StateAction{
		{
			Condition: nil,
			Update:    lib.MustGetAction("prog"),
		},
	}

	if err := st.SetValue("a", "1"); err != nil {
		t.Fatal(err)
	}

	val, err := st.GetValue("b")
	if err != nil {
		t.Fatal()
	}
	if val.Value != true {
		t.Fatal()
	}
}

func TestLibraryFromDir3(t *testing.T) {
	lib := &Library{}
	err := lib.LoadFromDir("./test/prog3")
	if err != nil {
		t.Fatal(err)
	}
	if lib == nil {
		t.Fatal()
	}

	if err := lib.Compile(); err != nil {
		t.Fatal(err)
	}

	st, err := NewMemState("Prog", &ValueHolder{
		Type: lib.MustGetType("prog"),
	})
	if err != nil {
		t.Fatal(err)
	}

	st.Actions = []*StateAction{
		{
			Update: lib.MustGetAction("prog"),
		},
	}

	if err := st.SetValue("a1", "1"); err != nil {
		t.Fatal(err)
	}

	val, err := st.GetValue("b1")
	if err != nil {
		t.Fatal()
	}
	if val.Value != "2" {
		t.Fatal()
	}
}
