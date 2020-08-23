package gee

type Node struct {
	pattern string // 待匹配路由，如 /p/:lang
	part string // 路由中的一部分 如 :lange
	children []*Node // 子节点
	isWild bool // 是否精准匹配 yes:模糊匹配 /:lang
}

func (n *Node) matchChild(part string) *Node {
	for _, child := range n.children {
		if child.part == part {
			return child
		}
	}
	return nil
}

func (n *Node) matchChildren(part string) []*Node {
	nodes := make([]*Node, 0)
	for _, child := range n.children {
		if child.part == part {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *Node) insert(pattern string, parts []string, height int) {
	if height == len(parts) {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &Node{part: part, isWild: part[0] == ':' || part[1] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

