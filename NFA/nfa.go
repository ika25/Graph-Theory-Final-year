package main
import ("fmt")


type state struct {
    symbol rune
    edge1 *state
    edge2 *state
}

func poregtonenfa(pofix string){

}

func main(){
	nfa := poregtonenfa("ab.c*|")
	fmt.Println(nfa)
}