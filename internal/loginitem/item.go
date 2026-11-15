package loginitem

type loginItem struct {
	Name   string `yaml:"name"`
	Path   string `yaml:"path"`
	Hidden bool   `yaml:"hidden"`
}

// Represents the set of macOS login items managed by the application.
type LoginItems []*loginItem
