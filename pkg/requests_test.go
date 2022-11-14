// Test
package pkg

import (
        "encoding/json"
        "testing"
)


func getTests() (tests map[string]struct) {
    tests := map[string]struct {
        content string
        metadata Metadata
    }{
        "affirmative test": {
            content: "this is the request body",
            metadata: Metadata{
                time: time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
                tag: "salsa"
            }
        },
        "failing test content": {
            data: "this field is wrongly named",
            metadata: Metadata{
                time: time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
                tag: "salsa"
            }
        },
        "failing request metadata": {
            content: "this is the request body",
            metadata: Metadata{
                time: time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
                dog: "this field is named wrongly"
            }
        },
    }
    return
}



func TestParse(t *testing.T) {
	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		expectation := Message{content: test.content, metadata: test.metadata}
		jsonStr, _ := json.Marshal(expectation)
		res := parse(byte(jsonStr))
		if res != expectation{
			t.Errorf("Expected result: %s, but got: %s", expectation, res)
		}
	}
}


func TestGetResponse(t *testing.T) {
	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		input := Message{content: test.content, metadata: test.metadata}
		res := getResponse(input)
		if res.content != input.content{
			t.Errorf("Expected content: %s, but got: %s", input.content, res.content)
	    if res.metadata.tag != "mambo" {
			t.Errorf("Expected metadata tag: %s, but got: %s", "mambo", res.metadata.tag)
		}
	}
}
