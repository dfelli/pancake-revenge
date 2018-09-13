package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MAX_NUM_TEST_CASES = 100
const MIN_NUM_TEST_CASES = 1
const MAX_TEST_CASE_SIZE = 100

// expects an input file with the first line containing a integer number representing the number of
// test cases that follow (one per line). Test cases involve a series of '+' or '-' (1 to 100) signs
// representing a pancake with a smiley face facing upwards or one with the smiley face facing
// downward respectively.
// returns these test cases as an array of strings
// returns an error if the files is not present, readable, or formatted incorrectly
func handleUserInput(args []string) ([]string, error) {
	if len(args) == 1 {
		return nil, fmt.Errorf("Please provide an input file containing a line with an integer" +
			" denoting the number test cases (one per line) then lines containing 1 to 100 '+' " +
			"and '-' signs. Exiting")
	}

	inputFile, err := os.Open(args[1])
	if err != nil {
		return nil, fmt.Errorf("Error on opening file : %v", err)
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	if scanner.Scan() != true {
		return nil, fmt.Errorf("Error reading first line of file : %v", err)
	}
	numberOfTestCases, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, fmt.Errorf("Error integer not given as first line in file : %v", err)
	}
	// since we know how big testCases should be we implement a capacity with 0 size
	testCases := make([]string, 0, numberOfTestCases)
	for scanner.Scan() {
		line := scanner.Text()
		testCases = append(testCases, line)
	}

	return validateTestCases(numberOfTestCases, testCases)
}

func validateTestCases(numberOfTestCases int, testCases []string) ([]string, error) {
	// checks that the number of test cases is within the specification
	if numberOfTestCases > MAX_NUM_TEST_CASES || numberOfTestCases < MIN_NUM_TEST_CASES {
		return nil, fmt.Errorf("File must provide number of test cases between %d and %d "+
			"inclusive. Provided : %d", MIN_NUM_TEST_CASES, MAX_NUM_TEST_CASES, numberOfTestCases)
	}

	// makes sure the number of test cases provided matches the number denoting the number of test cases
	if len(testCases) != numberOfTestCases {
		return nil, fmt.Errorf("Error: improper number of test cases provided. "+
			"expected : %d given :%d", numberOfTestCases, len(testCases))
	}

	// check the test cases are valid (not empty), are the right size, only only "+/-" signs
	for _, testCase := range testCases {
		// make sure the test case is not empty
		if testCase == "" {
			return nil, fmt.Errorf("Error: empty test case provided")
		}
		//check the size/length of a test case is within the specification
		if len(testCase) > MAX_TEST_CASE_SIZE {
			return nil, fmt.Errorf("Test case size should be less then %d. Provided test case"+
				" with length : %d", MAX_TEST_CASE_SIZE, len(testCase))
		}

		// check that each test case only contains '+' and '-' signs
		for _, character := range testCase {
			if character != '+' && character != '-' {
				return nil, fmt.Errorf("only '+' and '-' allowed. File "+
					"contains : %v in test cases", character)
			}
		}
	}

	return testCases, nil
}

func flip(pancakesToFlip int, pancakes *string) {
	array := []byte(*pancakes)
	for index := 0; index < pancakesToFlip; index++ {
		// switch '+' for '-' and '-' for '+'
		if array[index] == '+' {
			array[index] = '-'
		} else {
			array[index] = '+'
		}
	}
	*pancakes = string(array)
}

// solves the revenge of the pancake game through modeling and returns the minimum number
// of flips to orient all the pancakes upright. can only flip collectively a selected
// number of pancakes from the top.
func modelSolve(pancakes string) int {
	// we will take a simple approach of flipping together only the top pancakes with
	// same orientation. This will make them match the next set of things oriented in the
	// opposite orientation. continued flips in a like manner will get us to the quickest solution.
	pancakesToFlip := 0
	numOfFlips := 0
	loopCount := 0
	// to avoid an potential infinite loop, we limit the max times we loop
	for loopCount < MAX_TEST_CASE_SIZE {
		loopCount++
		for index := pancakesToFlip; index < len(pancakes); index++ {
			//if this current pacake is not oriented the same way, flip what we have before this
			if pancakes[index] != pancakes[0] {
				break
			}
			//pancake is the same as the previous, add this to be flipped
			pancakesToFlip++
		}

		if pancakesToFlip == len(pancakes) {
			break
		}

		numOfFlips++
		flip(pancakesToFlip, &pancakes)
	}

	// perform one last flip if all the pancakes are downward
	if pancakes[0] == '-' {
		numOfFlips++
	}

	return numOfFlips
}

// Alternate solver to the revenge of the pancake game that doesn't involve modeling.
// Returns the minimum number of flips to orient all the pancakes upright. can only
// flip collectively a selected number of pancakes from the top.
func cleverSolve(pancakes string) int {
	// we will take a clever approach of counting the times that sequences of '+' and '-' signs
	// switch. We can always flip the beginning common symbols sequence to match next symbol
	// If the last character is '-' we add 1 because we need to flip the entire sequence once more.
	numOfFlips := 0
	//counts the number of switches or alteranting pancake sides
	for index := 1; index < len(pancakes); index++ {
		if pancakes[index] != pancakes[index-1] {
			numOfFlips++
		}
	}

	// add 1 if the last pancake is facing down
	if pancakes[len(pancakes)-1] == '-' {
		numOfFlips++
	}

	return numOfFlips
}

func solvePancakeRevenge(args []string, printError bool) int {
	// handles what the user gives as input
	// gets an array of test cases each representing a stack of pancakes
	// and their upward/downwards orientation.
	testCases, err := handleUserInput(args)
	if err != nil {
		if printError {
			fmt.Printf("%v\n", err)
		}
		return 1
	}

	// open a file to output the results
	outFile, err := os.Create(args[1] + ".out")
	if err != nil {
		fmt.Errorf("Error on opening output file : %v", err)
		return 1
	}
	defer outFile.Close()

	// for each test case solve for the minium number of flips to orient all pancakes upright
	// write this solution formatted to a file
	for caseNumber, testCase := range testCases {
		// currently we are not using a modeling approach. but if we wanted to we could use this.
		// outFile.WriteString(fmt.Sprintf("Case #%d: %d\n", caseNumber+1, modelSolve(testCase)))
		outFile.WriteString(fmt.Sprintf("Case #%d: %d\n", caseNumber+1, cleverSolve(testCase)))
	}
	return 0
}
func main() {
	solvePancakeRevenge(os.Args, true)
}
