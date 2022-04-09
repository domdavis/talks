package main

import "C"
import (
	"fmt"
	"strings"
)

type Node map[string]interface{}

type Link struct {
	From *Node
	To   *Node
}

type Graph struct {
	Nodes []*Node
	Links []*Link
}

func NewNode(data map[string]interface{}) *Node {
	n := Node(data)

	return &n
}

func (n *Node) String() string {
	return fmt.Sprintf("(%s)", map[string]interface{}(*n))
}

func (l *Link) String() string {
	return fmt.Sprintf("%s-->%s", l.From, l.To)
}

func (n *Graph) String() string {
	s := make([]string, len(n.Links))

	for i, l := range n.Links {
		s[i] = l.String()
	}

	return strings.Join(s, ", ")
}

func graph() {
	a := NewNode(map[string]interface{}{"ID": "A", "Name": "foo"})
	b := NewNode(map[string]interface{}{"ID": "B", "Name": "bar"})

	g := &Graph{Nodes: []*Node{a, b}}

	fmt.Println(g)

}
