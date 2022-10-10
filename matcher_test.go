package go_vyze

import (
	"testing"
)

func TestNewStructMatcher(t *testing.T) {
	m := NewInterfaceMatcher(nil, nil)
	if m.Template() == nil {
		t.Fatal()
	}
	if m.Generics() == nil {
		t.Fatal()
	}
}

func TestStructMatcher_Match(t *testing.T) {
	g1 := Interface{
		Type: SchemaTypeGeneric,
		Generic: &GenericInterface{
			Name: "A",
		},
	}
	s2 := Interface{
		Type: SchemaTypePrimitive,
		Primitive: &PrimitiveInterface{
			Value: PrimitiveTypeFloat,
		},
	}
	m := NewInterfaceMatcher([]GenericSlot{{Name: "A"}}, nil)
	err := m.Match(g1, s2)
	if err != nil {
		t.Fatal(err)
	}
	gv := m.Generics()
	if gvs, ok := gv["A"]; !ok {
		t.Fatal()
	} else if gvs.Type != SchemaTypePrimitive {
		t.Fatal()
	} else if gvs.Primitive.Value != PrimitiveTypeFloat {
		t.Fatal()
	}
}

func TestStructMatcher_MatchTemplate1(t *testing.T) {
	g1 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "{a}",
					Expansion: &TemplateExpansion{
						Source: "as",
						Target: "a",
					},
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeBoolean,
						},
					},
				},
			},
		},
	}
	s2 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "a",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeBoolean,
						},
					},
				},
				{
					Key: "b",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeBoolean,
						},
					},
				},
			},
		},
	}

	m := NewInterfaceMatcher(nil, []TemplateSlot{{Type: TemplateTypeList, Name: "as"}})
	err := m.MatchTemplate(g1, s2)
	if err != nil {
		t.Fatal(err)
	}
	tv := m.Template()
	if _, ok := tv["as"]; !ok {
		t.Fatal()
	}
	if tv["as"].Type != TemplateTypeList {
		t.Fatal()
	}
	if len(tv["as"].Value.([]string)) != 2 {
		t.Fatal()
	}

	s22, err := g1.ApplyTemplate(m.Template())
	if err != nil {
		t.Fatal(err)
	}
	if !s2.Equals(s22) {
		t.Fatal()
	}
}

func TestStructMatcher_MatchTemplate2(t *testing.T) {
	g1 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "{a}",
					Expansion: &TemplateExpansion{
						Source: "as",
						Target: "a",
					},
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeBoolean,
						},
					},
				},
				{
					Key: "{b}",
					Expansion: &TemplateExpansion{
						Source: "bs",
						Target: "b",
					},
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeInteger,
						},
					},
				},
			},
		},
	}
	s2 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "a",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeBoolean,
						},
					},
				},
				{
					Key: "b",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeBoolean,
						},
					},
				},
				{
					Key: "c",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeInteger,
						},
					},
				},
			},
		},
	}

	m := NewInterfaceMatcher(nil, []TemplateSlot{{Type: TemplateTypeList, Name: "as"}, {Type: TemplateTypeList, Name: "bs"}})
	err := m.MatchTemplate(g1, s2)
	if err != nil {
		t.Fatal(err)
	}
	tv := m.Template()
	if _, ok := tv["as"]; !ok {
		t.Fatal()
	}
	if tv["as"].Type != TemplateTypeList {
		t.Fatal()
	}
	if len(tv["as"].Value.([]string)) != 2 {
		t.Fatal()
	}
	if _, ok := tv["bs"]; !ok {
		t.Fatal()
	}
	if tv["bs"].Type != TemplateTypeList {
		t.Fatal()
	}
	if len(tv["bs"].Value.([]string)) != 1 {
		t.Fatal()
	}

	s22, err := g1.ApplyTemplate(m.Template())
	if err != nil {
		t.Fatal(err)
	}
	if !s2.Equals(s22) {
		t.Fatal()
	}
}

func TestStructMatcher_MatchTemplate3(t *testing.T) {
	g1 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "{a}",
					Expansion: &TemplateExpansion{
						Source: "as",
						Target: "a",
					},
					Schema: Interface{
						Type: SchemaTypeGeneric,
						Generic: &GenericInterface{
							Name: "{a}",
						},
					},
				},
			},
		},
	}
	s2 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "a",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeBoolean,
						},
					},
				},
				{
					Key: "b",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeBoolean,
						},
					},
				},
				{
					Key: "c",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeInteger,
						},
					},
				},
			},
		},
	}

	m := NewInterfaceMatcher(nil, []TemplateSlot{{Type: TemplateTypeList, Name: "as"}, {Type: TemplateTypeList, Name: "bs"}})
	err := m.MatchTemplate(g1, s2)
	if err != nil {
		t.Fatal(err)
	}
	tv := m.Template()
	if _, ok := tv["as"]; !ok {
		t.Fatal()
	}
	if tv["as"].Type != TemplateTypeList {
		t.Fatal()
	}
	if len(tv["as"].Value.([]string)) != 3 {
		t.Fatal()
	}

	m2 := NewInterfaceMatcher(nil, []TemplateSlot{{Type: TemplateTypeList, Name: "as"}, {Type: TemplateTypeList, Name: "bs"}})
	err = m2.Match(g1, s2)
	if err != nil {
		t.Fatal(err)
	}

	s22, _ := s2.ApplyTemplate(m2.ts)
	g22, _ := s22.ApplyGenerics(m2.Generics())
	if !g22.Equals(s2) {
		t.Fatal()
	}
}

func TestStructMatcher_MatchGenerics(t *testing.T) {
	s1 := Interface{
		Type: SchemaTypePrimitive,
		Primitive: &PrimitiveInterface{
			Value: PrimitiveTypeString,
		},
	}
	s2 := Interface{
		Type: SchemaTypeList,
		List: &ListInterface{
			Entry: Interface{
				Type: SchemaTypePrimitive,
				Primitive: &PrimitiveInterface{
					Value: PrimitiveTypeString,
				},
			},
		},
	}
	m := NewInterfaceMatcher(nil, nil)
	err := m.MatchGenerics(s1, s2)
	if err == nil {
		t.Fatal()
	}
}

func TestStructMatcher_MatchGenerics__Primitive(t *testing.T) {
	g1 := Interface{
		Type: SchemaTypeGeneric,
		Generic: &GenericInterface{
			Name: "A",
		},
	}
	s2 := Interface{
		Type: SchemaTypePrimitive,
		Primitive: &PrimitiveInterface{
			Value: PrimitiveTypeFloat,
		},
	}
	m := NewInterfaceMatcher([]GenericSlot{{Name: "A"}}, nil)
	err := m.MatchGenerics(g1, s2)
	if err != nil {
		t.Fatal(err)
	}
	gv := m.Generics()
	if gvs, ok := gv["A"]; !ok {
		t.Fatal()
	} else if gvs.Type != SchemaTypePrimitive {
		t.Fatal()
	} else if gvs.Primitive.Value != PrimitiveTypeFloat {
		t.Fatal()
	}
}

func TestStructMatcher_MatchGenerics__Map(t *testing.T) {
	g1 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "a",
					Schema: Interface{
						Type: SchemaTypeGeneric,
						Generic: &GenericInterface{
							Name: "A",
						},
					},
				},
				{
					Key: "b",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeString,
						},
					},
				},
			},
		},
	}
	s2 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "a",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeBoolean,
						},
					},
				},
				{
					Key: "b",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeString,
						},
					},
				},
			},
		},
	}
	m := NewInterfaceMatcher([]GenericSlot{{Name: "A"}}, nil)
	err := m.MatchGenerics(g1, s2)
	if err != nil {
		t.Fatal(err)
	}
	gv := m.Generics()
	if gvs, ok := gv["A"]; !ok {
		t.Fatal()
	} else if gvs.Type != SchemaTypePrimitive {
		t.Fatal()
	} else if gvs.Primitive.Value != PrimitiveTypeBoolean {
		t.Fatal()
	}
}

func TestStructMatcher_MatchGenerics__List(t *testing.T) {
	g1 := Interface{
		Type: SchemaTypeList,
		List: &ListInterface{
			Entry: Interface{
				Type: SchemaTypeGeneric,
				Generic: &GenericInterface{
					Name: "A",
				},
			},
		},
	}
	s2 := Interface{
		Type: SchemaTypeList,
		List: &ListInterface{
			Entry: Interface{
				Type: SchemaTypePrimitive,
				Primitive: &PrimitiveInterface{
					Value: PrimitiveTypeBoolean,
				},
			},
		},
	}
	m := NewInterfaceMatcher([]GenericSlot{{Name: "A"}}, nil)
	err := m.MatchGenerics(g1, s2)
	if err != nil {
		t.Fatal(err)
	}
	gv := m.Generics()
	if gvs, ok := gv["A"]; !ok {
		t.Fatal()
	} else if gvs.Type != SchemaTypePrimitive {
		t.Fatal()
	} else if gvs.Primitive.Value != PrimitiveTypeBoolean {
		t.Fatal()
	}
}
