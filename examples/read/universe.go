package main

import (
	"github.com/Bitspark/go-vyze"
)

// Universe interfaces

type Document struct {
	ID      go_vyze.ID `json:"id"`
	Name    string     `json:"name"`
	Content string     `json:"content"`
}

// Universe endpoints

// Example is a client for the universe example providing its endpoints as methods
type Example struct{ Client go_vyze.Client }

func (u *Example) getDocument(id go_vyze.ID) (Document, error) {
	return go_vyze.GetNode[Document](u.Client, "getDocument", id)
}
