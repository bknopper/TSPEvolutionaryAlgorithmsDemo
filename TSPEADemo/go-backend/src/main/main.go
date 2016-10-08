package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

// TODO investigate project structure
// TODO investigate dependency management (no go get?)
// TODO investigate error handling and logging
// TODO investigate unit testing
// TODO investigate naming conventions (methods lower?)

// TODO implement runner using Channels : http://guzalexander.com/2013/12/06/golang-channels-tutorial.html

var algorithmRunner *AlgorithmRunner

type AlgorithmOptions struct {
	MutationProbability int
	PopulationSize int
	NrOfGenerations int
	FitnessThreshold float64
	ParentSelectionSize int
	ParentPoolSize int
	AlgorithmStyle string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/cities", ListCities)

	// TODO reconsider api structure when all works
	/* TODO suggest /api/travel
		POST will start traveling
		GET will return current and finished (true|false) until finished
		DELETE will stop
	 */

	router.HandleFunc("/api/currentBest", CurrentBest)
	router.HandleFunc("/api/latestBest", LatestBest)
	router.HandleFunc("/api/stillRunning", StillRunning)
	router.HandleFunc("/api/startAlgorithm", StartAlgorithm)
	router.HandleFunc("/api/stopAlgorithm", StopAlgorithm)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/app")))

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}



func ListCities(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	cities := getAllCities()
	json.NewEncoder(response).Encode(cities)
}

func CurrentBest(response http.ResponseWriter, request *http.Request) {
	if(algorithmRunner.running()) {
	    json.NewEncoder(response).Encode(algorithmRunner.getCurrentBest())
	} else {
	    http.Error(response, "No Algorithm running at this point in time. Please start one.", http.StatusInternalServerError)
	}
}

func LatestBest(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(algorithmRunner.getCurrentBest())
}
func StillRunning(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(algorithmRunner != nil && algorithmRunner.running())
}

func StartAlgorithm(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var algorithmOptions AlgorithmOptions
	decoder.Decode(&algorithmOptions)
	if (algorithmRunner != nil && algorithmRunner.running()) {
		log.Println("Trying to start an running Traveler, skipping")
	} else {
		algorithmRunner = startAlgorithmRunner(algorithmOptions)
	}
}

func StopAlgorithm(response http.ResponseWriter, request *http.Request) {
	if (algorithmRunner == nil || !algorithmRunner.running()) {
		log.Println("Trying to stop an non-running Traveler, skipping")
	} else {
		algorithmRunner.stop()
	}
}
