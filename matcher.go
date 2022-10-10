package go_vyze

import (
	"errors"
	"fmt"
)

type InterfaceMatcher struct {
	ts       TemplateValuation
	gv       GenericValuation
	generics []GenericSlot
	template []TemplateSlot
}

func NewInterfaceMatcher(generics []GenericSlot, template []TemplateSlot) InterfaceMatcher {
	return InterfaceMatcher{
		ts:       TemplateValuation{},
		gv:       GenericValuation{},
		generics: generics,
		template: template,
	}
}

func (m InterfaceMatcher) Template() TemplateValuation {
	return m.ts
}

func (m InterfaceMatcher) Generics() GenericValuation {
	return m.gv
}

func (m InterfaceMatcher) Match(ss, ts Interface) error {
	if err := m.MatchTemplate(ss, ts); err != nil {
		return err
	}
	ss2, _ := ss.ApplyTemplate(m.Template())
	if err := m.MatchGenerics(ss2, ts); err != nil {
		return err
	}
	return nil
}

// TEMPLATE

// MatchTemplate attempts to find a template valuation which transforms gs into ts when the returned valuation is
// applied to gs via gs.ApplyTemplate
func (m *InterfaceMatcher) MatchTemplate(gs, ts Interface) error {
	if gs.Type != ts.Type {
		if gs.Type != SchemaTypeGeneric {
			return fmt.Errorf("cannot match types %s and %s", gs.Type, ts.Type)
		}
		if !gs.Generic.Name.Resolved() {
			genericName := gs.Generic.Name.Apply(m.ts)
			m.generics = append(m.generics, GenericSlot{
				Name:        genericName,
				Description: fmt.Sprintf("Created from template string '%s'.", gs.Generic.Name),
			})
		}
		// We stop here, because the generic matcher needs to figure this out later
		return nil
	}
	switch gs.Type {
	case SchemaTypeNamed:
		return m.MatchTemplateNamed(*gs.Named, *ts.Named)
	case SchemaTypePrimitive:
		return m.MatchTemplatePrimitive(*gs.Primitive, *ts.Primitive)
	case SchemaTypeList:
		return m.MatchTemplateList(*gs.List, *ts.List)
	case SchemaTypeMap:
		return m.MatchTemplateMap(*gs.Map, *ts.Map)
	case SchemaTypeGeneric:
		return m.MatchTemplateGeneric(*gs.Generic, *ts.Generic)
	case SchemaTypeReference:
		return m.MatchTemplateReference(*gs.Reference, *ts.Reference)
	}
	return errors.New("not implemented")
}

func (m *InterfaceMatcher) MatchTemplateNamed(gs, ts NamedInterface) error {
	return m.MatchTemplate(gs.Schema, ts.Schema)
}

func (m *InterfaceMatcher) MatchTemplatePrimitive(gs, ts PrimitiveInterface) error {
	// TODO: Also check models in case value is an ID
	if gs.Value != ts.Value {
		return errors.New("incompatible primitive values")
	}
	return nil
}

func (m *InterfaceMatcher) MatchTemplateList(gs, ts ListInterface) error {
	return m.MatchTemplate(gs.Entry, ts.Entry)
}

func (m *InterfaceMatcher) MatchTemplateMap(gs, ts MapInterface) error {
	for _, targetEntry := range ts.Entries {
		keyValue := string(targetEntry.Key)
		if entry, ok := gs.GetEntry(keyValue); ok {
			// We find this very key so we descend without expanding
			if err := m.MatchTemplate(entry.Schema, targetEntry.Schema); err != nil {
				return err
			}
			continue
		}
		for _, tplEntry := range gs.Entries {
			if tplEntry.Key.Resolved() || tplEntry.Expansion == nil {
				continue
			}
			// We have an expandable key
			var tplSlot *TemplateSlot
			for _, slot := range m.template {
				if slot.Name == tplEntry.Expansion.Source {
					tplSlot = &slot
					break
				}
			}
			if tplSlot == nil {
				return fmt.Errorf("unknown template %s", tplEntry.Expansion.Source)
			}
			if tplSlot.Type != TemplateTypeList {
				return fmt.Errorf("expected template %s to be a list", tplEntry.Expansion.Source)
			}
			// TODO: Use copy
			m.ts[tplEntry.Expansion.Target] = TemplateValue{
				Type:  TemplateTypePrimitive,
				Value: keyValue,
			}
			if err := m.MatchTemplate(tplEntry.Schema, targetEntry.Schema); err != nil {
				continue
			}
			val, ok := m.ts[tplEntry.Expansion.Source]
			if !ok {
				m.ts[tplEntry.Expansion.Source] = TemplateValue{
					Type:  TemplateTypeList,
					Value: []string{keyValue},
				}
			} else {
				val.Value = append(val.Value.([]string), keyValue)
				m.ts[tplEntry.Expansion.Source] = val
			}
			goto done
		}
		return fmt.Errorf("could not match %s", keyValue)
	done:
	}
	return nil
}

func (m *InterfaceMatcher) MatchTemplateGeneric(gs, ts GenericInterface) error {
	panic("implement me")
}

func (m *InterfaceMatcher) MatchTemplateReference(gs, ts ReferenceInterface) error {
	if gs.Name != ts.Name {
		return fmt.Errorf("unequal references: %s and %s", gs.Name, ts.Name)
	}
	return nil
}

// GENERICS

// MatchGenerics attempts to find a generic valuation which transforms gs into ts when the returned valuation is
// applied to gs via gs.ApplyGenerics
func (m *InterfaceMatcher) MatchGenerics(gs, ts Interface) error {
	if gs.Type != ts.Type {
		if gs.Type != SchemaTypeGeneric {
			return fmt.Errorf("cannot match types %s and %s", gs.Type, ts.Type)
		}
		var genericSlot *GenericSlot
		for _, slot := range m.generics {
			if slot.Name == gs.Generic.Name {
				genericSlot = &slot
				break
			}
		}
		if genericSlot == nil {
			return fmt.Errorf("unknown generic %s", gs.Generic.Name)
		}
		if ev, ok := m.gv[string(gs.Generic.Name)]; ok {
			// Target type is generic, and we already have a valuation for that generic
			if !ev.Equals(ts) {
				// TODO: We should try to match them (test this properly!)
				return fmt.Errorf("cannot match types %s and %s", ev.Type, ts.Type)
			} else {
				// They are equal, so everything is fine
				return nil
			}
		} else {
			// Target type is generic, and we do not have a valuation for that, so we store the target type
			m.gv[string(gs.Generic.Name)] = ts
			return nil
		}
	}
	switch gs.Type {
	case SchemaTypeNamed:
		return m.MatchGenericsNamed(*gs.Named, *ts.Named)
	case SchemaTypePrimitive:
		return m.MatchGenericsPrimitive(*gs.Primitive, *ts.Primitive)
	case SchemaTypeList:
		return m.MatchGenericsList(*gs.List, *ts.List)
	case SchemaTypeMap:
		return m.MatchGenericsMap(*gs.Map, *ts.Map)
	case SchemaTypeGeneric:
		return m.MatchGenericsGeneric(*gs.Generic, *ts.Generic)
	case SchemaTypeReference:
		return m.MatchGenericsReference(*gs.Reference, *ts.Reference)
	}
	return errors.New("not implemented")
}

func (m *InterfaceMatcher) MatchGenericsNamed(gs, ts NamedInterface) error {
	return m.MatchGenerics(gs.Schema, ts.Schema)
}

func (m *InterfaceMatcher) MatchGenericsPrimitive(gs, ts PrimitiveInterface) error {
	// TODO: Also check models in case value is an ID
	if gs.Value != ts.Value {
		return errors.New("incompatible primitive values")
	}
	return nil
}

func (m *InterfaceMatcher) MatchGenericsList(gs, ts ListInterface) error {
	return m.MatchGenerics(gs.Entry, ts.Entry)
}

func (m *InterfaceMatcher) MatchGenericsMap(gs, ts MapInterface) error {
	for _, e := range gs.Entries {
		if e2, ok := ts.GetEntry(string(e.Key)); ok {
			if err := m.MatchGenerics(e.Schema, e2.Schema); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("missing entry %s", e.Key)
		}
	}
	for _, e2 := range ts.Entries {
		if _, ok := gs.GetEntry(string(e2.Key)); !ok {
			return fmt.Errorf("missing entry %s", e2.Key)
		}
	}
	return nil
}

func (m *InterfaceMatcher) MatchGenericsGeneric(gs, ts GenericInterface) error {
	panic("implement me")
}

func (m *InterfaceMatcher) MatchGenericsReference(gs, ts ReferenceInterface) error {
	if gs.Name != ts.Name {
		return fmt.Errorf("unequal references: %s and %s", gs.Name, ts.Name)
	}
	return nil
}
