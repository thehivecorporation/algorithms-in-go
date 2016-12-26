package graph

import "github.com/GaizkaRubio/algorithms-in-go/common"

var result common.Node
var cola = common.NewQueue(1)

func breadthFirstSearch(n common.Node, id int) common.Node{
	cola.Push(&n)
	for (cola.Size()>0) {
		nod := cola.Pop()
		if(nod.Id != id) {
			for i := 0; i < len(nod.Children); i++ {
				cola.Push(nod.Children[i])
			}
		} else {
			result = *nod
			break
		}
	}
	return result
}