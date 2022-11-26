package state

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type WebsocketServer struct {
	addr      string
	path      string
	factories map[string]*Factory
	clients   []*WebState
}

func NewWebsocketServer(addr string, path string) (*WebsocketServer, error) {
	ws := &WebsocketServer{
		addr:      addr,
		path:      path,
		factories: map[string]*Factory{},
	}
	return ws, nil
}

func (ws *WebsocketServer) Start() error {
	http.HandleFunc("/"+ws.path, ws.connect)
	return http.ListenAndServe(ws.addr, nil)
}

func (ws *WebsocketServer) AddFactory(st *MemState, factory string, token string, binding *TwoBinding, value *ValueHolder) (*Factory, error) {
	f := &Factory{
		Name:    factory,
		Token:   token,
		st:      st,
		binding: binding,
		value:   value,
	}
	ws.factories[factory] = f
	return f, nil
}

func (ws *WebsocketServer) Shutdown() error {
	return nil
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ws *WebsocketServer) connect(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	f, err := ws.getFactory(r.URL.Query().Get("f"), r.URL.Query().Get("t"))
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	chSt, err := f.Produce()
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	wst := &WebState{
		server: true,
		conn:   c,
		st:     make(chan *MemState),
	}

	f.st.muxChildren.Lock()
	chBS := f.st.Children[chSt.Name]
	f.st.Children[chSt.Name] = &BoundState{
		Binding: chBS.Binding,
		State:   wst,
	}
	f.st.muxChildren.Unlock()

	go func(st *MemState) {
		wst.st <- st
		wst.stType = st.Type()
	}(chSt)

	ws.clients = append(ws.clients, wst)

	wst.Run()

	if wst.err != nil {
		_, _ = w.Write([]byte(wst.err.Error()))
	}
}

func (ws *WebsocketServer) getFactory(name string, token string) (*Factory, error) {
	f, ok := ws.factories[name]
	if !ok {
		return nil, fmt.Errorf("factory not found: %s", name)
	}
	if f.Token != token {
		// Do not return token here (we don't know where the error goes)
		return nil, fmt.Errorf("wrong token: %s", name)
	}
	return f, nil
}
