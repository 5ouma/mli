package lib

import (
	"fmt"
	"runtime/debug"
)

var version = "HEAD"

func Version() string {
	rev := "unknown"
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				rev = fmt.Sprintf("#%s", setting.Value[:7])
				break
			}
		}
	}
	return fmt.Sprintf("%s (%s)", version, rev)
}
