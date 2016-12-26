package graph

import "testing"

func TestDepthFirstSearch(t *testing.T) {
	var nod5 Node = Node{"5", false, nil}
	var nod4 Node = Node{"4", false, nil}
	var nod3 Node = Node{"3", false, nil}
	var nod2 Node = Node{"2", false, []*Node{&nod4, &nod5}}
	var nod1 Node = Node{"1", false, []*Node{&nod2, &nod3}}

	if depthFirstSearch(nod1, "2").Id != "2" {
		t.Fail()
	}
}
