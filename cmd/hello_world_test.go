package cmd

import (
	"bytes"
	"io"
	"testing"
)


func TestHelloWorldCommand(t *testing.T) {
	rootCmd := NewRootCmd()
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"scoop", "--output-case", "upper"})
	rootCmd.Execute()
	out, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "SCOOP" {
		t.Fatalf("expected \"%s\" got \"%s\"", "SCOOP", string(out))
	}
}
