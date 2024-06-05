package vdf

import (
	generated "source.quilibrium.com/quilibrium/cermonyclient/vdf/generated/vdf"
)

//go:generate ./generate.sh

const intSizeBits = uint16(2048)

// WesolowskiSolve Solve and prove with the Wesolowski VDF using the given parameters.
// Outputs the concatenated solution and proof (in this order).
func WesolowskiSolve(challenge []uint8, difficulty uint64) []uint8 {
	return generated.WesolowskiSolve(intSizeBits, challenge, difficulty)
}

// WesolowskiVerify Verify with the Wesolowski VDF using the given parameters.
// `allegedSolution` is the output of `WesolowskiSolve`.
func WesolowskiVerify(challenge []uint8, difficulty uint64, allegedSolution []uint8) bool {
	return generated.WesolowskiVerify(intSizeBits, challenge, difficulty, allegedSolution)
}
