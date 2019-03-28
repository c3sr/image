package types

type ResizeAlgorithm int

const (
	ResizeAlgorithmLinear ResizeAlgorithm = iota
	ResizeAlgorithmBilinear
	ResizeAlgorithmNearestNeighbor
)
