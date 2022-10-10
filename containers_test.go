package go_vyze

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNewID(t *testing.T) {
	id := NewID()
	if id.Hex() == strings.Repeat("00", IDSize) {
		t.Fatal("wrong id")
	}
	id2 := NewID()
	if id.Equals(id2) {
		t.Fatal("id collision")
	}
}

func TestID_Hex(t *testing.T) {
	id := ID{}
	if id.Hex() != strings.Repeat("00", IDSize) {
		t.Fatal("wrong id string")
	}
}

func TestID_Hex_Lowercase(t *testing.T) {
	id := NewID()
	if id.Hex() != strings.ToLower(id.Hex()) {
		t.Fatal("should be lowercase")
	}
}

func TestID_ParseHex(t *testing.T) {
	id := NewID()
	idStr := id.Hex()
	id2, err := ParseIDHex(idStr)
	if err != nil {
		t.Fatal(err)
	}
	if !id.Equals(id2) {
		t.Fatal()
	}
}

func TestID_ParseBase32(t *testing.T) {
	id := NewID()
	idStr := id.Base32()
	id2, err := ParseIDBase32(idStr)
	if err != nil {
		t.Fatal(err)
	}
	if !id.Equals(id2) {
		t.Fatal()
	}
}

func TestID_Base64(t *testing.T) {
	id := ID{}
	if id.Base64() != strings.Repeat("A", (IDSize*4-1)/3+1) {
		t.Fatal(id.Base64())
	}
}

func TestID_ParseBase64(t *testing.T) {
	id := NewID()
	idStr := id.Base64()
	id2, err := ParseIDBase64(idStr)
	if err != nil {
		t.Fatal(err)
	}
	if !id.Equals(id2) {
		t.Fatal()
	}
}

func TestID_Base62_1(t *testing.T) {
	id := ID{}
	if id.Base62() != strings.Repeat("A", (IDSize*4-1)/3+1) {
		t.Fatal(id.Base62())
	}
}

func TestID_Base62_2(t *testing.T) {
	id := NewID()
	id62 := id.Base62()
	id64 := id.Base64()
	if id62 == id64 {
		t.Fatal(id62, id64)
	}
}

func TestID_ParseBase62(t *testing.T) {
	id := NewID()
	idStr := id.Base62()
	id2, err := ParseIDBase62(idStr)
	if err != nil {
		t.Fatal(err)
	}
	if !id.Equals(id2) {
		t.Fatal()
	}
}

func TestID_IsNull(t *testing.T) {
	id := ID{}
	if !id.IsNull() {
		t.Fatal("should be null")
	}
	id = NewID()
	if id.IsNull() {
		t.Fatal("should not be null")
	}
}

func TestID_Invert(t *testing.T) {
	id := NewID()
	idInv := id.Invert()

	if id.Equals(idInv) {
		t.Fatal()
	}
	if !id.Invert().Equals(idInv) {
		t.Fatal()
	}
}

func TestID_LessThan(t *testing.T) {
	id := NewID()

	// Not reflexive
	if id.LessThan(id) {
		t.Fatal()
	}

	// Correctly related with NULL id
	if id.LessThan(ID{}) {
		t.Fatal()
	}

	// Transitivity
	for i := 0; i < 32; i++ {
		id2 := NewID()
		id3 := NewID()

		if id.LessThan(id2) {
			if id2.LessThan(id3) {
				if !id.LessThan(id3) {
					t.Fatal()
				}
			} else {
				if !id.LessThan(id2) {
					t.Fatal()
				}
			}
		} else {
			if id.LessThan(id3) {
				if !id2.LessThan(id3) {
					t.Fatal()
				}
			} else {
				if !id2.LessThan(id) {
					t.Fatal()
				}
			}
		}
	}
}

func TestB2ID(t *testing.T) {
	b := make(Binary, IDSize)
	id := B2ID(b)
	if !id.IsNull() {
		t.Fatal()
	}

	id2 := NewID()
	b2 := Binary(id2[:])
	id22 := B2ID(b2)
	if !id22.Equals(id2) {
		t.Fatal()
	}
}

func TestMarshalEmptyID(t *testing.T) {
	str := struct {
		ID ID `json:"id"`
	}{}
	bts, err := json.Marshal(str)
	if err != nil {
		t.Fatal(err)
	}
	if string(bts) != "{\"id\":null}" {
		t.Fatal()
	}
	var mp any
	err = json.Unmarshal(bts, &mp)
	idr := mp.(map[string]any)["id"]
	if idr != nil {
		t.Fatal(idr)
	}
}

func TestObjectContainer_MarshalJSON(t *testing.T) {
	oc := IDSet{}
	id1 := NewID()
	oc[id1] = true

	b, err := json.Marshal(oc)
	if err != nil {
		t.Fatal(err)
	}

	oc2 := IDSet{}
	err = json.Unmarshal(b, &oc2)
	if err != nil {
		t.Fatal(err)
	}

	if len(oc2) != 1 || oc2[id1] != true {
		t.Fatal("unexpected content")
	}
}

func TestIDList_MarshalJSON(t *testing.T) {
	l := IDList{NewID(), NewID(), NewID()}

	lBytes, err := json.Marshal(l)
	if err != nil {
		t.Fatal(err)
	}

	l2 := IDList{}
	err = json.Unmarshal(lBytes, &l2)
	if err != nil {
		t.Fatal(err)
	}

	if len(l) != len(l2) {
		t.Fatal()
	}
}

func TestIDMap_MarshalYAML(t *testing.T) {
	l := IDMap{NewID(): NewID(), NewID(): NewID(), NewID(): NewID(), NewID(): NewID()}

	lBytes, err := yaml.Marshal(l)
	if err != nil {
		t.Fatal(err)
	}

	l2 := IDMap{}
	err = yaml.Unmarshal(lBytes, &l2)
	if err != nil {
		t.Fatal(err)
	}

	if len(l) != len(l2) {
		t.Fatal()
	}
}

func TestIDBounds_Contains(t *testing.T) {
	b := IDBounds{}
	if !b.Contains(MinID) {
		t.Fatal()
	}
	if !b.Contains(MaxID) {
		t.Fatal()
	}
	if !b.Contains(NewID()) {
		t.Fatal()
	}
}

func TestIDBounds_Split(t *testing.T) {
	b := IDBounds{}
	bs := b.Split()
	if !bs[0].Contains(MinID) {
		t.Fatal()
	}
	if bs[1].Contains(MinID) {
		t.Fatal()
	}
	if bs[0].Contains(MaxID) {
		t.Fatal()
	}
	if !bs[1].Contains(MaxID) {
		t.Fatal()
	}
	if bs[0].Contains(bs[1].Lower) {
		t.Fatal()
	}
	if !bs[1].Contains(bs[0].Upper) {
		t.Fatal()
	}
	if bs[0].Contains(bs[0].Upper) {
		t.Fatal()
	}
	if bs[1].Contains(bs[1].Upper) {
		t.Fatal()
	}
}

func TestIDBounds_NewID(t *testing.T) {
	b := IDBounds{}
	bs := b.Split()
	id := bs[1].NewID()
	if id.LessThan(bs[1].Lower) {
		t.Fatal()
	}
}

func TestIDBounds_Parsing(t *testing.T) {
	idb := IDBounds{}
	idbs := idb.Split()
	str := idbs[0].String()
	idb2, err := ParseIDBounds(str)
	if err != nil {
		t.Fatal(err)
	}
	if idb2.Lower != idbs[0].Lower || idb2.Upper != idbs[0].Upper {
		t.Fatal()
	}
}

func TestIDBounds_JSON(t *testing.T) {
	idb := IDBounds{}
	idbs := idb.Split()
	j1, err := json.Marshal(idbs[0])
	if err != nil {
		t.Fatal(err)
	}
	idb2 := IDBounds{}
	err = json.Unmarshal(j1, &idb2)
	if err != nil {
		t.Fatal(err)
	}
	if idb2.Lower != idbs[0].Lower || idb2.Upper != idbs[0].Upper {
		t.Fatal()
	}
}

func TestIDBounds_YAML(t *testing.T) {
	idb := IDBounds{}
	idbs := idb.Split()
	j1, err := yaml.Marshal(idbs[0])
	if err != nil {
		t.Fatal(err)
	}
	idb2 := IDBounds{}
	err = yaml.Unmarshal(j1, &idb2)
	if err != nil {
		t.Fatal(err)
	}
	if idb2.Lower != idbs[0].Lower || idb2.Upper != idbs[0].Upper {
		t.Fatal()
	}
}

func TestRichID_Hex(t *testing.T) {
	id := NewID()
	rid := WrapID(id, 123)
	h := rid.Hex()
	if len(h) != 32+1+2 {
		t.Fatal()
	}
}

func TestParseRID(t *testing.T) {
	id := NewID()
	rid := WrapID(id, 123)
	h := rid.Hex()
	rid2, err := ParseRID(h)
	if err != nil {
		t.Fatal(err)
	}
	if rid2 != rid {
		t.Fatal()
	}
}

func TestRichID_Flag(t *testing.T) {
	id := NewID()
	rid := WrapID(id, 0)

	if rid.HasFlag(1) {
		t.Fatal()
	}
	rid.SetFlag(1)
	if !rid.HasFlag(1) {
		t.Fatal()
	}

	if rid.HasFlag(2) {
		t.Fatal()
	}
	rid.SetFlag(2)
	if !rid.HasFlag(2) {
		t.Fatal()
	}

	rid.UnsetFlag(1)

	if rid.HasFlag(1) {
		t.Fatal()
	}
	if !rid.HasFlag(2) {
		t.Fatal()
	}

	rid.SetFlag(1)
	rid.UnsetFlag(2)

	if !rid.HasFlag(1) {
		t.Fatal()
	}
	if rid.HasFlag(2) {
		t.Fatal()
	}
}
