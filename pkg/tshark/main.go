package tshark

import (
	"regexp"
	"strings"
	"runtime"
	"local.packages/cfg"
)

func GetTsharkCommand() string {
	if cfg.Param.Profile.Path.Wireshark == "default" {
		switch(runtime.GOOS) {
			case "windows":
				return "C:/Program Files/Wireshark/tshark.exe"
			case "linux":
				return "tshark"
		}
	}
	return cfg.Param.Profile.Path.Wireshark + "\\tshark.exe"
}

func ConvertOutputResultIntoArray(out string) []string {
	for regexp.MustCompile(",,").Match([]byte(out)) {
		out = regexp.MustCompile(",{2}").ReplaceAllString(out, ",\"\",")
	}
	return strings.Split(out, "\n")
}

func CleanLineStrings(line string) string {
	if regexp.MustCompile(",\r$").Match([]byte(line)) {
		line = regexp.MustCompile(",\r$").ReplaceAllString(line, ",\"\"")
	}

	if regexp.MustCompile(",$").Match([]byte(line)) {
		line = regexp.MustCompile(",$").ReplaceAllString(line, ",\"\"")
	}
	return line
}
