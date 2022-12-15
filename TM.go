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

func (tr trans) showTransition() {
	var tmp rune
	if tr.d == 0 {
		tmp = 'L'
	} else {
		tmp = 'R'
	}

	fmt.Printf("  %v === '%v' / '%v' %c ===> %v\n", tr.is, tr.it, tr.nt, tmp, tr.ns)
}

func (t tm) showTM() {
	fmt.Println("States:", t.states)
	fmt.Printf("Alphabet: %c\n", t.inputs)
	fmt.Printf("Tape symbols: %c\n", t.tapesyms)
	fmt.Printf("Blank: '%c'\n", t.blank)
	fmt.Printf("Leftend '%c'\n", t.leftend)
	fmt.Println("Transitions:")
	for _, tr := range t.trans {
		tr.showTransition()
	}
	fmt.Printf("Start state: %v\n", t.start)
	fmt.Printf("Final state: %v\n", t.final)
}

func (c config) showConfig() {
	fmt.Printf("[%v: %c %c]", c.currentState, c.lefttrev, c.right)
}

func (h history) showHistory() {
	for _, configs := range h {
		fmt.Printf("[")
		configs[0].showConfig()
		for i := 1; i < len(configs); i++ {
			fmt.Printf(",")
			configs[i].showConfig()
		}
		fmt.Printf("]\n")
	}
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
	output := t.configsLazy(inputString)[:n]
	return output
}

func (t tm) accepting(inputString []tape) config {
	htmp := t.configsLazy(inputString)
	for _, val := range htmp {
		for _, val2 := range val {
			for _, val3 := range t.final {
				if val2.currentState == val3 {
					return val2
				}
			}
		}
	}
	return config{currentState: t.start}
}

func (t tm) accepts(inputString []tape) bool {
	for _, val := range t.final {
		if val == t.start {
			fmt.Println("true")
			return true
		}
	}
	x := t.accepting(inputString)
	fmt.Println(x.currentState != t.start)
	return x.currentState != t.start
}

func goRight(initialState state, initialTape tape, newTape tape, newState state) []trans {
	return []trans{trans{initialState, initialTape, moveRight, newState, newTape}}
}

func checkRight(initialState state, initialTape tape, newState state) []trans {
	return goRight(initialState, initialTape, initialTape, newState)
}

func goLeft(initialState state, initialTape tape, newTape tape, newState state) []trans {
	return []trans{trans{initialState, initialTape, moveLeft, newState, newTape}}
}

func checkLeft(initialState state, initialTape tape, newState state) []trans {
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

func combineArr(arr [][]trans) []trans {
	output := []trans{}
	for _, val := range arr {
		output = append(output, val...)
	}
	return output
}
