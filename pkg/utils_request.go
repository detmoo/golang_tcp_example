// Test
package pkg

import (
        "encoding/json"
        "testing"
        "time"
)


type TestCase struct {
        testTimeout time.Duration
        listenerTimeout time.Duration
        expected string
}

var tests = map[string]TestCase{
    "expect timeout": TestCase{
        testTimeout: (12 * time.Second), // greater than the listener timeout
        listenerTimeout: (4 * time.Second),
        expected: "banana",
    },
    "expect signal": TestCase{
        testTimeout: (4 * time.Second), // less than the listener timeout
        listenerTimeout: (12 * time.Second),
        expected: "grape",
    },
}


func TestDeferUserInterrupt(t *testing.T) {
	for testName, test := range tests {
		t.Logf("Running test case %s", testName)

		listener, err := net.Listen("tcp", host+":"+port)
        if err != nil {
            log.Fatal(err)
            return err

		ctx := context.Background()
		ctx, cancelCtx := context.WithTimeout(ctx, test.testTimeout)
		action := DeferCloseListener(&listener, listenerTimeout, ctx)
	    if action != expected{
			t.Errorf("Expected result: %s, but got: %s", expected, action)
		}
	}
}
