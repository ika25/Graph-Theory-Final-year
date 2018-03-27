package main

import (
	"fmt"
)

func intopost(infix string) string{
	pofix := []rune{}

	return string(profix)
}

func main(){
	// answer: ab.c*.
	fmt.Prinln("index:	","a.b.c")
	fmt.Prinln("Postfix: ", intopost("a.b.c"))

	// Answer: abd|.*
	fmt.Prinln("index:	 ", "(a.(b|d))*")
	fmt.Prinln("index:	 ", intopost("(a.(b|d))*"))

	// Answer: abd|.c*.
	fmt.Println("Index:	", "a.(b|d).c*")
	fnt.Println("Index:	", intopost("a.(b|d).c*"))

	// Answer: abb.+.c.
	fmt.Println("Index:	", "a.(b.b)+.c")
	fmt.Println("Index:	", intopost("a.(b.b)+.c"))

}
