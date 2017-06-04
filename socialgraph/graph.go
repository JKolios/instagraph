package socialgraph

import (
	"github.com/JKolios/instagraph/instagram"
	"github.com/gonum/graph"
)

type InstagramUserNode struct {
	instagram.User
}

func (n InstagramUserNode) ID() int {
	return n.InstagramId
}

type InstagramUserGraph struct {
	nodes map[int]graph.Node
	from  map[int]map[int]graph.Edge
	to    map[int]map[int]graph.Edge

	self, absent float64
}

func NewInstagramUserGraph() *InstagramUserGraph {
	return &InstagramUserGraph{
		nodes: map[int]graph.Node{},
		from:  map[int]map[int]graph.Edge{},
		to:    map[int]map[int]graph.Edge{},
		self: 1.0,
		absent: 0.0,
	}
}

func (g *InstagramUserGraph) AddNode(n InstagramUserNode) {

	if _, exists := g.nodes[n.ID()]; exists {
		return
	}

	g.nodes[n.ID()] = n
	g.from[n.ID()] = make(map[int]graph.Edge)
	g.to[n.ID()] = make(map[int]graph.Edge)
}

func (g *InstagramUserGraph) Has(n graph.Node) bool {

	_, ok := g.nodes[n.ID()]

	return ok
}

func (g *InstagramUserGraph) Nodes() []graph.Node {

	nodes := make([]graph.Node, len(g.from))
	i := 0
	for _, n := range g.nodes {
		nodes[i] = n
		i++
	}

	return nodes
}
