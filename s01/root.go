package s01

import "fmt"

var Challenges = []func([]string){challenge01}

func challenge01(args []string) {
	fmt.Printf("Challenge 1: %v\n", args)
}
