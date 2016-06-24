package graph

var Result node

func depthFirstSearch(n node, id int) node{
	if(id!=n.id) {
		for i := 0; i < len(n.children); i++ {
			if (n.children[i].id != id) {
				if (n.children[i].visited == false) {
					n.children[i].visited = true
					depthFirstSearch(*n.children[i], id)
				}
			} else {
				Result = *n.children[i]
				break
			}
		}
	} else {
		Result = n
	}
	return Result
}
