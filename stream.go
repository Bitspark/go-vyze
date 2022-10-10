package vyze

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type CmdType string

const (
	CmdInfo                CmdType = "info"
	CmdRegisterLayerTokens CmdType = "registerLayerTokens"
	CmdGetLayerTokens      CmdType = "getLayerTokens"
	CmdSubscribe           CmdType = "subscribe"
	CmdUnsubscribe         CmdType = "unsubscribe"
	CmdEvent               CmdType = "event"
)

type EventType string

const (
	EventName     EventType = "name"
	EventData     EventType = "data"
	EventSpecials EventType = "specials"
	EventValue    EventType = "value"
	EventResource EventType = "resource"
)

type ChangeType string

const (
	CngChange ChangeType = "change"
	CngAdd    ChangeType = "add"
	CngRemove ChangeType = "remove"
)

// Structs

// StreamSubscriptionParams contains generic parameters
type StreamSubscriptionParams struct {
	Object   ID              `json:"object"`
	Specials bool            `json:"specials"`
	Payload  bool            `json:"payload"`  // Payload specifies if payload should be included in triggered events
	Instant  bool            `json:"instant"`  // Instant specifies if the current state should instantly be transmitted
	Interval int             `json:"interval"` // Interval specifies how long changes should be collected before firing an update
	Params   json.RawMessage `json:"params"`
}

type StreamMessage struct {
	// MessageID is the ID of this message
	MessageID ID `json:"messageId"`

	// ReferenceID is the ID of the referenced message
	ReferenceID ID `json:"referenceId"`

	// Command is the type of the message
	Command CmdType `json:"command"`

	// Payload contains the command-specific parameters
	Payload json.RawMessage `json:"payload"`
}

type StreamLayerTokens struct {
	Tokens []string `json:"tokens"`
}

type StreamUnsubscribe struct{}

type MessageInfo struct {
	// APIVersion represents the API version
	APIVersion string `json:"apiVersion"`

	// ServerVersion represents the version of the server
	ServerVersion string `json:"serverVersion"`

	// UnixTime contains the server time
	UnixTime int64 `json:"unixTime"`

	// Session provides information about the session
	Session SessionInfo `json:"session"`

	// Subscriptions contains the number of currently open subscriptions
	Subscriptions int `json:"subscriptions"`
}

type LayerTokens struct {
	Tokens []LayerTokenInfo `json:"tokens"`
}

type MessageEvent struct {
	Event json.RawMessage `json:"event"`
}

// MessageSubscribe binds a stream message and stream subscription request together
type MessageSubscribe struct {
	Event  EventType                `json:"event"`
	Params StreamSubscriptionParams `json:"params"`
}

type SessionInfo struct {
	Created  time.Time `json:"created"`
	RsvCmds  int       `json:"commandsReceived"`
	RsvBytes int       `json:"bytesReceived"`
	RsvLast  time.Time `json:"lastReceived"`
	SntCmds  int       `json:"commandsSent"`
	SntBytes int       `json:"bytesSent"`
	SntLast  time.Time `json:"lastSent"`
}

type StreamResourceParams struct {
	Object ResourceType    `json:"object"`
	Schema ResourceSchema  `json:"schema"`
	Filter ResourceFilters `json:"filter,omitempty"`
}

type StreamResourceEvent struct {
	ID      ID          `json:"id"`
	Type    ChangeType  `json:"type"`
	Removed interface{} `json:"removed,omitempty"`
	Added   interface{} `json:"added,omitempty"`
}

// Client

type Callback func(message StreamMessage)

type Subscription struct {
	ref ID
	cbs []Callback
}

func (sub *Subscription) Subscribe(cb Callback) {
	sub.cbs = append(sub.cbs, cb)
}

func (sub Subscription) next(val StreamMessage) {
	for _, cb := range sub.cbs {
		cb(val)
	}
}

type ResourceCallback func(message StreamResourceEvent)

type resourceSubscription interface {
	Subscribe(cb ResourceCallback)
	next(message StreamResourceEvent)
}

type ResourceStreamInstance struct {
	mux   sync.Mutex
	res   *ResourceInstance
	ref   ID
	cbs   []ResourceCallback
	value any
}

func addValues(schema ResourceSchema, ins any, val any) error {
	if val == nil {
		return nil
	}
	insMp := ins.(map[string]any)
	valMp := val.(map[string]any)
	for k, v := range valMp {
		field := schema.GetField(k)
		if field == nil {
			continue
		}
		if field.Mapping == MappingTypePrimitive {
			insMp[k] = v
			continue
		}
		if field.Mapping == MappingTypeList {
			eVals, _ := insMp[k]
			if eVals == nil {
				eVals = []any{}
			}
			eValArr := eVals.([]any)
			eValArr = append(eValArr, v)
			insMp[k] = eValArr
			continue
		}
	}
	return nil
}

func removeValues(schema ResourceSchema, ins any, val any) error {
	if val == nil {
		return nil
	}
	insMp := ins.(map[string]any)
	valMp := val.(map[string]any)
	for k, v := range valMp {
		field := schema.GetField(k)
		if field == nil {
			continue
		}
		if field.Mapping == MappingTypePrimitive {
			delete(insMp, k)
			continue
		}
		if field.Mapping == MappingTypeList {
			if insMp[k] == nil || v == nil {
				continue
			}
			for _, vr := range v.([]any) {
				// TODO: This can be done more efficiently
				newVals := []any{}
				removed := false
				for _, ve := range insMp[k].([]any) {
					if removed {
						newVals = append(newVals, ve)
						continue
					}
					switch field.Format {
					case FormatTypeString, FormatTypeHex, FormatTypeBase64:
						if ve.(string) == vr.(string) {
							removed = true
						}
					case FormatTypeBoolean:
						if ve.(bool) == vr.(bool) {
							removed = true
						}
					case FormatTypeInteger:
						if ve.(int) == vr.(int) {
							removed = true
						}
					case FormatTypeFloat:
						if ve.(float64) == vr.(float64) {
							removed = true
						}
					}
					if removed {
						continue
					}
					newVals = append(newVals, ve)
				}
				insMp[k] = newVals
			}
			continue
		}
	}
	return nil
}

var _ resourceSubscription = &ResourceStreamInstance{}

func (sub *ResourceStreamInstance) Subscribe(cb ResourceCallback) {
	sub.cbs = append(sub.cbs, cb)
}

func (sub *ResourceStreamInstance) Value() any {
	return sub.value
}

func (sub *ResourceStreamInstance) addValues(val any) error {
	if sub.value == nil {
		return errors.New("not found")
	}
	sub.mux.Lock()
	defer sub.mux.Unlock()
	return addValues(sub.res.Schema, sub.value, val)
}

func (sub *ResourceStreamInstance) removeValues(val any) error {
	if sub.value == nil {
		return errors.New("not found")
	}
	sub.mux.Lock()
	defer sub.mux.Unlock()
	return removeValues(sub.res.Schema, sub.value, val)
}

func (sub *ResourceStreamInstance) remove() error {
	sub.mux.Lock()
	defer sub.mux.Unlock()
	sub.value = nil
	return nil
}

func (sub *ResourceStreamInstance) add() (any, error) {
	sub.mux.Lock()
	defer sub.mux.Unlock()
	if sub.value != nil {
		return nil, errors.New("already exists")
	}
	sub.value = map[string]any{}
	return sub.value, nil
}

func (sub *ResourceStreamInstance) next(message StreamResourceEvent) {
	for _, cb := range sub.cbs {
		cb(message)
	}
}

type ResourceStreamList struct {
	mux     sync.Mutex
	res     *ResourceSpecials
	ref     ID
	cbs     []ResourceCallback
	values  []any
	ids     []ID
	indexes map[ID]int
}

var _ resourceSubscription = &ResourceStreamList{}

func (sub *ResourceStreamList) Subscribe(cb ResourceCallback) {
	sub.cbs = append(sub.cbs, cb)
}

func (sub *ResourceStreamList) GetValue(id ID) any {
	sub.mux.Lock()
	defer sub.mux.Unlock()
	idx, ok := sub.indexes[id]
	if !ok {
		return nil
	}
	return sub.values[idx]
}

func (sub *ResourceStreamList) Values() []any {
	return sub.values
}

func (sub *ResourceStreamList) addValues(id ID, val any) error {
	ins := sub.GetValue(id)
	if ins == nil {
		return errors.New("not found")
	}
	sub.mux.Lock()
	defer sub.mux.Unlock()
	return addValues(sub.res.Schema, ins, val)
}

func (sub *ResourceStreamList) removeValues(id ID, val any) error {
	ins := sub.GetValue(id)
	if ins == nil {
		return errors.New("not found")
	}
	sub.mux.Lock()
	defer sub.mux.Unlock()
	return removeValues(sub.res.Schema, ins, val)
}

func (sub *ResourceStreamList) add(id ID) (any, error) {
	sub.mux.Lock()
	defer sub.mux.Unlock()
	if _, ok := sub.indexes[id]; ok {
		return nil, errors.New("already exists")
	}
	ins := map[string]any{}
	sub.indexes[id] = len(sub.ids)
	sub.ids = append(sub.ids, id)
	sub.values = append(sub.values, ins)
	return ins, nil
}

func (sub *ResourceStreamList) remove(id ID) error {
	sub.mux.Lock()
	defer sub.mux.Unlock()
	idx, ok := sub.indexes[id]
	if !ok {
		return errors.New("not fonud")
	}
	newIDs := []ID{}
	newVals := []any{}
	for i, val := range sub.values {
		if i == idx {
			delete(sub.indexes, id)
			continue
		}
		sub.indexes[sub.ids[i]] = len(newIDs)
		newIDs = append(newIDs, sub.ids[i])
		newVals = append(newVals, val)
	}
	sub.ids = newIDs
	sub.values = newVals
	return nil
}

func (sub *ResourceStreamList) next(message StreamResourceEvent) {
	for _, cb := range sub.cbs {
		cb(message)
	}
}

type ResourceOptions struct {
	Instant  bool `json:"instant"`
	Interval int  `json:"interval"`
}

type SystemStreamClient struct {
	endpoint      string
	subscriptions map[ID]*Subscription
	conn          *websocket.Conn
	sendMux       sync.Mutex
}

func NewSystemStreamClient(endpoint string) *SystemStreamClient {
	c := &SystemStreamClient{
		endpoint:      endpoint,
		subscriptions: map[ID]*Subscription{},
	}
	return c
}

func (sc *SystemStreamClient) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(sc.endpoint+"/v1/stream", nil)
	if err != nil {
		return err
	}
	sc.conn = conn
	go func() {
		for {
			msg := StreamMessage{}
			err = sc.conn.ReadJSON(&msg)
			if err != nil {
				return
			}
			sub, ok := sc.subscriptions[msg.ReferenceID]
			if !ok {
				continue
			}
			sub.next(msg)
		}
	}()
	return nil
}

func (sc *SystemStreamClient) Disconnect() error {
	if sc.conn == nil {
		return errors.New("not connected")
	}
	return sc.conn.Close()
}

func (sc *SystemStreamClient) Connected() bool {
	return sc.conn != nil
}

func (sc *SystemStreamClient) Info() (chan MessageInfo, error) {
	if sc.conn == nil {
		return nil, errors.New("not connected")
	}
	sub := sc.subscribe()
	infoChan := make(chan MessageInfo)
	sub.Subscribe(func(event StreamMessage) {
		infoMsg := MessageInfo{}
		_ = json.Unmarshal(event.Payload, &infoMsg)
		infoChan <- infoMsg
		close(infoChan)
	})
	go func() {
		msg := StreamMessage{
			MessageID: sub.ref,
			Command:   CmdInfo,
		}
		sc.sendMux.Lock()
		_ = sc.conn.WriteJSON(msg)
		sc.sendMux.Unlock()
	}()
	return infoChan, nil
}

func (sc *SystemStreamClient) RegisterLayerToken(token LayerToken) error {
	if sc.conn == nil {
		return errors.New("not connected")
	}
	sub := sc.subscribe()
	go func() {
		msg := StreamMessage{
			MessageID: sub.ref,
			Command:   CmdRegisterLayerTokens,
		}
		st := StreamLayerTokens{
			Tokens: []string{token.Token},
		}
		msg.Payload, _ = json.Marshal(st)
		sc.sendMux.Lock()
		_ = sc.conn.WriteJSON(msg)
		sc.sendMux.Unlock()
	}()
	return nil
}

func (sc *SystemStreamClient) RegisterLayerProfile(profile LayerProfile) error {
	if sc.conn == nil {
		return errors.New("not connected")
	}
	sub := sc.subscribe()
	go func() {
		msg := StreamMessage{
			MessageID: sub.ref,
			Command:   CmdRegisterLayerTokens,
		}
		st := StreamLayerTokens{
			Tokens: []string{},
		}
		for _, ag := range profile {
			for _, tk := range ag.Tokens {
				st.Tokens = append(st.Tokens, tk.Token)
			}
		}
		msg.Payload, _ = json.Marshal(st)
		sc.sendMux.Lock()
		_ = sc.conn.WriteJSON(msg)
		sc.sendMux.Unlock()
	}()
	return nil
}

func (sc *SystemStreamClient) LayerTokens() (chan LayerTokens, error) {
	if sc.conn == nil {
		return nil, errors.New("not connected")
	}
	sub := sc.subscribe()
	stChan := make(chan LayerTokens)
	sub.Subscribe(func(event StreamMessage) {
		stMsg := LayerTokens{}
		_ = json.Unmarshal(event.Payload, &stMsg)
		stChan <- stMsg
		close(stChan)
	})
	go func() {
		msg := StreamMessage{
			MessageID: sub.ref,
			Command:   CmdGetLayerTokens,
		}
		sc.sendMux.Lock()
		_ = sc.conn.WriteJSON(msg)
		sc.sendMux.Unlock()
	}()
	return stChan, nil
}

func (sc *SystemStreamClient) GetResourceInstance(resource *ResourceInstance, options ResourceOptions) (*ResourceStreamInstance, error) {
	if sc.conn == nil {
		return nil, errors.New("not connected")
	}
	sub := sc.subscribe()
	resSub := &ResourceStreamInstance{
		res: resource,
		ref: sub.ref,
		cbs: []ResourceCallback{},
	}
	resSub.Subscribe(func(message StreamResourceEvent) {
		if message.Type == CngRemove {
			_ = resSub.remove()
			return
		}
		if message.Type == CngAdd {
			_, _ = resSub.add()
		}
		_ = resSub.removeValues(message.Removed)
		_ = resSub.addValues(message.Added)
	})
	sub.Subscribe(func(event StreamMessage) {
		res := StreamResourceEvent{}
		_ = json.Unmarshal(event.Payload, &res)
		resSub.next(res)
	})
	go func() {
		params := StreamSubscriptionParams{
			Object:   resource.ObjectID,
			Specials: false,
			Payload:  true,
			Instant:  options.Instant,
			Interval: options.Interval,
		}
		params.Params, _ = json.Marshal(StreamResourceParams{
			Object: ResourceTypeInstance,
			Schema: resource.Schema,
		})
		subMsg := MessageSubscribe{
			Event:  EventResource,
			Params: params,
		}
		msg := StreamMessage{
			MessageID: sub.ref,
			Command:   CmdSubscribe,
		}
		msg.Payload, _ = json.Marshal(subMsg)
		sc.sendMux.Lock()
		_ = sc.conn.WriteJSON(msg)
		sc.sendMux.Unlock()
	}()
	return resSub, nil
}

func (sc *SystemStreamClient) GetResourceList(resource *ResourceSpecials, options ResourceOptions) (*ResourceStreamList, error) {
	if sc.conn == nil {
		return nil, errors.New("not connected")
	}
	sub := sc.subscribe()
	resSub := &ResourceStreamList{
		res:     resource,
		ref:     sub.ref,
		indexes: map[ID]int{},
		values:  []any{},
		cbs:     []ResourceCallback{},
	}
	resSub.Subscribe(func(message StreamResourceEvent) {
		if message.Type == CngRemove {
			_ = resSub.remove(message.ID)
			return
		}
		if message.Type == CngAdd {
			_, _ = resSub.add(message.ID)
		}
		_ = resSub.removeValues(message.ID, message.Removed)
		_ = resSub.addValues(message.ID, message.Added)
	})
	sub.Subscribe(func(event StreamMessage) {
		res := StreamResourceEvent{}
		_ = json.Unmarshal(event.Payload, &res)
		resSub.next(res)
	})
	go func() {
		params := StreamSubscriptionParams{
			Object:   resource.ObjectID,
			Specials: false,
			Payload:  true,
			Instant:  options.Instant,
			Interval: options.Interval,
		}
		params.Params, _ = json.Marshal(StreamResourceParams{
			Object: ResourceTypeSpecials,
			Schema: resource.Schema,
			Filter: resource.Filter,
		})
		subMsg := MessageSubscribe{
			Event:  EventResource,
			Params: params,
		}
		msg := StreamMessage{
			MessageID: sub.ref,
			Command:   CmdSubscribe,
		}
		msg.Payload, _ = json.Marshal(subMsg)
		sc.sendMux.Lock()
		_ = sc.conn.WriteJSON(msg)
		sc.sendMux.Unlock()
	}()
	return resSub, nil
}

func (sc *SystemStreamClient) subscribe() *Subscription {
	msgID := NewID()
	sub := &Subscription{
		ref: msgID,
	}
	sc.subscriptions[msgID] = sub
	return sub
}
