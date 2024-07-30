package main

type InnerLayerMatrix struct {
	FromLayer *Layer
	ToLayer   *Layer
	Matrix    [][]float64
}

func MakeMultidimensionalFloat64(sizeA int, sizeB int) [][]float64 {
	ret := make([][]float64, sizeA)
	for i := 0; i < sizeA; i++ {
		ret[i] = make([]float64, sizeB)
	}
	return ret
}

func NewInnerLayerMatrix(fromLayer *Layer, toLayer *Layer) InnerLayerMatrix {
	return InnerLayerMatrix{
		FromLayer: fromLayer,
		ToLayer:   toLayer,
		Matrix:    MakeMultidimensionalFloat64(fromLayer.GetLength(), toLayer.GetLength()),
	}
}

func (matrix *InnerLayerMatrix) ApplyMatricesForward() {
	// weight matrix 1st dimension is left/input layer (rows), 2nd dimension is right/output layer (column)
	// (note: opposite of youtube video I'm watching!)
	columns := len(matrix.Matrix[0])
	matrix.ToLayer.Reset()
	for resultIdx := 0; resultIdx < columns; resultIdx++ {
		var h_pre float64
		for pix := 0; pix < len(matrix.Matrix); pix++ {
			// multiply them by the input -> hidden layer 1, then apply sigmoid function to that
			h_pre += matrix.FromLayer.Neurons[pix] * matrix.Matrix[pix][resultIdx]
		}
		matrix.ToLayer.Neurons[resultIdx] = Sigmoid(h_pre)
	}
}

func (matrix *InnerLayerMatrix) CalculateBackpropagationMatrixDelta(deltaOutput []float64) (ret []float64) {
	deltaHiddenLayer1a := make([]float64, matrix.FromLayer.GetLength())
	for i := 0; i < matrix.FromLayer.GetLength(); i++ {
		deltaHiddenLayer1a[i] = (matrix.FromLayer.Neurons[i] * (1 - matrix.FromLayer.Neurons[i]))
	}
	deltaHiddenLayer1b := MatrixMultiply(matrix.Matrix, deltaOutput)
	deltaHiddenLayer1c := make([]float64, matrix.FromLayer.GetLength())
	for i := 0; i < len(deltaHiddenLayer1c); i++ {
		deltaHiddenLayer1c[i] = deltaHiddenLayer1b[i] * deltaHiddenLayer1a[i]
	}
	return deltaHiddenLayer1c
}

func (matrix *InnerLayerMatrix) ApplyBackPropagationMatrix(correctionMatrix [][]float64) {
	// should be the same shape (dimensions) so easy!
	for i := 0; i < len(correctionMatrix); i++ {
		for j := 0; j < len(correctionMatrix[0]); j++ {
			matrix.Matrix[i][j] += correctionMatrix[i][j]
		}
	}
}
