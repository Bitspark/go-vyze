package vyze

import (
	"testing"
)

func TestCreateMapping1(t *testing.T) {
	res := struct {
		ID     ID     `json:"id1" vyze:"id"`
		Name   string `json:"name1" vyze:"name"`
		IDs    IDList `json:"ids" vyze:"id,relation"`
		String string `json:"string" vyze:"data,relation,mapping123(change:ignore)"`
	}{}

	rs, err := ExtractSchema(res, nil, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(rs.Fields) != 4 {
		t.Fatal()
	}
	if rs.Fields[0].Name != "id1" {
		t.Fatal()
	}
	if rs.Fields[1].Name != "name1" {
		t.Fatal()
	}
	if rs.Fields[2].Name != "ids" {
		t.Fatal()
	}
	if rs.Fields[3].Name != "string" {
		t.Fatal()
	}
	if rs.Fields[3].Mapping != "mapping123" {
		t.Fatal()
	}
	if rs.Fields[3].GetMappingParam("change", "").(string) != "ignore" {
		t.Fatal()
	}

	rs, err = ExtractSchema(&res, nil, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(rs.Fields) != 4 {
		t.Fatal()
	}
	if rs.Fields[0].Name != "id1" {
		t.Fatal()
	}
	if rs.Fields[1].Name != "name1" {
		t.Fatal()
	}
	if rs.Fields[2].Name != "ids" {
		t.Fatal()
	}
	if rs.Fields[3].Name != "string" {
		t.Fatal()
	}
	if rs.Fields[3].Mapping != "mapping123" {
		t.Fatal()
	}
	if rs.Fields[3].GetMappingParam("change", "").(string) != "ignore" {
		t.Fatal()
	}
}

func TestWriteResource(t *testing.T) {
	id := NewID()
	i := int64(123)
	st := struct {
		ID      *ID      `vyze:"id" json:"id"`
		Name    string   `vyze:"name" json:"name"`
		Size    *int64   `vyze:"size" json:"size"`
		Strings []string `vyze:"data,rel" json:"strings"`
	}{
		ID:      &id,
		Name:    "test123",
		Size:    &i,
		Strings: []string{"a", "b", "c"},
	}
	vall, err := WriteResource(&st)
	val := vall.(map[string]any)
	if err != nil {
		t.Fatal(err)
	}
	if val["id"].(string) != id.Hex() {
		t.Fatal()
	}
	if val["name"].(string) != "test123" {
		t.Fatal()
	}
	if val["size"].(float64) != 123 {
		t.Fatal()
	}
	if val["strings"].([]interface{})[0].(string) != "a" {
		t.Fatal()
	}
}

func TestReadResource(t *testing.T) {
	st := struct {
		ID   ID     `vyze:"id" json:"id"`
		Name string `vyze:"name" json:"name"`
		Size int64  `vyze:"size" json:"size"`
	}{}
	id := NewID()
	err := ReadResource(map[string]any{"id": id.Hex(), "name": "test123", "size": float64(123)}, &st)
	if err != nil {
		t.Fatal(err)
	}
	if st.ID != id {
		t.Fatal()
	}
	if st.Name != "test123" {
		t.Fatal()
	}
	if st.Size != 123 {
		t.Fatal()
	}
}

func TestMustReadResource(t *testing.T) {
	type stType struct {
		ID   ID     `vyze:"id" json:"id"`
		Name string `vyze:"name" json:"name"`
		Size int64  `vyze:"size" json:"size"`
	}
	id := NewID()
	st := MustReadResource(map[string]any{"id": id.Hex(), "name": "test123", "size": float64(123)}, stType{})
	if st.ID != id {
		t.Fatal()
	}
	if st.Name != "test123" {
		t.Fatal()
	}
	if st.Size != 123 {
		t.Fatal()
	}
}

func TestWriteReadResource1(t *testing.T) {
	id := NewID()
	i := int64(123)
	type ts struct {
		ID   ID     `vyze:"id" json:"id"`
		Name string `vyze:"name" json:"name"`
		Size int64  `vyze:"size" json:"size"`
	}
	st := ts{
		ID:   id,
		Name: "test123",
		Size: i,
	}
	val, _ := WriteResource(&st)
	st2 := ts{}
	_ = ReadResource(val, &st2)
	if st.ID != st2.ID {
		t.Fatal()
	}
}

func TestWriteReadResource2(t *testing.T) {
	id := NewID()
	i := int64(123)
	type ts struct {
		ID      *ID       `vyze:"id" json:"id"`
		Name    string    `vyze:"name" json:"name"`
		Size    int64     `vyze:"size" json:"size"`
		Strings []*string `vyze:"data,rel" json:"strings"`
	}
	a := "a"
	b := "b"
	c := "c"
	st := ts{
		ID:      &id,
		Name:    "test123",
		Size:    i,
		Strings: []*string{&a, &b, &c},
	}
	val, _ := WriteResource([]ts{st, st, st})
	st22 := []ts{}
	_ = ReadResource(val, &st22)
	if len(st22) != 3 {
		t.Fatal()
	}
	st2 := st22[0]
	if st.ID.Hex() != st2.ID.Hex() {
		t.Fatal()
	}
	if len(st2.Strings) != 3 {
		t.Fatal()
	}
	if *st2.Strings[0] != "a" {
		t.Fatal()
	}
}

func TestWriteReadResource3(t *testing.T) {
	id := NewID()
	i1 := int64(123)
	i2 := uint8(255)
	type ts struct {
		ID      *ID       `vyze:"id" json:"id"`
		IDs1    []ID      `vyze:"id,rel5" json:"ids1"`
		IDs2    []*ID     `vyze:"id,rel6" json:"ids2"`
		Name    string    `vyze:"name" json:"name"`
		Int1    int64     `vyze:"data,rel" json:"int1"`
		Int2    *uint8    `vyze:"data,rel" json:"int2"`
		Strings []*string `vyze:"data,rel" json:"strings"`
		Bool    bool      `vyze:"data,rel2" json:"bool"`
	}
	a := "a"
	b := "b"
	c := "c"
	st := ts{
		ID:      &id,
		IDs1:    []ID{id, id},
		IDs2:    []*ID{&id, &id},
		Name:    "test123",
		Int1:    i1,
		Int2:    &i2,
		Strings: []*string{&a, &b, &c},
	}
	val, _ := WriteResource([]ts{st, st, st})
	st22 := []ts{}
	_ = ReadResource(val, &st22)
	if len(st22) != 3 {
		t.Fatal()
	}
	st2 := st22[0]
	if st.ID.Hex() != id.Hex() {
		t.Fatal()
	}
	if len(st2.Strings) != 3 {
		t.Fatal()
	}
	if *st2.Strings[0] != "a" {
		t.Fatal()
	}
	if len(st2.IDs1) != 2 {
		t.Fatal()
	}
	if st2.IDs1[0] != id {
		t.Fatal()
	}
	if len(st2.IDs2) != 2 {
		t.Fatal()
	}
	if st2.IDs2[0].Hex() != id.Hex() {
		t.Fatal()
	}
}
