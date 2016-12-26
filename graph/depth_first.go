package graph

var Result Node

func depthFirstSearch(n Node, id string) Node{
	if(id!=n.Id) {
		for i := 0; i < len(n.Children); i++ {
			if (n.Children[i].Id != id) {
				if (n.Children[i].Visited == false) {
					n.Children[i].Visited = true
					depthFirstSearch(*n.Children[i], id)
				}
			} else {
				Result = *n.Children[i]
				break
			}
		}
	} else {
		Result = n
	}
	return Result
}
