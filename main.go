package main

import (
	"os"

	"github.com/SeanDunford/simpleGraphGo/simplegraph"
)

const (
	apple    = `{"name":"Apple Computer Company","type":["company","start-up"],"founded":"April 1, 1976","id":"1"}`
	woz      = `{"id":"2","name":"Steve Wozniak","type":["person","engineer","founder"]}`
	wozNick  = `{"name":"Steve Wozniak","type":["person","engineer","founder"],"nickname":"Woz","id":"2"}`
	jobs     = `{"id":"3","name":"Steve Jobs","type":["person","designer","founder"]}`
	wayne    = `{"id":"4","name":"Ronald Wayne","type":["person","administrator","founder"]}`
	markkula = `{"name":"Mike Markkula","type":["person","investor"]}`
	founded  = `{"action":"founded"}`
	invested = `{"action":"invested","equity":80000,"debt":170000}`
	divested = `{"action":"divested","amount":800,"date":"April 12, 1976"}`
)

func main() {
	println("Hello, World!")

	file := "importTest.sqlite3"
	simplegraph.Initialize(file)
	defer os.Remove(file)

	_, _ = simplegraph.AddNode("1", []byte(apple), file)
	_, _ = simplegraph.AddNode("2", []byte(woz), file)
	_, _ = simplegraph.AddNode("3", []byte(jobs), file)
	_, _ = simplegraph.AddNode("4", []byte(wayne), file)
	_, _ = simplegraph.AddNode("5", []byte(markkula), file)
	_, _ = simplegraph.AddNode("1", []byte(apple), file)
	_, _ = simplegraph.AddNode("2", []byte(woz), file)
	_, _ = simplegraph.ConnectNodesWithProperties("2", "1", []byte(founded), file)
	_, _ = simplegraph.ConnectNodesWithProperties("3", "1", []byte(founded), file)
	_, _ = simplegraph.ConnectNodesWithProperties("4", "1", []byte(founded), file)
	_, _ = simplegraph.ConnectNodesWithProperties("5", "1", []byte(invested), file)
	_, _ = simplegraph.ConnectNodesWithProperties("1", "4", []byte(divested), file)
	_, _ = simplegraph.ConnectNodes("2", "3", file)
	_, _ = simplegraph.FindNode("1", file)
	_, _ = simplegraph.FindNode("7", file)
	kvNameLike := simplegraph.GenerateWhereClause(&simplegraph.WhereClause{KeyValue: true, Key: "name", Predicate: "LIKE"})
	statement := simplegraph.GenerateSearchStatement(&simplegraph.SearchQuery{ResultColumn: "body", SearchClauses: []string{kvNameLike}})
	_, _ = simplegraph.FindNodes(statement, []string{"Steve%"}, file)
	_ = simplegraph.UpdateNodeBody("2", wozNick, file)
	_ = simplegraph.UpsertNode("1", apple, file)
	_, _ = simplegraph.FindNode("2", file)
	arrayType := simplegraph.GenerateWhereClause(&simplegraph.WhereClause{Tree: true, Predicate: "="})
	statement = simplegraph.GenerateSearchStatement(&simplegraph.SearchQuery{ResultColumn: "body", Tree: true, Key: "type", SearchClauses: []string{arrayType}})
	_, _ = simplegraph.FindNodes(statement, []string{"founder"}, file)
	basicTraversal := simplegraph.GenerateTraversal(&simplegraph.Traversal{WithBodies: false, Inbound: true, Outbound: true})
	_, _ = simplegraph.TraverseFromTo("2", "3", basicTraversal, file)
	basicTraversalInbound := simplegraph.GenerateTraversal(&simplegraph.Traversal{WithBodies: false, Inbound: true, Outbound: false})
	_, _ = simplegraph.TraverseFrom("5", basicTraversalInbound, file)
	basicTraversalOutbound := simplegraph.GenerateTraversal(&simplegraph.Traversal{WithBodies: false, Inbound: false, Outbound: true})
	_, _ = simplegraph.TraverseFrom("5", basicTraversalOutbound, file)
	_, _ = simplegraph.TraverseFrom("5", basicTraversal, file)
	basicTraversalWithBodies := simplegraph.GenerateTraversal(&simplegraph.Traversal{WithBodies: true, Inbound: true, Outbound: true})
	_ = simplegraph.NodeData{Identifier: nil, Body: nil}
	_, _ = simplegraph.TraverseWithBodiesFromTo("2", "3", basicTraversalWithBodies, file)
	_, _ = simplegraph.ConnectionsIn("1", file)
	_ = []simplegraph.EdgeData{{Source: "1", Target: "4", Label: divested}}
	_, _ = simplegraph.ConnectionsOut("1", file)
	_, _ = simplegraph.Connections("1", file)
	_, _ = simplegraph.FindNode("2", file)
	_, _ = simplegraph.FindNode("4", file)

	println("Hello moto")
	println("fin")
}
