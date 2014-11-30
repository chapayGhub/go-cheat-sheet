package main

import (
	"fmt"
	"os"
)

var commands [2]CMD
var WAIT, BLOCK CMD
func initCmd() {
	WAIT = CMD("WAIT")
	BLOCK = CMD("BLOCK")
	commands[0] = WAIT
	commands[1] = BLOCK
}

var LEFT, RIGHT Direction

func initDirections() {
	LEFT = Direction{"LEFT", -1}
	RIGHT = Direction{"RIGHT", 1}
}

var nbFloors, width, nbRounds, exitFloor, exitPos, nbTotalClones, nbAdditionalElevators, nbElevators, nbClones int

func main() {
	initCmd()
	initDirections()

	// nbFloors: number of floors
	// width: width of the area
	// nbRounds: maximum number of rounds
	// exitFloor: floor on which the exit is found
	// exitPos: position of the exit on its floor
	// nbTotalClones: number of generated clones
	// nbAdditionalElevators: ignore (always zero)
	// nbElevators: number of elevators
	fmt.Scan(&nbFloors, &width, &nbRounds, &exitFloor, &exitPos, &nbTotalClones, &nbAdditionalElevators, &nbElevators)
	nbClones = nbTotalClones

	elevatorMap := make(map[int]int)
	for i := 0; i < nbElevators; i++ {
		// elevatorFloor: floor on which this elevator is found
		// elevatorPos: position of the elevator on its floor
		var elevatorFloor, elevatorPos int
		fmt.Scan(&elevatorFloor, &elevatorPos)
		elevatorMap[elevatorFloor] = elevatorPos
	}

	exitPoint := Point{exitPos, exitFloor}

	for {
		// cloneFloor: floor of the leading clone
		// clonePos: position of the leading clone on its floor
		// direction: direction of the leading clone: LEFT or RIGHT
		var cloneFloor, clonePos int
		var dir string
		fmt.Scan(&cloneFloor, &clonePos, &dir)

		direction := getDirection(dir)
		clonePoint := Point{clonePos, cloneFloor}

		if clonePos == -1 && cloneFloor == -1 {
			fmt.Println(WAIT)
		} else {
			fmt.Println(bfs(direction, clonePoint, exitPoint, elevatorMap)) // action: WAIT or BLOCK
		}
	}
}

func bfs(direction Direction, clonePoint, exitPoint Point, elevatorMap map[int]int) CMD {
	current := NewNode(clonePoint, direction)
	edge := make([]Node, 0)
	edge = append(edge, *current)

	goalList := make([]Node, 0)
	visited := make([]Point, 0)
	visited = append(visited, clonePoint)

	for len(edge) > 0 {
		toCheck := make([]Node, 0)
		for i, _ := range edge {
			for _, cmd := range commands {
				node := edge[i]
				if !cmd.canExecute() {
					continue;
				}
				newP, newDir := node.point.newPoint(node.direction, cmd)

				// If its an elevator, then increment the floor
				elevatorX, found := elevatorMap[newP.y]
				if found && elevatorX == newP.x {
					newP.y++
				}

				if newP.canMoveHere(visited) {
					node.next = NewNode(newP, newDir)
					node.next.length = node.length + 1
					node.next.cmd = cmd
					node.next.previous = &node
					toCheck = append(toCheck, *node.next)
					visited = append(visited, newP)
					if node.next.point.isWinningPoint(exitPoint) {
						goalList = append(goalList, node)
					}
				}
			}
		}
		edge = make([]Node, 0)
		if len(toCheck) > 0 {
			for _, nodeToCheck := range toCheck {
				edge = append(edge, nodeToCheck)
			}
		}
	}

	if len(goalList) > 0 {
		result := goalList[0]
		return result.getFirstCmd()
	}
	return WAIT
}

func debug(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
}

// ---------------------------
type CMD string

func (cmd CMD) canExecute() bool {
	if cmd == "BLOCK" {
		return nbClones > 0
	}
	return true
}
func (cmd CMD) execute() CMD {
	if cmd == "BLOCK" {
		nbClones--
	}
	return cmd;
}

// ---------------------------
type Direction struct {
	name  string
	value int
}

func getDirection(dir string) Direction {
	if dir == "RIGHT" {
		return RIGHT
	} else {
		return LEFT
	}
}

func (d *Direction) opposite() Direction {
	if d.name == "RIGHT" {
		return LEFT
	} else {
		return RIGHT
	}
}

// ---------------------------
type Point struct {
	x, y int
}
func (p Point) equals(p1 Point) bool {
	return p.x == p1.x && p.y == p1.y
}

func (p *Point) add(direction Direction) Point {
	return Point{p.x + direction.value, p.y}
}

func (p *Point) canMoveHere(visited []Point) bool {
	for _, visitedP := range visited {
		if p.equals(visitedP) {
			return false
		}
	}
	if p.x >= 0 && p.x < width && p.y < nbFloors {
		return true
	}
	return false
}
func (p *Point) newPoint(direction Direction, cmd CMD) (Point, Direction) {
	var newP Point
	var newDirection Direction
	if cmd == "WAIT" {
		newDirection = direction
	} else {
		newDirection = direction.opposite()
	}
	newP = p.add(newDirection)
	return newP, newDirection
}
func (p *Point) isWinningPoint(exitPoint Point) bool {
	return p.y == exitPoint.y && p.x == exitPoint.x
}

// ---------------------------
type Node struct {
	point      Point
	length     int
	direction  Direction
	cmd        CMD
	next, previous *Node
}
func NewNode(point Point, direction Direction) *Node {
	node := new(Node)
	node.point = point
	node.length = 0
	node.direction = direction
	node.cmd = WAIT
	return node
}
func (n Node) getFirstCmd() CMD {
	current := n
	prev := current
	for current.previous != nil {
		prev = current
		current = *current.previous
	}
	return prev.cmd
}
