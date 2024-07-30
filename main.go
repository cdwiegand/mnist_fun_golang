package main

import (
	"fmt"
	"os"
)

func main() {
	config := NewRuntimeConfig(os.Args[1:])
	trainer := NewTrainer(config)
	source := SourceData{}
	source.LoadTrainingSource(config.TrainingPath, config)
	chains := trainer.BuildTrainingLayers(source)

	fmt.Printf("Loaded %d model data, now analysing...", len(source.List))
	trainer.RunTraining(chains, source)
}
