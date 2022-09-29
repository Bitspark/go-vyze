package vyze

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

// Resolver

// A Resolver maps identifiers onto object IDs
type Resolver interface {
	Resolve(ident string, base string) ID
}

// A Universe serves as universal exchange format and captures a subpart of the VYZE graph
type Universe struct {
	ID           ID                    `json:"id" yaml:"id"`
	Name         string                `json:"name" yaml:"name"`
	Description  string                `json:"description" yaml:"description"`
	Bases        []string              `json:"bases" yaml:"bases"`
	Dependencies []string              `json:"dependencies" yaml:"dependencies"`
	Identifiers  []UniverseIdentifier  `json:"objects" yaml:"objects"`
	Models       []UniverseObjectInfo  `json:"info" yaml:"info"`
	Relations    []UniverseRelation    `json:"relations" yaml:"relations"`
	Abstractions []UniverseAbstraction `json:"abstractions" yaml:"abstractions"`
	Endpoints    []EndpointNode        `json:"endpoints" yaml:"endpoints"`
	Interfaces   []NamedInterface      `json:"interfaces" yaml:"interfaces"`
}

// Ensure that Universe implements the Resolver interface
var _ Resolver = Universe{}

type UniverseIdentifier struct {
	// Base is the base universe to which this object belongs
	Base string `json:"universe,omitempty" yaml:"universe,omitempty"`

	// Name is the name of the object
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Target indicates if this object is inside or outside the universe
	Target string `json:"base,omitempty" yaml:"base,omitempty"`
}

type UniverseObjectInfo struct {
	// Mapping contains the identifier
	Mapping UniverseIdentifier `json:"mapping" yaml:"mapping"`

	// Type specifies whether this object is a model, relation or data type
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// ObjectID contains the ID of the associated object
	ObjectID *ID `json:"object,omitempty" yaml:"object,omitempty"`

	// Description contains the description of the model
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

type UniverseAbstraction struct {
	Abstract UniverseIdentifier `json:"abstract" yaml:"abstract"`
	Special  UniverseIdentifier `json:"special" yaml:"special"`
}

type UniverseRelation struct {
	Relation UniverseIdentifier `json:"relation" yaml:"relation"`
	Origin   UniverseIdentifier `json:"origin" yaml:"origin"`
	Target   UniverseIdentifier `json:"target" yaml:"target"`
}

func (rf Universe) GetModel(ident string, base string) *UniverseObjectInfo {
	ri := ParseUniverseObjectIdentifier(ident)
	for _, bi := range rf.Models {
		if bi.Mapping.Canonical(rf.Name, rf.Name) == ri.Canonical(base, rf.Name) {
			return &bi
		}
	}
	return nil
}

func (rf Universe) GetEndpoint(name string) *EndpointNode {
	for _, nd := range rf.Endpoints {
		if nd.Name == name {
			return &nd
		}
	}
	return nil
}

func (rf Universe) GetInterface(name string) *NamedInterface {
	for _, s := range rf.Interfaces {
		if s.Name == name {
			return &s
		}
	}
	return nil
}

func (rf Universe) GetTarget(model UniverseIdentifier) *UniverseIdentifier {
	for _, f := range rf.Relations {
		if f.Relation.Canonical(rf.Name, rf.Name) == model.Canonical(rf.Name, rf.Name) {
			return &f.Target
		}
	}
	return nil
}

func (rf Universe) GetOrigin(model UniverseIdentifier) *UniverseIdentifier {
	for _, f := range rf.Relations {
		if f.Relation.Canonical(rf.Name, rf.Name) == model.Canonical(rf.Name, rf.Name) {
			return &f.Origin
		}
	}
	return nil
}

func (rf Universe) Resolve(ident string, base string) ID {
	id, _ := ParseID(ident)
	if !id.IsNull() {
		return id
	}
	model := rf.GetModel(ident, base)
	if model == nil || model.ObjectID == nil {
		return ID{}
	}
	return *model.ObjectID
}

func (rf Universe) hasIdentifier(id UniverseIdentifier, bases bool, deps bool) bool {
	if id.Target == "" || id.Target == rf.Name {
		return true
	}
	if bases {
		for _, base := range rf.Bases {
			if id.Target == base {
				return true
			}
		}
	}
	if deps {
		for _, dep := range rf.Dependencies {
			if id.Target == dep {
				return true
			}
		}
	}
	return false
}

func (rf Universe) Copy(base string, bases bool, deps bool, objects bool) Universe {
	rfCpy := Universe{
		Name:        rf.Name,
		Description: rf.Description,
	}

	for _, obj := range rf.Bases {
		rfCpy.Bases = append(rfCpy.Bases, obj)
	}

	for _, obj := range rf.Dependencies {
		rfCpy.Dependencies = append(rfCpy.Dependencies, obj)
	}

	for _, obj := range rf.Identifiers {
		if !rf.hasIdentifier(obj, bases, deps) {
			continue
		}
		if obj.Base == "" {
			obj.Base = base
		}
		rfCpy.Identifiers = append(rfCpy.Identifiers, obj)
	}

	for _, info := range rf.Models {
		if !rf.hasIdentifier(info.Mapping, bases, deps) {
			continue
		}
		if info.Mapping.Base == "" {
			info.Mapping.Base = base
		}
		if !objects {
			info.ObjectID = nil
		}
		rfCpy.Models = append(rfCpy.Models, info)
	}

	for _, rel := range rf.Relations {
		if !rf.hasIdentifier(rel.Relation, bases, deps) || !rf.hasIdentifier(rel.Origin, bases, deps) || !rf.hasIdentifier(rel.Target, bases, deps) {
			continue
		}
		if rel.Target.Base == "" {
			rel.Target.Base = base
		}
		if rel.Relation.Base == "" {
			rel.Relation.Base = base
		}
		if rel.Origin.Base == "" {
			rel.Origin.Base = base
		}
		rfCpy.Relations = append(rfCpy.Relations, rel)
	}

	for _, abs := range rf.Abstractions {
		if !rf.hasIdentifier(abs.Abstract, bases, deps) || !rf.hasIdentifier(abs.Special, bases, deps) {
			continue
		}
		if abs.Abstract.Base == "" {
			abs.Abstract.Base = base
		}
		if abs.Special.Base == "" {
			abs.Special.Base = base
		}
		rfCpy.Abstractions = append(rfCpy.Abstractions, abs)
	}

	for _, node := range rf.Endpoints {
		rfCpy.Endpoints = append(rfCpy.Endpoints, node)
	}

	for _, str := range rf.Interfaces {
		rfCpy.Interfaces = append(rfCpy.Interfaces, str)
	}

	return rfCpy
}

// Extend transform the universe file into an extension
func (rf Universe) Extend(base string) Universe {
	ef := rf.Copy(base, true, false, false)
	infoMap := map[string]UniverseObjectInfo{}
	for _, info := range ef.Models {
		infoMap[info.Mapping.String()] = info
	}
	for _, obj := range ef.Identifiers {
		if obj.Target != "" {
			continue
		}
		info, ok := infoMap[obj.String()]
		if !ok {
			continue
		}
		aoExt := UniverseIdentifier{
			Base:   obj.Base,
			Name:   obj.Name,
			Target: base,
		}
		infExt := UniverseObjectInfo{
			Mapping:     aoExt,
			Type:        info.Type,
			Description: info.Description,
		}
		ef.Identifiers = append(ef.Identifiers, aoExt)
		ef.Models = append(ef.Models, infExt)
		ef.Abstractions = append(ef.Abstractions, UniverseAbstraction{
			Abstract: aoExt,
			Special:  obj,
		})
	}
	//ef.Objects = append(ef.Objects, root, rootRelation)
	return ef
}

func (rf Universe) DumpFile(filepath string) error {
	b := &bytes.Buffer{}
	var err error
	if strings.HasSuffix(filepath, ".vyu.zlib") {
		err = rf.Dump(b, true)
	}
	if strings.HasSuffix(filepath, ".vyu") {
		err = rf.Dump(b, false)
	}
	if err != nil {
		return err
	}
	bts := b.Bytes()
	if err := os.WriteFile(filepath, bts, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (rf Universe) Dump(w io.Writer, compressed bool) error {
	bts, err := yaml.Marshal(rf)
	if err != nil {
		return err
	}
	if compressed {
		z := zlib.NewWriter(w)
		defer z.Close()
		w = z
	}
	if _, err := w.Write(bts); err != nil {
		return err
	}
	return nil
}

func (rf *Universe) LoadFile(filepath string) error {
	bts, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	b := bytes.NewBuffer(bts)
	if strings.HasSuffix(filepath, ".vyu.zlib") {
		return rf.Load(b, true)
	}
	if strings.HasSuffix(filepath, ".vyu") {
		return rf.Load(b, false)
	}
	return errors.New(fmt.Sprintf("invalid ending: %s", filepath))
}

func (rf *Universe) Load(r io.Reader, compressed bool) error {
	if compressed {
		z, err := zlib.NewReader(r)
		if err != nil {
			return err
		}
		defer z.Close()
		r = z
	}
	bts, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(bts, rf); err != nil {
		return err
	}
	return nil
}

func (rf *Universe) Clean() {
	newAbses := []UniverseAbstraction{}
	for _, abs := range rf.Abstractions {
		if _, ok := rf.AbstractsOf(abs.Special)[abs.Abstract.String()]; !ok {
			newAbses = append(newAbses, abs)
		}
	}
	rf.Abstractions = newAbses

	usedObjs := map[string]bool{}

	newObjs := []UniverseIdentifier{}
	for _, obj := range rf.Identifiers {
		if rf.usedIdentifier(obj) {
			newObjs = append(newObjs, obj)
			usedObjs[obj.String()] = true
		}
	}
	rf.Identifiers = newObjs

	newInfos := []UniverseObjectInfo{}
	for _, info := range rf.Models {
		if usedObjs[info.Mapping.String()] {
			newInfos = append(newInfos, info)
		}
	}
	rf.Models = newInfos

	// Checks if an abstraction is redundant in the sense that it is a composition of two other abstractions
	isAbsRedundant := func(a UniverseAbstraction) bool {
		for _, a1 := range rf.Abstractions {
			if a1.Abstract != a.Abstract {
				continue
			}
			for _, a2 := range rf.Abstractions {
				if a1 == a2 {
					continue
				}
				if a1.Special != a2.Abstract {
					continue
				}
				if a2.Special == a.Special {
					return true
				}
			}
		}
		return false
	}

	newAbses = []UniverseAbstraction{}
	for _, abs := range rf.Abstractions {
		if !isAbsRedundant(abs) {
			newAbses = append(newAbses, abs)
		}
	}
	rf.Abstractions = newAbses

	newNodes := []EndpointNode{}
	for _, nd := range rf.Endpoints {
		newNodes = append(newNodes, nd)
	}
	rf.Endpoints = newNodes

	newStructs := []NamedInterface{}
	for _, str := range rf.Interfaces {
		newStructs = append(newStructs, str)
	}
	rf.Interfaces = newStructs

	sort.Slice(rf.Bases, func(i, j int) bool {
		return rf.Bases[i] < rf.Bases[j]
	})
	sort.Slice(rf.Dependencies, func(i, j int) bool {
		return rf.Dependencies[i] < rf.Dependencies[j]
	})
	sort.Slice(rf.Identifiers, func(i, j int) bool {
		return rf.Identifiers[i].String() < rf.Identifiers[j].String()
	})
	sort.Slice(rf.Models, func(i, j int) bool {
		return rf.Models[i].Mapping.String() < rf.Models[j].Mapping.String()
	})
	sort.Slice(rf.Relations, func(i, j int) bool {
		if rf.Relations[i].Relation.String() != rf.Relations[j].Relation.String() {
			return rf.Relations[i].Relation.String() < rf.Relations[j].Relation.String()
		}
		if rf.Relations[i].Origin.String() != rf.Relations[j].Origin.String() {
			return rf.Relations[i].Origin.String() < rf.Relations[j].Origin.String()
		}
		return rf.Relations[i].Target.String() < rf.Relations[j].Target.String()
	})
	sort.Slice(rf.Abstractions, func(i, j int) bool {
		if rf.Abstractions[i].Abstract.String() != rf.Abstractions[j].Abstract.String() {
			return rf.Abstractions[i].Abstract.String() < rf.Abstractions[j].Abstract.String()
		}
		return rf.Abstractions[i].Special.String() < rf.Abstractions[j].Special.String()
	})
	sort.Slice(rf.Endpoints, func(i, j int) bool {
		return rf.Endpoints[i].Name < rf.Endpoints[j].Name
	})
	sort.Slice(rf.Interfaces, func(i, j int) bool {
		return rf.Interfaces[i].Name < rf.Interfaces[j].Name
	})
}

func (rf Universe) usedIdentifier(ri UniverseIdentifier) bool {
	for _, abs := range rf.Abstractions {
		if abs.Abstract.Equals(ri) {
			return true
		}
		if abs.Special.Equals(ri) {
			return true
		}
	}
	for _, abs := range rf.Relations {
		if abs.Relation.Equals(ri) {
			return true
		}
		if abs.Origin.Equals(ri) {
			return true
		}
		if abs.Target.Equals(ri) {
			return true
		}
	}
	return false
}

func (rf Universe) AbstractsOf(ri UniverseIdentifier) map[string]bool {
	mp := map[string]bool{}
	for _, abs := range rf.Abstractions {
		if !abs.Special.Equals(ri) {
			continue
		}
		for a2 := range rf.abstractsOf(abs.Abstract) {
			mp[a2] = true
		}
	}
	return mp
}

func (rf Universe) abstractsOf(ri UniverseIdentifier) map[string]bool {
	mp := map[string]bool{}
	for _, abs := range rf.Abstractions {
		if !abs.Special.Equals(ri) {
			continue
		}
		mp[abs.Abstract.String()] = true
		for a2 := range rf.abstractsOf(abs.Abstract) {
			mp[a2] = true
		}
	}
	return mp
}

func (rf Universe) Library() Library {
	lib := Library{
		Interfaces: map[string]NamedInterface{},
		Name:       rf.Name,
	}
	for _, ni := range rf.Interfaces {
		lib.Interfaces[ni.Name] = ni
	}
	return lib
}

func ParseUniverseObjectIdentifier(str string) (ri UniverseIdentifier) {
	targetSplit := strings.Split(str, "/")
	objectId := targetSplit[0]

	objSplit := strings.Split(objectId, ".")

	if len(objSplit) == 2 {
		ri.Base = objSplit[0]
		ri.Name = objSplit[1]
	} else {
		ri.Name = objSplit[0]
	}

	if len(targetSplit) == 2 {
		ri.Target = targetSplit[1]
	} else {
		ri.Target = ri.Base
	}

	ri.Base = strings.ToLower(ri.Base)
	ri.Name = strings.ToLower(ri.Name)
	ri.Target = strings.ToLower(ri.Target)

	return
}

/*
func ParseUniverseObjectIdentifier(str string, uname string) (ri UniverseIdentifier) {
	metaSplit := strings.Split(str, ":")
	targetSplit := strings.Split(metaSplit[0], "/")
	baseSplit := strings.Split(targetSplit[0], ".")

	oi := UniverseIdentifier{
		Base:   uname,
		Name:   targetSplit[0],
		Target: uname,
	}

	if len(baseSplit) >= 2 {
		oi.Base = baseSplit[0]
		oi.Name = baseSplit[1]
	}

	if len(targetSplit) >= 2 {
		if targetSplit[1] != "" {
			oi.Target = targetSplit[1]
		} else {
			oi.Target = uname
		}
	} else {
		oi.Target = oi.Base
	}

	return oi
}
*/

func ParseUniverseFileRelation(str string) (ri UniverseRelation) {
	identSplit := strings.Split(str, ":")
	if len(identSplit) != 3 {
		return
	}

	ri.Relation = ParseUniverseObjectIdentifier(identSplit[0])
	ri.Origin = ParseUniverseObjectIdentifier(identSplit[1])
	ri.Target = ParseUniverseObjectIdentifier(identSplit[2])

	return
}

func ParseUniverseFileAbstraction(str string) (ri UniverseAbstraction) {
	identSplit := strings.Split(str, ":")
	if len(identSplit) != 2 {
		return
	}

	ri.Abstract = ParseUniverseObjectIdentifier(identSplit[0])
	ri.Special = ParseUniverseObjectIdentifier(identSplit[1])

	return
}

func (ri UniverseIdentifier) Equals(ri2 UniverseIdentifier) bool {
	return ri.Name == ri2.Name && ri.Target == ri2.Target && ri.Base == ri2.Base
}

func (ri UniverseIdentifier) Canonical(base string, target string) string {
	objName := ri.Name
	objBase := ri.Base
	objTarget := ri.Target
	if objTarget == "" {
		objTarget = target
	}
	if objBase == "" {
		if base == "" {
			objBase = objTarget
		} else {
			objBase = base
		}
	}
	return objBase + "." + objName + "/" + objTarget
}

func (ri UniverseIdentifier) String() string {
	str := ""
	if ri.Base == "" {
		if ri.Target == ri.Base {
			str = ri.Name
		} else {
			str = ri.Name + "/" + ri.Target
		}
	} else {
		if ri.Target == ri.Base {
			str = ri.Base + "." + ri.Name
		} else {
			str = ri.Base + "." + ri.Name + "/" + ri.Target
		}
	}
	return str
}

func (ri UniverseIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(ri.String())
}

func (ri *UniverseIdentifier) UnmarshalJSON(b []byte) error {
	var str *string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	if str == nil {
		return errors.New("empty")
	}
	*ri = ParseUniverseObjectIdentifier(*str)
	return nil
}

func (ri UniverseIdentifier) MarshalYAML() (interface{}, error) {
	return ri.String(), nil
}

func (ri *UniverseIdentifier) UnmarshalYAML(value *yaml.Node) error {
	var str string
	err := value.Decode(&str)
	if err != nil {
		return err
	}
	*ri = ParseUniverseObjectIdentifier(str)
	return nil
}

func (ri UniverseRelation) MarshalYAML() (interface{}, error) {
	rel := UniverseIdentifier{
		Base:   ri.Relation.Base,
		Name:   ri.Relation.Name,
		Target: ri.Relation.Target,
	}
	orig := UniverseIdentifier{
		Base:   ri.Origin.Base,
		Name:   ri.Origin.Name,
		Target: ri.Origin.Target,
	}
	target := UniverseIdentifier{
		Base:   ri.Target.Base,
		Name:   ri.Target.Name,
		Target: ri.Target.Target,
	}
	return rel.String() + ":" + orig.String() + ":" + target.String(), nil
}

func (ri *UniverseRelation) UnmarshalYAML(value *yaml.Node) error {
	var str string
	err := value.Decode(&str)
	if err != nil {
		return err
	}
	*ri = ParseUniverseFileRelation(str)
	return nil
}

func (ri UniverseAbstraction) MarshalYAML() (interface{}, error) {
	abs := UniverseIdentifier{
		Base:   ri.Abstract.Base,
		Name:   ri.Abstract.Name,
		Target: ri.Abstract.Target,
	}
	spec := UniverseIdentifier{
		Base:   ri.Special.Base,
		Name:   ri.Special.Name,
		Target: ri.Special.Target,
	}
	return abs.String() + ":" + spec.String(), nil
}

func (ri *UniverseAbstraction) UnmarshalYAML(value *yaml.Node) error {
	var str string
	err := value.Decode(&str)
	if err != nil {
		return err
	}
	*ri = ParseUniverseFileAbstraction(str)
	return nil
}
