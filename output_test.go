package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"ascii-art/src"
)

func runMainAndGetStdout(t *testing.T, args ...string) []byte {
	t.Helper()

	cmd := exec.Command(os.Args[0], "-test.run=TestActualMain")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")

	out, err := cmd.Output()
	if err != nil {
		// If it exited non-zero, include stderr in message via CombinedOutput
		combined, _ := cmd.CombinedOutput()
		t.Fatalf("process failed: %v\noutput:\n%s", err, combined)
	}
	return out
}

func TestOutput_NewlineOnly(t *testing.T) {
	golden, err := os.ReadFile("testdata/newline.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "\\n")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}

func TestOutput_HelloDoubleNewline(t *testing.T) {
	golden, err := os.ReadFile("testdata/hello_double_newline.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "Hello\\n\\nThere")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}
func TestOutput_EmptyInput(t *testing.T) {
	golden, err := os.ReadFile("testdata/empty.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%q\n--- want (golden) ---\n%q", got, golden)
	}
}
func TestOutput_HelloTrailingNewline(t *testing.T) {
	golden, err := os.ReadFile("testdata/hello_trailing_newline.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "Hello\\n")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}
func TestOutput_HelloThere(t *testing.T) {
	golden, err := os.ReadFile("testdata/hello_there.golden")
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	got := runMainAndGetStdout(t, "Hello There")

	if !bytes.Equal(got, golden) {
		t.Fatalf("stdout mismatch\n--- got ---\n%s\n--- want (golden) ---\n%s", got, golden)
	}
}

func TestOutput_WriteToFile(t *testing.T) {
	filePath := filepath.Join(t.TempDir(), "banner.txt")
	expected := runMainAndGetStdout(t, "hello", "standard")

	cmd := exec.Command(os.Args[0], "-test.run=TestActualMain")
	cmd.Env = append(
		os.Environ(),
		"GO_WANT_HELPER_PROCESS=1",
		"ASCII_ART_TEST_ARGS=--output="+filePath+"\x1fhello\x1fstandard",
	)

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("expected success, got error: %v, output: %s", err, out)
	}

	gotFile, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}
	if !bytes.Equal(gotFile, expected) {
		t.Fatalf("output file mismatch\n--- got ---\n%s\n--- want ---\n%s", gotFile, expected)
	}
}

func TestOutput_InvalidOutputFormat_ShowsUsage(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-test.run=TestActualMain")
	cmd.Env = append(
		os.Environ(),
		"GO_WANT_HELPER_PROCESS=1",
		"ASCII_ART_TEST_ARGS=--output\x1fbanner.txt\x1fhello\x1fstandard",
	)

	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("expected non-zero exit, got success. output: %s", out)
	}
	if !strings.Contains(string(out), src.OutputUsage) {
		t.Fatalf("expected usage message, got: %s", out)
	}
}
