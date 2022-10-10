package go_vyze

import (
	"fmt"
)

// NodeType defines the type of node
type NodeType string

const (
	NodeTypeEndpoint  NodeType = "endpoint"
	NodeTypeContext   NodeType = "context"
	NodeTypeRelation  NodeType = "relation"
	NodeTypeSpecials  NodeType = "specials"
	NodeTypeValue     NodeType = "value"
	NodeTypeInstance  NodeType = "instance"
	NodeTypeGroup     NodeType = "group"
	NodeTypeList      NodeType = "list"
	NodeTypeFilter    NodeType = "filter"
	NodeTypeSort      NodeType = "sort"
	NodeTypeSlice     NodeType = "slice"
	NodeTypeAggregate NodeType = "aggregate"
	NodeTypeMap       NodeType = "map"
	NodeTypeReference NodeType = "reference"
)

// ContextType defines the type of context
type ContextType string

const (
	ContextTypeBlank    ContextType = "blank"
	ContextTypeSpecials ContextType = "specials"
	ContextTypeRelation ContextType = "relation"
)

// AggregateType defines the type of aggregation
type AggregateType string

const (
	AggregateTypeCount   AggregateType = "count"
	AggregateTypeMinimum AggregateType = "minimum"
	AggregateTypeMaximum AggregateType = "maximum"
	AggregateTypeSum     AggregateType = "sum"
	AggregateTypeAverage AggregateType = "average"
)

// EnvironmentType defines the type of environment
type EnvironmentType string

const (
	EnvironmentTypePrimitive EnvironmentType = "primitive"
	EnvironmentTypeList      EnvironmentType = "list"
	EnvironmentTypeKeyedList EnvironmentType = "keyed_list"
)

// OperatorType defines the operator to be used for filters
type OperatorType string

type SourceType string

const (
	SourceTypeNode SourceType = "node"
	SourceTypeKey  SourceType = "key"
)

const (
	OperatorTypeEqual        OperatorType = "eq"
	OperatorTypeGreater      OperatorType = "gt"
	OperatorTypeGreaterEqual OperatorType = "ge"
	OperatorTypeLess         OperatorType = "lt"
	OperatorTypeLessEqual    OperatorType = "le"
	OperatorTypeNotEqual     OperatorType = "ne"
)

// FILTER

type Filter struct {
	Source   ValueSource  `json:"source"`
	Operator OperatorType `json:"operator"`
	Value    any          `json:"value"`
}

// ORDER

type Order struct {
	Source     ValueSource `json:"source"`
	Descending bool        `json:"descending"`
}

// CONTEXT

type Context struct {
	Environment Environment `json:"environment"`
	Value       any         `json:"value"`
}

// NODE

type Node struct {
	Type      NodeType       `json:"type"`
	Endpoint  *EndpointNode  `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Reference *ReferenceNode `json:"reference,omitempty" yaml:"reference,omitempty"`
	Context   *ContextNode   `json:"context,omitempty" yaml:"context,omitempty"`
	Relation  *RelationNode  `json:"relation,omitempty" yaml:"relation,omitempty"`
	Specials  *SpecialsNode  `json:"specials,omitempty" yaml:"specials,omitempty"`
	Value     *ValueNode     `json:"value,omitempty" yaml:"value,omitempty"`
	Instance  *InstanceNode  `json:"instance,omitempty" yaml:"instance,omitempty"`
	Group     *GroupNode     `json:"group,omitempty" yaml:"group,omitempty"`
	List      *ListNode      `json:"list,omitempty" yaml:"list,omitempty"`
	Filter    *FilterNode    `json:"filter,omitempty" yaml:"filter,omitempty"`
	Sort      *SortNode      `json:"sort,omitempty" yaml:"sort,omitempty"`
	Slice     *SliceNode     `json:"slice,omitempty" yaml:"slice,omitempty"`
	Aggregate *AggregateNode `json:"aggregate,omitempty" yaml:"aggregate,omitempty"`
	Map       *MapNode       `json:"map,omitempty" yaml:"map,omitempty"`
}

type Environment struct {
	Type  EnvironmentType `json:"type"`
	Model string          `json:"model,omitempty"`
}

type Parameter struct {
	Name        string    `json:"name" yaml:"name"`
	Description string    `json:"description" yaml:"description"`
	Interface   Interface `json:"interface" yaml:"interface"`
}

type EndpointType string

const (
	EndpointTypeGet EndpointType = "get"
	EndpointTypePut EndpointType = "put"
)

// EndpointNode injects a constant context into the first node
type EndpointNode struct {
	Type       EndpointType `json:"type" yaml:"type"`
	ID         ID           `json:"id" yaml:"id"`
	Name       string       `json:"name" yaml:"name"`
	Node       Node         `json:"node" yaml:"node"`
	Context    Context      `json:"context,omitempty" yaml:"context,omitempty"`
	Interface  *Interface   `json:"interface,omitempty" yaml:"interface,omitempty"`
	Parameters []Parameter  `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// ReferenceNode is replaced by its reference, thereby accepting and producing the same contexts as the referenced node
type ReferenceNode struct {
	Name string `json:"name"`
}

type ContextNode struct {
	Context Context `json:"context" yaml:"context"`
	Node    Node    `json:"node" yaml:"node"`
}

type SpecialsNode struct {
	Type     EnvironmentType `json:"type" yaml:"type"`
	Direct   bool            `json:"direct" yaml:"direct"`
	Indirect bool            `json:"indirect" yaml:"indirect"`
	Node     Node            `json:"node" yaml:"node"`
}

type RelationNode struct {
	Type     EnvironmentType `json:"type" yaml:"type"`
	Relation string          `json:"relation" yaml:"relation"`
	Reverse  bool            `json:"reverse" yaml:"reverse"`
	Node     Node            `json:"node" yaml:"node"`

	PutAdd        bool `json:"putAdd" yaml:"putAdd"`
	PutDuplicates bool `json:"putDuplicates" yaml:"putDuplicates"`
}

// ValueNode turns a primitive context into a single value, based on direct information (id, name, value, user, ...)
// about the context value
type ValueNode struct {
	Field  FieldType  `json:"field"`
	Format FormatType `json:"format"`
}

type InstanceSwitch struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}

// InstanceNode turns a primitive context into a single value, based on the abstracts of the context value
type InstanceNode struct {
	Format   FormatType       `json:"format"`
	Switches []InstanceSwitch `json:"switches"`
}

// GroupNode turns a list into a group context
type GroupNode struct {
	Source ValueSource `json:"source"`
	Node   Node        `json:"node"`
}

// ListNode turns a list or a keyed list context into a primitive context
type ListNode struct {
	Entry     Node       `json:"entry"`
	KeyFormat FormatType `json:"keyFormat,omitempty" yaml:"keyFormat,omitempty"`
	KeyName   string     `json:"keyName,omitempty" yaml:"keyName,omitempty"`
	ValueName string     `json:"valueName,omitempty" yaml:"valueName,omitempty"`
}

type ValueSource struct {
	Type   SourceType `json:"type" yaml:"type"`
	Format FormatType `json:"format" yaml:"format"`
	Node   *Node      `json:"node" yaml:"node"`
}

// FilterNode modifies a list or a keyed list context by removing elements
type FilterNode struct {
	Filter Filter `json:"filter"`
	Node   Node   `json:"node"`
}

// SortNode modifies a list or a keyed list context by changing the order of elements
type SortNode struct {
	Order Order `json:"order"`
	Node  Node  `json:"node"`
}

// SliceNode modifies a list or a keyed list context by taking a slice of successive elements from the list
type SliceNode struct {
	Offset *int `json:"offset,omitempty" yaml:"offset,omitempty"`
	Limit  *int `json:"limit,omitempty" yaml:"limit,omitempty"`
	Node   Node `json:"node"`
}

// AggregateNode turns a list or a keyed list context into a single value
type AggregateNode struct {
	Source   ValueSource   `json:"source"`
	Function AggregateType `json:"function"`
}

type MapNodeEntry struct {
	Name string `json:"name" yaml:"name"`
	Node Node   `json:"node" yaml:"node"`
}

// MapNode turns a primitive context into a map of multiple primitive contexts
type MapNode struct {
	Entries []MapNodeEntry `json:"entries"`
}

func (en MapNode) GetEntry(name string) *MapNodeEntry {
	for _, e := range en.Entries {
		if e.Name == name {
			return &e
		}
	}
	return nil
}

func (n Node) Resolve(universe Universe) (Node, error) {
	resNode := Node{
		Type: n.Type,
	}
	switch n.Type {
	case NodeTypeContext:
		childNode, err := n.Context.Node.Resolve(universe)
		if err != nil {
			return Node{}, err
		}
		resNode.Context = &ContextNode{
			Context: n.Context.Context,
			Node:    childNode,
		}
		if str, ok := resNode.Context.Context.Value.(string); ok {
			resNode.Context.Context.Value = universe.Resolve(str, universe.Name)
		} else if strs, ok := resNode.Context.Context.Value.([]string); ok {
			context := IDList{}
			for _, str := range strs {
				context = append(context, universe.Resolve(str, universe.Name))
			}
			resNode.Context.Context.Value = context
		}
	case NodeTypeRelation:
		childNode, err := n.Relation.Node.Resolve(universe)
		if err != nil {
			return Node{}, err
		}
		relationID := universe.Resolve(n.Relation.Relation, "")
		if relationID.IsNull() {
			return Node{}, fmt.Errorf("could not resolve '%s'", n.Relation.Relation)
		}
		resNode.Relation = &RelationNode{
			Type:     n.Relation.Type,
			Relation: relationID.Hex(),
			Reverse:  n.Relation.Reverse,
			Node:     childNode,
		}
	case NodeTypeSpecials:
		childNode, err := n.Specials.Node.Resolve(universe)
		if err != nil {
			return Node{}, err
		}
		resNode.Specials = &SpecialsNode{
			Type:     n.Specials.Type,
			Direct:   n.Specials.Direct,
			Indirect: n.Specials.Indirect,
			Node:     childNode,
		}
	case NodeTypeValue:
		resNode.Value = &ValueNode{
			Field:  n.Value.Field,
			Format: n.Value.Format,
		}
	case NodeTypeInstance:
		resNode.Instance = &InstanceNode{
			Switches: []InstanceSwitch{},
		}
		for _, s := range n.Instance.Switches {
			typeID := universe.Resolve(s.Type, "")
			resNode.Instance.Switches = append(resNode.Instance.Switches, InstanceSwitch{
				Type:  typeID.Hex(),
				Value: s.Value,
			})
		}
	case NodeTypeGroup:
		childNode, err := n.Group.Node.Resolve(universe)
		if err != nil {
			return Node{}, err
		}
		groupNode := &GroupNode{
			Source: ValueSource{
				Type:   n.Filter.Filter.Source.Type,
				Format: n.Filter.Filter.Source.Format,
			},
			Node: childNode,
		}
		if n.Group.Source.Node != nil {
			sourceNode, err := n.Group.Source.Node.Resolve(universe)
			if err != nil {
				return Node{}, err
			}
			groupNode.Source.Node = &sourceNode
		}
		resNode.Group = groupNode
	case NodeTypeList:
		childNode, err := n.List.Entry.Resolve(universe)
		if err != nil {
			return Node{}, err
		}
		resNode.List = &ListNode{
			Entry:     childNode,
			KeyFormat: n.List.KeyFormat,
			KeyName:   n.List.KeyName,
			ValueName: n.List.ValueName,
		}
	case NodeTypeFilter:
		childNode, err := n.Filter.Node.Resolve(universe)
		if err != nil {
			return Node{}, err
		}
		filterNode := &FilterNode{
			Filter: Filter{
				Source: ValueSource{
					Type:   n.Filter.Filter.Source.Type,
					Format: n.Filter.Filter.Source.Format,
				},
				Operator: n.Filter.Filter.Operator,
				Value:    n.Filter.Filter.Value,
			},
			Node: childNode,
		}
		if n.Filter.Filter.Source.Node != nil {
			sourceNode, err := n.Filter.Filter.Source.Node.Resolve(universe)
			if err != nil {
				return Node{}, err
			}
			filterNode.Filter.Source.Node = &sourceNode
		}
		resNode.Filter = filterNode
	case NodeTypeSort:
		childNode, err := n.Sort.Node.Resolve(universe)
		if err != nil {
			return Node{}, err
		}
		sortNode := &SortNode{
			Order: Order{
				Source: ValueSource{
					Type:   n.Sort.Order.Source.Type,
					Format: n.Sort.Order.Source.Format,
				},
				Descending: n.Sort.Order.Descending,
			},
			Node: childNode,
		}
		if n.Sort.Order.Source.Node != nil {
			sourceNode, err := n.Sort.Order.Source.Node.Resolve(universe)
			if err != nil {
				return Node{}, err
			}
			sortNode.Order.Source.Node = &sourceNode
		}
		resNode.Sort = sortNode
	case NodeTypeSlice:
		childNode, err := n.Slice.Node.Resolve(universe)
		if err != nil {
			return Node{}, err
		}
		resNode.Slice = &SliceNode{
			Offset: n.Slice.Offset,
			Limit:  n.Slice.Limit,
			Node:   childNode,
		}
	case NodeTypeAggregate:
		aggNode := &AggregateNode{
			Source: ValueSource{
				Type:   n.Aggregate.Source.Type,
				Format: n.Aggregate.Source.Format,
			},
		}
		if n.Sort.Order.Source.Node != nil {
			sourceNode, err := n.Aggregate.Source.Node.Resolve(universe)
			if err != nil {
				return Node{}, err
			}
			aggNode.Source.Node = &sourceNode
		}
		resNode.Aggregate = aggNode
	case NodeTypeMap:
		resNode.Map = &MapNode{Entries: []MapNodeEntry{}}
		for _, v := range n.Map.Entries {
			childNode, err := v.Node.Resolve(universe)
			if err != nil {
				return Node{}, err
			}
			resNode.Map.Entries = append(resNode.Map.Entries, MapNodeEntry{Name: v.Name, Node: childNode})
		}
	case NodeTypeReference:
		return universe.GetEndpoint(n.Reference.Name).Node.Resolve(universe)
	case NodeTypeEndpoint:
		childNode, err := n.Endpoint.Node.Resolve(universe)
		if err != nil {
			return Node{}, err
		}
		resNode.Endpoint = &EndpointNode{
			ID:   n.Endpoint.ID,
			Name: n.Endpoint.Name,
			Node: childNode,
			Context: Context{
				Environment: Environment{
					Type:  n.Endpoint.Context.Environment.Type,
					Model: universe.Resolve(n.Endpoint.Context.Environment.Model, "").Hex(),
				},
				Value: n.Endpoint.Context.Value,
			},
		}
	}
	return resNode, nil
}

// Requests

type GetNodeRequest struct {
	Node Node `json:"node"`
}

type PutNodeRequest struct {
	Node  Node `json:"node"`
	Value any  `json:"value"`
	Layer ID   `json:"layer"`
}
