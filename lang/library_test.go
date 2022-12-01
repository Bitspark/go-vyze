package lang

import (
	"github.com/Bitspark/go-vyze/service"
	"github.com/Bitspark/go-vyze/system"
	"github.com/Bitspark/go-vyze/vyze"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestLibrary_File1(t *testing.T) {
	univ := &system.Universe{}
	err := univ.LoadFile("./test/vergleichsportal.vyu")
	if err != nil {
		t.Fatal(err)
	}

	l := NewLibrary(univ)
	source, _ := os.ReadFile("./test/queries1.vy")
	errs := l.Parse(string(source))
	if errs != nil {
		t.Fatal(errs[0].Err)
	}

	if len(l.Pipes) != 8 {
		t.Fatal()
	}
}

func TestLibrary_Pipe1(t *testing.T) {
	l := NewLibrary(&system.Universe{})
	pipe, errs := l.ParsePipe("on test -> @id")
	if len(errs) != 1 {
		t.Fatal(errs)
	}
	if pipe == nil || pipe.Node == nil {
		t.Fatal()
	}

	if pipe.Node.Type != system.NodeTypeValue {
		t.Fatal()
	}
	if pipe.Node.Value.Field != system.FieldTypeID {
		t.Fatal()
	}
	if pipe.Node.Value.Format != system.FormatTypeHex {
		t.Fatal()
	}
}

func TestLibrary_Pipe2(t *testing.T) {
	l := NewLibrary(&system.Universe{})
	pipe, errs := l.ParsePipe("on test -> {@id, @name}")
	if len(errs) != 1 {
		t.Fatal(errs)
	}
	if pipe == nil || pipe.Node == nil {
		t.Fatal()
	}

	if pipe.Node.Type != system.NodeTypeMap {
		t.Fatal()
	}
	if len(pipe.Node.Map.Entries) != 2 {
		t.Fatal()
	}

	e1 := pipe.Node.Map.Entries[0]
	if e1.Name != "id" {
		t.Fatal()
	}
	if e1.Node.Type != system.NodeTypeValue || e1.Node.Value.Field != system.FieldTypeID || e1.Node.Value.Format != system.FormatTypeHex {
		t.Fatal()
	}

	e2 := pipe.Node.Map.Entries[1]
	if e2.Name != "name" {
		t.Fatal()
	}
	if e2.Node.Type != system.NodeTypeValue || e2.Node.Value.Field != system.FieldTypeName || e2.Node.Value.Format != system.FormatTypeString {
		t.Fatal()
	}
}

func TestLibrary_Pipe3(t *testing.T) {
	l := NewLibrary(&system.Universe{})
	pipe, errs := l.ParsePipe("on test -> {id2: @id, name2: @name}")
	if len(errs) != 1 {
		t.Fatal(errs)
	}
	if pipe == nil || pipe.Node == nil {
		t.Fatal()
	}

	if pipe.Node.Type != system.NodeTypeMap {
		t.Fatal()
	}
	if len(pipe.Node.Map.Entries) != 2 {
		t.Fatal()
	}

	e1 := pipe.Node.Map.Entries[0]
	if e1.Name != "id2" {
		t.Fatal()
	}
	if e1.Node.Type != system.NodeTypeValue || e1.Node.Value.Field != system.FieldTypeID || e1.Node.Value.Format != system.FormatTypeHex {
		t.Fatal()
	}

	e2 := pipe.Node.Map.Entries[1]
	if e2.Name != "name2" {
		t.Fatal()
	}
	if e2.Node.Type != system.NodeTypeValue || e2.Node.Value.Field != system.FieldTypeName || e2.Node.Value.Format != system.FormatTypeString {
		t.Fatal()
	}
}

func TestLibrary_Pipe4(t *testing.T) {
	l := NewLibrary(&system.Universe{})
	pipe, errs := l.ParsePipe("on test -> {@id, field1 -> {@id, @name, createdAt: @created}}")
	if len(errs) != 1 {
		t.Fatal(errs)
	}
	if pipe == nil || pipe.Node == nil {
		t.Fatal(errs)
	}

	if pipe.Node.Type != system.NodeTypeMap {
		t.Fatal()
	}
	if len(pipe.Node.Map.Entries) != 2 {
		t.Fatal()
	}

	e1 := pipe.Node.Map.Entries[0]
	if e1.Name != "id" {
		t.Fatal()
	}
	if e1.Node.Type != system.NodeTypeValue || e1.Node.Value.Field != system.FieldTypeID || e1.Node.Value.Format != system.FormatTypeHex {
		t.Fatal()
	}

	e2 := pipe.Node.Map.Entries[1]
	if e2.Name != "field1" {
		t.Fatal()
	}
	if e2.Node.Type != system.NodeTypeRelation {
		t.Fatal()
	}
	if e2.Node.Relation.Type != system.EnvironmentTypePrimitive {
		t.Fatal()
	}
	if e2.Node.Relation.Node.Type != system.NodeTypeMap {
		t.Fatal()
	}
	if len(e2.Node.Relation.Node.Map.Entries) != 3 {
		t.Fatal()
	}
}

func TestLibrary_Pipe5(t *testing.T) {
	l := NewLibrary(&system.Universe{})
	pipe, errs := l.ParsePipe("on base.object/ -> {@id, fieldNew: base.object#field1/ -> [] {@id, @name, createdAt: @created}}")
	if len(errs) != 1 {
		t.Fatal(errs)
	}
	if pipe == nil || pipe.Node == nil {
		t.Fatal()
	}

	if pipe.Node.Type != system.NodeTypeMap {
		t.Fatal()
	}
	if len(pipe.Node.Map.Entries) != 2 {
		t.Fatal()
	}

	e1 := pipe.Node.Map.Entries[0]
	if e1.Name != "id" {
		t.Fatal()
	}
	if e1.Node.Type != system.NodeTypeValue || e1.Node.Value.Field != system.FieldTypeID || e1.Node.Value.Format != system.FormatTypeHex {
		t.Fatal()
	}

	e2 := pipe.Node.Map.Entries[1]
	if e2.Name != "fieldNew" {
		t.Fatal()
	}
	if e2.Node.Type != system.NodeTypeRelation {
		t.Fatal()
	}
	if e2.Node.Relation.Type != system.EnvironmentTypeList {
		t.Fatal()
	}
	if e2.Node.Relation.Relation != "" {
		t.Fatal()
	}
	if e2.Node.Relation.Reverse {
		t.Fatal()
	}
	if e2.Node.Relation.Node.Type != system.NodeTypeList {
		t.Fatal(errs)
	}
	if e2.Node.Relation.Node.List.Entry.Type != system.NodeTypeMap {
		t.Fatal(errs)
	}
	if len(e2.Node.Relation.Node.List.Entry.Map.Entries) != 3 {
		t.Fatal()
	}

	t.Log(l.models.String())
}

func TestLibrary_Pipe6(t *testing.T) {
	l := NewLibrary(&system.Universe{})
	pipe, errs := l.ParsePipe("on base.object/ -> {@id, fieldNew: <- base.object2#field2/ [] {@id, @name, createdAt: @created}}")
	if len(errs) != 1 {
		t.Fatal(errs)
	}
	if pipe == nil || pipe.Node == nil {
		t.Fatal()
	}

	if pipe.Node.Type != system.NodeTypeMap {
		t.Fatal()
	}
	if len(pipe.Node.Map.Entries) != 2 {
		t.Fatal()
	}

	e1 := pipe.Node.Map.Entries[0]
	if e1.Name != "id" {
		t.Fatal()
	}
	if e1.Node.Type != system.NodeTypeValue || e1.Node.Value.Field != system.FieldTypeID || e1.Node.Value.Format != system.FormatTypeHex {
		t.Fatal()
	}

	e2 := pipe.Node.Map.Entries[1]
	if e2.Name != "fieldNew" {
		t.Fatal()
	}
	if e2.Node.Type != system.NodeTypeRelation {
		t.Fatal()
	}
	if e2.Node.Relation.Type != system.EnvironmentTypeList {
		t.Fatal()
	}
	if e2.Node.Relation.Relation != "" {
		t.Fatal()
	}
	if !e2.Node.Relation.Reverse {
		t.Fatal()
	}
	if e2.Node.Relation.Node.Type != system.NodeTypeList {
		t.Fatal(errs)
	}
	if e2.Node.Relation.Node.List.Entry.Type != system.NodeTypeMap {
		t.Fatal(errs)
	}
	if len(e2.Node.Relation.Node.List.Entry.Map.Entries) != 3 {
		t.Fatal()
	}

	t.Log(l.models.String())
}

func TestLibrary_Uni1(t *testing.T) {
	univ := &system.Universe{}
	err := univ.LoadFile("./test/vergleichsportal.vyu")
	if err != nil {
		t.Fatal(err)
	}

	lib := NewLibrary(univ)
	if lib == nil {
		t.Fatal()
	}

	pipe, errs := lib.ParsePipe("on distributorprice -> {@id, @name}")
	if errs != nil {
		t.Fatal(errs)
	}
	ymlNode, _ := yaml.Marshal(pipe.Node)
	t.Log(string(ymlNode))
}

func TestLibrary_Uni2(t *testing.T) {
	client := vyze.NewClient(
		service.NewServiceClient("https://api.vyze.io/service"),
		system.NewSystemClient("https://api.vyze.io/system"),
	)

	client.Service.SetToken("mt82e3R0cSsMOiRS9TLGpFRAhWMAAAAAAI0nAAEABQAAAAUqLyovKmFF0rv8uE0sB_7aG9xpoUr2WpGq")
	lp, _ := system.ReadLayerProfile("main_full:1ffffff:9adf367b7474712b0c3a2452f532c6a4962031f02a7ec76482a9a3b1ba4c890901ffffff0000000000000000000000006385405400278cff6aa81f46257a30f270a296b7c37c89b958245650;main_read:4924ea:9adf367b7474712b0c3a2452f532c6a48ce700d22de0697b5f22dacd84842941004924ea0000000000000000000000006385405400278cff9886d3816a193c56eb4628189ea679165e8b4abb,9adf367b7474712b0c3a2452f532c6a4962031f02a7ec76482a9a3b1ba4c8909004924ea0000000000000000000000006385405400278cffc4f0d01fda6c1b13b565b5ac7a090489847382be;model_extend:492cea:9adf367b7474712b0c3a2452f532c6a457780ea25d5c139b1132dd66ecaa910a00492cea0000000000000000000000006385405400278cff69b28871abd93a7891a13301a0f1725bf32a5ed1")
	client.System.SetLayerProfile(lp)
	client.System.SetDefaultOptions(&system.AccessOptions{
		Access:      "main_full",
		AccessNames: []string{"main_full", "main_read", "model_read", "model_extend"},
	})

	univ, err := client.LoadUniverse("vergleichsportal")
	if err != nil {
		t.Fatal(err)
	}

	lib := NewLibrary(univ)
	if lib == nil {
		t.Fatal()
	}

	pipe, errs := lib.ParsePipe(`on distributorprice -> {
		@id, 
		@name, 
		@created,
		price -> @value,
		item -> {@id, @name}, 
		distributor -> @name
	}`)
	if errs != nil {
		t.Fatal(errs)
	}
	query := QueryNode[any](client, *pipe, "get")
	rlts, err := query.GetObjects()
	if err != nil {
		t.Fatal(err)
	}
	for _, rlt := range rlts {
		t.Log(rlt)
	}

	t.Log(">>>")
	pipe, errs = lib.ParsePipe(`on item -> {
		@id, 
		itemId: item_id -> @value,
		<- distributorprice#item [] {
			distributor -> @name,
			price -> @value
		}
	}`)
	//yamlPipe, _ := yaml.Marshal(pipe)
	//t.Log(string(yamlPipe))
	query = QueryNode[any](client, *pipe, "get")
	rlts, err = query.GetObjects()
	if err != nil {
		t.Fatal(err)
	}
	for _, rlt := range rlts {
		t.Log(rlt)
	}
	t.Log("<<<")

	pipe, errs = lib.ParsePipe(`on item -> {
		@id,
		user: @id,
		item_id -> @string,
		prices: <- distributorprice#item [] price -> @value
	}`)
	query = QueryNode[any](client, *pipe, "get")
	rlts, err = query.GetObjects()
	if err != nil {
		t.Fatal(err)
	}
	for _, rlt := range rlts {
		t.Log(rlt)
	}
}

func TestLibrary_Uni3(t *testing.T) {
	client := vyze.NewClient(
		service.NewServiceClient("https://api.vyze.io/service"),
		system.NewSystemClient("https://api.vyze.io/system"),
	)

	client.Service.SetToken("mt82e3R0cSsMOiRS9TLGpFRAhWMAAAAAAI0nAAEABQAAAAUqLyovKmFF0rv8uE0sB_7aG9xpoUr2WpGq")
	lp, _ := system.ReadLayerProfile("main_full:1ffffff:9adf367b7474712b0c3a2452f532c6a4962031f02a7ec76482a9a3b1ba4c890901ffffff0000000000000000000000006385405400278cff6aa81f46257a30f270a296b7c37c89b958245650;main_read:4924ea:9adf367b7474712b0c3a2452f532c6a48ce700d22de0697b5f22dacd84842941004924ea0000000000000000000000006385405400278cff9886d3816a193c56eb4628189ea679165e8b4abb,9adf367b7474712b0c3a2452f532c6a4962031f02a7ec76482a9a3b1ba4c8909004924ea0000000000000000000000006385405400278cffc4f0d01fda6c1b13b565b5ac7a090489847382be;model_extend:492cea:9adf367b7474712b0c3a2452f532c6a457780ea25d5c139b1132dd66ecaa910a00492cea0000000000000000000000006385405400278cff69b28871abd93a7891a13301a0f1725bf32a5ed1")
	client.System.SetLayerProfile(lp)
	client.System.SetDefaultOptions(&system.AccessOptions{
		Access:      "main_full",
		AccessNames: []string{"main_full", "main_read", "model_read", "model_extend"},
	})

	univ, _ := client.LoadUniverse("vergleichsportal")
	lib := NewLibrary(univ)

	pipe, errs := lib.ParsePipe(`on crawljob -> {
		id,
		item -> [] {
			item_id -> @auto
		}
	}`)
	if errs != nil {
		t.Fatal(errs)
	}
	query := QueryNode[any](client, *pipe, "get")
	rlts, err := query.GetObjects()
	if err != nil {
		t.Fatal(err)
	}
	for _, rlt := range rlts {
		t.Log(rlt)
	}
}
