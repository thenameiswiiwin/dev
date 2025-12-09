package osdetect

import (
	"testing"
)

func TestDetect(t *testing.T) {
	info, err := Detect()
	if err != nil {
		t.Fatalf("Detect() failed: %v", err)
	}

	if info.OS == "" {
		t.Error("OS should not be empty")
	}

	if info.PackageManager == "" {
		t.Error("PackageManager should not be empty")
	}

	t.Logf("Detected OS: %s", info.OS)
	t.Logf("Package Manager: %s", info.PackageManager)
}

func TestOSString(t *testing.T) {
	tests := []struct {
		os   OS
		want string
	}{
		{MacOS, "macos"},
		{Arch, "arch"},
		{Debian, "debian"},
		{Ubuntu, "ubuntu"},
		{Linux, "linux"},
	}

	for _, tt := range tests {
		t.Run(string(tt.os), func(t *testing.T) {
			if got := tt.os.String(); got != tt.want {
				t.Errorf("OS.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCommandExists(t *testing.T) {
	// Test with a command that should exist
	if !commandExists("ls") {
		t.Error("ls command should exist")
	}

	// Test with a command that shouldn't exist
	if commandExists("this-command-definitely-does-not-exist-12345") {
		t.Error("Non-existent command should return false")
	}
}
