package sudoku_solver

func solveSudoku(board [][]byte) {
	g := newGame(board)
	g.estimateOptions()
	g.reduceOptions()

	if g.cellsToSolve > 0 {
		g.quickSave()

		for g.cellsToSolve > 0 {
			r, c, d := g.bestGuess()
			g.board[r][c] = d
			g.options[r][c] = map[byte]bool{}
			g.optionsTotal--
			g.cellsToSolve--
			g.updateOptions()
			g.reduceOptions()

			if g.inDeadEnd() {
				g.quickLoad()
				delete(g.options[r][c], d)
				g.updateOptions()
				g.quickSave()
			}
		}
	}
}

type game struct {
	board        [][]byte
	options      [][]map[byte]bool
	optionsTotal int
	cellsToSolve int
	save         *game
}

func newGame(board [][]byte) game {
	return game{board: board, options: make([][]map[byte]bool, 9)}
}

func (g *game) estimateOptions() {
	for r := 0; r < 9; r++ {
		g.options[r] = make([]map[byte]bool, 9)
		for c := 0; c < 9; c++ {
			if g.board[r][c] != byte('.') {
				continue
			}
			g.cellsToSolve++
			g.options[r][c] = map[byte]bool{
				'1': true,
				'2': true,
				'3': true,
				'4': true,
				'5': true,
				'6': true,
				'7': true,
				'8': true,
				'9': true,
			}
			g.optionsTotal += 9

			g.excludeOptionsFromRowAndColumn(r, c)
			g.excludeOptionsFromSubBox(r, c)
		}
	}
}

func (g *game) excludeOptionsFromRowAndColumn(r, c int) {
	for rr := 0; rr < 9; rr++ {
		if rr == r {
			for cc := 0; cc < 9; cc++ {
				if g.options[r][c][g.board[rr][cc]] {
					delete(g.options[r][c], g.board[rr][cc])
					g.optionsTotal--
				}
			}
		} else {
			if g.options[r][c][g.board[rr][c]] {
				delete(g.options[r][c], g.board[rr][c])
				g.optionsTotal--
			}
		}
	}
}

func (g *game) excludeOptionsFromSubBox(r, c int) {
	var sr, er int
	switch {
	case r < 3:
		sr, er = 0, 3
	case r > 5:
		sr, er = 6, 9
	default:
		sr, er = 3, 6
	}

	var sc, ec int
	switch {
	case c < 3:
		sc, ec = 0, 3
	case c > 5:
		sc, ec = 6, 9
	default:
		sc, ec = 3, 6
	}

	for rr := sr; rr < er; rr++ {
		for cc := sc; cc < ec; cc++ {
			if g.options[r][c][g.board[rr][cc]] {
				delete(g.options[r][c], g.board[rr][cc])
				g.optionsTotal--
			}
		}
	}

}

func (g *game) updateOptions() {
	for r := range g.options {
		for c := range g.options[r] {
			g.excludeOptionsFromRowAndColumn(r, c)
			g.excludeOptionsFromSubBox(r, c)
		}
	}
}

func (g *game) reduceOptions() {
	for {
		optionsAtStart := g.optionsTotal

		for r := range g.options {
			for c := range g.options[r] {
				if len(g.options[r][c]) == 1 {
					for d := range g.options[r][c] {
						g.board[r][c] = d
						g.options[r][c] = map[byte]bool{}
						g.optionsTotal--
						g.cellsToSolve--
					}
				}
			}
		}

		if g.optionsTotal != 0 && g.optionsTotal != optionsAtStart {
			g.updateOptions()
		}

		g.reduceOptionsInSubBoxes()

		if g.optionsTotal == 0 || g.optionsTotal == optionsAtStart {
			break
		}
	}
}

func (g *game) reduceOptionsInSubBoxes() {
	needUpdate := false
	type check struct{ i, r, c int }

	for rs := 0; rs < 3; rs++ {
		for cs := 0; cs < 3; cs++ {
			optionsInSubBox := make(map[byte]*check, 9)

			for r := rs * 3; r < rs*3+3; r++ {
				for c := cs * 3; c < cs*3+3; c++ {
					for d := range g.options[r][c] {
						if optionsInSubBox[d] == nil {
							optionsInSubBox[d] = &check{r: r, c: c}
						}
						optionsInSubBox[d].i++
					}
				}
			}

			for d, ch := range optionsInSubBox {
				if ch.i == 1 {
					g.board[ch.r][ch.c] = d
					g.options[ch.r][ch.c] = map[byte]bool{}
					g.optionsTotal--
					g.cellsToSolve--
					needUpdate = true
				}
			}
		}
	}

	if needUpdate {
		g.updateOptions()
	}
}

func (g *game) quickSave() {
	g.save = &game{board: make([][]byte, 9), options: make([][]map[byte]bool, 9)}

	for r := 0; r < 9; r++ {
		g.save.board[r] = make([]byte, 9)
		copy(g.save.board[r], g.board[r])

		g.save.options[r] = make([]map[byte]bool, 9)
		for c := range g.options[r] {
			g.save.options[r][c] = make(map[byte]bool, 9)
			for d := range g.options[r][c] {
				g.save.options[r][c][d] = g.options[r][c][d]
			}
		}
	}

	g.save.cellsToSolve, g.save.optionsTotal = g.cellsToSolve, g.optionsTotal
}

func (g *game) quickLoad() {
	for r := 0; r < 9; r++ {
		copy(g.board[r], g.save.board[r])
	}
	g.options, g.cellsToSolve, g.optionsTotal = g.save.options, g.save.cellsToSolve, g.save.optionsTotal
}

func (g *game) bestGuess() (r, c int, d byte) {
	minOptions := 10
	for rr := 0; rr < 9; rr++ {
		for cc := 0; cc < 9; cc++ {
			if len(g.options[rr][cc]) < minOptions {
				for dd := range g.options[rr][cc] {
					minOptions, r, c, d = len(g.options[rr][cc]), rr, cc, dd
					continue
				}
			}
		}
	}
	return r, c, d
}

func (g *game) inDeadEnd() bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if g.board[r][c] == byte('.') && len(g.options[r][c]) == 0 {
				return true
			}
		}
	}
	return false
}
