package cfg

import (
	"os"
	"fmt"
	"errors"
	"runtime"
	"strconv"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

func (c configPath) configLoad() profile {
	profile := profile{}
	b, _ := os.ReadFile(c.path)
	yaml.Unmarshal(b, &profile)
	return profile
}

func (d disassembledCharacter) stringJoin() configPath {
	c := configPath{path: d.homedir + d.separate + d.filename}
	return c
}

func getConfigParameter() profile {
	ds := disassembledCharacter{
		homedir:  getHomedir(),
		separate: getSeparate(),
		filename: getConfigName(),
	}
	return ds.stringJoin().configLoad()
}

func getHomedir() string {
	h, err := os.UserHomeDir()
	if err != nil {
		os.Exit(0)
	}
	return h
}

func getSeparate() string {
	switch runtime.GOOS {
		case "windows":
			return "\\"
		case "linux":
			return "/"
		default:
			fmt.Println("Your OS is not support")
			os.Exit(0)
	}
	return ""
}

func getConfigName() string {
	return ".makeplantuml.yml"
}

func InitializeConfig(initConfig Config) error {
	homeDir, _ := os.UserHomeDir()
	var fileName string

	if runtime.GOOS == "windows" {
		str := homeDir + "\\" + getConfigName()
		fileName = str

	} else if runtime.GOOS == "linux" {
		str := homeDir + "/" + getConfigName()
		fileName = str

	} else {
		return errors.New("Your OS is not supported.")
	}

	fp, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fp.Close()
	data := [] string{
		"version: 1.0\n",
		"profile: \n",
		"  path: \n",
		"    java: " + initConfig.Java + "\n",
		"    wireshark: " + initConfig.Wireshark + "\n",
		"    plantuml: " + initConfig.Plantuml + "\n",
		"  feature: \n",
		"    timestamp: " + strconv.FormatBool(initConfig.Timestamp) + "\n",
		"    nameResolution: " + strconv.FormatBool(initConfig.NameResolution) + "\n",
	}
	err = writeConfig(data, fileName)
	return err
}

func writeConfig(data []string, fileName string) error {
	b := []byte{}
	for _, line := range data {
		ll := []byte(line)
		for _, l := range ll {
			b = append(b, l)
		}
	}

	err := ioutil.WriteFile(fileName, b, 0666)
	return err
}

func ExistInitConfig() bool {
	ds := disassembledCharacter{
		homedir:  getHomedir(),
		separate: getSeparate(),
		filename: getConfigName(),
	}
	f := ds.homedir + ds.separate + ds.filename
	_, err := os.Stat(f)
	return err != nil
}