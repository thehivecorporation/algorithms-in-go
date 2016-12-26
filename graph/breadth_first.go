package graph

func BreadthFirstSearch(n Node, id string) Node {
	var result Node
	queue := NewQueue(1)

	queue.Push(&n)

	for queue.Size() > 0 {
		nod := queue.Pop()
		if nod.Id != id {
			for i := 0; i < len(nod.Children); i++ {
				queue.Push(nod.Children[i])
			}
		} else {
			result = *nod
			break
		}
	}
	return result
}
