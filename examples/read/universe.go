package main

import "github.com/Bitspark/go-vyze/pkg/vyze"

// Universe interfaces

type Document struct {
	ID      vyze.ID `json:"id"`
	Name    string  `json:"name"`
	Content string  `json:"content"`
}

// Universe endpoints

// Example is a client for the universe example providing its endpoints as methods
type Example struct{ Client vyze.Client }

func (u *Example) getDocument(id vyze.ID) (Document, error) {
	return vyze.GetNode[Document](u.Client, "getDocument", id)
}
