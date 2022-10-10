package vyze

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Client struct {
	Service          *ServiceClient
	System           *SystemClient
	Universes        map[string]*Universe
	SelectedUniverse string
}

func NewClient(serviceClient *ServiceClient, systemClient *SystemClient) Client {
	return Client{
		Service:   serviceClient,
		System:    systemClient,
		Universes: map[string]*Universe{},
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

func (v Client) Universe() *Universe {
	return v.Universes[v.SelectedUniverse]
}

func GetNode[N any](v Client, ep string, args ...any) (N, error) {
	var n N
	endpointNode := v.Universe().GetEndpoint(ep)
	if endpointNode == nil {
		return n, fmt.Errorf("endpoint not found: %s", ep)
	}
	if endpointNode.Context.Value == nil && len(args) > 0 {
		if endpointNode.Context.Environment.Type == EnvironmentTypePrimitive {
			endpointNode.Context.Value = args[0].(ID)
		} else if endpointNode.Context.Environment.Type == EnvironmentTypeList {
			endpointNode.Context.Value = args[0].(IDList)
		}
	}
	nd1 := Node{
		Type: NodeTypeContext,
		Context: &ContextNode{
			Context: endpointNode.Context,
			Node:    endpointNode.Node,
		},
	}
	nd2, err := nd1.Resolve(*v.Universe())
	if err != nil {
		return n, err
	}
	val, err := v.System.GetNode(nd2, nil)
	if err != nil {
		return n, err
	}
	bts, err := json.Marshal(val)
	if err != nil {
		return n, err
	}
	if err := json.Unmarshal(bts, &n); err != nil {
		return n, err
	}
	return n, nil
}

func PutNode[N any](v Client, ep string, n *N, args ...any) error {
	endpointNode := v.Universe().GetEndpoint(ep)
	if endpointNode == nil {
		return fmt.Errorf("endpoint not found: %s", ep)
	}
	if endpointNode.Context.Value == nil && len(args) > 0 {
		if endpointNode.Context.Environment.Type == EnvironmentTypePrimitive {
			endpointNode.Context.Value = args[0].(ID)
		} else if endpointNode.Context.Environment.Type == EnvironmentTypeList {
			endpointNode.Context.Value = args[0].(IDList)
		}
	}
	nd1 := Node{
		Type: NodeTypeContext,
		Context: &ContextNode{
			Context: endpointNode.Context,
			Node:    endpointNode.Node,
		},
	}
	nd2, err := nd1.Resolve(*v.Universe())
	if err != nil {
		return err
	}
	val, err := v.System.PutNode(nd2, *n, "", nil)
	if err != nil {
		return err
	}
	bts, err := json.Marshal(val)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bts, n); err != nil {
		return err
	}
	return nil
}

func GetSpecials[N any](v Client, ep string) ([]N, error) {
	var n []N
	nd := v.Universe().GetEndpoint(ep)
	if nd == nil {
		return n, errors.New("could not find node definition")
	}
	entryNode, err := nd.Node.Resolve(*v.Universe())
	if err != nil {
		return n, err
	}
	modelID := v.Universe().Resolve(nd.Context.Environment.Model, "")
	if modelID.IsNull() {
		return n, errors.New("could not find entry model definition")
	}
	outerNode := Node{
		Type: NodeTypeList,
		List: &ListNode{
			Entry: entryNode,
		},
	}
	nd2 := Node{
		Type: NodeTypeContext,
		Context: &ContextNode{
			Context: Context{Value: modelID},
			Node: Node{
				Type: NodeTypeSpecials,
				Specials: &SpecialsNode{
					Type:     EnvironmentTypeList,
					Direct:   true,
					Indirect: true,
					Node:     outerNode,
				},
			},
		},
	}
	val, err := v.System.GetNode(nd2, nil)
	if err != nil {
		return n, err
	}
	bts, err := json.Marshal(val)
	if err != nil {
		return n, err
	}
	if err := json.Unmarshal(bts, &n); err != nil {
		return n, err
	}
	return n, nil
}

func PutSpecial[N any](v Client, ep string, n *N, args ...any) error {
	nd := v.Universe().GetEndpoint(ep)
	if nd == nil {
		return errors.New("could not find node definition")
	}
	epNode, err := nd.Node.Resolve(*v.Universe())
	if err != nil {
		return err
	}
	modelID := v.Universe().Resolve(nd.Context.Environment.Model, "")
	if modelID.IsNull() {
		return errors.New("could not find entry model definition")
	}
	nd2 := Node{
		Type: NodeTypeContext,
		Context: &ContextNode{
			Context: Context{Value: modelID},
			Node: Node{
				Type: NodeTypeSpecials,
				Specials: &SpecialsNode{
					Type:     EnvironmentTypePrimitive,
					Direct:   true,
					Indirect: true,
					Node:     epNode,
				},
			},
		},
	}
	val, err := v.System.PutNode(nd2, n, "", nil)
	if err != nil {
		return err
	}
	bts, err := json.Marshal(val)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bts, n); err != nil {
		return err
	}
	return nil
}

func PutSpecials[N any](v Client, ep string, n *[]N, args ...any) error {
	nd := v.Universe().GetEndpoint(ep)
	if nd == nil {
		return errors.New("could not find node definition")
	}
	entryNode, err := nd.Node.Resolve(*v.Universe())
	if err != nil {
		return err
	}
	modelID := v.Universe().Resolve(nd.Context.Environment.Model, "")
	if modelID.IsNull() {
		return errors.New("could not find entry model definition")
	}
	nd2 := Node{
		Type: NodeTypeContext,
		Context: &ContextNode{
			Context: Context{Value: modelID},
			Node: Node{
				Type: NodeTypeSpecials,
				Specials: &SpecialsNode{
					Type:     EnvironmentTypeList,
					Direct:   true,
					Indirect: true,
					Node: Node{
						Type: NodeTypeList,
						List: &ListNode{
							Entry: entryNode,
						},
					},
				},
			},
		},
	}
	val, err := v.System.PutNode(nd2, n, "", nil)
	if err != nil {
		return err
	}
	bts, err := json.Marshal(val)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bts, n); err != nil {
		return err
	}
	return nil
}
