package state

import "github.com/Bitspark/go-vyze/core"

type Factory struct {
	Name  string `json:"name"`
	Token string `json:"token"`

	st      *MemState
	binding *TwoBinding
	value   *ValueHolder
}

func (f *Factory) Produce() (*MemState, error) {
	chSt, err := f.st.NewChild(f.Name+"-"+core.NewID().Hex()[:8], f.value.Copy(), f.binding)
	if err != nil {
		return nil, err
	}
	return chSt, nil
}
