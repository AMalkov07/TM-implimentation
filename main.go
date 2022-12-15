package main

func main() {
	transitionsarr := [][]trans{checkRight(1, ' ', 6),
		loopRight(1, []tape{'*'}),
		goRight(1, 'a', '*', 2),
		loopRight(2, []tape{'a', '*'}),
		goRight(2, 'b', '*', 3),
		loopRight(3, []tape{'b', '*'}),
		goRight(3, 'c', '*', 4),
		loopRight(4, []tape{'c', '*'}),
		checkLeft(4, ' ', 5),
		loopLeft(5, []tape{'a', 'b', 'c', '*'}),
		checkRight(5, '!', 1)}
	transitions := combineArr(transitionsarr)
	tripletm := createTM([]state{1: 6}, []input{'a', 'b', 'c'}, []tape{'a', 'b', 'c', '*', '!', ' '}, ' ', '!', transitions, 1, []state{6})
	tripletm.configs(35, []tape{'a', 'a', 'b', 'b', 'c', 'c'}).showHistory()
	tripletm.accepts([]tape{'a', 'a', 'b', 'b', 'c', 'c'})
}
