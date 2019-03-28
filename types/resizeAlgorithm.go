package types

type ResizeAlgorithm int

const (
	ResizeAlgorithmLinear ResizeAlgorithm = iota
	ResizeAlgorithmLinearASM
	ResizeAlgorithmHermite
	ResizeAlgorithmNearestNeighbor
	ResizeAlgorithmBiLinear    = ResizeAlgorithmLinear
	ResizeAlgorithmBiLinearASM = ResizeAlgorithmLinearASM
)
