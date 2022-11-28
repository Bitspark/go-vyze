package vyze

import (
	"encoding/json"
	"fmt"
	"github.com/Bitspark/go-vyze/core"
	"github.com/Bitspark/go-vyze/service"
	"github.com/Bitspark/go-vyze/system"
	"strings"
)

type Client struct {
	Service          *service.ServiceClient
	System           *system.SystemClient
	Universes        map[string]*system.Universe
	SelectedUniverse string
}

func NewClient(serviceClient *service.ServiceClient, systemClient *system.SystemClient) Client {
	return Client{
		Service:   serviceClient,
		System:    systemClient,
		Universes: map[string]*system.Universe{},
	}
}

func (v *Client) LoadUniverse(name string) (*system.Universe, error) {
	univID, err := v.Service.ResolveUniverse(name)
	if err != nil {
		return nil, err
	}
	univ, err := v.Service.LoadUniverse(univID)
	if err != nil {
		return nil, err
	}
	v.Universes[name] = &univ
	v.SelectedUniverse = name
	return &univ, nil
}

func (v *Client) Universe() *system.Universe {
	return v.Universes[v.SelectedUniverse]
}

// Objects

type Q[N any] struct {
	Type      system.EndpointType
	Client    Client
	Universe  *system.Universe
	ModelID   core.ID
	ModelNode system.Node
	objectID  core.ID
	orders    []system.Order
	filters   []system.Filter
	limit     *int
	offset    *int
	err       error
}

func QueryReference[N any](c Client, ep string) Q[N] {
	u := c.Universe()
	ndEp := u.GetEndpoint(ep)
	if ndEp == nil {
		return Q[N]{
			err: fmt.Errorf("endpoint not found: %s", ep),
		}
	}
	q := Q[N]{
		Type:      ndEp.Type,
		Client:    c,
		Universe:  u,
		ModelNode: ndEp.Node,
	}
	q = q.Model(ndEp.Context.Environment.Model)
	if id, ok := ndEp.Context.Value.(core.ID); ok {
		q.ModelID = id
	}
	return q
}

func (q Q[N]) Object(id core.ID) Q[N] {
	if q.err != nil {
		return q
	}
	q.objectID = id
	return q
}

func (q Q[N]) Model(model string) Q[N] {
	if q.err != nil {
		return q
	}
	m := q.Universe.GetModel(model, q.Universe.Name)
	if m != nil && m.ObjectID != nil {
		q.ModelID = *m.ObjectID
	} else {
		q.err = fmt.Errorf("model not found: %s", model)
	}
	return q
}

func (q Q[N]) Filter(path string, operator system.OperatorType, value any) Q[N] {
	if q.Type != system.EndpointTypeGet {
		q.err = fmt.Errorf("wrong endpoint type: %s (filter requires get)", q.Type)
		return q
	}
	pathSource, err := q.pathSource(path)
	if err != nil {
		q.err = err
		return q
	}
	if q.err != nil {
		return q
	}
	q.filters = append(q.filters, system.Filter{
		Source:   *pathSource,
		Operator: operator,
		Value:    value,
	})
	return q
}

func (q Q[N]) Equals(path string, value any) Q[N] {
	return q.Filter(path, system.OperatorTypeEqual, value)
}

func (q Q[N]) Sort(path string, asc bool) Q[N] {
	if q.Type != system.EndpointTypeGet {
		q.err = fmt.Errorf("wrong endpoint type: %s (sort requires get)", q.Type)
		return q
	}
	pathSource, err := q.pathSource(path)
	if err != nil {
		q.err = err
		return q
	}
	if q.err != nil {
		return q
	}
	q.orders = append(q.orders, system.Order{
		Source:     *pathSource,
		Descending: !asc,
	})
	return q
}

func (q Q[N]) Limit(limit int) Q[N] {
	if q.Type != system.EndpointTypeGet {
		q.err = fmt.Errorf("wrong endpoint type: %s (limit requires get)", q.Type)
		return q
	}
	q.limit = &limit
	return q
}

func (q Q[N]) Offset(offset int) Q[N] {
	if q.Type != system.EndpointTypeGet {
		q.err = fmt.Errorf("wrong endpoint type: %s (offset requires get)", q.Type)
		return q
	}
	q.offset = &offset
	return q
}

func (q Q[N]) Slice(offset int, limit int) Q[N] {
	return q.Offset(offset).Limit(limit)
}

func (q Q[N]) GetObject() (N, error) {
	var respObj N
	if q.err != nil {
		return respObj, q.err
	}
	if q.Type != system.EndpointTypeGet {
		return respObj, fmt.Errorf("wrong endpoint type: %s (expected get)", q.Type)
	}
	nd, err := q.objectNode().Resolve(*q.Universe)
	if err != nil {
		return respObj, err
	}
	val, err := q.Client.System.GetNode(nd, nil)
	if err != nil {
		return respObj, err
	}
	bts, err := json.Marshal(val)
	if err != nil {
		return respObj, err
	}
	if err := json.Unmarshal(bts, &respObj); err != nil {
		return respObj, err
	}
	return respObj, nil
}

func (q Q[N]) GetObjects() ([]N, error) {
	if q.err != nil {
		return nil, q.err
	}
	if q.Type != system.EndpointTypeGet {
		return nil, fmt.Errorf("wrong endpoint type: %s (expected get)", q.Type)
	}
	nd, err := q.listNode().Resolve(*q.Universe)
	if err != nil {
		return nil, err
	}
	val, err := q.Client.System.GetNode(nd, nil)
	if err != nil {
		return nil, err
	}
	bts, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	var respObjs []N
	if err := json.Unmarshal(bts, &respObjs); err != nil {
		return nil, err
	}
	return respObjs, nil
}

func (q Q[N]) PutObject(obj N) (N, error) {
	var respObj N
	if q.err != nil {
		return respObj, q.err
	}
	if q.Type != system.EndpointTypePut {
		return respObj, fmt.Errorf("wrong endpoint type: %s (expected put)", q.Type)
	}
	nd, err := q.objectNode().Resolve(*q.Universe)
	if err != nil {
		return respObj, err
	}
	val, err := q.Client.System.PutNode(nd, obj, "", nil)
	if err != nil {
		return respObj, err
	}
	bts, err := json.Marshal(val)
	if err != nil {
		return respObj, err
	}
	if err := json.Unmarshal(bts, &respObj); err != nil {
		return respObj, err
	}
	return respObj, nil
}

func (q Q[N]) PutObjects(objs []N) ([]N, error) {
	if q.err != nil {
		return nil, q.err
	}
	if q.Type != system.EndpointTypePut {
		return nil, fmt.Errorf("wrong endpoint type: %s (expected put)", q.Type)
	}
	nd, err := q.listNode().Resolve(*q.Universe)
	if err != nil {
		return nil, err
	}
	val, err := q.Client.System.PutNode(nd, objs, "", nil)
	if err != nil {
		return nil, err
	}
	bts, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	var respObjs []N
	if err := json.Unmarshal(bts, &respObjs); err != nil {
		return nil, err
	}
	return respObjs, nil
}

func (q Q[N]) pathSource(path string) (*system.ValueSource, error) {
	node, format := cutMapNodes(strings.Split(path, "."), q.ModelNode)
	return &system.ValueSource{
		Type:   system.SourceTypeNode,
		Format: format,
		Node:   &node,
	}, nil
}

func (q Q[N]) objectNode() system.Node {
	// Embed into context
	var node system.Node

	if !q.objectID.IsNull() {
		node = system.Node{
			Type: system.NodeTypeContext,
			Context: &system.ContextNode{
				Context: system.Context{
					Value: q.objectID,
				},
				Node: q.ModelNode,
			},
		}
	} else {
		node = system.Node{
			Type: system.NodeTypeContext,
			Context: &system.ContextNode{
				Context: system.Context{
					Value: q.ModelID,
				},
				Node: system.Node{
					Type: system.NodeTypeSpecials,
					Specials: &system.SpecialsNode{
						Type:     system.EnvironmentTypePrimitive,
						Direct:   true,
						Indirect: true,
						Node:     q.ModelNode,
					},
				},
			},
		}
	}

	return node
}

func (q Q[N]) listNode() system.Node {
	// The following node nesting will be executed in reverse order

	// Prepare list node (last piece)
	node := system.Node{
		Type: system.NodeTypeList,
		List: &system.ListNode{
			Entry: q.ModelNode,
		},
	}

	// Slicing
	if q.limit != nil || q.offset != nil {
		node = system.Node{
			Type: system.NodeTypeSlice,
			Slice: &system.SliceNode{
				Offset: q.offset,
				Limit:  q.limit,
				Node:   node,
			},
		}
	}

	// Ordering
	for _, o := range q.orders {
		node = system.Node{
			Type: system.NodeTypeSort,
			Sort: &system.SortNode{
				Order: o,
				Node:  node,
			},
		}
	}

	// Filtering
	for _, f := range q.filters {
		node = system.Node{
			Type: system.NodeTypeFilter,
			Filter: &system.FilterNode{
				Filter: f,
				Node:   node,
			},
		}
	}

	// Embed into context and specials node (first piece)
	node = system.Node{
		Type: system.NodeTypeContext,
		Context: &system.ContextNode{
			Context: system.Context{
				Value: q.ModelID,
			},
			Node: system.Node{
				Type: system.NodeTypeSpecials,
				Specials: &system.SpecialsNode{
					Type:     system.EnvironmentTypeList,
					Direct:   true,
					Indirect: true,
					Node:     node,
				},
			},
		},
	}

	return node
}

func cutMapNodes(path []string, iterNode system.Node) (system.Node, system.FormatType) {
	iterType := iterNode.Type

	if iterType == "map" {
		if len(path) == 0 {
			panic("Path leads to a map. Add a field by attaching it with a dot (.)")
		}

		for _, e := range iterNode.Map.Entries {
			if e.Name == path[0] {
				return cutMapNodes(path[1:], e.Node)
			}
		}
		panic("Entry not found: {path[0]}")
	} else if iterType == "value" {
		if len(path) != 0 {
			panic("Path exceeds the node structure. Leftover path: ...{" + strings.Join(path, "."))
		}
		return iterNode, iterNode.Value.Format
	} else if iterType == "relation" {
		returnNode, format := cutMapNodes(path, iterNode.Relation.Node)
		iterNode = system.Node{
			Type: system.NodeTypeRelation,
			Relation: &system.RelationNode{
				Type:     iterNode.Relation.Type,
				Relation: iterNode.Relation.Relation,
				Reverse:  iterNode.Relation.Reverse,
				Node:     returnNode,
			},
		}
		return iterNode, format
	} else if iterType == "specials" {
		returnNode, format := cutMapNodes(path, iterNode.Specials.Node)
		iterNode = system.Node{
			Type: system.NodeTypeRelation,
			Specials: &system.SpecialsNode{
				Type:     iterNode.Specials.Type,
				Direct:   iterNode.Specials.Direct,
				Indirect: iterNode.Specials.Indirect,
				Node:     returnNode,
			},
		}
		return iterNode, format
	}

	panic("Unsupported node type: {iterType}")
}
