package socialgraph

import (
	"fmt"
	"github.com/JKolios/instagraph/instagram"
	"github.com/gonum/graph"
	"github.com/gonum/graph/encoding/dot"
	"io/ioutil"
	"strings"
)

type InstagramUserNode struct {
	instagram.User
}

func (n InstagramUserNode) ID() int {
	return n.InstagramId
}

func (n InstagramUserNode) DOTID() string {
	return fmt.Sprintf("%s_%d", strings.Replace(n.FullName, " ", "_", -1) , n.InstagramId)
}

type InstagramUserGraph struct {
	nodes map[int]graph.Node
	from  map[int]map[int]graph.Edge
	to    map[int]map[int]graph.Edge

	self, absent float64
}

func NewInstagramUserGraph() *InstagramUserGraph {
	return &InstagramUserGraph{
		nodes:  map[int]graph.Node{},
		from:   map[int]map[int]graph.Edge{},
		to:     map[int]map[int]graph.Edge{},
		self:   1.0,
		absent: 0.0,
	}
}

func (g *InstagramUserGraph) AddNode(n graph.Node) {

	if _, exists := g.nodes[n.ID()]; exists {
		return
	}

	g.nodes[n.ID()] = n
	g.from[n.ID()] = make(map[int]graph.Edge)
	g.to[n.ID()] = make(map[int]graph.Edge)
}

func (g InstagramUserGraph) Has(n graph.Node) bool {

	_, ok := g.nodes[n.ID()]

	return ok
}

func (g InstagramUserGraph) Nodes() []graph.Node {

	nodes := make([]graph.Node, len(g.from))
	i := 0
	for _, n := range g.nodes {
		nodes[i] = n
		i++
	}

	return nodes
}

func (g *InstagramUserGraph) SetEdge(e graph.Edge) {
	var (
		from = e.From()
		fid  = from.ID()
		to   = e.To()
		tid  = to.ID()
	)

	if fid == tid {
		panic("simple: adding self edge")
	}

	if !g.Has(from) {
		g.AddNode(from)
	}
	if !g.Has(to) {
		g.AddNode(to)
	}

	g.from[fid][tid] = e
	g.to[tid][fid] = e
}

func (g InstagramUserGraph) Edge(u, v graph.Node) graph.Edge {
	if _, ok := g.nodes[u.ID()]; !ok {
		return nil
	}
	if _, ok := g.nodes[v.ID()]; !ok {
		return nil
	}
	edge, ok := g.from[u.ID()][v.ID()]
	if !ok {
		return nil
	}
	return edge
}

func (g *InstagramUserGraph) Edges() []graph.Edge {

	var edges []graph.Edge
	for _, u := range g.nodes {
		for _, e := range g.from[u.ID()] {
			edges = append(edges, e)
		}
	}
	return edges

}

func (g InstagramUserGraph) HasEdgeBetween(x, y graph.Node) bool {
	xid := x.ID()
	yid := y.ID()
	if _, ok := g.nodes[xid]; !ok {
		return false
	}
	if _, ok := g.nodes[yid]; !ok {
		return false
	}
	if _, ok := g.from[xid][yid]; ok {
		return true
	}
	_, ok := g.from[yid][xid]
	return ok
}

// HasEdgeFromTo returns whether an edge exists in the graph from u to v.
func (g InstagramUserGraph) HasEdgeFromTo(u, v graph.Node) bool {
	if _, ok := g.nodes[u.ID()]; !ok {
		return false
	}
	if _, ok := g.nodes[v.ID()]; !ok {
		return false
	}
	if _, ok := g.from[u.ID()][v.ID()]; !ok {
		return false
	}
	return true
}

// From returns all nodes in g that can be reached directly from n.
func (g InstagramUserGraph) From(n graph.Node) []graph.Node {
	if _, ok := g.from[n.ID()]; !ok {
		return nil
	}

	from := make([]graph.Node, len(g.from[n.ID()]))
	i := 0
	for id := range g.from[n.ID()] {
		from[i] = g.nodes[id]
		i++
	}

	return from
}

func (g InstagramUserGraph) To(n graph.Node) []graph.Node {

	if _, ok := g.from[n.ID()]; !ok {
		return nil
	}

	to := make([]graph.Node, len(g.to[n.ID()]))
	i := 0
	for id := range g.to[n.ID()] {
		to[i] = g.nodes[id]
		i++
	}

	return to
}

func (g *InstagramUserGraph) ExportDOT() ([]byte, error) {

	exportBytes, err := dot.Marshal(graph.Directed(*g), "", "", "", false)
	return exportBytes, err

}

func (g *InstagramUserGraph) ExportDOTToFile(filename string) error {

	exportBytes, err := g.ExportDOT()

	if err != nil {
		return err
	}

	if ioutil.WriteFile(filename, exportBytes, 0755); err != nil {
		return err
	}

	return nil

}
