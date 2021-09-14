package user

func (t toWritingUser) new() UserMethod {
	return &toWritingUser{}
}

func (t toWritingUser) RenderingE(fileName string) error {
	return nil
}


func (f fromRenderingUser) new() UserMethod {
	return &fromRenderingUser{}
}

func (f fromRenderingUser) SetCmd() {
	return
}

func (f fromRenderingUser) SetArgs(fileName string) {
	return
}

func (f fromRenderingUser) RunE() error {
	return nil
}

func (f fromRenderingUser) Parse() {
	return
}

func (f fromRenderingUser) NameResE(fileName string) error {
	return nil
}

func (f fromRenderingUser) Display() {
	return
}

func (f fromRenderingUser) CreateE(title string) error {
	return nil
}

func (f fromRenderingUser) WritingE(timestamp bool) error {
	return nil
}
