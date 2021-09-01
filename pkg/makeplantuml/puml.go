package makeplantuml

import (
	"os"
	"fmt"
)

func CreateTemplate() {
	file, err := os.Create("./.tmp.puml")

	if err != nil {
		fmt.Println(os.Stderr, err)
	}
	defer file.Close()

	file.Write(([]byte)("skinparam Monochrome true\n"))
	file.Write(([]byte)("skinparam shadowing false\n"))
	file.Write(([]byte)("hide footbox\n"))
	file.Write(([]byte)("title xxxxxxxx\n"))
	file.Write(([]byte)("skinparam note {\n"))
	file.Write(([]byte)("	BackgroundColor white\n"))
	file.Write(([]byte)("	BorderColor white\n"))
	file.Write(([]byte)("}\n"))
}