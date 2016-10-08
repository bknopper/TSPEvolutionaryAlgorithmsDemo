package main

import (
	"fmt"
)

type AlgorithmRunner struct {
	Algorithm Algorithm
}

// TODO implement multi threading here

// Start using a pointer so we can run a go routine on the actual travelr
func startAlgorithmRunner(algorithmOptions AlgorithmOptions) *AlgorithmRunner {

	// Create new traveller and set to running
	// TODO Should accept the config from the client side
	// TODO Refactor
	fmt.Println("Starting Algorithm")
  	algorithm := NewAlgorithm(algorithmOptions)
	traveler := AlgorithmRunner{}
	fmt.Println("start")
	traveler.Algorithm = algorithm
	traveler.Algorithm.start()
	return &traveler

}

func (traveler *AlgorithmRunner) stop() {
	traveler.Algorithm.stop()
}

func (traveler *AlgorithmRunner) running() bool {
	return traveler.Algorithm.running
}

func (traveler *AlgorithmRunner) getCurrentBest() CandidateSolution {
	return traveler.Algorithm.getCurrentBest()
}
