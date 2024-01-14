package api

type loginItem struct {
	Name   string
	Path   string
	Hidden bool
}

type LoginItems []*loginItem
