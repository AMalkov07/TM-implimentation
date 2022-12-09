package main

import "fmt"

const (
	moveLeft = iota
	moveRight
)

type direction int
type state int
type tape rune
type input rune

type trans struct {
	//t  *trans
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
	blank    tape
	leftend  tape
	trans    []trans
	start    state
	final    []state
}

func createTM(as []state, in []input, ts []tape, bl tape, le tape, trns []trans, st state, fnl []state) *tm {
	newTM := new(tm)
	newTM.states = as
	newTM.inputs = in
	newTM.tapesyms = ts
	newTM.blank = bl
	newTM.leftend = le
	newTM.trans = trns
	newTM.start = st
	newTM.final = fnl
	return newTM
}

type config struct {
	currentState state
	blank        tape
	lefttrev     []tape
	right        []tape
}

type configs []config
type history []configs

func (c config) shiftConfig(d direction) config {
	if d == moveRight {
		if len(c.right) == 0 {
			return config{c.currentState, c.blank, append(c.lefttrev, c.blank), c.right} //possible this line can cause problems
		}
		return config{c.currentState, c.blank, append(c.lefttrev, c.right[0]), c.right[1:]} //possible this line can cause problems
	}
	return config{c.currentState, c.blank, c.lefttrev[:len(c.lefttrev)-1], append([]tape{c.lefttrev[len(c.lefttrev)-1]}, c.right...)} //possible this line can cause problems
}

//not sure if this function is correct
func (c config) updateConfig(s state, t tape, d direction) config {
	arrCopy := make([]tape, len(c.lefttrev))
	copy(arrCopy, c.lefttrev)
	newConfig := config{s, c.blank, append(arrCopy[:len(arrCopy)-1], t), c.right}
	return newConfig.shiftConfig(d)
}

func (t tm) newConfig(c config) configs {
	var output configs
	for _, tr := range t.trans {
		if tr.is == c.currentState && tr.it == c.lefttrev[len(c.lefttrev)-1] {
			output = append(output, c.updateConfig(tr.ns, tr.nt, tr.d))
		}
	}
	return output
}

func (t tm) initialConfig(inputString []tape) config { //changed type input to type tape
	return config{t.start, t.blank, []tape{t.leftend, inputString[0]}, inputString[1:]} // haskel version uses an infinite length slice, mine doesn't
}

func (t tm) configsLazy(inputString []tape) history {
	var output history
	configsStack := configs{t.initialConfig(inputString)}
	output = append(output, configsStack)
	for {
		if len(configsStack) == 0 {
			break
		}
		configsStack = append(configsStack, t.newConfig(configsStack[0])...)
		configsStack = configsStack[1:]
		output = append(output, configsStack)
	}
	return output
}

func (t tm) configs(n int, inputString []tape) history {
	output := make(history, n)
	output = t.configsLazy(inputString)
	return output

}

func (t tm) accepts(inputString []tape) config {
	var output config

	return output
}

func goRight(initialState state, initialTape tape, newTape tape, newState state) trans {
	return trans{initialState, initialTape, moveRight, newState, newTape}
}

func checkRight(initialState state, initialTape tape, newState state) trans {
	return goRight(initialState, initialTape, initialTape, newState)
}

func goLeft(initialState state, initialTape tape, newTape tape, newState state) trans {
	return trans{initialState, initialTape, moveLeft, newState, newTape}
}

func checkLeft(initialState state, initialTape tape, newState state) trans {
	return goLeft(initialState, initialTape, initialTape, newState)
}

func loop(d direction, st state, tapes []tape) []trans {
	output := []trans{}
	for _, val := range tapes {
		output = append(output, trans{st, val, d, st, val})
	}
	return output
}

func loopRight(st state, tapes []tape) []trans {
	return loop(moveRight, st, tapes)
}

func loopLeft(st state, tapes []tape) []trans {
	return loop(moveLeft, st, tapes)
}

func main() {
	transitions := []trans{}
	transitions = append(transitions, checkRight(1, ' ', 6))
	transitions = append(transitions, loopRight(1, []tape{'*'})...)
	transitions = append(transitions, goRight(1, 'a', '*', 2))
	transitions = append(transitions, goRight(1, 'a', '?', 3)) // added for testing
	transitions = append(transitions, loopRight(2, []tape{'a', '*'})...)
	transitions = append(transitions, goRight(2, 'b', '*', 3))
	transitions = append(transitions, loopRight(3, []tape{'b', '*'})...)
	transitions = append(transitions, goRight(3, 'c', '*', 4))
	transitions = append(transitions, loopRight(4, []tape{'c', '*'})...)
	transitions = append(transitions, checkLeft(4, ' ', 5))
	transitions = append(transitions, loopLeft(5, []tape{'a', 'b', 'c', '*'})...)
	transitions = append(transitions, checkRight(5, '!', 1))

	tripletm := createTM([]state{1: 6}, []input{'a', 'b', 'c'}, []tape{'a', 'b', 'c', '*', '!', ' '}, ' ', '!', transitions, 1, []state{6})

	//ic := tripletm.initialConfig([]tape{'a', 'b', 'c'})
	//fmt.Println(ic)
	//uc := ic.updateConfig(2, '*', moveRight).updateConfig(3, 'b', moveRight).updateConfig(4, '*', moveRight)
	//nc := tripletm.newConfig(ic)
	//fmt.Println(nc)
	x := tripletm.configs(35, []tape{'a', 'a', 'b', 'b', 'c', 'c'})
	for _, val := range x {
		fmt.Println(val)
	}

}
