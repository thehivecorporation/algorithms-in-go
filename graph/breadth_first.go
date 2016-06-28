package graph

import "../common"

var result common.Node
var cola = common.NewQueue(1)

func breadthFirstSearch(n common.Node, id int) common.Node{
	cola.Push(n)
	for (cola.Size()>0) {
		nod := cola.Pop()
		if(nod.Id != id) {
			for i := 0; i < len(nod.Children); i++ {
				cola.Push(nod.Children[i])
			}
		} else {
			result = nod
			break
		}
	}
	return result
}


////BFS uses Queue data structure
//Queue q=new LinkedList();
//q.add(this.rootNode);
//printNode(this.rootNode);
//rootNode.visited=true;
//while(!q.isEmpty())
//{
//Node n=(Node)q.remove();
//Node child=null;
//while((child=getUnvisitedChildNode(n))!=null)
//{
//child.visited=true;
//printNode(child);
//q.add(child);
//}
//}
////Clear visited property of nodes
//clearNodes();
//}