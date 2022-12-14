package lang

import (
	"github.com/Bitspark/go-vyze/system"
	"os"
	"testing"
)

func TestLibrary_File1(t *testing.T) {
	univ := &system.Universe{}
	err := univ.LoadFile("./test/vergleichsportal.vyu")
	if err != nil {
		t.Fatal(err)
	}

	l := NewLibrary(univ)
	source, _ := os.ReadFile("./test/queries1.vy")
	errs := l.Parse(string(source))
	if errs != nil {
		t.Fatal(errs[0].Err)
	}

	if len(l.Pipes) != 8 {
		t.Fatal()
	}
}

func TestLibrary_Pipe1(t *testing.T) {
	l := NewLibrary(&system.Universe{})
	pipe, errs := l.ParsePipe("on test -> @id")
	if len(errs) != 1 {
		t.Fatal(errs)
	}
	if pipe == nil || pipe.Node == nil {
		t.Fatal()
	}

	if pipe.Node.Type != system.NodeTypeValue {
		t.Fatal()
	}
	if pipe.Node.Value.Field != system.FieldTypeID {
		t.Fatal()
	}
	if pipe.Node.Value.Format != system.FormatTypeHex {
		t.Fatal()
	}
}

func TestLibrary_Pipe2(t *testing.T) {
	l := NewLibrary(&system.Universe{})
	pipe, errs := l.ParsePipe("on test -> {@id, @name}")
	if len(errs) != 1 {
		t.Fatal(errs)
	}
	if pipe == nil || pipe.Node == nil {
		t.Fatal()
	}

	if pipe.Node.Type != system.NodeTypeMap {
		t.Fatal()
	}
	if len(pipe.Node.Map.Entries) != 2 {
		t.Fatal()
	}

	e1 := pipe.Node.Map.Entries[0]
	if e1.Name != "id" {
		t.Fatal()
	}
	if e1.Node.Type != system.NodeTypeValue || e1.Node.Value.Field != system.FieldTypeID || e1.Node.Value.Format != system.FormatTypeHex {
		t.Fatal()
	}

	e2 := pipe.Node.Map.Entries[1]
	if e2.Name != "name" {
		t.Fatal()
	}
	if e2.Node.Type != system.NodeTypeValue || e2.Node.Value.Field != system.FieldTypeName || e2.Node.Value.Format != system.FormatTypeString {
		t.Fatal()
	}
}

func TestLibrary_Pipe3(t *testing.T) {
	l := NewLibrary(&system.Universe{})
	pipe, errs := l.ParsePipe("on test -> {id2: @id, name2: @name}")
	if len(errs) != 1 {
		t.Fatal(errs)
	}
	if pipe == nil || pipe.Node == nil {
		t.Fatal()
	}

	if pipe.Node.Type != system.NodeTypeMap {
		t.Fatal()
	}
	if len(pipe.Node.Map.Entries) != 2 {
		t.Fatal()
	}

	e1 := pipe.Node.Map.Entries[0]
	if e1.Name != "id2" {
		t.Fatal()
	}
	if e1.Node.Type != system.NodeTypeValue || e1.Node.Value.Field != system.FieldTypeID || e1.Node.Value.Format != system.FormatTypeHex {
		t.Fatal()
	}

	e2 := pipe.Node.Map.Entries[1]
	if e2.Name != "name2" {
		t.Fatal()
	}
	if e2.Node.Type != system.NodeTypeValue || e2.Node.Value.Field != system.FieldTypeName || e2.Node.Value.Format != system.FormatTypeString {
		t.Fatal()
	}
}

func TestLibrary_Pipe4(t *testing.T) {
	l := NewLibrary(&system.Universe{})
	pipe, errs := l.ParsePipe("on test -> {@id, field1 -> {@id, @name, createdAt: @created}}")
	if len(errs) != 1 {
		t.Fatal(errs)
	}
	if pipe == nil || pipe.Node == nil {
		t.Fatal(errs)
	}

	if pipe.Node.Type != system.NodeTypeMap {
		t.Fatal()
	}
	if len(pipe.Node.Map.Entries) != 2 {
		t.Fatal()
	}

	e1 := pipe.Node.Map.Entries[0]
	if e1.Name != "id" {
		t.Fatal()
	}
	if e1.Node.Type != system.NodeTypeValue || e1.Node.Value.Field != system.FieldTypeID || e1.Node.Value.Format != system.FormatTypeHex {
		t.Fatal()
	}

	e2 := pipe.Node.Map.Entries[1]
	if e2.Name != "field1" {
		t.Fatal()
	}
	if e2.Node.Type != system.NodeTypeRelation {
		t.Fatal()
	}
	if e2.Node.Relation.Type != system.EnvironmentTypePrimitive {
		t.Fatal()
	}
	if e2.Node.Relation.Node.Type != system.NodeTypeMap {
		t.Fatal()
	}
	if len(e2.Node.Relation.Node.Map.Entries) != 3 {
		t.Fatal()
	}
}

func TestLibrary_Pipe5(t *testing.T) {
	l := NewLibrary(&system.Universe{})
	pipe, errs := l.ParsePipe("on base.object/ -> {@id, fieldNew: base.object#field1/ -> [] {@id, @name, createdAt: @created}}")
	if len(errs) != 1 {
		t.Fatal(errs)
	}
	if pipe == nil || pipe.Node == nil {
		t.Fatal()
	}

	if pipe.Node.Type != system.NodeTypeMap {
		t.Fatal()
	}
	if len(pipe.Node.Map.Entries) != 2 {
		t.Fatal()
	}

	e1 := pipe.Node.Map.Entries[0]
	if e1.Name != "id" {
		t.Fatal()
	}
	if e1.Node.Type != system.NodeTypeValue || e1.Node.Value.Field != system.FieldTypeID || e1.Node.Value.Format != system.FormatTypeHex {
		t.Fatal()
	}

	e2 := pipe.Node.Map.Entries[1]
	if e2.Name != "fieldNew" {
		t.Fatal()
	}
	if e2.Node.Type != system.NodeTypeRelation {
		t.Fatal()
	}
	if e2.Node.Relation.Type != system.EnvironmentTypeList {
		t.Fatal()
	}
	if e2.Node.Relation.Relation != "" {
		t.Fatal()
	}
	if e2.Node.Relation.Reverse {
		t.Fatal()
	}
	if e2.Node.Relation.Node.Type != system.NodeTypeList {
		t.Fatal(errs)
	}
	if e2.Node.Relation.Node.List.Entry.Type != system.NodeTypeMap {
		t.Fatal(errs)
	}
	if len(e2.Node.Relation.Node.List.Entry.Map.Entries) != 3 {
		t.Fatal()
	}

	t.Log(l.models.String())
}

func TestLibrary_Pipe6(t *testing.T) {
	l := NewLibrary(&system.Universe{})
	pipe, errs := l.ParsePipe("on base.object/ -> {@id, fieldNew: <- base.object2#field2/ [] {@id, @name, createdAt: @created}}")
	if len(errs) != 1 {
		t.Fatal(errs)
	}
	if pipe == nil || pipe.Node == nil {
		t.Fatal()
	}

	if pipe.Node.Type != system.NodeTypeMap {
		t.Fatal()
	}
	if len(pipe.Node.Map.Entries) != 2 {
		t.Fatal()
	}

	e1 := pipe.Node.Map.Entries[0]
	if e1.Name != "id" {
		t.Fatal()
	}
	if e1.Node.Type != system.NodeTypeValue || e1.Node.Value.Field != system.FieldTypeID || e1.Node.Value.Format != system.FormatTypeHex {
		t.Fatal()
	}

	e2 := pipe.Node.Map.Entries[1]
	if e2.Name != "fieldNew" {
		t.Fatal()
	}
	if e2.Node.Type != system.NodeTypeRelation {
		t.Fatal()
	}
	if e2.Node.Relation.Type != system.EnvironmentTypeList {
		t.Fatal()
	}
	if e2.Node.Relation.Relation != "" {
		t.Fatal()
	}
	if !e2.Node.Relation.Reverse {
		t.Fatal()
	}
	if e2.Node.Relation.Node.Type != system.NodeTypeList {
		t.Fatal(errs)
	}
	if e2.Node.Relation.Node.List.Entry.Type != system.NodeTypeMap {
		t.Fatal(errs)
	}
	if len(e2.Node.Relation.Node.List.Entry.Map.Entries) != 3 {
		t.Fatal()
	}

	t.Log(l.models.String())
}
