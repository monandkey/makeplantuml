package makeplantuml

import (
	"os"
	"fmt"
	"regexp"
	"strings"
)

type tsharkHeader struct {
	number   string
	time     string
	srcAddr  string
	srcPort  string
	dstAddr  string
	dstPort  string
	protocol string
	message  string
	checksum string
}

func setAddress(v4 string, v6 string) string {
	v4 = regexp.MustCompile("\"").ReplaceAllString(v4, "")
	v6 = regexp.MustCompile("\"").ReplaceAllString(v6, "")

	if v4 != "" && v6 == "" {
		return v4
	}

	if v4 == "" && v6 != "" {
		return v6
	}

	return ""
}

func setPortAndCheckSum(u string, t string, s string) string {
	u = regexp.MustCompile("\"").ReplaceAllString(u, "")
	t = regexp.MustCompile("\"").ReplaceAllString(t, "")
	s = regexp.MustCompile("\"").ReplaceAllString(s, "")

	if u != "" && t == "" && s == "" {
		return u
	}

	if u == "" && t != "" && s == "" {
		return t
	}

	if u == "" && t == "" && s != "" {
		return s
	}

	return ""
}

func setMessage(msg string, protocol string) string {
	msg = regexp.MustCompile("\"").ReplaceAllString(msg, "")

	if protocol == "GTPv2" {
		return msg
	}

	if protocol == "DIAMETER" {
		msg = regexp.MustCompile("Request.*$").ReplaceAllString(msg, "Request")
		msg = regexp.MustCompile("Answer.*$").ReplaceAllString(msg, "Answer")
		msg = regexp.MustCompile("^.*cmd=3GPP-").ReplaceAllString(msg, "")
		msg = regexp.MustCompile("^.*cmd=").ReplaceAllString(msg, "")
		return msg
	}

	if regexp.MustCompile("S1AP").Match([]byte(protocol)) {
		return msg
	}

	if protocol == "HTTP" || protocol == "HTTP2" {
		msg = regexp.MustCompile("^.*: ").ReplaceAllString(msg, "")
		msg = regexp.MustCompile(",.*$").ReplaceAllString(msg, "")
		msg = regexp.MustCompile("\\?.*$").ReplaceAllString(msg, "")
		msg = regexp.MustCompile("nf-instances.*$").ReplaceAllString(msg, "nf-instances")

		if regexp.MustCompile("(GET)|(HEAD)|(POST)|(PUT)|(DELETE)|(CONNECT)|(OPTIONS)|(TRACE)|(PATCH)|[0-9]{3}").Match([]byte(msg)) {
			return msg
		}
		return ""
	}

	if protocol == "TCP" || protocol == "SCTP" {
		return ""
	}
	return msg
}

type tsharkHeaders []tsharkHeader

func (t tsharkHeaders) setHeader(out string) tsharkHeaders {
	for regexp.MustCompile(",,").Match([]byte(out)) {
		out = regexp.MustCompile(",{2}").ReplaceAllString(out, ",\"\",")
	}

	lines := strings.Split(out, "\n")

	for _, line := range lines {
		if regexp.MustCompile(",\r$").Match([]byte(line)) {
			line = regexp.MustCompile(",\r$").ReplaceAllString(line, ",\"\"")
		}

		if regexp.MustCompile(",$").Match([]byte(line)) {
			line = regexp.MustCompile(",$").ReplaceAllString(line, ",\"\"")
		}

		column := regexp.MustCompile("\".*?\"").FindAllStringSubmatch(line, -1)
		if len(column) == 0 {
			continue
		}

		tNumber := regexp.MustCompile("\"").ReplaceAllString(column[0][0], "")
		tTime := regexp.MustCompile("\"").ReplaceAllString(column[1][0], "")
		tSrcAddr := setAddress(column[2][0], column[3][0])
		tSrcPort := setPortAndCheckSum(column[4][0], column[5][0], column[6][0])
		tDstAddr := setAddress(column[7][0], column[8][0])
		tDstPort := setPortAndCheckSum(column[9][0], column[10][0], column[11][0])
		tProtocol := regexp.MustCompile("\"").ReplaceAllString(column[12][0], "")
		tMessage := setMessage(column[13][0], tProtocol)
		tChecksum := setPortAndCheckSum(column[14][0], column[15][0], column[16][0])

		if tMessage == "" {
			continue
		}

		th := tsharkHeader{
			number:   tNumber,
			time:     tTime,
			srcAddr:  tSrcAddr,
			srcPort:  tSrcPort,
			dstAddr:  tDstAddr,
			dstPort:  tDstPort,
			protocol: tProtocol,
			message:  tMessage,
			checksum: tChecksum,
		}
		t = append(t, th)
	}
	return t
}

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
	file, err := os.OpenFile(pumlLocation.path + "/tmp.puml", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	str := "participant " + t
	fmt.Fprintln(file, str)

	r = append(r, t)
	return r
}


func NameResolution(t tsharkHeaders) {
	hosts := hostInfoFormating(fileRead("./docs/hosts"))

	var resolvedAddress []string
	for _, host := range hosts {
		label1:
		for i, v := range t {
			if v.srcAddr == host["address"] && v.srcPort == host["port"] {
				t[i].srcAddr = host["name"]
				resolvedAddress = checkResolution(resolvedAddress, host["name"])
				continue
			}

			if v.dstAddr == host["address"] && v.dstPort == host["port"] {
				t[i].dstAddr = host["name"]
				resolvedAddress = checkResolution(resolvedAddress, host["name"])
				continue
			}

			for _, w := range hosts {
				if host["address"] == w["address"] && v.srcPort == w["port"] && host["port"] != "" && host["name"] != w["name"] {
					break label1
				}

				if host["address"] == w["address"] && v.dstPort == w["port"] && host["port"] != "" && host["name"] != w["name"] {
					break label1
				}
			}

			if v.srcAddr == host["address"] {
				t[i].srcAddr = host["name"]
				resolvedAddress = checkResolution(resolvedAddress, host["name"])
				continue
			}

			if v.dstAddr == host["address"] {
				t[i].dstAddr = host["name"]
				resolvedAddress = checkResolution(resolvedAddress, host["name"])
				continue
			}
		}
	}
}
