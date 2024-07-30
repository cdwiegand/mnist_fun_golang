package main

type Model struct {
	Matrices      []InnerLayerMatrix
	Layers        []*Layer
	TrainingLoops []TrainingLoopResult
}

func NewModel() Model {
	return Model{}
}

func (m *Model) GetOutput() *Layer {
	return m.Layers[len(m.Layers)-1]
}

func (m *Model) Add(layer *Layer) {
	m.Layers = append(m.Layers, layer)
	if len(m.Layers) >= 2 {
		m.Matrices = append(m.Matrices, NewInnerLayerMatrix(m.Layers[len(m.Layers)-2], layer))
	}
}

func (m *Model) SetInputNeurons(input []float64) {
	m.Layers[0].Neurons = input
}

func (m *Model) ApplyMatricesForward() {
	for _, m := range m.Matrices {
		m.ApplyMatricesForward()
	}
}

func (m *Model) GenerateOutputDelta(properMatchingOutputNeuron int) []float64 {
	output := m.Layers[len(m.Layers)-1]
	deltaOutput := make([]float64, output.GetLength())

	for i := 0; i < output.GetLength(); i++ {
		if i == properMatchingOutputNeuron {
			deltaOutput[i] = output.Neurons[i] - 1
		} else {
			deltaOutput[i] = output.Neurons[i]
		}
	}
	return deltaOutput
}

func (m *Model) ApplyOutputDelta(deltaOutput []float64) {
	hidden1 := m.Layers[len(m.Layers)-2]
	hidden1ToOutput := m.Matrices[len(m.Matrices)-1]

	// backpropagate from output layer to hidden layer 1 learning adjustments
	outputToHiddenLayer1Correction := hidden1.CalculateBackpropagationMatrix(deltaOutput, -0.01)
	hidden1ToOutput.ApplyBackPropagationMatrix(outputToHiddenLayer1Correction)
}

func (m *Model) BackpropagateDelta(deltaOutput []float64) {
	for i := len(m.Matrices) - 1; i > 0; i-- {

		backSourceMatrix := m.Matrices[i]
		backDestMatrix := m.Matrices[i-1]

		// harder math, can't cheat
		deltaOutput = backSourceMatrix.CalculateBackpropagationMatrixDelta(deltaOutput)
		// ok, now adjust input weights
		corrections := backDestMatrix.FromLayer.CalculateBackpropagationMatrix(deltaOutput, -0.01)
		backDestMatrix.ApplyBackPropagationMatrix(corrections)
	}
}
