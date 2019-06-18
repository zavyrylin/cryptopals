package main

import "os"
import "fmt"
import "regexp"
import "strconv"
import "cryptopals/s01"

var Sets = [][]func([]string){
	s01.Challenges,
}

func errexit(msg string, xs ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", xs...)
	os.Exit(1)
}

func parseSetAndChallenge(s string) (int, int) {
	re, _ := regexp.Compile("s(\\d{1,2})ch(\\d{1,2})")

	groups := re.FindStringSubmatch(s)
	if len(groups) < 3 {
		errexit("Unsupported challenge name format, given: %v", s)
	}

	set, _ := strconv.ParseInt(groups[1], 10, 32)
	ch, _ := strconv.ParseInt(groups[2], 10, 32)

	return int(set), int(ch)
}

func main() {
	if len(os.Args) < 2 {
		errexit("Expected at least 2 arguments")
	}

	sxxchxx := os.Args[1]
	challengeArgs := os.Args[2:]

	setNo, chNo := parseSetAndChallenge(sxxchxx)

	if len(Sets) < setNo {
		errexit("There are %v sets only", len(Sets))
	}

	challenges := Sets[setNo-1]
	if len(challenges) < chNo {
		errexit("There are only %v challenges in set %v", len(challenges), setNo)
	}

	challenges[chNo-1](challengeArgs)
}
