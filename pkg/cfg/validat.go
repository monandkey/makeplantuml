package cfg

import (
	"fmt"
	"errors"
	"runtime"
	"os/exec"
	"local.packages/util"
)

func validationJava() error {
	var cmd string

	if Param.Profile.Path.Java == "default" {
		switch(runtime.GOOS) {
			case "windows":
				cmd, _ = exec.LookPath("java")
			case "linux":
				cmd, _ = exec.LookPath("java")
			default:
				return errors.New("Your OS not support.")
		}

		} else {
		cmd = Param.Profile.Path.Java
	}

	if util.FileExist(cmd) {
		return errors.New("")
	}
	return nil
}

func validationWireshark() error {
	var cmd string

	if Param.Profile.Path.Wireshark == "default" {
		switch(runtime.GOOS) {
			case "windows":
				cmd = "C:/Program Files/Wireshark/tshark.exe"
			case "linux":
				cmd, _ = exec.LookPath("tshark")
			default:
				return errors.New("Your OS not support.")
		}

		} else {
		cmd = Param.Profile.Path.Wireshark + "\\tshark.exe"
	}

	if util.FileExist(cmd) {
		return errors.New("")
	}
	return nil
}

func validationPlantuml() error {
	var cmd string

	if Param.Profile.Path.Plantuml == "default" {
		cmd = PlantumlShortPath
	
	} else {
		cmd = Param.Profile.Path.Plantuml
	}

	if util.FileExist(cmd) {
		return errors.New("")
	}
	return nil
}

func ValidationConfig() error {
	const errMsg string = "does not exist in the specified location."

	if err := validationJava(); err != nil {
		return errors.New("Error: Java " + errMsg)
	}

	if err := validationWireshark(); err != nil {
		return errors.New("Error: Wireshark " + errMsg)
	}

	if err := validationPlantuml(); err != nil {
		return errors.New("Error: PlantUML " + errMsg)
	}

	fmt.Println("OK")
	return nil
}