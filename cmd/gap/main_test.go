package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	cmd := newRoot()
	var out bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&out)
	cmd.SetArgs([]string{"--help"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("help: %v", err)
	}
	got := out.String()
	for _, want := range []string{"gap", "Usage:"} {
		if !strings.Contains(got, want) {
			t.Fatalf("help missing %q\n%s", want, got)
		}
	}
}
