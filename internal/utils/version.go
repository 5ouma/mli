package utils

import (
	"fmt"
	"runtime/debug"
	"strings"
)

func Version(info *debug.BuildInfo, ok bool) string {
	if !ok {
		return "unknown"
	}
	versions := strings.Split(info.Main.Version, "-")
	if len(versions) < 3 {
		if versions[0] == "(devel)" {
			return "unknown"
		}
		return versions[0]
	}
	return fmt.Sprintf("%s (#%s)", versions[0], versions[2][:7])

}
