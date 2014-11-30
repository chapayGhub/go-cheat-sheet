package main

import "fmt"
//import "os"

const (
	SPEED = "SPEED"
	SLOW = "SLOW"
	JUMP = "JUMP"
	WAIT = "WAIT"
)

func main() {
	// R: the length of the road before the gap.
	var R int
	fmt.Scan(&R)

	// G: the length of the gap.
	var G int
	fmt.Scan(&G)

	// L: the length of the landing platform.
	var L int
	fmt.Scan(&L)

	for {
		// S: the motorbike's speed.
		var S int
		fmt.Scan(&S)

		// X: the position on the road of the motorbike.
		var X int
		fmt.Scan(&X)

		var cmd string
		if isBeforeGap(R, X) {
			if shouldJump(R, X, S, G) {
				cmd = JUMP
			} else {
				if isSpeedEnough(S, G) {
					if (shouldSlowDow(S, G)) {
						cmd = SLOW
					} else {
						cmd = WAIT
					}
				} else {
					cmd = SPEED
				}
			}
		} else {
			cmd = SLOW
		}

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		fmt.Println(cmd) // A single line containing one of 4 keywords: SPEED, SLOW, JUMP, WAIT.
	}
}

func isBeforeGap(R, X int) bool {
	return R - X > 0
}

func isSpeedEnough(S, G int) bool {
	return S > G
}

func shouldSlowDow(S, G int) bool {
	return S > G + 1
}

func shouldJump(R, X, S, G int) bool {
	nextR := X + S
	roadPlusGap := R + G
	return nextR - roadPlusGap >= 0
}
