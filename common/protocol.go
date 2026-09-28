package common

// Protocol identifies one specification implemented by a package.
type Protocol struct {
	Name    string
	Version string
	Docs    []Doc
}

// Doc is a link to a specification or a companion document.
type Doc struct {
	Title string
	URL   string
}
