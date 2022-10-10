package vyze

import (
	"testing"
)

func TestTemplate_Apply(t *testing.T) {
	tm := TemplateValuation{
		"a": TemplateValue{
			Type:  TemplateTypePrimitive,
			Value: "test",
		},
		"b": TemplateValue{
			Type:   TemplateTypePrimitive,
			Source: "b",
		},
	}
	ts := TemplateValuation{
		"b": TemplateValue{
			Type:  TemplateTypePrimitive,
			Value: "test2",
		},
	}
	rt := tm.Apply(ts)
	if rt["a"].Value != "test" {
		t.Fatal()
	}
	if rt["b"].Value != "test2" {
		t.Fatal()
	}
}

func TestTemplate_Expand(t *testing.T) {
	tpl := TemplateValuation{
		"var1": {
			Type:  TemplateTypeList,
			Value: []string{"v1", "v2", "v3"},
		},
	}
	tpls, err := tpl.Expand(TemplateExpansion{Source: "var1", Target: "var2"})
	if err != nil {
		t.Fatal(err)
	}
	if len(tpls) != 3 {
		t.Fatal()
	}
	if tpls[0]["var2"].Type != TemplateTypePrimitive {
		t.Fatal()
	}
	if tpls[0]["var2"].Value != "v1" {
		t.Fatal()
	}
}

func TestTemplate_Camel(t *testing.T) {
	tpl := TemplateValuation{
		"var1": {
			Type:  TemplateTypeList,
			Value: []string{"v1", "v2", "v3"},
		},
		"var2": {
			Type:  TemplateTypeList,
			Value: []string{"single"},
		},
	}
	if tpl.Camel() != "V1V2V3Single" {
		t.Fatal()
	}
}

func TestTemplateString_Apply(t *testing.T) {
	tpl := TemplateValuation{
		"var1": TemplateValue{
			Type:  TemplateTypePrimitive,
			Value: "val1",
		},
	}
	ts := TemplateString("v{var1}")
	ts2 := ts.Apply(tpl)
	if ts2 != "vval1" {
		t.Fatal()
	}
}

func TestTemplateString_Resolved1(t *testing.T) {
	ts := TemplateString("v{var1}")
	if ts.Resolved() {
		t.Fatal()
	}
}

func TestTemplateString_Resolved2(t *testing.T) {
	ts := TemplateString("vval1")
	if !ts.Resolved() {
		t.Fatal()
	}
}

func TestTemplateValue_Camel1(t *testing.T) {
	tv := TemplateValue{
		Type:  TemplateTypePrimitive,
		Value: "val",
	}
	if tv.Camel() != "Val" {
		t.Fatal()
	}
}

func TestTemplateValue_Camel2(t *testing.T) {
	tv := TemplateValue{
		Type:  TemplateTypePrimitive,
		Value: "val1 val2",
	}
	if tv.Camel() != "Val1Val2" {
		t.Fatal()
	}
}

func TestTemplateValue_Camel3(t *testing.T) {
	tv := TemplateValue{
		Type:  TemplateTypeList,
		Value: []string{"val1", "val2"},
	}
	if tv.Camel() != "Val1Val2" {
		t.Fatal()
	}
}

func TestMapSchema_ApplyTemplate1(t *testing.T) {
	schema := MapInterface{
		Entries: []MapInterfaceEntry{
			{
				Key: "anInt",
				Schema: Interface{
					Type: SchemaTypePrimitive,
					Primitive: &PrimitiveInterface{
						Value: PrimitiveTypeInteger,
					},
				},
			},
		},
	}
	schema2, _ := schema.ApplyTemplate(TemplateValuation{})
	schemaJSON, err := schema2.JSONSchema(nil)
	if err != nil {
		t.Fatal(err)
	}
	if schemaJSON["type"] != "object" {
		t.Fatal()
	}
	if len(schemaJSON["properties"].(map[string]any)) != 1 {
		t.Fatal()
	}
	if schemaJSON["properties"].(map[string]any)["anInt"].(map[string]any)["type"] != "integer" {
		t.Fatal()
	}
}

func TestMapSchema_ApplyTemplate2(t *testing.T) {
	schema := MapInterface{
		Entries: []MapInterfaceEntry{
			{
				Key: "{key}",
				Expansion: &TemplateExpansion{
					Source: "keys",
					Target: "key",
				},
				Schema: Interface{
					Type: SchemaTypePrimitive,
					Primitive: &PrimitiveInterface{
						Value: PrimitiveTypeInteger,
					},
				},
			},
		},
	}
	schema2, _ := schema.ApplyTemplate(TemplateValuation{
		"keys": {
			Type:  TemplateTypeList,
			Value: []string{"int1", "int2", "int3"},
		},
	})
	schemaJSON, err := schema2.JSONSchema(nil)
	if err != nil {
		t.Fatal(err)
	}
	if schemaJSON["type"] != "object" {
		t.Fatal()
	}
	if len(schemaJSON["properties"].(map[string]any)) != 3 {
		t.Fatal()
	}
	if schemaJSON["properties"].(map[string]any)["int1"].(map[string]any)["type"] != "integer" {
		t.Fatal()
	}
}

func TestGenericSchema_ApplyTemplate1(t *testing.T) {
	schema := GenericInterface{
		Name: "gen1",
	}
	schema2, _ := schema.ApplyTemplate(TemplateValuation{})
	schemaJSON, err := schema2.JSONSchema(nil)
	if err != nil {
		t.Fatal(err)
	}
	if schemaJSON["type"] != nil {
		t.Fatal()
	}
	if schemaJSON["title"] != "gen1 (generic)" {
		t.Fatal()
	}
	if len(schemaJSON["description"].(string)) == 0 {
		t.Fatal()
	}
}

func TestGenericSchema_ApplyTemplate2(t *testing.T) {
	schema := GenericInterface{
		Name: "g_{var1}",
	}
	schema2, _ := schema.ApplyTemplate(TemplateValuation{
		"var1": TemplateValue{
			Type:  TemplateTypePrimitive,
			Value: "gen1",
		},
	})
	schemaJSON, err := schema2.JSONSchema(nil)
	if err != nil {
		t.Fatal(err)
	}
	if schemaJSON["type"] != nil {
		t.Fatal()
	}
	if schemaJSON["title"] != "g_gen1 (generic)" {
		t.Fatal()
	}
	if len(schemaJSON["description"].(string)) == 0 {
		t.Fatal()
	}
}
