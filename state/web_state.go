package state

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type WebState struct {
	server  bool
	conn    *websocket.Conn
	name    string
	st      chan *MemState
	stType  *Type
	err     error
	connMux sync.Mutex
}

var _ State = &WebState{}

func (wst *WebState) String() string {
	return wst.Mem().String()
}

func (wst *WebState) Run() {
	defer wst.conn.Close()

	wst.init()
	if wst.err != nil {
		return
	}
	wst.Mem()

	for {
		msg := wst.nextMessage()
		if wst.err != nil {
			return
		}
		if err := wst.handleMessage(msg); err != nil {
			wst.err = err
			log.Println(wst.err)
			return
		}
	}
}

func (wst *WebState) FullName() string {
	return wst.Mem().FullName()
}

func (wst *WebState) Interface() NativeInterface {
	return &nativeInterface{
		state: wst,
	}
}

func (wst *WebState) SetValue(name string, value any) error {
	msg := setValueWebMsg{
		Name:  name,
		Value: value,
	}
	msgBts, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if err := wst.sendMessage("set", msgBts); err != nil {
		return err
	}
	return wst.Mem().SetValue(name, value)
}

func (wst *WebState) GetValue(name string) (*ValueHolder, error) {
	return wst.Mem().GetValue(name)
}

func (wst *WebState) RemoveChild(name string) error {
	return wst.Mem().RemoveChild(name)
}

func (wst *WebState) Mem() *MemState {
	st := <-wst.st
	go func() {
		wst.st <- st
	}()
	return st
}

func (wst *WebState) Type() *Type {
	return wst.stType
}

func (wst *WebState) Sync(name string) error {
	msg := getValueWebMsg{
		Name: name,
	}
	msgBts, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	if err := wst.sendMessage("get", msgBts); err != nil {
		return err
	}
	return nil
}

func (wst *WebState) handleMessage(msg []byte) error {
	wsMsg := webMessage{}
	if err := json.Unmarshal(msg, &wsMsg); err != nil {
		return err
	}
	switch wsMsg.Cmd {
	case "get":
		return wst.handleGetMessage(wsMsg.Payload)
	case "set":
		return wst.handleSetMessage(wsMsg.Payload)
	default:
		log.Println(fmt.Errorf("unknown command: %s", wsMsg.Cmd))
		return nil
	}
}

func (wst *WebState) handleGetMessage(payload json.RawMessage) error {
	msgGet := getValueWebMsg{}
	if err := json.Unmarshal(payload, &msgGet); err != nil {
		return err
	}
	// Get value
	val, err := wst.GetValue(msgGet.Name)
	if err != nil {
		return err
	}
	// Send the value
	msgSet := setValueWebMsg{
		Name:  msgGet.Name,
		Value: val.Value,
	}
	msgBts, err := json.Marshal(msgSet)
	if err != nil {
		return err
	}
	if err := wst.sendMessage("set", msgBts); err != nil {
		return err
	}
	return nil
}

func (wst *WebState) handleSetMessage(payload json.RawMessage) error {
	msg := setValueWebMsg{}
	if err := json.Unmarshal(payload, &msg); err != nil {
		return err
	}
	err := wst.Mem().SetValue(msg.Name, msg.Value)
	return err
}

func (wst *WebState) sendMessage(cmd string, msg json.RawMessage) error {
	if cmd == "" {
		wst.connMux.Lock()
		err := wst.conn.WriteMessage(websocket.TextMessage, msg)
		wst.connMux.Unlock()
		return err
	}
	msgBts, err := json.Marshal(webMessage{
		Cmd:     cmd,
		Payload: msg,
	})
	if err != nil {
		return err
	}
	wst.connMux.Lock()
	err = wst.conn.WriteMessage(websocket.TextMessage, msgBts)
	wst.connMux.Unlock()
	return err
}

type stateInit struct {
	Name  string       `json:"name"`
	Value *ValueHolder `json:"values"`
}

func (wst *WebState) init() {
	if wst.server {
		st := wst.Mem()
		stInit := stateInit{
			Name:  st.Name,
			Value: st.Values,
		}
		msg, err := json.Marshal(stInit)
		if err != nil {
			wst.err = err
			return
		}
		if err := wst.sendMessage("", msg); err != nil {
			wst.err = err
			return
		}
	} else {
		time.Sleep(200 * time.Millisecond)
		msg := wst.nextMessage()
		if wst.err != nil {
			return
		}
		stInit := stateInit{}
		if err := json.Unmarshal(msg, &stInit); err != nil {
			wst.err = err
			return
		}
		wst.name = stInit.Name
		st, err := NewMemState(stInit.Name, stInit.Value)
		if err != nil {
			wst.err = err
			return
		}
		st.self = wst
		wst.st <- st
		wst.stType = st.Type()
	}
}

func (wst *WebState) nextMessage() []byte {
	_, msg, err := wst.conn.ReadMessage()
	if err != nil {
		wst.err = err
		return nil
	}
	return msg
}

// Structs

type webMessage struct {
	Cmd     string          `json:"cmd"`
	Payload json.RawMessage `json:"payload"`
}

type setValueWebMsg struct {
	Name  string `json:"name,omitempty"`
	Value any    `json:"value"`
}

type getValueWebMsg struct {
	Name string `json:"name,omitempty"`
}
