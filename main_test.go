package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestMainProcess(t *testing.T) {

	tests := []struct {
		name     string
		args     []string
		wantExit int
	}{
		{
			name:     "empty arguments",
			args:     []string{},
			wantExit: 1,
		},
		{
			name:     "too many arguments",
			args:     []string{"sub string", "input string", "something"},
			wantExit: 1,
		},
		{
			name:     "Successful execution",
			args:     []string{"hello"},
			wantExit: 0,
		},
		{
			name:     "Successful execution",
			args:     []string{"l", "hello\nworld"},
			wantExit: 0,
		},
		{
			name:     "invalid data",
			args:     []string{"invalid\tdata"},
			wantExit: 1,
		},
		{
			name:     "valid input with new line",
			args:     []string{"A\nB\nC"},
			wantExit: 0,
		},
		{
			name:     "not ascii",
			args:     []string{"ψάρι"},
			wantExit: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This trick tells the test to run the 'main' function
			// within this specific test binary.
			cmd := exec.Command(os.Args[0], "-test.run=TestActualMain")

			// We pass the file arguments to the command
			cmd.Args = append(cmd.Args, tt.args...)

			// We use an environment variable to tell the test
			// to actually execute the real logic.
			cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")

			out, err := cmd.CombinedOutput()
			if tt.wantExit == 0 && err != nil {
				t.Fatalf("expected success, got error: %v, output: %s", err, out)
			}

			if tt.wantExit != 0 {
				if err == nil {
					t.Fatalf("expected exit code %d, but process succeeded", tt.wantExit)
				}
				if e, ok := err.(*exec.ExitError); ok {
					if e.ExitCode() != tt.wantExit {
						t.Errorf("got exit code %d, want %d. Output: %s", e.ExitCode(), tt.wantExit, out)
					}
				}
			}
		})
	}
}

// Helper process that runs the real main() for integration tests.
func TestActualMain(t *testing.T) {
	// Run only when invoked as helper process
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	// Execute CLI logic
	main()

	// Exit immediately to avoid printing "PASS" to stdout
	os.Exit(0)
}
