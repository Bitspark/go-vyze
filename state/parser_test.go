package state

import (
	"testing"
)

func TestParser_ParseTypes1(t *testing.T) {
	lib := NewLibrary()
	p := NewParser(lib)
	err := p.ParseFile("./test/lang/types1.vy")
	if err != nil {
		t.Fatal(err)
	}

	vt, err := lib.GetType("string1")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}

	if vt.Name != "string1" {
		t.Fatal()
	}
	if vt.Leaf != LeafString {
		t.Fatal()
	}
}

func TestParser_ParseTypes2(t *testing.T) {
	lib := NewLibrary()
	p := NewParser(lib)
	err := p.ParseFile("./test/lang/types2.vy")
	if err != nil {
		t.Fatal(err)
	}

	vt, err := lib.GetType("map1")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}

	if vt.Name != "map1" {
		t.Fatal()
	}
	if vt.MapOf == nil {
		t.Fatal()
	}
	if vt.MapOf["a"].Leaf != LeafString {
		t.Fatal()
	}
	if vt.MapOf["b"].Leaf != LeafInteger {
		t.Fatal()
	}

	vt, err = lib.GetType("list1")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}

	if vt.Name != "list1" {
		t.Fatal()
	}
	if vt.ListOf == nil {
		t.Fatal()
	}
	if vt.ListOf.Leaf != LeafBoolean {
		t.Fatal()
	}

	vt, err = lib.GetType("ref1")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}

	if vt.Name != "ref1" {
		t.Fatal()
	}
	if vt.Reference != "list1" {
		t.Fatal()
	}
}

func TestParser_ParseTypes3(t *testing.T) {
	lib := NewLibrary()
	p := NewParser(lib)
	err := p.ParseFile("./test/lang/types3.vy")
	if err != nil {
		t.Fatal(err)
	}

	vt, err := lib.GetType("complex1")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}
	if vt.Name != "complex1" {
		t.Fatal()
	}
	if vt.MapOf == nil {
		t.Fatal()
	}

	vt, err = lib.GetType("complex2")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}
	if vt.Name != "complex2" {
		t.Fatal()
	}
	if vt.ListOf == nil {
		t.Fatal()
	}
	if vt.ListOf.MapOf == nil {
		t.Fatal()
	}
}

func TestParser_ParseTypes4(t *testing.T) {
	lib := NewLibrary()
	p := NewParser(lib)
	err := p.ParseFile("./test/lang/types4.vy")
	if err != nil {
		t.Fatal(err)
	}

	vt, err := lib.GetType("string1")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}
	if vt.Name != "string1" {
		t.Fatal()
	}
	if vt.Leaf != LeafString {
		t.Fatal()
	}

	vt, err = lib.GetType("string2")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}
	if vt.Name != "string2" {
		t.Fatal()
	}
	if vt.Leaf != LeafString {
		t.Fatal()
	}
	if vt.Initial != "testA" {
		t.Fatal()
	}

	vt, err = lib.GetType("string3")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}
	if vt.Name != "string3" {
		t.Fatal()
	}
	if vt.Leaf != LeafString {
		t.Fatal()
	}
	if len(vt.Options) != 3 {
		t.Fatal()
	}
	if vt.Options[0] != "testB" || vt.Options[1] != "testB2" || vt.Options[2] != "testB3" {
		t.Fatal()
	}

	vt, err = lib.GetType("string4")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}
	if vt.Name != "string4" {
		t.Fatal()
	}
	if vt.Leaf != LeafString {
		t.Fatal()
	}
	if vt.Initial != "testC" {
		t.Fatal()
	}
	if len(vt.Options) != 2 {
		t.Fatal()
	}
	if vt.Options[0] != "testC" || vt.Options[1] != "testC2" {
		t.Fatal()
	}
}

func TestParser_ParseTypes5(t *testing.T) {
	lib := NewLibrary()
	p := NewParser(lib)
	err := p.ParseFile("./test/lang/types5.vy")
	if err != nil {
		t.Fatal(err)
	}

	vt, err := lib.GetType("bool1")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}
	if vt.Name != "bool1" {
		t.Fatal()
	}
	if vt.Leaf != LeafBoolean {
		t.Fatal()
	}

	vt, err = lib.GetType("bool2")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}
	if vt.Name != "bool2" {
		t.Fatal()
	}
	if vt.Leaf != LeafBoolean {
		t.Fatal()
	}
	if vt.Initial != true {
		t.Fatal()
	}

	vt, err = lib.GetType("bool3")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}
	if vt.Name != "bool3" {
		t.Fatal()
	}
	if vt.Leaf != LeafBoolean {
		t.Fatal()
	}
	if len(vt.Options) != 3 {
		t.Fatal()
	}
	if vt.Options[0] != true || vt.Options[1] != false || vt.Options[2] != true {
		t.Fatal()
	}

	vt, err = lib.GetType("bool4")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}
	if vt.Name != "bool4" {
		t.Fatal()
	}
	if vt.Leaf != LeafBoolean {
		t.Fatal()
	}
	if vt.Initial != true {
		t.Fatal()
	}
	if len(vt.Options) != 2 {
		t.Fatal()
	}
	if vt.Options[0] != true || vt.Options[1] != false {
		t.Fatal()
	}
}

func TestParser_ParseTypes6(t *testing.T) {
	lib := NewLibrary()
	p := NewParser(lib)
	err := p.ParseFile("./test/lang/types6.vy")
	if err != nil {
		t.Fatal(err)
	}

	vt, err := lib.GetType("map1")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}
	if vt.Name != "map1" {
		t.Fatal()
	}
	if vt.MapOf == nil {
		t.Fatal()
	}
	if vt.Initial == nil {
		t.Fatal()
	}
	if len(vt.Initial.(map[string]any)) != 2 {
		t.Fatal()
	}
	if vt.Initial.(map[string]any)["a"] != "test" {
		t.Fatal()
	}
	if vt.Initial.(map[string]any)["b"] != true {
		t.Fatal()
	}
	if len(vt.Options) != 2 {
		t.Fatal()
	}
	if len(vt.Options[0].(map[string]any)) != 2 || len(vt.Options[1].(map[string]any)) != 2 {
		t.Fatal()
	}
	if vt.Options[0].(map[string]any)["a"] != "test" {
		t.Fatal()
	}
	if vt.Options[0].(map[string]any)["b"] != true {
		t.Fatal()
	}
	if vt.Options[1].(map[string]any)["a"] != "test2" {
		t.Fatal()
	}
	if vt.Options[1].(map[string]any)["b"] != false {
		t.Fatal()
	}

	vt, err = lib.GetType("video")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}
	if vt.Name != "video" {
		t.Fatal()
	}
}

func TestParser_ParseAction1(t *testing.T) {
	lib := NewLibrary()
	p := NewParser(lib)
	err := p.ParseFile("./test/lang/action1.vy")
	if err != nil {
		t.Fatal(err)
	}

	va, err := lib.GetAction("action1")
	if err != nil {
		t.Fatal(err)
	}
	if va == nil {
		t.Fatal()
	}
	if va.Type.MapOf == nil || va.Type.MapOf["a"] == nil {
		t.Fatal()
	}

	va, err = lib.GetAction("action2")
	if err != nil {
		t.Fatal(err)
	}
	if va == nil {
		t.Fatal()
	}
	if va.Type.MapOf == nil || va.Type.MapOf["a"] == nil {
		t.Fatal()
	}
	if va.Sequence == nil {
		t.Fatal()
	}

	va, err = lib.GetAction("action3")
	if err != nil {
		t.Fatal(err)
	}
	if va == nil {
		t.Fatal()
	}
	if va.Parallel == nil {
		t.Fatal()
	}

	va, err = lib.GetAction("action4")
	if err != nil {
		t.Fatal(err)
	}
	if va == nil {
		t.Fatal()
	}
	if va.If == nil || va.If.Then == nil {
		t.Fatal()
	}

	va, err = lib.GetAction("action5")
	if err != nil {
		t.Fatal(err)
	}
	if va == nil {
		t.Fatal()
	}
	if va.If == nil || va.If.Then == nil || va.If.Else == nil {
		t.Fatal()
	}
}

func TestParser_ParseState1(t *testing.T) {
	lib := NewLibrary()
	p := NewParser(lib)
	err := p.ParseFile("./test/lang/state1.vy")
	if err != nil {
		t.Fatal(err)
	}

	vt, err := lib.GetType("state")
	if err != nil {
		t.Fatal(err)
	}
	if vt == nil {
		t.Fatal()
	}
}
