package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"path/filepath"
)

func TestHandleUserInput(t *testing.T) {
	result, err := handleUserInput([]string{"programName", "testing-input-files/goodInput"})
	assert.NotNil(t, result)
	assert.NoError(t, err)
	result, err = handleUserInput([]string{"programName", "testing-input-files/goodLargeInput"})
	assert.NotNil(t, result)
	assert.NoError(t, err)

	_, err = handleUserInput([]string{"programName", "testing-input-files/badInput"})
	assert.Error(t, err)
	_, err = handleUserInput([]string{"programName"}) //no arguement
	assert.Error(t, err)
	_, err = handleUserInput([]string{"programName", "non-existant-file"})
	assert.Error(t, err)
	_, err = handleUserInput([]string{"programName", "testing-input-files/badLargeInput"})
	assert.Error(t, err)
}

func TestValidateTestCases(t *testing.T) {
	type test struct {
		numOfTestCases int
		testCases      []string
		expected       []string
	}

	tests := []test{
		{1, []string{"+++++-"}, []string{"+++++-"}},
		{4, []string{"--+-++-+", "--+--++++--++", "--+-+-+++--+", "--+-++--+-+---+"}, []string{"--+-++-+", "--+--++++--++", "--+-+-+++--+", "--+-++--+-+---+"}},
	}

	for _, test := range tests {
		result, _ := validateTestCases(test.numOfTestCases, test.testCases)
		assert.Equal(t, test.expected, result)
	}

	tests = []test{
		{1, []string{"+A+++-"}, nil},                                    // contains bad character
		{4, []string{"--+-++-+", "--+--++++--++", "--+-+-+++--+"}, nil}, //not enough test cases
		{2, []string{"--+-++-+", "--+--++++--++", "--+-+-+++--+"}, nil}, //too many test cases
		{1, []string{""}, nil},                                          //empty test case
	}

	for _, test := range tests {
		result, err := validateTestCases(test.numOfTestCases, test.testCases)
		assert.Nil(t, result, nil, "Expected nil, but got : %v", result)
		assert.NotNil(t, err, "Expected error, but didn't get one")
	}
}

func TestFlip(t *testing.T) {
	type test struct {
		pancakes string
		numFlip  int
		expected string
	}

	tests := []test{
		{"---++-", 3, "+++++-"},
		{"----", 1, "+---"},
		{"+++", 0, "+++"},
		{"-+-+++++++-----+---+------++---+-------++-", 15, "+-+-------++++++---+------++---+-------++-"},
	}

	for _, test := range tests {
		flip(test.numFlip, &test.pancakes)
		assert.Equal(t, test.expected, test.pancakes)
	}
}

func TestModelSolve(t *testing.T) {
	type test struct {
		input    string
		expected int
	}

	tests := []test{
		{"---++-", 3},
		{"----", 1},
		{"+++", 0},
		{"-+-+++++++-----+---+------++---+-------++-", 15},
	}

	for _, test := range tests {
		result := cleverSolve(test.input)
		assert.Equal(t, test.expected, result)
	}
}

func TestCleverSolve(t *testing.T) {
	type test struct {
		input    string
		expected int
	}

	tests := []test{
		{"---++-", 3},
		{"----", 1},
		{"+++", 0},
		{"-+-+++++++-----+---+------++---+-------++-", 15},
	}

	for _, test := range tests {
		result := cleverSolve(test.input)
		assert.Equal(t, test.expected, result)
	}
}

func TestSolvePancakeRevenge(t *testing.T) {
	path := filepath.FromSlash("testing-input-files/goodInput")
	os.Remove(path + ".out")
	assert.Equal(t, 0, solvePancakeRevenge([]string{"programName", string(path)}, false))
	f, err := os.Open(path + ".out")
	f.Close()
	assert.NoError(t, err)

	path = filepath.FromSlash("testing-input-files/goodLargeInput")
	os.Remove(path + ".out")
	assert.Equal(t, 0, solvePancakeRevenge([]string{"programName", string(path)}, false))
	f, err = os.Open(path + ".out")
	f.Close()
	assert.NoError(t, err)

	assert.Equal(t, 1, solvePancakeRevenge([]string{"programName", "testing-input-files/badInput"}, false))
	assert.Equal(t, 1, solvePancakeRevenge([]string{"programName"}, false)) //no arguement
	assert.Equal(t, 1, solvePancakeRevenge([]string{"programName", "non-existant-file"}, false))
}
