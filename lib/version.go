package lib

import (
	"fmt"
	"runtime/debug"
)

func Version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}
	return fmt.Sprintf("%s (#%s)", info.Main.Version, info.Main.Sum[:7])
}
