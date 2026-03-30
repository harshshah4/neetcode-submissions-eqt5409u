/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visit := make(map[*Node]*Node)
	queue := []*Node{node}
	curr := &Node{
		Val: node.Val,
		Neighbors: []*Node{},
	}
	visit[node] = curr
    for len(queue) > 0 {
		for _, each := range queue {
			// fmt.Printf("visit[each].Val %d, each.Val %d\n", visit[each].Val, each.Val)
			if len(queue) > 1 {
				queue = queue[1:]
			} else {
				queue = []*Node{}
			}
			for _, neighbor := range each.Neighbors {
				newNode, exists := visit[neighbor]
				// fmt.Printf("neighbor.Val %d exists %t\n", neighbor.Val, exists)
				if !exists {
					queue = append(queue, neighbor)
					newNode = &Node{
						Val: neighbor.Val,
						Neighbors: []*Node{},
					}
					visit[neighbor] = newNode
				}
				visit[each].Neighbors = append(visit[each].Neighbors, newNode)
			}
			// fmt.Printf("neighbors visit[each].Neighbors %v each.Neighbors %v\n", visit[each].Neighbors, each.Neighbors)
		}
	}
	return curr
}
