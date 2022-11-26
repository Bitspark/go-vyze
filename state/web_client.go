package state

import (
	"fmt"
	"github.com/gorilla/websocket"
)

func NewWebState(addr string, path string, factory string, token string) (*WebState, error) {
	c, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s/%s?f=%s&t=%s", addr, path, factory, token), nil)
	if err != nil {
		return nil, err
	}

	wst := &WebState{
		server: false,
		conn:   c,
		st:     make(chan *MemState),
	}

	go wst.Run()

	return wst, nil
}
