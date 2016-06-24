package common

type node struct {
	id int
	visited bool
	children []*node
}