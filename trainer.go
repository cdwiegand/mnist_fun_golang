package main

import (
	"fmt"
	"time"
)

type Trainer struct {
	Config RuntimeConfig
}

func NewTrainer(config RuntimeConfig) Trainer {
	return Trainer{
		Config: config,
	}
}

func (t *Trainer) BuildTrainingLayers(sourceData SourceData) Model {
	maxCountPixels := 0
	for _, j := range sourceData.List {
		if maxCountPixels < len(j.Pixels) {
			maxCountPixels = len(j.Pixels)
		}
	}
	// should be 784

	// now, run through ML..
	model := NewModel()
	model.Add(NewLayer(maxCountPixels))

	model.Add(NewLayer(80))
	model.Add(NewLayer(28))

	distinctValues := make([]byte, 0)
	for _, j := range sourceData.List {
		foundJ := false
		for _, j2 := range distinctValues {
			if distinctValues[j2] == j.Character {
				foundJ = true
			}
		}
		if !foundJ {
			distinctValues = append(distinctValues, j.Character)
		}
	}
	model.Add(NewLayer(len(distinctValues))) // 10 if mnist if MNIST

	return model
}

func (t *Trainer) RunTraining(model Model, sourceData SourceData) {
	epoch := len(model.TrainingLoops)

	for epoch < t.Config.Loops {

		res := NewTrainingLoopResult(epoch) //  StartTime = DateTime.UtcNow

		// FIXME: random ordering please!
		for _, item := range sourceData.List {
			properMatchingOutputNeuron := t.Config.VectorizeKey(item.Character)
			model.SetInputNeurons(item.Pixels)
			model.ApplyMatricesForward()

			// ok, figure out error cost
			deltaOutput := model.GenerateOutputDelta(properMatchingOutputNeuron)
			model.ApplyOutputDelta(deltaOutput)
			model.BackpropagateDelta(deltaOutput)

			// ok, did the output match?
			matchedOutputNeuron := model.GetOutput().FindHighestValueOutputNeuron()
			matchedOutputChar := t.Config.DevectorizeKey(matchedOutputNeuron)
			if matchedOutputChar == item.Character {
				res.CountRight(item.Character)
			} else {
				res.CountWrong(item.Character)
			}
		}

		res.EndTime = time.Now().UTC().Unix()
		model.TrainingLoops = append(model.TrainingLoops, res)
		fmt.Println(res.ToString())
		epoch++
	}
}
