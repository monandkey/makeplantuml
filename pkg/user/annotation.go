package user

func (a annotationUser) new() UserMethod {
	return &annotationUser{}
}

func (a *annotationUser) SetArgs(fileName string) {
	a.Args = a.GetArgs("annotation", fileName)
}

func (a *annotationUser) Parse() {
	a.Header = a.ParserOutput("annotation")
}

func (a *annotationUser) WritingE(timestamp bool) error {
	if err := a.WriteUml("annotation", a.Header, timestamp); err != nil {
		return err
	}
	return nil
}
