package main

import "fmt"

//thompsons "./thompsons"
//utils "./utils"

func shunt(infix string) string { //function to convert inflix regExp to postfix regExp

	specials := map[rune]int{'*': 10, '.': 9, '|': 8} //create a map for special characters which will be used in this program

	pofix := []rune{}
	s := []rune{}

	for _, r := range infix {
		switch {
		case r == "(":
			s = append(s, r)

		case r == ")":
			for s[len(s)-1] != "(" {
				postfix = append(psotfix, s[len(s)-1])
				s = s[:len(s)-1]
			}

			s = s[:len(s)-1]

		case prec[r] > 0:
			for len(s) > 0 && prec[r] <= prec[s[len(s)-1]] {
				postfix = append(pofix, s[len(s)-1])
				s = s[:len(s)-1]
			}

			s = append(s, r)
		default:
			postfix = append(postfix, r)
		}

	}

	for len(s) > 0 {
		postfix = append(postfix, s[len(s)-1])
		s = s[:len(s)-1]
	}

	return string(postfix)
}


