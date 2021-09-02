package makeplantuml

import (
	"os"
	"fmt"
	"runtime"
	"os/exec"
)

func CreateTemplate() {
	file, err := os.Create("./.tmp.puml")

	if err != nil {
		fmt.Println(os.Stderr, err)
	}
	defer file.Close()

	file.Write(([]byte)("@startuml\n"))
	file.Write(([]byte)("skinparam Monochrome true\n"))
	file.Write(([]byte)("skinparam shadowing false\n"))
	file.Write(([]byte)("hide footbox\n"))
	file.Write(([]byte)("title xxxxxxxx\n"))
	file.Write(([]byte)("skinparam note {\n"))
	file.Write(([]byte)("	BackgroundColor white\n"))
	file.Write(([]byte)("	BorderColor white\n"))
	file.Write(([]byte)("}\n"))
}

func RenderingUml() {
	var cmd string
	switch(runtime.GOOS) {
		case "windows":
			cmd = "java"
		case "linux":
			cmd = "java"
		default:
			fmt.Println("Your OS not support.")
			os.Exit(0)
	}

	out, err := exec.Command(cmd,
		"-jar", "./docs/plantuml.jar",
		".tmp.puml",
		"-o", "./",
	).Output()

	fmt.Println("Exec")

	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(0)
	}

	fmt.Println("No Error")
	fmt.Println(out)
}