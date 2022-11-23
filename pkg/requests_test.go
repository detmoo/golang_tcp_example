// Test
package pkg

import (
        "encoding/json"
        "testing"
        "time"
)


type requestsTestCase struct {
        Content string
        Metadata MetadataSchema
        expectedTag string
}

var requestsTests = map[string]requestsTestCase{
    "affirmative test": requestsTestCase{
        Content: "this is the request body",
        Metadata: MetadataSchema{
            Timestamp: time.Now().Format("Monday, 02-Jan-06 15:04:05 MST"),
            Tag: "salsa",
            },
        expectedTag: "mambo",
    },
}


func TestParse(t *testing.T) {
	for testName, test := range requestsTests {
		t.Logf("Running test case %s", testName)
		expectation := Message(test)
		jsonStr, _ := json.Marshal(expectation)
		msg := new(Message)
		msg.parse(jsonStr)
		if *msg != expectation{
			t.Errorf("Expected result: %s, but got: %s", expectation, msg)
		}
	}
}


func TestGetResponse(t *testing.T) {
	for testName, test := range requestsTests {
		t.Logf("Running test case %s", testName)
		input := Message(test)
		res := getResponse(&input)
		if res.Content != input.Content{
			t.Errorf("Expected content: %s, but got: %s", input.Content, res.Content)
		}
	    if res.Metadata.Tag != test.expectedTag {
			t.Errorf("Expected metadata tag: %s, but got: %s", test.expectedTag, res.Metadata.Tag)
		}
	}
}
