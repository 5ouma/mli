package utils

import (
	"runtime/debug"
	"testing"
)

func TestVersion(t *testing.T) {
	tests := []struct {
		name string
		info *debug.BuildInfo
		ok   bool
		want string
	}{
		{name: "build info unavailable", want: "unknown"},
		{name: "development build", info: &debug.BuildInfo{Main: debug.Module{Version: "(devel)"}}, ok: true, want: "unknown"},
		{name: "release version", info: &debug.BuildInfo{Main: debug.Module{Version: "v1.2.3"}}, ok: true, want: "v1.2.3"},
		{
			name: "revision version",
			info: &debug.BuildInfo{Main: debug.Module{Version: "v1.2.3-0.20260101-abcdef123456"}},
			ok:   true,
			want: "v1.2.3 (#abcdef1)",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Version(test.info, test.ok); got != test.want {
				t.Fatalf("Version() = %q, want %q", got, test.want)
			}
		})
	}

	info, ok := debug.ReadBuildInfo()
	if Version(info, ok) == "" {
		t.Fatal("Version() returned empty string for current build info")
	}
}
