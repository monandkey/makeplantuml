package user

func (h handsonUser) new() UserMethod {
	return &handsonUser{}
}

func (h *handsonUser) SetArgs(fileName string) {
	h.Args = h.GetArgs("handson", fileName)
}
