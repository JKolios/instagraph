package socialgraph

import (
	"testing"

	"github.com/JKolios/instagraph/instagram"
	"github.com/gonum/graph"
	"github.com/gonum/graph/simple"
)

var testGraph *InstagramUserGraph
var nodes []InstagramUserNode

func TestInstagramUserGraph(t *testing.T) {
	users := []instagram.User{
		{InstagramId: 100, UserName: "john", FullName: "John Doe"},
		{InstagramId: 101, UserName: "jane", FullName: "Jane Doe"},
		{InstagramId: 102, UserName: "jamie", FullName: "Jamie Doe"},
		{InstagramId: 100, UserName: "wontadd", FullName: "Won't add"},
	}

	testGraph = NewInstagramUserGraph()

	var _ graph.Directed = (*InstagramUserGraph)(nil)

	for _, user := range users {
		nodes = append(nodes, InstagramUserNode{user})
	}

	t.Run("Nodes", _TestNodes)
	t.Run("Edges", _TestEdges)
	t.Run("Export", _TestExport)

}

func _TestNodes(t *testing.T) {

	for _, node := range nodes {
		testGraph.AddNode(node)
	}

	if len(testGraph.Nodes()) != 3 {
		t.Error("Wrong number of nodes added")
	}

	for i, node := range nodes[:3] {
		if !testGraph.Has(node) {
			t.Errorf("Node %v was not added", i)
		}
	}

}

func _TestEdges(t *testing.T) {

	fromNode := testGraph.Nodes()[0]
	toNode := testGraph.Nodes()[1]

	testEdge := simple.Edge{F: fromNode, T: toNode, W: 1.0}

	testGraph.SetEdge(testEdge)

	if testGraph.Edge(fromNode, toNode) != testEdge {
		t.Error("Edge is not returned")
	}

	if !testGraph.HasEdgeBetween(fromNode, toNode) {
		t.Error("Edge between nodes cannot be detected")
	}

	if !testGraph.HasEdgeFromTo(fromNode, toNode) {
		t.Error("Edge from node to node cannot be detected")
	}

	testInverseEdge := simple.Edge{F: toNode, T: fromNode, W: 1.0}

	testGraph.SetEdge(testInverseEdge)

	if !testGraph.HasEdgeBetween(fromNode, toNode) {
		t.Error("Inverse Edge between nodes cannot be detected")
	}

	if !testGraph.HasEdgeFromTo(toNode, fromNode) {
		t.Error("Inverse Edge from node to node cannot be detected")
	}

}

func _TestExport(t *testing.T) {

	_, err := testGraph.ExportDOT()

	if err != nil {
		t.Error("DOT export failed")
	}

	if err = testGraph.ExportDOTToFile("test_export.gv"); err != nil {
		t.Error("Failed to write DOT export to file")
	}

}
