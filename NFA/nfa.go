package main

import (
	"fmt"
)

type state struct {
	// 	If edges are labeled with epsylon symbol = 0 (default value)
	// 	If edges are labeled with some other letter
	symbol rune
	edge1  *state
	edge2  *state
}

type nfa struct {
	// Keeps track of initial and accept states and builds a list of connected state structs representing nfa
	// If these states won't be kept track of, the final state would only be found when going through the whole linked list
	initial *state
	accept  *state
}

//converts a postfix regular expression to an nfa
func poregtonfa(pofix string) *nfa {
	// Thompson's algorithm keeps track of fragments of nfa's on a stack
	// []*nfa{} - It gets an array of empty nfa pointers
	nfastack := []*nfa{}

	// Loop through postfix and the stack is modified depending on the characters from the postfix
	for _, r := range pofix {
		switch r {
		case '.': //pops two fragments off the nfa stack, it joins them together and pushes the new fragment onto the nfa stack

			// pops two fragments off the nfa stack
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			// Join the two fragments together in a concatenated fragment and push it back to the nfa stack
			frag1.accept.edge1 = frag2.initial
			// Pushes new concatenated fragment onto the nfa stack
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

		case '|': //pops two fragments off the nfa stack,it joins them to new created states and pushes new fragmetn onto nfa stack
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			//create two new states
			accept := state{}
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '*': // shows one fragment off the nfa stack and creates two new states and pushes new fragment onto nfa stack
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		default: //diffrent non special characters (a,b,0,1) fragmetn pushed to stack.
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		}
	}
	return nfastack[0]
}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)
}
