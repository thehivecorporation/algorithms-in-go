package graph

type Node struct {
	Id       string
	Visited  bool
	Children []*Node
}
