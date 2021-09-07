package uml

import (
	"os"
	"fmt"
	"runtime"
	"os/exec"
	"local.packages/util"
	"local.packages/tshark"
	"local.packages/cfg"
)

func CreateTemplate(t string) error {
	if util.PumlLocation.ValidateLocation() {
		err := os.Mkdir(util.PumlLocation.Path, 0777)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(util.PumlLocation.Path + "/tmp.puml")
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(([]byte)("@startuml\n"))
	file.Write(([]byte)("skinparam Monochrome true\n"))
	file.Write(([]byte)("skinparam shadowing false\n"))
	file.Write(([]byte)("skinparam defaultFontName Courier\n"))
	file.Write(([]byte)("hide footbox\n"))

	if t != "" {
		file.Write(([]byte)("title " + t +"\n"))

	} else {
		file.Write(([]byte)("title \"\"\n"))
	}

	file.Write(([]byte)("skinparam note {\n"))
	file.Write(([]byte)("	BackgroundColor white\n"))
	file.Write(([]byte)("	BorderColor white\n"))
	file.Write(([]byte)("}\n\n"))
	return nil
}

func WriteUml(t tshark.TsharkHeaders, tf bool) {
	file, err := os.OpenFile(util.PumlLocation.Path + "/tmp.puml", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for i, v := range t {
		if i == 0 {
			fmt.Fprintln(file, "")
		}

		str := "\"" + v.SrcAddr + "\" -> \"" + v.DstAddr + "\" : ["+ v.Number + "][" + v.Protocol + "] " + v.Message
		fmt.Fprintln(file, str)

		if tf {
			rnote := "rnote left: " + v.Time
			fmt.Fprintln(file, rnote)
		}
	}
	fmt.Fprintln(file, "\n@enduml")
}

func RenderingUml() {
	var cmd string
	if cfg.CfgVal.Profile.Path.Java == "default" {
		switch(runtime.GOOS) {
			case "windows":
				cmd = "java"
			case "linux":
				cmd = "java"
			default:
				fmt.Println("Your OS not support.")
				os.Exit(0)
		}

		} else {
		cmd = cfg.CfgVal.Profile.Path.Java
	}

	if util.OutLocation.ValidateLocation() {
		err := os.Mkdir(util.OutLocation.Path, 0777)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(0)
		}
	}

	err := exec.Command(cmd,
		"-jar", "./docs/plantuml.jar",
		util.PumlLocation.Path + "/tmp.puml",
		"-o", "." + util.OutLocation.Path,
	).Run()

	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(0)
	}
}
