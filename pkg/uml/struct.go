package uml

type UmlArgs struct {
	Target string
	Output string
	Export string
}

type UmlMethod interface {
	CreateE(string) error
	WritingE(bool) error
	RenderingE(string) error
}
