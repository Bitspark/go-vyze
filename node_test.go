package go_vyze

import (
	"testing"
)

func TestNodeDef_ResolveRelation(t *testing.T) {
	univ := Universe{}
	_ = univ.LoadFile("./test/vision.vyu")

	id := univ.Resolve("vision.box#dim_x/vision", "")
	if id.IsNull() {
		t.Fatal()
	}

	node1 := Node{
		Type: NodeTypeRelation,
		Relation: &RelationNode{
			Relation: "vision.box#dim_x/vision",
			Node: Node{
				Type: NodeTypeReference,
				Reference: &ReferenceNode{
					Name: "Box",
				},
			},
		},
	}

	node2, err := node1.Resolve(univ)
	if err != nil {
		t.Fatal(err)
	}
	if node2.Relation.Relation != id.Hex() {
		t.Fatal()
	}
	if node2.Relation.Node.Type != NodeTypeMap {
		t.Fatal()
	}
}

func TestNodeDef_ResolveReference(t *testing.T) {
	univ := Universe{}
	err := univ.LoadFile("./test/vision.vyu")
	if err != nil {
		t.Fatal(err)
	}

	id := univ.Resolve("vision.box#dim_x/vision", "")
	if id.IsNull() {
		t.Fatal()
	}
	id2 := univ.Resolve("vision.box#dim_x/vision", "")
	if id2.IsNull() {
		t.Fatal()
	}

	node1 := Node{
		Type: NodeTypeReference,
		Reference: &ReferenceNode{
			Name: "Box",
		},
	}

	node2, err := node1.Resolve(univ)
	if err != nil {
		t.Fatal(err)
	}
	if node2.Type != NodeTypeMap {
		t.Fatal()
	}
	if node2.Map.GetEntry("dimX").Node.Relation.Relation != id2.Hex() {
		t.Fatal()
	}
}
