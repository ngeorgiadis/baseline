package graph

// EdgeType ...
type EdgeType int

const (

	// Undirected ...
	Undirected EdgeType = iota

	// Directed ...
	Directed
)

// Edge ...
type Edge struct {
	From  string
	To    string
	EType EdgeType
}

// NewEdge ...
func NewEdge(from string, to string, t EdgeType) *Edge {
	return &Edge{
		From:  from,
		To:    to,
		EType: t,
	}
}
