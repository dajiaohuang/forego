package main

import (
	"io"
	"path/filepath"
	"strings"
	"testing"
)

func TestProcessWorkingDirectory(t *testing.T) {
	dir := t.TempDir()
	process := NewProcess(dir, "cd", Env{}, false)
	stdout, err := process.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := process.Start(); err != nil {
		t.Fatal(err)
	}
	output, err := io.ReadAll(stdout)
	if err != nil {
		t.Fatal(err)
	}
	if err := process.Wait(); err != nil {
		t.Fatal(err)
	}
	if got := strings.TrimSpace(string(output)); !strings.EqualFold(filepath.Clean(got), filepath.Clean(dir)) {
		t.Fatalf("working directory = %q, want %q", got, dir)
	}
}
