package main

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

//Calculates a sequence of fibonacci numbers and
//returns a slice of integers containing the sequence.
func calculateFibonacci(n int) []int {
	//We must set the first two seed values
	slice := []int{0, 1}

	//If we're calculating more than 2 vals, perform the calculation.
	//Otherwise, simply return a subset of the first two seeds.
	if n >= 2 {
		for i := 2; i < n; i++ {
			newVal := slice[i-1] + slice[i-2]
			slice = append(slice, newVal)
		}
	} else {
		return slice[:n]
	}
	return slice
}

//This handles any requests that aren't to /fib
//It simply responds back with an incorrect usage error.
func rootPath(w http.ResponseWriter, r *http.Request) {
	log.Info("Received request to non-supported path " + r.URL.Path)
	fmt.Fprintf(w, "ERROR: Path not recognized. Use /fib.\n")
}

//Handles the request for a fibonacci sequence and calls for calculation.
func fibPath(w http.ResponseWriter, r *http.Request) {
	urlValues := r.URL.Query()
	numCalculations, err := strconv.Atoi(urlValues.Get("COUNT"))
	if err != nil {
		log.Info("Unable to retrieve COUNT input. Error is : %v", err)
		fmt.Fprintf(w, "ERROR: Unable to retrieve COUNT input. Please try again.\n")
	} else {
		if numCalculations <= 0 {
			log.Info("Count input was a negative or zero. Returning error.")
			fmt.Fprintf(w, "ERROR: Count must be > 0\n")
		} else if numCalculations > 93 {
			log.Info("This calculation for %d would cause 64bit integer overflow. Returning error.", numCalculations)
			fmt.Fprintf(w, "ERROR: This calculation would cause a 64bit integer overflow. Choose a value <= 93\n")
		} else {
			returnSlice := calculateFibonacci(numCalculations)
			fmt.Fprintf(w, "%v\n", returnSlice)
		}
	}
}

//Start webserver and listen.
func main() {
	log.Info("Starting web server on port 8080")
	http.HandleFunc("/", rootPath)
	http.HandleFunc("/fib", fibPath)
	http.ListenAndServe(":8080", nil)
}
