package system

import (
	"os"
	"testing"
)

func TestNewStructLibrary(t *testing.T) {
	lb := NewLibrary()
	if lb.Interfaces == nil {
		t.Fatal()
	}
}

func TestStructLibrary_AddSchema(t *testing.T) {
	lb := NewLibrary()
	ts1 := Interface{Type: SchemaTypePrimitive, Primitive: &PrimitiveInterface{Value: PrimitiveTypeString}}
	ts2 := Interface{Type: SchemaTypePrimitive, Primitive: &PrimitiveInterface{Value: PrimitiveTypeInteger}}
	lb.AddSchema(NamedInterface{Name: "test1", Schema: ts1})
	lb.AddSchema(NamedInterface{Name: "test2", Schema: ts2})
	if len(lb.Interfaces) != 2 {
		t.Fatal()
	}
}

func TestStructLibrary_GetSchema1(t *testing.T) {
	lb := NewLibrary()
	ts1 := Interface{Type: SchemaTypePrimitive, Primitive: &PrimitiveInterface{Value: PrimitiveTypeString}}
	lb.AddSchema(NamedInterface{Name: "test1", Schema: ts1})
	if lb.GetSchema("test1") == nil {
		t.Fatal()
	}
	if lb.GetSchema("test1").Schema.Type != SchemaTypePrimitive {
		t.Fatal()
	}
}

func TestStructLibrary_GetSchema2(t *testing.T) {
	lb := NewLibrary()
	if lb.GetSchema("test1") != nil {
		t.Fatal()
	}
}

func TestStructLibrary_Load(t *testing.T) {
	lb := Library{}
	lb.Load("does_not_exist.yml")
	if lb.Interfaces == nil {
		t.Fatal()
	}
}

func TestStructLibrary_Save(t *testing.T) {
	lb := NewLibrary()
	ts := Interface{Type: SchemaTypePrimitive, Primitive: &PrimitiveInterface{Value: PrimitiveTypeString}}
	lb.AddSchema(NamedInterface{Name: "test", Schema: ts})
	lb.Save("test.yml")
	defer os.Remove("test.yml")
	lb2 := Library{}
	lb2.Load("test.yml")
	if lb2.GetSchema("test") == nil {
		t.Fatal()
	}
}
