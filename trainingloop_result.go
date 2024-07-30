package main

import (
	"fmt"
	"time"
)

type TrainingLoopResult struct {
	LoopGeneration int
	Characters     map[byte]GuessResult
	StartTime      int64
	EndTime        int64
}

func NewTrainingLoopResult(gen int) TrainingLoopResult {
	return TrainingLoopResult{
		LoopGeneration: gen,
		StartTime:      time.Now().UTC().Unix(),
	}
}

func (res *TrainingLoopResult) CountedCorrect() (ret int) {
	ret = 0
	for _, j := range res.Characters {
		ret += j.CountedRight
	}
	return ret
}
func (res *TrainingLoopResult) CountedWrong() (ret int) {
	ret = 0
	for _, j := range res.Characters {
		ret += j.CountedWrong
	}
	return ret
}
func (res *TrainingLoopResult) CountedTotal() (ret int) {
	ret = 0
	for _, j := range res.Characters {
		ret += j.CountedRight + j.CountedWrong
	}
	return ret
}
func (res *TrainingLoopResult) Accuracy() (ret float64) {
	return float64(res.CountedCorrect()) / float64(res.CountedTotal())
}

func (res *TrainingLoopResult) CountRight(value byte) {
	val, ok := res.Characters[value]
	if !ok {
		res.Characters[value] = NewGuessResult(true)
	} else {
		val.CountedRight++
	}
}

func (res *TrainingLoopResult) CountWrong(value byte) {
	val, ok := res.Characters[value]
	if !ok {
		res.Characters[value] = NewGuessResult(false)
	} else {
		val.CountedWrong++
	}
}

func (res *TrainingLoopResult) Duration() int64 {
	return res.EndTime - res.StartTime
}

func (res *TrainingLoopResult) ToString() (ret string) {
	ret = fmt.Sprintf("\nEpoch: %d: %f @ %d\n", res.LoopGeneration, res.Accuracy(), res.Duration())
	for i, j := range res.Characters {
		ret = fmt.Sprintf(" %d: %f", i, j.Percent())
	}
	return
}

// percent is 0->1, so 100 is never the worst unless there is no data
func (res *TrainingLoopResult) WorstAccuracy() (ret float64) {
	ret = float64(100)
	for _, j := range res.Characters {
		if j.Percent() < ret {
			ret = j.Percent()
		}
	}
	return ret
}
