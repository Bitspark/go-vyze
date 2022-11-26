package system

import "testing"

func TestMapSchema_GetEntry1(t *testing.T) {
	s := MapInterface{
		Entries: []MapInterfaceEntry{
			{
				Key: "a",
				Schema: Interface{
					Type: SchemaTypePrimitive,
					Primitive: &PrimitiveInterface{
						Value: PrimitiveTypeString,
					},
				},
			},
		},
	}
	e, ok := s.GetEntry("a")
	if !ok {
		t.Fatal()
	}
	if e.Schema.Type != SchemaTypePrimitive {
		t.Fatal()
	}
}

func TestMapSchema_GetEntry2(t *testing.T) {
	s := MapInterface{
		Entries: []MapInterfaceEntry{
			{
				Key: "a",
				Schema: Interface{
					Type: SchemaTypePrimitive,
					Primitive: &PrimitiveInterface{
						Value: PrimitiveTypeString,
					},
				},
			},
		},
	}
	_, ok := s.GetEntry("b")
	if ok {
		t.Fatal()
	}
}

func TestSchema_Equals1(t *testing.T) {
	s1 := Interface{
		Type: SchemaTypePrimitive,
		Primitive: &PrimitiveInterface{
			Value: PrimitiveTypeInteger,
		},
	}
	s2 := Interface{
		Type: SchemaTypePrimitive,
		Primitive: &PrimitiveInterface{
			Value: PrimitiveTypeInteger,
		},
	}
	if !s1.Equals(s2) {
		t.Fatal()
	}
}

func TestSchema_Equals2(t *testing.T) {
	s1 := Interface{
		Type: SchemaTypePrimitive,
		Primitive: &PrimitiveInterface{
			Value: PrimitiveTypeInteger,
		},
	}
	s2 := Interface{
		Type: SchemaTypePrimitive,
		Primitive: &PrimitiveInterface{
			Value: PrimitiveTypeBoolean,
		},
	}
	if s1.Equals(s2) {
		t.Fatal()
	}
}

func TestSchema_Equals3(t *testing.T) {
	s1 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "a",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeString,
						},
					},
				},
				{
					Key: "b",
					Schema: Interface{
						Type: SchemaTypeList,
						List: &ListInterface{
							Entry: Interface{
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
	s2 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "a",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeString,
						},
					},
				},
				{
					Key: "b",
					Schema: Interface{
						Type: SchemaTypeList,
						List: &ListInterface{
							Entry: Interface{
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
	if !s1.Equals(s2) {
		t.Fatal()
	}
}

func TestSchema_Equals4(t *testing.T) {
	s1 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "a",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeString,
						},
					},
				},
				{
					Key: "b",
					Schema: Interface{
						Type: SchemaTypeList,
						List: &ListInterface{
							Entry: Interface{
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
	s2 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "a",
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
	if s1.Equals(s2) {
		t.Fatal()
	}
}

func TestSchema_Equals5(t *testing.T) {
	s1 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "a",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeString,
						},
					},
				},
				{
					Key: "b",
					Schema: Interface{
						Type: SchemaTypeList,
						List: &ListInterface{
							Entry: Interface{
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
	s2 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "a",
					Schema: Interface{
						Type: SchemaTypePrimitive,
						Primitive: &PrimitiveInterface{
							Value: PrimitiveTypeString,
						},
					},
				},
				{
					Key: "b",
					Schema: Interface{
						Type: SchemaTypeList,
						List: &ListInterface{
							Entry: Interface{
								Type: SchemaTypePrimitive,
								Primitive: &PrimitiveInterface{
									Value: PrimitiveTypeBoolean,
								},
							},
						},
					},
				},
			},
		},
	}
	if s1.Equals(s2) {
		t.Fatal()
	}
}

func TestSchema_Equals6(t *testing.T) {
	s1 := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "a",
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
							Value: PrimitiveTypeString,
						},
					},
				},
				{
					Key: "b",
					Schema: Interface{
						Type: SchemaTypeList,
						List: &ListInterface{
							Entry: Interface{
								Type: SchemaTypePrimitive,
								Primitive: &PrimitiveInterface{
									Value: PrimitiveTypeBoolean,
								},
							},
						},
					},
				},
			},
		},
	}
	if s1.Equals(s2) {
		t.Fatal()
	}
}

func TestSchema_JSONSchema_Named(t *testing.T) {
	schema := Interface{
		Type: SchemaTypeNamed,
		Named: &NamedInterface{
			Name: "String",
			Schema: Interface{
				Type: SchemaTypePrimitive,
				Primitive: &PrimitiveInterface{
					Value: PrimitiveTypeString,
				},
			},
		},
	}
	js, err := schema.JSONSchema(nil)
	if err != nil {
		t.Fatal(err)
	}
	if js["type"] != "string" {
		t.Fatal()
	}
}

func TestSchema_JSONSchema_Primitive(t *testing.T) {
	schema := Interface{
		Type: SchemaTypePrimitive,
		Primitive: &PrimitiveInterface{
			Value: PrimitiveTypeInteger,
		},
	}
	js, err := schema.JSONSchema(nil)
	if err != nil {
		t.Fatal(err)
	}
	if js["type"] != "integer" {
		t.Fatal()
	}
}

func TestSchema_JSONSchema_List(t *testing.T) {
	schema := Interface{
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
	js, err := schema.JSONSchema(nil)
	if err != nil {
		t.Fatal(err)
	}
	if js["type"] != "array" {
		t.Fatal()
	}
	if js["items"].(map[string]any)["type"] != "string" {
		t.Fatal()
	}
}

func TestSchema_JSONSchema_Map(t *testing.T) {
	schema := Interface{
		Type: SchemaTypeMap,
		Map: &MapInterface{
			Entries: []MapInterfaceEntry{
				{
					Key: "key",
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
	js, err := schema.JSONSchema(nil)
	if err != nil {
		t.Fatal(err)
	}
	if js["type"] != "object" {
		t.Fatal()
	}
	if js["properties"].(map[string]any)["key"].(map[string]any)["type"] != "string" {
		t.Fatal()
	}
}

func TestNode_JSONSchema_ID(t *testing.T) {
	entryNode := EndpointNode{
		Name: "test",
		Node: Node{
			Type: NodeTypeValue,
			Value: &ValueNode{
				Field:  "id",
				Format: "hex",
			},
		},
	}
	inf, err := entryNode.InterfaceSchema(Universe{})
	schema, err := inf.JSONSchema(nil)
	if err != nil {
		t.Fatal(err)
	}
	if schema == nil {
		t.Fatal()
	}
}

func TestNode_JSONSchema_List(t *testing.T) {
	entryNode := EndpointNode{
		Name: "test",
		Node: Node{
			Type: NodeTypeList,
			List: &ListNode{
				Entry: Node{
					Type: NodeTypeValue,
					Value: &ValueNode{
						Field:  "id",
						Format: "hex",
					},
				},
			},
		},
	}
	inf, err := entryNode.InterfaceSchema(Universe{})
	schema, err := inf.JSONSchema(nil)
	if err != nil {
		t.Fatal(err)
	}
	if schema == nil {
		t.Fatal()
	}
}

func TestNode_JSONSchema_Map(t *testing.T) {
	entryNode := EndpointNode{
		Name: "test",
		Node: Node{
			Type: NodeTypeMap,
			Map: &MapNode{
				Entries: []MapNodeEntry{
					{
						Name: "id",
						Node: Node{
							Type: NodeTypeValue,
							Value: &ValueNode{
								Field:  "id",
								Format: "hex",
							},
						},
					},
				},
			},
		},
	}
	inf, err := entryNode.InterfaceSchema(Universe{})
	schema, err := inf.JSONSchema(nil)
	if err != nil {
		t.Fatal(err)
	}
	if schema == nil {
		t.Fatal()
	}
}

func TestNode_StructSchema_ID1(t *testing.T) {
	entryNode := Node{
		Type: NodeTypeValue,
		Value: &ValueNode{
			Field:  "id",
			Format: "hex",
		},
	}
	nodeInterface, err := entryNode.InterfaceSchema(nil, Environment{Model: "testmodel"})
	if err != nil {
		t.Fatal(err)
	}
	if nodeInterface.Type != SchemaTypePrimitive || nodeInterface.Primitive.Value != PrimitiveTypeID || nodeInterface.Primitive.Model != "testmodel" {
		t.Fatal()
	}
}

func TestNode_StructSchema_ID2(t *testing.T) {
	univ := Universe{}
	_ = univ.LoadFile("./test/vision.vyu")
	entryNode := Node{
		Type: NodeTypeRelation,
		Relation: &RelationNode{
			Relation: "image#camera",
			Node: Node{
				Type: NodeTypeValue,
				Value: &ValueNode{
					Field:  "id",
					Format: "hex",
				},
			},
		},
	}
	nodeInterface, err := entryNode.InterfaceSchema(&univ, Environment{Model: "image"})
	if err != nil {
		t.Fatal(err)
	}
	if nodeInterface.Type != SchemaTypePrimitive || nodeInterface.Primitive.Value != PrimitiveTypeID || nodeInterface.Primitive.Model != "camera" {
		t.Fatal()
	}
}
