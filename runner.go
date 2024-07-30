package main

import "fmt"

func Run(sourceData SourceData, model Model, config RuntimeConfig) {
	for _, item := range sourceData.List {
		config.VectorizeKey(item.Character)
		model.SetInputNeurons(item.Pixels)
		model.ApplyMatricesForward()

		// ok, did the output match?
		matchedOutputNeuron := model.GetOutput().FindHighestValueOutputNeuron()
		matchedOutputChar := config.DevectorizeKey(matchedOutputNeuron)

		fmt.Printf("%c", matchedOutputChar)
	}
}
