package uml

type UmlMethod interface {
	CreateE(string) error
	WritingE(bool) error
	RenderingE() error
}
