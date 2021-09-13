package uml

import (
	"os"
	"fmt"
	"errors"
	"runtime"
	"os/exec"
	"local.packages/util"
	"local.packages/cfg"
)

func (u UmlArgs) CreateTemplate(t string) error {
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

func (u UmlArgs) WriteUml(headers []map[string]string, tf bool) error {
	file, err := os.OpenFile(util.PumlLocation.Path + "/tmp.puml", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	for i, header := range headers {
		if i == 0 {
			fmt.Fprintln(file, "")
		}

		str := "\"" + header["srcAddr"] + "\" -> \"" + header["dstAddr"] + "\" : ["+ header["number"] + "][" + header["protocol"] + "] " + header["message"]
		fmt.Fprintln(file, str)

		if tf {
			rnote := "rnote left: " + header["time"]
			fmt.Fprintln(file, rnote)
		}
	}
	fmt.Fprintln(file, "\n@enduml")
	return nil
}

func (u UmlArgs) RenderingUml(fileName string) error {
	var (
		cmd      string
		plantuml string
	)

	if cfg.Param.Profile.Path.Java == "default" {
		switch(runtime.GOOS) {
			case "windows":
				cmd = "java"
			case "linux":
				cmd = "java"
			default:
				return errors.New("Your OS not support.")
		}

		} else {
		cmd = cfg.Param.Profile.Path.Java
	}

	if cfg.Param.Profile.Path.Plantuml == "default" {
		plantuml = cfg.PlantumlShortPath
	
	} else {
		plantuml = cfg.Param.Profile.Path.Plantuml
	}

	if util.OutLocation.ValidateLocation() {
		err := os.Mkdir(util.OutLocation.Path, 0777)
		if err != nil {
			return err
		}
	}

	out, _ := exec.Command(cmd,
		"-jar", plantuml,
		fileName,
		"-o", "." + util.OutLocation.Path,
		"-tsvg",
	).CombinedOutput()

	if string(out) != "" {
		return errors.New("Failed to render the puml file.\n" + string(out))
	}

	return nil
}
