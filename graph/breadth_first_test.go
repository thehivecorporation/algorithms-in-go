package graph

import (
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	nod5 := Node{Id: "5"}
	nod4 := Node{Id: "4"}
	nod3 := Node{Id: "3"}
	nod2 := Node{Id: "2", Children: []*Node{&nod4, &nod5}}
	nod1 := Node{Id: "1", Children: []*Node{&nod2, &nod3}}

	if BreadthFirstSearch(nod1, "1").Id != "1" {
		t.Fail()
	}
}
