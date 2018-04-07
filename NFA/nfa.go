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
	if len(nfastack) != 1 {
		fmt.Println("Shoudl count 1 Thing", len(nfastack), nfastack)
	}

	return nfastack[0]
}

func addState(list []*state, single *state, accept *state) []*state { //list of pointers to states, single pointer to a state and accept state
	list = append(list, single)
	// Any state that has the rune value 0 then the state has empty arrows coming from it.
	if single != accept && single.symbol == 0 {
		list = addState(list, single.edge1, accept)
		if single.edge2 != nil {
			list = addState(list, single.edge2, aaccept)
		}

	}

	return list
}

func pomatchfix(pox string, s string) bool { //find out if pofix regexp matches the string
	ismatch := false
	pofnfa := poregtonfa(pox)

	current := []*state{} //array of pointers to state. List of states that we are currently in on NFA
	next := []*state{}    // states we can get to

	current = addState(current[:], pofnfa.initial, pofnfa.accept) //pass current state, pass initial state and accept state. [:] is slice, for when you pass an array and want to change it

	for _, r := range s { //r is rune. loop through s a character at a time
		for _, c := range current { //c is current state we are in. loop through the current states
			if c.symbol == r { // if the sybmol is the same as the one currently reading from s
				next = addState(next[:], c.edge1, pofnfa.accept)
			}
		}
		current, next = next, []*state{} //move current states to next states. Next array made blank
	}

	for _, c := range current { //loop through the state that you end up in at the very end
		if c == pofnfa.accept {
			ismatch = true
			break // from this point is match will always be true
		}
	}

	return ismatch
}

func main() {
	nfa := poregtonfa("ab.c*|")
	fmt.Println(nfa)

	fmt.Println(pofixMatch("ab.c*|", "cccc"))
}
