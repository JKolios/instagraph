package socialgraph

import (
	"github.com/JKolios/instagraph/instagram"
	"testing"
)

var graph *InstagramUserGraph
var nodes []InstagramUserNode

func TestInstagramUserGraph(t *testing.T) {
	users := []instagram.User{
		{InstagramId: 100, UserName: "john", FullName: "John Doe"},
		{InstagramId: 101, UserName: "jane", FullName: "Jane Doe"},
		{InstagramId: 102, UserName: "jamie", FullName: "Jamie Doe"},
		{InstagramId: 100, UserName: "wontadd", FullName: "Won't add"},
	}

	graph = NewInstagramUserGraph()

	nodes := []InstagramUserNode{}

	for _, user := range users {
		nodes = append(nodes, InstagramUserNode{user})
	}

	for _, node := range nodes {
		graph.AddNode(node)
	}

	t.Run(_TestAddNode)

}

func _TestAddNode(t *testing.T) {


	if len(graph.Nodes()) != 3 {
		t.Error("Wrong number of nodes added")
	}

	for i, node := range nodes[:3] {
		if !graph.Has(node) {
			t.Errorf("Node %v was not added", i)
		}
	}

}
