package graph

import (
	"github.com/GaizkaRubio/algorithms-in-go/common"
	"testing"
)


func TestBreadthFirstSearch(t *testing.T) {
	var nod5 node = node{5, false, nil}
	var nod4 node = node{4, false, nil}
	var nod3 node = node{3, false, nil}
	var nod2 node = node{2, false, []*node{&nod4, &nod5}}
	var nod1 node = node{1, false, []*node{&nod2, &nod3}}

	println(breadthFirstSearch(nod1, 2).id)
}
