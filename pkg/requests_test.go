// Test
package pkg

import (
        "encoding/json"
        "testing"
        "time"
)


type TestCase struct {
        Content string
        Metadata MetadataSchema
}

var tests = map[string]TestCase{
    "affirmative test A": TestCase{
        Content: "this is the request body",
        Metadata: MetadataSchema{
            Timestamp: time.Now().Format("Monday, 02-Jan-06 15:04:05 MST"),
            Tag: "salsa",
            },
    },
    "affirmative test B": TestCase{
        Content: "this is the request body",
        Metadata: MetadataSchema{
            Timestamp: time.Now().Format("Monday, 02-Jan-06 15:04:05 MST"),
            Tag: "salsa",
            },
    },
}


func TestParse(t *testing.T) {
	for testName, test := range tests {
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
	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		input := Message(test)
		res := getResponse(input)
		if res.content != input.content{
			t.Errorf("Expected content: %s, but got: %s", input.content, res.content)
		}
	    if res.metadata.tag != "mambo" {
			t.Errorf("Expected metadata tag: %s, but got: %s", "mambo", res.metadata.tag)
		}
	}
}
