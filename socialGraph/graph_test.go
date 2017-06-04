package socialGraph

import (
	"github.com/JKolios/instagraph/instagram"
	"testing"
)

func TestAddNode(t *testing.T) {
	users := []instagram.User{
		{InstagramId: 100, UserName: "john", FullName: "John Doe"},
		{InstagramId: 101, UserName: "jane", FullName: "Jane Doe"},
		{InstagramId: 102, UserName: "jamie", FullName: "Jamie Doe"},
		{InstagramId: 100, UserName: "wontadd", FullName: "Won't add"},
	}

	instagramGraph := NewInstagramUserGraph()

	nodes := []InstagramUserNode{}

	for _, user := range users {
		nodes = append(nodes, InstagramUserNode{user})
	}

	for _, node := range nodes {
		instagramGraph.AddNode(node)
	}

	t.Log(instagramGraph.Nodes())

	if len(instagramGraph.Nodes()) != 3 {
		t.Error("Wrong number of nodes added")
	}

	for i, node := range nodes[:3] {
		if !instagramGraph.Has(node) {
			t.Errorf("Node %v was not added", i)
		}
	}

}
