package vyze

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/jxskiss/base62"
	"gopkg.in/yaml.v3"
	"sort"
	"strings"
)

const IDSize = 16

type ID [IDSize]byte
type RichID [IDSize + 1]byte
type IDPair [2]ID
type KeyedIDPair struct {
	Key  Binary `json:"key"`
	Pair IDPair `json:"pair"`
}

var IDHexLen = len(ID{}.Hex())
var IDBase64Len = len(ID{}.Base64())
var IDStringLen = IDBase64Len

var noPadding32 = base32.HexEncoding.WithPadding(base32.NoPadding)

var MinID ID
var MaxID ID

func init() {
	for i := 0; i < IDSize; i++ {
		MinID[i] = 0
		MaxID[i] = 255
	}
}

func NewID() ID {
	id := ID{}
	_, err := rand.Read(id[:])
	if err != nil {
		panic(err)
	}
	return id
}

func (id ID) MarshalJSON() ([]byte, error) {
	if id.IsNull() {
		return []byte("null"), nil
	}
	buffer := bytes.NewBufferString("")
	buffer.WriteString("\"")
	buffer.WriteString(id.String())
	buffer.WriteString("\"")
	return buffer.Bytes(), nil
}

func (id *ID) UnmarshalJSON(b []byte) error {
	if b == nil || string(b) == "null" {
		*id = ID{}
		return nil
	}
	var idString string
	err := json.Unmarshal(b, &idString)
	if err != nil {
		return err
	}
	*id, err = ParseID(idString)
	if err != nil {
		return err
	}
	return nil
}

func (id ID) MarshalYAML() (interface{}, error) {
	return id.String(), nil
}

func (id *ID) UnmarshalYAML(value *yaml.Node) error {
	var err error
	*id, err = ParseID(value.Value)
	if err != nil {
		return err
	}
	return nil
}

func (id ID) Invert() ID {
	invID := ID{}
	for i := 0; i < len(id); i++ {
		invID[i] = ^id[i]
	}
	return invID
}

func (id ID) String() string {
	return id.Hex()
}

func (id ID) Hex() string {
	return hex.EncodeToString(id[:])
}

func (id ID) Base32() string {
	return noPadding32.EncodeToString(id[:])
}

func (id ID) Base62() string {
	return base62.EncodeToString(id[:])
}

func (id ID) Base64() string {
	return base64.RawURLEncoding.EncodeToString(id[:])
}

func (id ID) Equals(other ID) bool {
	return bytes.Compare(id[:], other[:]) == 0
}

func (id ID) LessThan(other ID) bool {
	for i := range id {
		if id[i] < other[i] {
			return true
		}
		if id[i] > other[i] {
			return false
		}
	}
	return false
}

func (id ID) IsNull() bool {
	return id.Equals(ID{})
}

func B2ID(b Binary) ID {
	id := ID{}
	copy((id)[:], b)
	return id
}

func ParseID(str string) (ID, error) {
	if len(str) == IDHexLen {
		return ParseIDHex(str)
	}
	if len(str) == IDBase64Len {
		return ParseIDBase64(str)
	}
	return ID{}, errors.New("unknown format (invalid length)")
}

func MustParseID(str string) ID {
	id, err := ParseID(str)
	if err != nil {
		panic(err)
	}
	return id
}

func ParseIDNC(i interface{}) ID {
	str, ok := i.(string)
	if !ok {
		return ID{}
	}
	if len(str) == IDBase64Len {
		id, _ := ParseIDBase64(str)
		return id
	}
	if len(str) == IDHexLen {
		id, _ := ParseIDHex(str)
		return id
	}
	return ID{}
}

func ParseIDHex(str string) (ID, error) {
	if len(str) != 2*IDSize {
		return ID{}, errors.New("wrong length")
	}
	idBytes, err := hex.DecodeString(str)
	if err != nil {
		return ID{}, err
	}
	if len(idBytes) != IDSize {
		return ID{}, errors.New("wrong length")
	}
	id := ID{}
	copy(id[:], idBytes)
	return id, nil
}

func ParseIDBase32(str string) (ID, error) {
	idBytes, err := noPadding32.DecodeString(str)
	if err != nil {
		return ID{}, err
	}
	if len(idBytes) != IDSize {
		return ID{}, errors.New("wrong length")
	}
	id := ID{}
	copy(id[:], idBytes)
	return id, nil
}

func ParseIDBase62(str string) (ID, error) {
	idBytes, err := base62.DecodeString(str)
	if err != nil {
		return ID{}, err
	}
	if len(idBytes) != IDSize {
		return ID{}, errors.New("wrong length")
	}
	id := ID{}
	copy(id[:], idBytes)
	return id, nil
}

func ParseIDBase64(str string) (ID, error) {
	idBytes, err := base64.RawURLEncoding.DecodeString(str)
	if err != nil {
		return ID{}, err
	}
	if len(idBytes) != IDSize {
		return ID{}, errors.New("wrong length")
	}
	id := ID{}
	copy(id[:], idBytes)
	return id, nil
}

// Rich ID

func WrapID(id ID, info byte) RichID {
	r := RichID{}
	copy(r[:], id[:])
	r[IDSize] = info
	return r
}

func UnwrapRID(rid RichID) (ID, byte) {
	id := ID{}
	copy(id[:], rid[:])
	return id, rid[IDSize]
}

func UnwrapID(rid RichID) ID {
	id, _ := UnwrapRID(rid)
	return id
}

func (id RichID) HasFlag(flag byte) bool {
	return id[IDSize]&flag == flag
}

func (id *RichID) SetFlag(flag byte) {
	id[IDSize] |= flag
}

func (id *RichID) UnsetFlag(flag byte) {
	id[IDSize] ^= flag
}

func (id RichID) Hex() string {
	return hex.EncodeToString(id[:IDSize]) + ":" + hex.EncodeToString(id[IDSize:])
}

func (id RichID) String() string {
	return id.Hex()
}

func ParseRID(str string) (RichID, error) {
	sp := strings.Split(str, ":")
	if len(sp) != 2 {
		return RichID{}, errors.New("malformed rid")
	}
	id, err := ParseIDHex(sp[0])
	if err != nil {
		return RichID{}, err
	}
	info, err := hex.DecodeString(sp[1])
	if err != nil {
		return RichID{}, err
	}
	if len(info) != 1 {
		return RichID{}, errors.New("wrong info length")
	}
	rid := RichID{}
	copy(rid[:], id[:])
	copy(rid[IDSize:], info)
	return rid, nil
}

// ID Set

type IDSet map[ID]bool

func NewIDSet(ids ...ID) IDSet {
	idSet := IDSet{}
	for _, id := range ids {
		idSet[id] = true
	}
	return idSet
}

func (ic IDSet) Copy() IDSet {
	s := IDSet{}
	for id := range ic {
		s[id] = true
	}
	return s
}

func (ic IDSet) Hash() ID {
	h := ID{}
	for id := range ic {
		for i := 0; i < IDSize; i++ {
			h[i] ^= id[i]
		}
	}
	return h
}

func (ic IDSet) Contains(id ID) bool {
	_, ok := ic[id]
	return ok
}

func (ic *IDSet) AddAll(set IDSet) {
	for id := range set {
		(*ic)[id] = true
	}
}

func (ic *IDSet) List() IDList {
	if ic == nil {
		return nil
	}
	l := IDList{}
	for id := range *ic {
		l = append(l, id)
	}
	return l
}

func (ic *IDSet) SortedList() IDList {
	l := IDList{}
	for id := range *ic {
		l = append(l, id)
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i].LessThan(l[j])
	})
	return l
}

func (ic IDSet) MarshalJSON() ([]byte, error) {
	ids := []ID{}
	for id := range ic {
		ids = append(ids, id)
	}
	return json.Marshal(&ids)
}

func (ic *IDSet) UnmarshalJSON(b []byte) error {
	if *ic == nil {
		*ic = IDSet{}
	}
	var ids []ID
	err := json.Unmarshal(b, &ids)
	if err != nil {
		return err
	}
	for _, id := range ids {
		(*ic)[id] = true
	}
	return nil
}

func (ic IDSet) Without(n IDSet) IDSet {
	newIDs := IDSet{}
	for e := range ic {
		if !n.Contains(e) {
			newIDs[e] = true
		}
	}
	return newIDs
}

// ID List

type IDList []ID

func (ic IDList) MarshalJSON() ([]byte, error) {
	ids := []ID{}
	for _, id := range ic {
		ids = append(ids, id)
	}
	return json.Marshal(&ids)
}

func (ic *IDList) UnmarshalJSON(b []byte) error {
	var ids []ID
	err := json.Unmarshal(b, &ids)
	if err != nil {
		return err
	}
	for _, id := range ids {
		*ic = append(*ic, id)
	}
	return nil
}

func (ic IDList) Len() int {
	return len(ic)
}

func (ic IDList) Less(i, j int) bool {
	return ic[i].LessThan(ic[j])
}

func (ic *IDList) Swap(i, j int) {
	(*ic)[i], (*ic)[j] = (*ic)[j], (*ic)[i]
}

func (ic IDList) Set() IDSet {
	s := IDSet{}
	for _, i := range ic {
		s[i] = true
	}
	return s
}

func (ic IDList) Contains(id ID) bool {
	for _, i := range ic {
		if i == id {
			return true
		}
	}
	return false
}

func (ic *IDList) Remove(id ID, remove int) int {
	newIDs := IDList{}
	removed := 0
	for _, i := range *ic {
		if i == id && (remove == -1 || removed < remove) {
			removed++
			continue
		}
		newIDs = append(newIDs, i)
	}
	*ic = newIDs
	return removed
}

func (ic IDList) Copy() IDList {
	ids := make(IDList, len(ic))
	copy(ids, ic)
	return ids
}

// ID Map

type IDMap map[ID]ID

func (rs IDMap) KeyByValue(id ID) *ID {
	for id2, ref := range rs {
		if ref.Equals(id) {
			return &id2
		}
	}
	return nil
}

func (rs IDMap) ContainsValue(id ID) bool {
	return rs.KeyByValue(id) != nil
}

func (rs IDMap) Values() []ID {
	ids := []ID{}
	for _, v := range rs {
		ids = append(ids, v)
	}
	return ids
}

func (rs IDMap) KeySet() IDSet {
	ids := IDSet{}
	for k := range rs {
		ids[k] = true
	}
	return ids
}

func (rs IDMap) ValueSet() IDSet {
	ids := IDSet{}
	for _, v := range rs {
		ids[v] = true
	}
	return ids
}

func (rs IDMap) MarshalJSON() ([]byte, error) {
	rels := map[string]string{}
	for key, val := range rs {
		rels[key.String()] = val.String()
	}
	return json.Marshal(rels)
}

func (rs *IDMap) UnmarshalJSON(b []byte) error {
	if *rs == nil {
		*rs = IDMap{}
	}
	rels := map[string]string{}
	err := json.Unmarshal(b, &rels)
	if err != nil {
		return err
	}
	for key, val := range rels {
		idK, _ := ParseID(key)
		idV, _ := ParseID(val)
		(*rs)[idK] = idV
	}
	return nil
}

func (rs IDMap) MarshalYAML() (interface{}, error) {
	rels := map[string]string{}
	for key, val := range rs {
		rels[key.String()] = val.String()
	}
	return rels, nil
}

func (rs *IDMap) UnmarshalYAML(value *yaml.Node) error {
	if *rs == nil {
		*rs = IDMap{}
	}
	rels := map[string]string{}
	err := value.Decode(&rels)
	if err != nil {
		return err
	}
	for key, val := range rels {
		idK, _ := ParseID(key)
		idV, _ := ParseID(val)
		(*rs)[idK] = idV
	}
	return nil
}

func (rs IDMap) Pick() ID {
	for k := range rs {
		return k
	}
	return ID{}
}

// ID Pair List

type IDPairList []IDPair

func (il IDPairList) Keys() IDList {
	ids := IDList{}
	for _, i := range il {
		ids = append(ids, i[0])
	}
	return ids
}

func (il IDPairList) Values() IDList {
	ids := IDList{}
	for _, i := range il {
		ids = append(ids, i[1])
	}
	return ids
}

func (il IDPairList) Map() IDMap {
	mp := IDMap{}
	for _, i := range il {
		mp[i[0]] = i[1]
	}
	return mp
}

// Keyed ID Pair List

type KeyedIDPairList []KeyedIDPair

func (il KeyedIDPairList) Keys() IDList {
	ids := IDList{}
	for _, i := range il {
		ids = append(ids, i.Pair[0])
	}
	return ids
}

func (il KeyedIDPairList) Values() IDList {
	ids := IDList{}
	for _, i := range il {
		ids = append(ids, i.Pair[1])
	}
	return ids
}

func (il KeyedIDPairList) Map() IDMap {
	mp := IDMap{}
	for _, i := range il {
		mp[i.Pair[0]] = i.Pair[1]
	}
	return mp
}

// ID Bounds

var MaxBounds IDBounds

func init() {
	MaxBounds.Lower = ID{}
	MaxBounds.Upper = MustParseID("ffffffffffffffffffffffffffffffff")
}

type IDBounds struct {
	Lower ID
	Upper ID
}

func (ib IDBounds) String() string {
	var i int
	for i = IDSize - 1; i > 1; i-- {
		if ib.Lower[i] != 0 || ib.Upper[i] != 0 {
			break
		}
	}
	return hex.EncodeToString(ib.Lower[:i]) + "_" + hex.EncodeToString(ib.Upper[:i])
}

func (ib IDBounds) Contains(id ID) bool {
	if id.LessThan(ib.Lower) {
		return false
	}
	// From here onwards id >= Lower
	if ib.Upper.IsNull() {
		// If b.Upper is null, we interpret it as infinitely large, hence b.Upper > id
		return true
	}
	if ib.Upper.LessThan(id) || ib.Upper == id {
		// The upper bound is not part of the span
		return false
	}
	// From here onwards b.Upper > id
	return true
}

func (ib IDBounds) Middle() ID {
	middle := ID{}
	lower := ib.Lower
	upper := ib.Upper
	if upper.IsNull() {
		upper = MaxID
	}
	for i := 0; i < IDSize; i++ {
		if lower[i] == upper[i] {
			middle[i] = lower[i]
		} else {
			middle[i] = byte((int(lower[i]) + int(upper[i])) / 2)
			break
		}
	}
	return middle
}

func (ib IDBounds) Split() [2]IDBounds {
	m := ib.Middle()
	b1 := IDBounds{Lower: ib.Lower, Upper: m}
	b2 := IDBounds{Lower: m, Upper: ib.Upper}
	return [2]IDBounds{b1, b2}
}

func (ib IDBounds) NewID() ID {
	for {
		id := NewID()
		if ib.Contains(id) {
			return id
		}
	}
}

func ParseIDBounds(ibstr string) (IDBounds, error) {
	ibs := strings.Split(ibstr, "_")
	if len(ibs) != 2 {
		return IDBounds{}, errors.New("malformed ID bounds")
	}
	var err error
	ib := IDBounds{}
	ib.Lower, err = ParseID(ibs[0] + strings.Repeat("0", 2*IDSize-len(ibs[0])))
	if err != nil {
		return IDBounds{}, err
	}
	ib.Upper, err = ParseID(ibs[1] + strings.Repeat("0", 2*IDSize-len(ibs[1])))
	if err != nil {
		return IDBounds{}, err
	}
	return ib, nil
}

func (ib IDBounds) MarshalJSON() ([]byte, error) {
	return json.Marshal(ib.String())
}

func (ib *IDBounds) UnmarshalJSON(b []byte) error {
	var ibstr string
	err := json.Unmarshal(b, &ibstr)
	if err != nil {
		return err
	}
	*ib, err = ParseIDBounds(ibstr)
	return err
}

func (ib IDBounds) MarshalYAML() (interface{}, error) {
	return ib.String(), nil
}

func (ib *IDBounds) UnmarshalYAML(value *yaml.Node) error {
	var ibstr string
	err := value.Decode(&ibstr)
	if err != nil {
		return err
	}
	*ib, err = ParseIDBounds(ibstr)
	return err
}

type IDBoundGroup []IDBounds

func (idg IDBoundGroup) Contains(id ID) bool {
	for _, idg := range idg {
		if idg.Contains(id) {
			return true
		}
	}
	return false
}

func (idg *IDBoundGroup) Remove(b IDBounds) bool {
	idgn := IDBoundGroup{}
	removed := false
	for _, c := range *idg {
		if c != b && !removed {
			removed = true
			continue
		}
		idgn = append(idgn, c)
	}
	*idg = idgn
	return removed
}

func (idg *IDBoundGroup) Add(b IDBounds) {
	*idg = append(*idg, b)
}

func (idg *IDBoundGroup) AddAll(bs IDBoundGroup) {
	for _, b := range bs {
		idg.Add(b)
	}
}
