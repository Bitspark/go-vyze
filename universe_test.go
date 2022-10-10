package go_vyze

import (
	"bytes"
	"testing"
)

func TestParseUniverseObjectIdentifier__NoType(t *testing.T) {
	ri := ParseUniverseObjectIdentifier("base.name/target")
	if ri.Base != "base" {
		t.Fatal()
	}
	if ri.Name != "name" {
		t.Fatal()
	}
	if ri.Target != "target" {
		t.Fatal()
	}

	ri = ParseUniverseObjectIdentifier("base.name/")
	if ri.Base != "base" {
		t.Fatal()
	}
	if ri.Name != "name" {
		t.Fatal()
	}
	if ri.Target != "" {
		t.Fatal()
	}

	ri = ParseUniverseObjectIdentifier("base.name")
	if ri.Base != "base" {
		t.Fatal()
	}
	if ri.Name != "name" {
		t.Fatal()
	}
	if ri.Target != "base" {
		t.Fatal()
	}

	ri = ParseUniverseObjectIdentifier("name")
	if ri.Base != "" {
		t.Fatal()
	}
	if ri.Name != "name" {
		t.Fatal()
	}
	if ri.Target != "" {
		t.Fatal()
	}

	ri = ParseUniverseObjectIdentifier("name/test")
	if ri.Base != "" {
		t.Fatal()
	}
	if ri.Name != "name" {
		t.Fatal()
	}
	if ri.Target != "test" {
		t.Fatal()
	}
}

func TestUniverseObjectIdentifier_String(t *testing.T) {
	if (UniverseIdentifier{Base: "universe", Name: "name", Target: "target"}).String() != "universe.name/target" {
		t.Fatal()
	}
	if (UniverseIdentifier{Base: "universe", Name: "name", Target: ""}).String() != "universe.name/" {
		t.Fatal()
	}
	if (UniverseIdentifier{Base: "universe", Name: "name", Target: "universe"}).String() != "universe.name" {
		t.Fatal()
	}
	if (UniverseIdentifier{Base: "", Name: "name", Target: ""}).String() != "name" {
		t.Fatal()
	}
	if (UniverseIdentifier{Base: "", Name: "name", Target: "target"}).String() != "name/target" {
		t.Fatal()
	}
}

func TestUniverseObjectIdentifier_FQString(t *testing.T) {
	if (UniverseIdentifier{Base: "universe", Name: "name", Target: "target"}).Canonical("", "t2") != "universe.name/target" {
		t.Fatal()
	}
	if (UniverseIdentifier{Base: "universe", Name: "name", Target: ""}).Canonical("", "t2") != "universe.name/t2" {
		t.Fatal()
	}
	if (UniverseIdentifier{Base: "universe", Name: "name", Target: "universe"}).Canonical("", "t2") != "universe.name/universe" {
		t.Fatal()
	}
	if (UniverseIdentifier{Base: "", Name: "name", Target: ""}).Canonical("", "t2") != "t2.name/t2" {
		t.Fatal()
	}
	if (UniverseIdentifier{Base: "", Name: "name", Target: "target"}).Canonical("", "t2") != "target.name/target" {
		t.Fatal()
	}
	if (UniverseIdentifier{Base: "universe", Name: "name", Target: "target"}).Canonical("t3", "t2") != "universe.name/target" {
		t.Fatal()
	}
	if (UniverseIdentifier{Base: "universe", Name: "name", Target: ""}).Canonical("t3", "t2") != "universe.name/t2" {
		t.Fatal()
	}
	if (UniverseIdentifier{Base: "universe", Name: "name", Target: "universe"}).Canonical("t3", "t2") != "universe.name/universe" {
		t.Fatal()
	}
	if (UniverseIdentifier{Base: "", Name: "name", Target: ""}).Canonical("t3", "t2") != "t3.name/t2" {
		t.Fatal()
	}
	if (UniverseIdentifier{Base: "", Name: "name", Target: "target"}).Canonical("t3", "t2") != "t3.name/target" {
		t.Fatal()
	}
}

func TestUniverseFile_DumpLoadZ1(t *testing.T) {
	rf := Universe{
		Name: "Test123",
	}

	b := &bytes.Buffer{}
	if err := rf.Dump(b, true); err != nil {
		t.Fatal(err)
	}
	rf2 := Universe{}
	if err := rf2.Load(b, true); err != nil {
		t.Fatal(rf2)
	}

	if rf.Name != rf2.Name {
		t.Fatal()
	}
}

func TestUniverseFile_DumpLoadZ2(t *testing.T) {
	rf := Universe{
		Identifiers: []UniverseIdentifier{
			{
				Base:   "a",
				Name:   "b",
				Target: "c",
			},
		},
		Models: []UniverseObjectInfo{
			{
				Mapping: UniverseIdentifier{
					Base:   "a",
					Name:   "b",
					Target: "c",
				},
				Type: "d",
			},
		},
	}

	b := &bytes.Buffer{}
	if err := rf.Dump(b, true); err != nil {
		t.Fatal(err)
	}
	rf2 := Universe{}
	if err := rf2.Load(b, true); err != nil {
		t.Fatal(rf2)
	}

	if len(rf2.Identifiers) != 1 || !rf.Identifiers[0].Equals(rf2.Identifiers[0]) {
		t.Fatal()
	}
	if len(rf2.Models) != 1 || rf.Models[0].Type != rf2.Models[0].Type {
		t.Fatal()
	}
}

func TestUniverseFile_DumpLoadZ3(t *testing.T) {
	rf := Universe{
		Relations: []UniverseRelation{
			{
				Relation: UniverseIdentifier{
					Base:   "a1",
					Name:   "b1",
					Target: "c1",
				},
				Origin: UniverseIdentifier{
					Base:   "a2",
					Name:   "b2",
					Target: "c2",
				},
				Target: UniverseIdentifier{
					Base:   "a3",
					Name:   "b3",
					Target: "c3",
				},
			},
		},
	}

	b := &bytes.Buffer{}
	if err := rf.Dump(b, true); err != nil {
		t.Fatal(err)
	}
	rf2 := Universe{}
	if err := rf2.Load(b, true); err != nil {
		t.Fatal(rf2)
	}

	if len(rf2.Relations) != 1 || !rf.Relations[0].Relation.Equals(rf2.Relations[0].Relation) {
		t.Fatal()
	}
	if !rf.Relations[0].Origin.Equals(rf2.Relations[0].Origin) {
		t.Fatal()
	}
	if !rf.Relations[0].Target.Equals(rf2.Relations[0].Target) {
		t.Fatal()
	}
}

func TestUniverseFile_DumpLoadZ4(t *testing.T) {
	rf := Universe{
		Abstractions: []UniverseAbstraction{
			{
				Abstract: UniverseIdentifier{
					Base:   "a1",
					Name:   "b1",
					Target: "c1",
				},
				Special: UniverseIdentifier{
					Base:   "a2",
					Name:   "b2",
					Target: "c2",
				},
			},
		},
	}

	b := &bytes.Buffer{}
	if err := rf.Dump(b, true); err != nil {
		t.Fatal(err)
	}
	rf2 := Universe{}
	if err := rf2.Load(b, true); err != nil {
		t.Fatal(rf2)
	}

	if len(rf2.Abstractions) != 1 || !rf.Abstractions[0].Abstract.Equals(rf2.Abstractions[0].Abstract) {
		t.Fatal()
	}
	if !rf.Abstractions[0].Special.Equals(rf2.Abstractions[0].Special) {
		t.Fatal()
	}
}

func TestResolveResourceSchema(t *testing.T) {
	rs := ResourceSchema{}
	rs.AddRelationField("model", "base.object#relation/", FieldTypeID, FormatTypeHex, NewPrimitiveMapping())
	univ := Universe{}
	_ = univ.LoadFile("./test/test1.vyu")
	relModel := univ.GetModel("base.object#relation/", "")
	if relModel == nil {
		t.Fatal()
	}
	err := rs.Resolve(univ, "")
	if err != nil {
		t.Fatal(err)
	}
	if *rs.Fields[0].Relation != relModel.ObjectID.Hex() {
		t.Fatal()
	}
}

func TestResolveResourceSchema__Reverse(t *testing.T) {
	rs := ResourceSchema{}
	rs.AddRelationField("model", "-base.object#relation/", FieldTypeID, FormatTypeHex, NewPrimitiveMapping())
	univ := Universe{}
	_ = univ.LoadFile("./test/test1.vyu")
	relModel := univ.GetModel("base.object#relation/", "")
	if relModel == nil {
		t.Fatal()
	}
	err := rs.Resolve(univ, "")
	if err != nil {
		t.Fatal(err)
	}
	if *rs.Fields[0].Relation != "-"+relModel.ObjectID.Hex() {
		t.Fatal()
	}
}

func TestUniverse_GetModel(t *testing.T) {
	rs := ResourceSchema{}
	rs.AddRelationField("model", "base.object#relation/", FieldTypeID, FormatTypeHex, NewPrimitiveMapping())
	univ := Universe{}
	_ = univ.LoadFile("./test/vision.vyu")
	relModel := univ.GetModel("image#camera", "vision")
	if relModel == nil {
		t.Fatal()
	}
	camModel := univ.GetTarget(relModel.Mapping)
	if camModel == nil {
		t.Fatal()
	}
}
