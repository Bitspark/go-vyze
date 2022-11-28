package lang

import (
	"github.com/Bitspark/go-vyze/system"
	"github.com/Bitspark/go-vyze/vyze"
)

type Pipe struct {
	Node  *system.Node
	Model system.UniverseObjectInfo
}

func QueryNode[N any](c vyze.Client, p Pipe, tp system.EndpointType) vyze.Q[N] {
	u := c.Universe()
	q := vyze.Q[N]{
		Type:      tp,
		Client:    c,
		Universe:  u,
		ModelNode: *p.Node,
	}
	m := p.Model
	if m.ObjectID != nil && !m.ObjectID.IsNull() {
		q.ModelID = *m.ObjectID
	}
	return q
}
