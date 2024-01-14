package ci

import "fmt"

var (
	ErrorItemsNotSame = fmt.Errorf("items are not the same")
)

func Err(err error) string {
	return fmt.Sprintf("🚨 %v", err)
}
