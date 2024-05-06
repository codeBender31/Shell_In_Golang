package builtins

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	sampleTests := []struct {
		name          string
		args          []string
		expectedOut   string
		expectedError bool
	}{
		{name: "zero args",
			args:          []string{},
			expectedOut:   "\n",
			expectedError: false,
		},
		{
			name:          "Only one argument",
			args:          []string{"hello"},
			expectedOut:   "hello\n",
			expectedError: false,
		},
		{
			name:          "Multiple arguments",
			args:          []string{"hello", "world"},
			expectedOut:   "hello world\n",
			expectedError: false,
		},
	}

	//Iterate through each test
	for _, testCase := range sampleTests {
		t.Run(testCase.name, func(t *testing.T) {
			var outPut bytes.Buffer
			err := Echo(&outPut, testCase.args...)
			if (err != nil) != testCase.expectedError {
				t.Errorf("%s: unexpected error status: %v, expected error: %v", testCase.name, err, testCase.expectedError)
			}
			//Checking the output
			if recievedOutPut := outPut.String(); recievedOutPut != testCase.expectedOut {
				t.Errorf("%s: expected output %q, recieved %q", testCase.name, testCase.expectedOut, &recievedOutPut)
			}
		})
	}
}
