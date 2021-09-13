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

func checkResolution(r []string, t string, f string) []string {
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

	if f != "" {
		str := "participant " + t + " as " + f
		fmt.Fprintln(file, str)

	} else {
		str := "participant " + t
		fmt.Fprintln(file, str)
	} 

	r = append(r, t)
	return r
}

func NameOrNfSelection(name string, nf string) string {
	if nf == "" {
		return name
	}
	return nf
}

func (t TsharkArgs) NameResolution(headers []map[string]string, hostsFile string) error {
	tmp, err := util.FileRead(hostsFile)
	if err != nil {
		return err
	}

	hosts := hostInfoFormating(tmp)

	var resolvedAddress []string
	for _, host := range hosts {
		label1:
		for i, header := range headers {
			if header["srcAddr"] == host["address"] && header["srcPort"] == host["port"] {
				headers[i]["srcAddr"] = NameOrNfSelection(host["name"], host["nf"])
				resolvedAddress = checkResolution(resolvedAddress, host["name"], host["nf"])
				continue
			}

			if header["dstAddr"] == host["address"] && header["dstPort"] == host["port"] {
				headers[i]["dstAddr"] = NameOrNfSelection(host["name"], host["nf"])
				resolvedAddress = checkResolution(resolvedAddress, host["name"], host["nf"])
				continue
			}

			for _, w := range hosts {
				if host["address"] == w["address"] && header["srcPort"] == w["port"] && host["port"] != "" && host["name"] != w["name"] {
					break label1
				}

				if host["address"] == w["address"] && header["dstPort"] == w["port"] && host["port"] != "" && host["name"] != w["name"] {
					break label1
				}
			}

			if header["srcAddr"] == host["address"] {
				headers[i]["srcAddr"] = NameOrNfSelection(host["name"], host["nf"])
				resolvedAddress = checkResolution(resolvedAddress, host["name"], host["nf"])
				continue
			}

			if header["dstAddr"] == host["address"] {
				headers[i]["dstAddr"] = NameOrNfSelection(host["name"], host["nf"])
				resolvedAddress = checkResolution(resolvedAddress, host["name"], host["nf"])
				continue
			}
		}
	}
	return nil
}
