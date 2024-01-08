package utils

import (
	"fmt"
)

func errAppNotFound(app string) error {
	return fmt.Errorf(`app not found: "%s"`, app)
}
