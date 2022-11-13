// Desc: TCP echo server and client
package pkg

import (
	"fmt"
	"strings"
)

// HelloInput returns the input string in the preferred case
func HelloInput(input string, outputCase string) (res string, err error) {
	res, err = mySetCase(input, outputCase)
	return res, err
}

func mySetCase(input string, outputCase string) (res string, err error) {
	switch strings.ToUpper(outputCase) {
	case "UPPER":
		res = strings.ToUpper(input)
	case "LOWER":
		res = strings.ToLower(input)
	default:
		err = &argError{outputCase, "unrecognised outputCase parameter"}
	}
	return res, err
}

type argError struct {
	arg  string
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%s - %s", e.arg, e.prob)
}


