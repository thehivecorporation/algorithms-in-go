package graph

var result node

func depthFirstSearch(n node, id int) node{
	if(id!=n.id) {
		for i := 0; i < len(n.children); i++ {
			if (n.children[i].id != id) {
				if (n.children[i].visited == false) {
					n.children[i].visited = true
					depthFirstSearch(*n.children[i], id)
				}
			} else {
				result = *n.children[i]
				break
			}
		}
	} else {
		result = n
	}
	return result
}
