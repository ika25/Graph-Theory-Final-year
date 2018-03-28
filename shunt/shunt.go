package main

import (
	"fmt"
)

func intopost(infix string) string { //function to convert inflix regExp to postfix regExp

	specials := map[rune]int{'*': 10, '.': 9, '|': 8} //create a map for special characters which will be used in this program

	pofix := []rune{}
	s := []rune{}

	for _, r := range infix { //loop through the range of infix (input) and return the index of the current character on each loop. _ ignores the index. r is the character, string converted to rune

		switch {

		case r == '(':
			s = append(s, r)

		case r == ')':
			for s[len(s)-1] != '(' { //while the last character isnt an open bracket. Pop off the stack and push onto pofix
				pofix = append(pofix, s[len(s)-1]) //append what is on the top of the stack
				s = s[:len(s)-1]                   //everything in s except the last character
			}
			s = s[:len(s)-1] //kick the round bracket off the end of the stack

		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] { //while there is still something on the stack and the precedence of the cuurent is less than the character at the top of the stack
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r) //character has less precedence than the current character
		default:
			pofix = append(pofix, r)

		}
	}

	for len(s) > 0 { //if there is anything on the top of the stack append it

		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1] //takes the top charcter of the stack and puts it as the top character of pofix

	}

	return string(pofix)
}

func main() {
	// answer: ab.c*.
	fmt.Println("index:	", "a.b.c*") //A followed by B, followed by zero or more C
	fmt.Println("Postfix: ", intopost("a.b.c*"))

	// Answer: abd|.*
	fmt.Println("index:	 ", "(a.(b|d))*")
	fmt.Println("index:	 ", intopost("(a.(b|d))*")) //Zero or more of A followed by B or D

	// Answer: abd|.c*.
	fmt.Println("Index:	", "a.(b|d).c*")
	fmt.Println("Index:	", intopost("a.(b|d).c*"))

	// Answer: abb.+.c.
	fmt.Println("Index:	", "a.(b.b)+.c")
	fmt.Println("Index:	", intopost("a.(b.b)+.c"))

}
