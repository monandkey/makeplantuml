package makeplantuml

import (
	"os"
	"fmt"
	"runtime"
	"strconv"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type profile struct {
	Path    path    `yaml:"path"`
	Feature feature `yaml:"feature"`
}

type path struct {
	Java      string `yaml:"java"`
	Wireshark bool   `yaml:"wireshark"`
}

type feature struct {
	Timestamp      int    `yaml:"timestamp"`
	NameResolution string `yaml:"nameResolution"`
}

type configPath struct {
	path string
}

func (c configPath) configLoad() profile {
	profile := profile{}
	b, _ := os.ReadFile(c.path)
	yaml.Unmarshal(b, &profile)
	return profile
}

type disassembledCharacter struct {
	homedir  string
	separate string
	filename string
}

func (d disassembledCharacter) stringJoin() configPath {
	c := configPath{path: d.homedir + d.separate + d.filename}
	return c
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

type Config struct {
	Java           string
	Wireshark      string
	Timestamp      bool
	NameResolution bool
}

func InitializeConfig(initConfig Config) {
	homeDir, _ := os.UserHomeDir()
	var fileName string

	if runtime.GOOS == "windows" {
		str := homeDir + "\\" + getConfigName()
		fileName = str

	} else if runtime.GOOS == "linux" {
		str := homeDir + "/" + getConfigName()
		fileName = str

	} else {
		fmt.Println("Your OS is not supported.")
		os.Exit(0)
	}

	fp, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	data := [] string{
		"version: 1.0\n",
		"profile: \n",
		"  path: \n",
		"    java: " + initConfig.Java + "\n",
		"    wireshark: " + initConfig.Wireshark + "\n",
		"  feature: \n",
		"    timestamp: " + strconv.FormatBool(initConfig.Timestamp) + "\n",
		"    nameResolution: " + strconv.FormatBool(initConfig.NameResolution) + "\n",
	}
	writeConfig(data, fileName)
}

func writeConfig(data []string, fileName string) {
	b := []byte{}
	for _, line := range data {
		ll := []byte(line)
		for _, l := range ll {
			b = append(b, l)
		}
	}

	err := ioutil.WriteFile(fileName, b, 0666)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(0)
	}
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