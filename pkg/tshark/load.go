package tshark

import (
	"os"
	"fmt"
	"regexp"
	"strings"
	"local.packages/util"
)

func hostInfoFormating(h string) []map[string]string {
	h = regexp.MustCompile("\r\n").ReplaceAllString(h, "\n")
	lines := strings.Split(h, "\n")

	var hosts []map[string]string
	for _, line := range lines {
		if regexp.MustCompile("^;").Match([]byte(line)) {
			continue
		}

		column := strings.Split(line, " ")

		if len(column) <= 1 {
			continue
		}

		for len(column) <= 4 {
			column = append(column, "")
		}

		if regexp.MustCompile("[a-zA-Z]").Match([]byte(column[2])) {
			// Do Nothing. 

		} else {
			tmp := column[3]
			column[3] = column[2]
			column[2] = tmp
		}

		hostMap := map[string]string{
			"address": column[0],
			"name":    column[1],
			"nf":      column[2],
			"port":    column[3],
		}

		hosts = append(hosts, []map[string]string{hostMap}...)
	}
	return hosts
}

func checkResolution(r []string, t string) []string {
	for _, v := range r {
		if v == t {
			return r
		}
	}
	file, err := os.OpenFile(util.PumlLocation.Path + "/tmp.puml", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	str := "participant " + t
	fmt.Fprintln(file, str)

	r = append(r, t)
	return r
}


func NameResolution(t TsharkHeaders) {
	hosts := hostInfoFormating(util.FileRead("./docs/hosts"))

	var resolvedAddress []string
	for _, host := range hosts {
		label1:
		for i, v := range t {
			if v.SrcAddr == host["address"] && v.SrcPort == host["port"] {
				t[i].SrcAddr = host["name"]
				resolvedAddress = checkResolution(resolvedAddress, host["name"])
				continue
			}

			if v.DstAddr == host["address"] && v.DstPort == host["port"] {
				t[i].DstAddr = host["name"]
				resolvedAddress = checkResolution(resolvedAddress, host["name"])
				continue
			}

			for _, w := range hosts {
				if host["address"] == w["address"] && v.SrcPort == w["port"] && host["port"] != "" && host["name"] != w["name"] {
					break label1
				}

				if host["address"] == w["address"] && v.DstPort == w["port"] && host["port"] != "" && host["name"] != w["name"] {
					break label1
				}
			}

			if v.SrcAddr == host["address"] {
				t[i].SrcAddr = host["name"]
				resolvedAddress = checkResolution(resolvedAddress, host["name"])
				continue
			}

			if v.DstAddr == host["address"] {
				t[i].DstAddr = host["name"]
				resolvedAddress = checkResolution(resolvedAddress, host["name"])
				continue
			}
		}
	}
}