package user

import (
	"fmt"
	"os/exec"
)

func (b *baseUser) SetCmd() {
	b.Cmd = b.GetTsharkCommand()
}

func (b *baseUser) SetArgs(fileName string) {
	b.Args = b.GetArgs("normal", fileName)
}

func (b *baseUser) RunE() error {
	var err error
	b.Out, err = exec.Command(b.Cmd, b.Args...).Output()

	if err != nil {
		return err
	}
	return nil
}

func (b *baseUser) Parse() {
	b.Header = b.ParserOutput("normal")
}

func (b *baseUser) NameResE(fileName string) error {
	if err := b.NameResolution(b.Header, fileName); err != nil {
		return err
	}
	return nil
}

func (b *baseUser) Display() {
	fmt.Println(b.Header)
}

func (b *baseUser) CreateE(title string) error {
	if err := b.CreateTemplate(title); err != nil {
		return err
	}
	return nil
}

func (b *baseUser) WritingE(timestamp bool) error {
	if err := b.WriteUml("normal", b.Header, timestamp); err != nil {
		return err
	}
	return nil
}

func (b *baseUser) RenderingE(fileName string) error {
	if err := b.RenderingUml(fileName); err != nil {
		return err
	}
	return nil
}
