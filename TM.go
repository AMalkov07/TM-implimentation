package main

import "fmt"

const (
	goLeft = iota
	goRight
)

type direction int
type state int
type tape rune
type input rune

type trans struct {
	t  *trans
	is state     //initial state
	it tape      //inital tape character
	d  direction //left or right
	ns state     //new state
	nt tape      //new tape character
}

type tm struct {
	states   []state
	inputs   []input
	tapesyms []tape
	leftend  tape
	trans    []trans
	start    state
	final    []state
}

func createTM(as []state, in []input, ts []tape, le tape, trns []trans, st state, fnl []state) *tm {
	newTM := new(tm)
	newTM.states = as
	newTM.inputs = in
	newTM.tapesyms = ts
	newTM.leftend = le
	newTM.trans = trns
	newTM.start = st
	newTM.final = fnl
	return newTM
}

func main() {
	//tripletm := tm{nil, "abc", "abc*! ", ' ', '!', 5, 1, []int{6}}
	tripletm := createTM([]state{1: 6}, []input{'a', 'b', 'c'}, []tape{'a', 'b', 'c', '*', '!', ' '}, '!', nil, 1, []state{6})
	fmt.Println(tripletm)
}
