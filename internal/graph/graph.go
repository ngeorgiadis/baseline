package graph

// Graph ...
type Graph []*Node

// IxGraph ...
type IxGraph map[string]*Node

// XGraph ...
type XGraph map[string]*XNode

// New ...
func New() Graph {
	return Graph{}
}

// NewIx ...
func NewIx() IxGraph {
	return IxGraph{}
}

// NewXGraph ...
func NewXGraph() XGraph {
	return XGraph{}
}

// AddNode ...
func (g Graph) AddNode(n *Node) {
	g = append(g, n)
}

// AddNode ...
func (g IxGraph) AddNode(n *Node) {
	g[n.ID] = n
}

// AddNodes ...
func (g IxGraph) AddNodes(nodes []*Node) {
	for _, n := range nodes {
		g[n.ID] = n
	}
}

// AddNode ...
func (g XGraph) AddNode(n *XNode) {
	g[n.ID] = n
}

// AddNodes ...
func (g XGraph) AddNodes(nodes []*XNode) {
	for _, n := range nodes {
		g[n.ID] = n
	}
}
