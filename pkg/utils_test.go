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
        expected: "banana",
    },
    "expect signal": utilsTestCase{
        testTimeout: (4 * time.Second), // less than the listener timeout
        listenerTimeout: (12 * time.Second),
        host: "localhost",
        port: "9003",
        expected: "grape",
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
		ctx := context.Background()
		ctx, _ = context.WithTimeout(ctx, test.testTimeout)
		action := DeferCloseListener(listener, test.listenerTimeout, ctx)
	    if action != test.expected{
			t.Errorf("Expected result: %s, but got: %s", test.expected, action)
		}
	}
}
