package makeplantuml

import (
	"os"
	"fmt"
	"runtime"
	"os/exec"
)

var pumlLocation = location{
	path: "./puml",
}

var outLocation = location{
	path: "./result",
}

func CreateTemplate() {
	if pumlLocation.validateLocation() {
		err := os.Mkdir(pumlLocation.path, 0777)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(0)
		}
	}

	file, err := os.Create(pumlLocation.path + "/tmp.puml")
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
	file.Write(([]byte)("}\n\n"))
}

func WriteUml(t tsharkHeaders) {
	file, err := os.OpenFile(pumlLocation.path + "/tmp.puml", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for i, v := range t {
		if i == 0 {
			fmt.Fprintln(file, "")
		}

		str := v.srcAddr + " -> " + v.dstAddr + " : " + v.message
		fmt.Fprintln(file, str)
	}
	fmt.Fprintln(file, "\n@enduml")
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

	if outLocation.validateLocation() {
		err := os.Mkdir(outLocation.path, 0777)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(0)
		}
	}

	err := exec.Command(cmd,
		"-jar", "./docs/plantuml.jar",
		pumlLocation.path + "/tmp.puml",
		"-o", "." + outLocation.path,
	).Run()

	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(0)
	}
}
