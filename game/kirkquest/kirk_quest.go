package main

import (
	"fmt"
	"os"
)

const (
	FIRE = "FIRE"
	HOLD = "HOLD"
	MAX_SX = 7
	MAX_SY = 10
	MAX_MH = 9
)

func main() {
	for {
		var SX, SY int
		fmt.Scan(&SX, &SY)

		var mountainsHeight [MAX_SX + 1]int
		for i := 0; i < MAX_SX + 1; i++ {
			// MH: represents the height of one mountain, from 9 to 0. Mountain heights are provided from left to right.
			var MH int
			fmt.Scan(&MH)
			mountainsHeight[i] = MH
		}
		var maxMH int
		for i := 0; i < MAX_SX + 1; i++ {
			if mountainsHeight[i] > mountainsHeight[maxMH] {
				maxMH = i
			}
		}

		if SX == maxMH {
			fmt.Println(FIRE)
		} else {
			fmt.Println(HOLD)
		}
	}
}

func debug(msg interface {}) {
	fmt.Fprintln(os.Stderr, msg)
}
