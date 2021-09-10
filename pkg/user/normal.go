package user

func (n normalUser) new() UserMethod {
	return &normalUser{}
}
