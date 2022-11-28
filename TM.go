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

func createTM() *tm {
	newTM := new(tm)
	newTM.states = []state{}
	newTM.inputs = []input{}
	newTM.tapesyms = []tape{}
	newTM.trans = []trans{}
	newTM.final = []state{}
	return newTM
}

func main() {
	fmt.Println("hello")
}
