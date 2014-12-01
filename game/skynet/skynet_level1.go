package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// N: the total number of nodes in the level, including the gateways
	// L: the number of links
	// E: the number of exit gateways
	var N, L, E int
	fmt.Scan(&N, &L, &E)

	g := NewGraph()

	for i := 0; i < L; i++ {
		// N1: N1 and N2 defines afm link between these nodes
		var N1, N2 int
		fmt.Scan(&N1, &N2)
		g.addLinks(N1, N2)
	}
	for i := 0; i < E; i++ {
		// EI: the index of a gateway node
		var EI int
		fmt.Scan(&EI)
		g.addGateWay(EI)
	}
	for {
		// SI: The index of the node on which the Skynet agent is positioned this turn
		var SI int
		fmt.Scan(&SI)

		goalList := make([]*Goal, 0)
		for _, gw := range g.gateWays {
			goal, err := bfs(gw.id, SI, g);
			if err != nil {
			} else {
				goalList = append(goalList, goal)
			}
		}

		goal := goalList[0]
		for _, nextGoal := range goalList {
			if goal.length > nextGoal.length {
				goal = nextGoal
			}
		}

		// Reverse because we come from the gateway to find the skynet
		execute(goal.to, goal.from)
	}
}

func debug(msg interface {}) {
	fmt.Fprintln(os.Stderr, msg)
}

func execute(from, to int) {
	fmt.Println(fmt.Sprintf("%d %d", from, to))
}
// -------------------

func bfs(start, end int, g *Graph) (*Goal, error) {
	current := NewGoal(start, start, 0)
	edge := make([]Goal, 0)
	edge = append(edge, *current)

	goalList := make([]Goal, 0)
	visited := make(map[int]bool)
	visited[start] = true
	for len(edge) > 0{
		toCheck := make([]Goal, 0)
		for _, goal := range edge {
			n, _ := g.getNode(goal.to)
			for _, linkedNode := range n.linkedNodes {
				_, isVisited := visited[linkedNode.id]
				if !isVisited {
					goal.next = NewGoal(goal.to, linkedNode.id, goal.length + 1)
					goal.next.previous = &goal;

					toCheck = append(toCheck, *goal.next)
					visited[linkedNode.id] = true

					if linkedNode.id == end {
						goalList = append(goalList, *goal.next)
					}
				}
			}
		}
		edge = make([]Goal, 0)
		if len(toCheck) > 0 {
			for _, goalToCheck := range toCheck {
				edge = append(edge, goalToCheck)
			}
		}
	}

	if len(goalList) > 0 {
		result := goalList[0]
		return &result, nil
	}

	return nil, NoGoalFound(start)
}

// -------------------
type Goal struct {
	from, to, length int
	next, previous *Goal
}
func NewGoal(from, to, length int) *Goal {
	goal := new(Goal)
	goal.from = from
	goal.to = to
	goal.length = length
	return goal
}
// -------------------
type NoGoalFound int

func (e NoGoalFound) Error() string {
	return "No goal found for node id " + fmt.Sprint("%v", e)
}

// -------------------
type Node struct {
	id int
	linkedNodes []*Node
	isGateWay bool
}

func NewNode(id int) *Node {
	node := new(Node)
	node.id = id
	node.linkedNodes = make([]*Node, 0)
	return node
}

func (n *Node) addNode(linkedNode *Node) {
	n.linkedNodes = append(n.linkedNodes, linkedNode)
}

func (n *Node) String() string {
	s := strconv.Itoa(n.id)
	if n.isGateWay {
		s += "*"
	}
	s += " - "
	for _, linkedNode := range n.linkedNodes {
		s += strconv.Itoa(linkedNode.id)
		if linkedNode.isGateWay {
			s += "*"
		}
		s += " "
	}
	s += "\n"
	return s
}
// -------------------
type Graph struct {
	nodes map[int]*Node
	gateWays []*Node
}

func NewGraph() *Graph {
	g := new(Graph)
	g.nodes = make(map[int]*Node)
	g.gateWays = make([]*Node, 0)
	return g
}

func (g *Graph) getNode(id int) (*Node, bool) {
	n, isPresent := g.nodes[id]
	return n, isPresent
}

func (g *Graph) addLinks(id1, id2 int) {
	n1, isN1Present := g.getNode(id1)
	if !isN1Present {
		n1 = NewNode(id1)
		g.nodes[id1] = n1
	}
	n2, isN2Present := g.getNode(id2)
	if !isN2Present {
		n2 = NewNode(id2)
		g.nodes[id2] = n2
	}
	n1.addNode(n2)
	n2.addNode(n1)
}

func (g *Graph) String() string {
	s := ""
	for _, node := range g.nodes {
		s += fmt.Sprintf("%v", node)
	}
	return s
}

func (g *Graph) addGateWay(EI int) {
	n, _ := g.getNode(EI)
	n.isGateWay = true
	g.gateWays = append(g.gateWays, n)
}
