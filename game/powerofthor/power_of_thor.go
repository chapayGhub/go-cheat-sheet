package main

import "fmt"
import "os"

type Direction struct {
	value, x, y int
	cmd string
}
func (d Direction) String() string {
	return d.cmd
}

var N, NE, E, SE, S, SW, W, NW Direction
var directions [8]Direction

func initDirections() {
	N = Direction{-10, 0, -1, "N"}
	E = Direction{1, 1, 0, "E"}
	S = Direction{10, 0, 1, "S"}
	W = Direction{-1, -1, 0, "W"}

	NE = Direction{N.value + E.value, N.x + E.x, N.y + E.y, N.cmd + E.cmd}
	SE = Direction{S.value + E.value, S.x + E.x, S.y + E.y, S.cmd + E.cmd}
	SW = Direction{S.value + W.value, S.x + W.x, S.y + W.y, S.cmd + W.cmd}
	NW = Direction{N.value + W.value, N.x + W.x, N.y + W.y, N.cmd + W.cmd}

	directions[0] = N
	directions[1] = E
	directions[2] = S
	directions[3] = W
	directions[4] = NE
	directions[5] = SE
	directions[6] = SW
	directions[7] = NW
}

func getDirections(value int) Direction {
	for _, direction := range directions {
		if direction.value == value {
			return direction
		}
	}
	return directions[0]
}

func computeDir(L, T, multi int) int {
	if L > T {
		return 1 * multi
	}
	if L < T {
		return -1 * multi
	}
	return 0
}

func main() {
	initDirections()
	// LX: the X position of the light of power
	// LY: the Y position of the light of power
	// TX: Thor's starting X position
	// TY: Thor's starting Y position
	var LX, LY, TX, TY int
	fmt.Scan(&LX, &LY, &TX, &TY)

	for {
		// E: The level of Thor's remaining energy, representing the number of moves he can still make.
		var energy int
		fmt.Scan(&energy)

		x, y := computeDir(LX, TX, 1), computeDir(LY, TY, 10)

		// fmt.Fprintln(os.Stderr, "Debug messages...")
		move := getDirections(x + y)

		TX += move.x
		TY += move.y
		debug(TX)
		debug(TY)

		fmt.Println(move) // A single line providing the move to be made: N NE E SE S SW W or NW
	}
}

func debug(msg int) {
	fmt.Fprintln(os.Stderr, msg)
}
