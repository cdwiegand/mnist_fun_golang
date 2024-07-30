package main

type GuessResult struct {
	CountedRight int
	CountedWrong int
}

func NewGuessResult(wasCorrect bool) GuessResult {
	if wasCorrect {
		return GuessResult{CountedRight: 1}
	}
	return GuessResult{CountedWrong: 1}
}

func (g GuessResult) Total() int {
	return g.CountedRight + g.CountedWrong
}

func (g GuessResult) Percent() float64 {
	return float64(g.CountedRight) / float64(g.CountedRight+g.CountedWrong)
}
