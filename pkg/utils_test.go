// Test
package pkg

import (
        "context"
        "fmt"
        "log"
        "net"
        "testing"
        "time"
)


type utilsTestCase struct {
        testTimeout time.Duration
        listenerTimeout time.Duration
        host string
        port string
        expected string
}

var utilsTests = map[string]utilsTestCase{
    "expect timeout": utilsTestCase{
        testTimeout: (12 * time.Second), // greater than the listener timeout
        listenerTimeout: (4 * time.Second),
        host: "localhost",
        port: "9002",
        expected: "timeout",
    },
    "expect signal": utilsTestCase{
        testTimeout: (4 * time.Second), // less than the listener timeout
        listenerTimeout: (12 * time.Second),
        host: "localhost",
        port: "9003",
        expected: "interrupted",
    },
}


func TestDeferUserInterrupt(t *testing.T) {
	for testName, test := range utilsTests {
		t.Logf("Running test case %s", testName)

		listener, err := net.Listen("tcp", test.host+":"+test.port)
        if err != nil {
            log.Fatal(err)
            fmt.Println(err)
        }
        closureChannel := make(chan error)
		ctx := context.Background()
		ctx, _ = context.WithTimeout(ctx, test.testTimeout)
		clsErr := DeferCloseListener(listener, test.listenerTimeout, closureChannel, ctx)
	    if clsErr.Reason != test.expected{
			t.Errorf("Expected result: %s, but got: %s", test.expected, err.Reason)
		}
	}
}
