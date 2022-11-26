package system

import "testing"

func TestStructSchema_ApplyGenerics1(t *testing.T) {
	s := Interface{
		Type: SchemaTypeList,
		List: &ListInterface{
			Entry: Interface{
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
			},
		},
	}
	gv := GenericValuation{
		"A": Interface{
			Type: SchemaTypePrimitive,
			Primitive: &PrimitiveInterface{
				Value: PrimitiveTypeInteger,
			},
		},
	}
	s2, err := s.ApplyGenerics(gv)
	if err != nil {
		t.Fatal(err)
	}
	if s2.List.Entry.Map.Entries[0].Schema.Type != SchemaTypePrimitive {
		t.Fatal()
	}
	if s2.List.Entry.Map.Entries[0].Schema.Primitive.Value != PrimitiveTypeInteger {
		t.Fatal()
	}
}
