package lib

import (
	"fmt"
	"runtime/debug"
)

const version = "HEAD"

func Version() string {
	rev := "unknown"
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, settings := range info.Settings {
			if settings.Key == "vcs.revision" {
				rev = fmt.Sprintf("#%s", settings.Value[:7])
				break
			}
		}
	}
	return fmt.Sprintf("%s (%s)", version, rev)
}
