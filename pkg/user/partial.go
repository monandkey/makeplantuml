package user

func (t toWritingUser) new() UserMethod {
	return &toWritingUser{}
}

func (t toWritingUser) RenderingE() error {
	return nil
}
