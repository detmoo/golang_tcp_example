// Test
package pkg

import (
        "encoding/json"
        "testing"
        "time"
)


type requestsTestCase struct {
        requestContent string
        requestTag string
        expectedResponseContent string
        expectedResponseTag string
}


var requestsTests = map[string]requestsTestCase{
    "affirmative test": requestsTestCase{
        requestContent: "this is the request body",
        requestTag: "untagged-tcp-endpoint",
        expectedResponseContent: "TCP listener received: this is the request body",
        expectedResponseTag: "untagged-tcp-endpoint",
    },
}


func TestParse(t *testing.T) {
	for testName, test := range requestsTests {
		t.Logf("Running test case %s", testName)
		expectation := new(Message)
		expectation.compose(test.expectedResponseContent, test.expectedResponseTag)
		jsonStr, _ := json.Marshal(expectation)
		result := new(Message)
		result.parse(jsonStr)
		if result != expectation{
			t.Errorf("Expected result: %s, but got: %s", expectation, result)
		}
	}
}


func TestGetResponse(t *testing.T) {
	for testName, test := range requestsTests {
		t.Logf("Running test case %s", testName)
	    input := new(Message)
		input.compose(test.requestContent, test.requestTag)

		res := getResponse(input.Content)
		if res.Content != "TCP listener received: "+input.Content {
			t.Errorf("Expected content: %s, but got: %s", "TCP listener received Message.Content: "+input.Content, res.Content)
		}
	    if res.Metadata.Tag != "untagged-tcp-server" {
			t.Errorf("Expected metadata tag: %s, but got: %s", input.Metadata.Tag, res.Metadata.Tag)
		}
	}
}
