package main

type Layer struct {
	Neurons []float64
}

func NewLayer(size int) *Layer {
	ret := new(Layer)
	ret.SetLength(size)
	return ret
}
func (l *Layer) GetLength() int {
	return len(l.Neurons)
}
func (l *Layer) SetLength(size int) {
	l.Neurons = make([]float64, size)
}

func (l *Layer) CalculateBackpropagationMatrix(errorCost []float64, learnRate float64) (ret [][]float64) {
	if learnRate > 0 {
		learnRate = 0 - learnRate // needs to be negative
	}
	countNeurons := l.GetLength()
	ret = make([][]float64, countNeurons, len(errorCost))

	for neuron := 0; neuron < countNeurons; neuron++ {
		for errorIdx := 0; errorIdx < len(errorCost); errorIdx++ {
			ret[neuron][errorIdx] = l.Neurons[neuron] * errorCost[errorIdx] * learnRate // learn rate should be like -0.01
		}
	}

	return ret
}
func (l *Layer) Reset() {
	for idx := 0; idx < len(l.Neurons); idx++ {
		l.Neurons[idx] = 0
	}
}
func (l *Layer) FindHighestValueOutputNeuron() int {
	highestValue := l.Neurons[0]
	for _, mc := range l.Neurons {
		if highestValue < mc {
			highestValue = mc
		}
	}

	idx := 0
	for l.Neurons[idx] < highestValue {
		idx++ // have to iterate
	}
	return idx
}
