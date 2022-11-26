package state

import "testing"

func TestValueType_Accepts__Incompatible1(t *testing.T) {
	v1 := MustParseType(`["string"]`)
	v2 := MustParseType(`"string"`)
	if ok, err := v1.Accepts(v2); ok || err == nil {
		t.Fatal()
	}
	if ok, err := v2.Accepts(v1); ok || err == nil {
		t.Fatal()
	}
}

func TestValueType_Accepts__Incompatible2(t *testing.T) {
	v1 := MustParseType(`{"a": "string"}`)
	v2 := MustParseType(`"string"`)
	if ok, err := v1.Accepts(v2); ok || err == nil {
		t.Fatal()
	}
	if ok, err := v2.Accepts(v1); ok || err == nil {
		t.Fatal()
	}
}

func TestValueType_Accepts__Incompatible3(t *testing.T) {
	v1 := MustParseType(`{"a": "string"}`)
	v2 := MustParseType(`["string"]`)
	if ok, err := v1.Accepts(v2); ok || err == nil {
		t.Fatal()
	}
	if ok, err := v2.Accepts(v1); ok || err == nil {
		t.Fatal()
	}
}

func TestValueType_Accepts__Leaf1(t *testing.T) {
	v1 := MustParseType(`"string"`)
	v2 := MustParseType(`"integer"`)
	if ok, err := v1.Accepts(v2); ok || err == nil {
		t.Fatal()
	}
	if ok, err := v2.Accepts(v1); ok || err == nil {
		t.Fatal()
	}
}

func TestValueType_Accepts__Leaf2(t *testing.T) {
	v1 := MustParseType(`"integer"`)
	v2 := MustParseType(`"integer"`)
	if ok, err := v1.Accepts(v2); !ok || err != nil {
		t.Fatal(err)
	}
	if ok, err := v2.Accepts(v1); !ok || err != nil {
		t.Fatal(err)
	}
}

func TestValueType_Accepts__Map1(t *testing.T) {
	v1 := MustParseType(`{"a": "string"}`)
	v2 := MustParseType(`{"a": "string", "b": "string"}`)
	if ok, err := v1.Accepts(v2); ok || err == nil {
		t.Fatal()
	}
	if ok, err := v2.Accepts(v1); !ok || err != nil {
		t.Fatal(err)
	}
}

func TestValueType_Accepts__Map2(t *testing.T) {
	v1 := MustParseType(`{"a": "raw", "b": "boolean"}`)
	v2 := MustParseType(`{"a": "string", "b": "boolean"}`)
	if ok, err := v1.Accepts(v2); ok || err == nil {
		t.Fatal()
	}
	if ok, err := v2.Accepts(v1); ok || err == nil {
		t.Fatal()
	}
}

func TestValueType_Accepts__List1(t *testing.T) {
	v1 := MustParseType(`["string"]`)
	v2 := MustParseType(`["integer"]`)
	if ok, err := v1.Accepts(v2); ok || err == nil {
		t.Fatal()
	}
	if ok, err := v2.Accepts(v1); ok || err == nil {
		t.Fatal()
	}
}

func TestValueType_Accepts__List2(t *testing.T) {
	v1 := MustParseType(`["float"]`)
	v2 := MustParseType(`["float"]`)
	if ok, err := v1.Accepts(v2); !ok || err != nil {
		t.Fatal(err)
	}
	if ok, err := v2.Accepts(v1); !ok || err != nil {
		t.Fatal(err)
	}
}
