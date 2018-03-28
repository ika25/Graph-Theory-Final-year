package main

import (
	"fmt"
)

func intopost(infix string) string {
	pofix := []rune{}

	return string(pofix)
}

func main() {
	// answer: ab.c*.
	fmt.Println("index:	", "a.b.c")
	fmt.Println("Postfix: ", intopost("a.b.c"))

	// Answer: abd|.*
	fmt.Println("index:	 ", "(a.(b|d))*")
	fmt.Println("index:	 ", intopost("(a.(b|d))*"))

	// Answer: abd|.c*.
	fmt.Println("Index:	", "a.(b|d).c*")
	fmt.Println("Index:	", intopost("a.(b|d).c*"))

	// Answer: abb.+.c.
	fmt.Println("Index:	", "a.(b.b)+.c")
	fmt.Println("Index:	", intopost("a.(b.b)+.c"))

}
