package loginitem

type loginItem struct {
	Name   string `yaml:"name"`
	Path   string `yaml:"path"`
	Hidden bool   `yaml:"hidden"`
}

type LoginItems []*loginItem
