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

func (v *Client) LoadUniverse(name string) error {
	univID, err := v.Service.ResolveUniverse(name)
	if err != nil {
		return err
	}
	univ, err := v.Service.LoadUniverse(univID)
	if err != nil {
		return err
	}
	v.Universes[name] = &univ
	v.SelectedUniverse = name
	return nil
}

func (v *Client) Universe() *system.Universe {
	return v.Universes[v.SelectedUniverse]
}

// Objects

type Q[N any] struct {
	epType    system.EndpointType
	client    Client
	universe  *system.Universe
	modelID   core.ID
	modelNode system.Node
	objectID  core.ID
	orders    []system.Order
	filters   []system.Filter
	limit     *int
	offset    *int
	err       error
}

func Query[N any](c Client, ep string) Q[N] {
	u := c.Universe()
	ndEp := u.GetEndpoint(ep)
	if ndEp == nil {
		return Q[N]{
			err: fmt.Errorf("endpoint not found: %s", ep),
		}
	}
	q := Q[N]{
		epType:    ndEp.Type,
		client:    c,
		universe:  u,
		modelNode: ndEp.Node,
	}
	q = q.Model(ndEp.Context.Environment.Model)
	if id, ok := ndEp.Context.Value.(core.ID); ok {
		q.modelID = id
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
	m := q.universe.GetModel(model, q.universe.Name)
	if m != nil && m.ObjectID != nil {
		q.modelID = *m.ObjectID
	} else {
		q.err = fmt.Errorf("model not found: %s", model)
	}
	return q
}

func (q Q[N]) Filter(path string, operator system.OperatorType, value any) Q[N] {
	if q.epType != system.EndpointTypeGet {
		q.err = fmt.Errorf("wrong endpoint type: %s (filter requires get)", q.epType)
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
	if q.epType != system.EndpointTypeGet {
		q.err = fmt.Errorf("wrong endpoint type: %s (sort requires get)", q.epType)
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
	if q.epType != system.EndpointTypeGet {
		q.err = fmt.Errorf("wrong endpoint type: %s (limit requires get)", q.epType)
		return q
	}
	q.limit = &limit
	return q
}

func (q Q[N]) Offset(offset int) Q[N] {
	if q.epType != system.EndpointTypeGet {
		q.err = fmt.Errorf("wrong endpoint type: %s (offset requires get)", q.epType)
		return q
	}
	q.offset = &offset
	return q
}

func (q Q[N]) Slice(offset int, limit int) Q[N] {
	return q.Offset(offset).Limit(limit)
}

func (q Q[N]) GetObject() (*N, error) {
	if q.err != nil {
		return nil, q.err
	}
	if q.epType != system.EndpointTypeGet {
		return nil, fmt.Errorf("wrong endpoint type: %s (expected get)", q.epType)
	}
	nd, err := q.objectNode().Resolve(*q.universe)
	if err != nil {
		return nil, err
	}
	val, err := q.client.System.GetNode(nd, nil)
	if err != nil {
		return nil, err
	}
	bts, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	var respObj N
	if err := json.Unmarshal(bts, &respObj); err != nil {
		return nil, err
	}
	return &respObj, nil
}

func (q Q[N]) GetObjects() ([]N, error) {
	if q.err != nil {
		return nil, q.err
	}
	if q.epType != system.EndpointTypeGet {
		return nil, fmt.Errorf("wrong endpoint type: %s (expected get)", q.epType)
	}
	nd, err := q.listNode().Resolve(*q.universe)
	if err != nil {
		return nil, err
	}
	val, err := q.client.System.GetNode(nd, nil)
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

func (q Q[N]) PutObject(obj *N) (*N, error) {
	if q.err != nil {
		return nil, q.err
	}
	if q.epType != system.EndpointTypePut {
		return nil, fmt.Errorf("wrong endpoint type: %s (expected put)", q.epType)
	}
	nd, err := q.objectNode().Resolve(*q.universe)
	if err != nil {
		return nil, err
	}
	val, err := q.client.System.PutNode(nd, *obj, "", nil)
	if err != nil {
		return nil, err
	}
	bts, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	var respObj N
	if err := json.Unmarshal(bts, &respObj); err != nil {
		return nil, err
	}
	return &respObj, nil
}

func (q Q[N]) PutObjects(objs []N) ([]N, error) {
	if q.err != nil {
		return nil, q.err
	}
	if q.epType != system.EndpointTypePut {
		return nil, fmt.Errorf("wrong endpoint type: %s (expected put)", q.epType)
	}
	nd, err := q.listNode().Resolve(*q.universe)
	if err != nil {
		return nil, err
	}
	val, err := q.client.System.PutNode(nd, objs, "", nil)
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
	pathNode := q.modelNode
	var sourceNode *system.Node
	var format system.FormatType
	curNode := &system.Node{}
	paths := strings.Split(path, ".")

	if len(paths) > 0 && paths[0] != "" {
		paths = append(paths, "")
	}

	descendPath := func(nd system.Node, ndPtr *system.Node) {
		if sourceNode == nil {
			sourceNode = &nd
			curNode = ndPtr
		} else {
			*curNode = nd
			curNode = ndPtr
		}
	}

	for _, p := range paths {
		if p != "" {
			// Map
			if pathNode.Type == system.NodeTypeMap {
				for _, e := range pathNode.Map.Entries {
					if e.Name == p {
						pathNode = e.Node
						goto descendMap
					}
				}
				return nil, fmt.Errorf("entry not found: %s", p)
			}
		}

		for {
			// Value
			if pathNode.Type == system.NodeTypeValue {
				nd := system.Node{
					Type: system.NodeTypeValue,
					Value: &system.ValueNode{
						Field:  pathNode.Value.Field,
						Format: pathNode.Value.Format,
					},
				}
				format = pathNode.Value.Format
				descendPath(nd, nil)
				break
			}

			// Aggregate
			if pathNode.Type == system.NodeTypeAggregate {
				nd := system.Node{
					Type: system.NodeTypeAggregate,
					Aggregate: &system.AggregateNode{
						Source:   pathNode.Aggregate.Source,
						Function: pathNode.Aggregate.Function,
					},
				}
				format = pathNode.Aggregate.Source.Format
				descendPath(nd, nil)
				break
			}

			// Instance
			if pathNode.Type == system.NodeTypeInstance {
				nd := system.Node{
					Type: system.NodeTypeInstance,
					Instance: &system.InstanceNode{
						Format:   pathNode.Instance.Format,
						Switches: pathNode.Instance.Switches,
					},
				}
				format = pathNode.Instance.Format
				descendPath(nd, nil)
				break
			}

			// Relation
			if pathNode.Type == system.NodeTypeRelation {
				nd := system.Node{
					Type: system.NodeTypeRelation,
					Relation: &system.RelationNode{
						Type:     pathNode.Relation.Type,
						Relation: pathNode.Relation.Relation,
						Reverse:  pathNode.Relation.Reverse,
					},
				}
				pathNode = pathNode.Relation.Node
				descendPath(nd, &nd.Relation.Node)
				continue
			}

			// Specials
			if pathNode.Type == system.NodeTypeSpecials {
				nd := system.Node{
					Type: system.NodeTypeSpecials,
					Specials: &system.SpecialsNode{
						Type:     "",
						Direct:   false,
						Indirect: false,
					},
				}
				pathNode = pathNode.Specials.Node
				descendPath(nd, &nd.Specials.Node)
				continue
			}

			return nil, fmt.Errorf("cannot descend further: %s, %v", p, pathNode)
		}

	descendMap:
	}

	return &system.ValueSource{
		Type:   system.SourceTypeNode,
		Format: format,
		Node:   sourceNode,
	}, nil
}

func (q Q[N]) objectNode() system.Node {
	// Embed into context
	node := system.Node{
		Type: system.NodeTypeContext,
		Context: &system.ContextNode{
			Context: system.Context{
				Value: q.objectID,
			},
			Node: q.modelNode,
		},
	}

	return node
}

func (q Q[N]) listNode() system.Node {
	// The following node nesting will be executed in reverse order

	// Prepare list node (last piece)
	node := system.Node{
		Type: system.NodeTypeList,
		List: &system.ListNode{
			Entry: q.modelNode,
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
				Value: q.modelID,
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
