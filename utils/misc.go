package utils

import "path/filepath"

func getFileName(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
