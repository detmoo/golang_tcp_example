// Test
package pkg

import (
        "testing"
)


// TODO: var  nilerror *Error = nil


// HelloInput returns the input string in the preferred case
func TestHelloInput(t *testing.T) {
	tests := map[string]struct {
		input string
		outputCase string
		expectation string
	}{
		"successful uppercase": {input: "scoop", outputCase: "Upper", expectation: "SCOOP"},
		"successful lowercase": {input: "sCoOp", outputCase: "Lower", expectation: "scoop"},
		"error scenario": {input: "scoop", outputCase: "banana", expectation: ""},
	}
	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		res, _ := HelloInput(test.input, test.outputCase)
		if res != test.expectation{
			t.Errorf("Expected result: %s, but got: %s", test.expectation, res)
		}
	}
}


