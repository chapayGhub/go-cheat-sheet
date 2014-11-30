package main

import (
	"fmt"
	"os"
	"math"
)

var W, H int
var U, UR, R, DR, D, DL, L, UL BombDir
var BOMB_DIR_LIST [8]BombDir

func initBombDir() {
	U := BombDir{"U", 0, -1}
	R := BombDir{"R", 1, 0}
	D := BombDir{"D", 0, 1}
	L := BombDir{"L", -1, 0}

	UR := BombDir{U.name + R.name, U.x + R.x, U.y + R.y}
	DR := BombDir{D.name + R.name, D.x + R.x, D.y + R.y}
	DL := BombDir{D.name + L.name, D.x + L.x, D.y + L.y}
	UL := BombDir{U.name + L.name, U.x + L.x, U.y + L.y}

	BOMB_DIR_LIST[0] = U
	BOMB_DIR_LIST[1] = R
	BOMB_DIR_LIST[2] = D
	BOMB_DIR_LIST[3] = L
	BOMB_DIR_LIST[4] = UR
	BOMB_DIR_LIST[5] = DR
	BOMB_DIR_LIST[6] = DL
	BOMB_DIR_LIST[7] = UL
}

func main() {
	initBombDir()

	// W: width of the building.
	// H: height of the building.
	fmt.Scan(&W, &H)

	// N: maximum number of turns before game over.
	var N int
	fmt.Scan(&N)

	var X0, Y0 int
	fmt.Scan(&X0, &Y0)

	X1 := X0
	Y1 := Y0
	minX := 0
	minY := 0
	maxX := W
	maxY := H

	round := 1
	for {
		// BOMB_DIR: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
		var BOMB_DIR string
		fmt.Scan(&BOMB_DIR)
		bombDir := getBombDir(BOMB_DIR)

		X0 = X1
		Y0 = Y1

		X1 = computeNewValue(X1, bombDir.x, minX, maxX)
		Y1 = computeNewValue(Y1, bombDir.y, minY, maxY)

		minX, maxX = computeMax(X0, X1, minX, maxX)
		minY, maxY = computeMax(Y0, Y1, minY, maxY)

		fmt.Println(newLocation(X1, Y1)) // the location of the next window Batman should jump to.

		round++
	}
}

func valueToAdd(bombDirValue int) int {
	if bombDirValue > 0 {
		return 1
	} else if bombDirValue < 0 {
		return -1
	}
	return 0
}
func computeMax(value0, value1, oldMinValue, oldMaxValue int) (minValue, maxValue int) {
	if value1 > value0 {
		minValue = value0
		maxValue = oldMaxValue
	} else if value1 < value0 {
		maxValue = value0
		minValue = oldMinValue
	} else {
		minValue = oldMinValue
		maxValue = oldMaxValue
	}
	return
}

func computeNewValue(value, bombDirValue, minValue, maxValue int) int {
	debug(fmt.Sprintf("value = %d, bombDirValue = %d, minValue = %d, maxValue = %d", value, bombDirValue, minValue, maxValue))
	newValue := value
	switch  {
	case bombDirValue > 0:
		// Going R
		newValue = int(math.Ceil(float64((maxValue + value) / 2)))
	case bombDirValue < 0:
		// Going L
		newValue = int(math.Ceil(float64((value + minValue) / 2)))
	default:
		// Staying in the same column
		newValue = value
	}
	return newValue
}

func newLocation(X, Y int) string {
	return fmt.Sprintf("%d %d", X, Y)
}

func debug(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
}

// -------------------
type BombDir struct {
	name string
	x, y int
}
func getBombDir(BOMB_DIR string) BombDir {
	for _, bombDir := range BOMB_DIR_LIST {
		if BOMB_DIR == bombDir.name {
			return bombDir
		}
	}
	return BOMB_DIR_LIST[0]
}
// -------------------
