package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

/*
This function just handles the submission of http requests and returns
the response as a string. It also checks that a 200 was returned by the server.
Figured it was easier to make this a dedicated function since I'll be using it
several times below.
*/
func retrieveHTTPResponse(t *testing.T, handlerFunc http.HandlerFunc, path string) string {
	request, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Errorf("Failed to create new HTTP request: %v", err)
	}
	recorder := httptest.NewRecorder()
	root := http.HandlerFunc(handlerFunc)
	root.ServeHTTP(recorder, request)
	response, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
	}
	if recorder.Code != http.StatusOK {
		t.Errorf("Request failed to return a 200 error code. Got %d", recorder.Code)
	}
	return string(response)
}

/*
Test the calculateFibonacci function with various inputs
First, ensure two seed values are returning properly, then check for a low, known value.
We're not error checking for high numbers here since we only handle that case while processing
the HTTP request.
*/
func TestFib(t *testing.T) {
	var tests = []struct {
		inputNumber  int
		outputValues []int
	}{
		{1, []int{0}},
		{2, []int{0, 1}},
		{5, []int{0, 1, 1, 2, 3}},
	}

	for i, test := range tests {
		t.Logf("Running TestFib test #%d", i)
		fibResult := calculateFibonacci(test.inputNumber)
		//We're making use of deepequal to check these values. A possible speedup could be to
		//iterate of the slices ourselves. Appears that deepequal is slow, but this is just a
		//test anyways.
		if !reflect.DeepEqual(fibResult, test.outputValues) {
			t.Errorf("Values returned from calculateFibonacci were not correct. Return value was %v, expected value was %v", fibResult, test.outputValues)
		}
	}
}

/*
Test the webserver we created to handle fibonacci requests. We need to test:
 - Handling of a random path, to make sure it returns an "incorrect path" error
 - Handling of negative numbers passed to /fib
 - Handling of no COUNT input passed to /fib
 - Handling of values that would overflow 64bit ints
 - Handling of a normal, valid request
*/
func TestIncorrectPath(t *testing.T) {

	response := retrieveHTTPResponse(t, rootPath, "/foo")

	if response != "ERROR: Path not recognized. Use /fib.\n" {
		t.Errorf("Incorrect response to /foo path request. Got %s", response)
	}
}

func TestNegativesAndZero(t *testing.T) {
	response := retrieveHTTPResponse(t, fibPath, "/fib?COUNT=-1")
	if response != "ERROR: Count must be > 0\n" {
		t.Errorf("Incorrect response to negative value. Got %s", response)
	}

	response = retrieveHTTPResponse(t, fibPath, "/fib?COUNT=0")
	if response != "ERROR: Count must be > 0\n" {
		t.Errorf("Incorrect response to zero value. Got %s", response)
	}
}

func TestNoCountInput(t *testing.T) {
	response := retrieveHTTPResponse(t, fibPath, "/fib")
	if response != "ERROR: Unable to retrieve COUNT input. Please try again.\n" {
		t.Errorf("Failed to handle no COUNT input. Got %s", response)
	}
}

func TestOverflow(t *testing.T) {
	response := retrieveHTTPResponse(t, fibPath, "/fib?COUNT=94")
	if response != "ERROR: This calculation would cause a 64bit integer overflow. Choose a value <= 93\n" {
		t.Errorf("Failed to handle 64 bit integer overflow. Got %s", response)
	}
}

func TestNormalInput(t *testing.T) {
	response := retrieveHTTPResponse(t, fibPath, "/fib?COUNT=5")
	if response != "[0 1 1 2 3]\n" {
		t.Errorf("Failed to return normal output for sequence of n=5. Got %s", response)
	}
}
