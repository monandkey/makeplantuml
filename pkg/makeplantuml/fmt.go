package makeplantuml

import (
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
		return t
	}

	return ""
}

type tsharkHeaders []tsharkHeader

func (t tsharkHeaders) setHeader(out string) tsharkHeaders {
	for regexp.MustCompile(",,").Match([]byte(out)) {
		out = regexp.MustCompile(",{2}").ReplaceAllString(out, ",\"\",")
	}

	lines := strings.Split(out, "\n")

	for _, line := range lines {
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
		tMessage := regexp.MustCompile("\"").ReplaceAllString(column[13][0], "")
		tChecksum := setPortAndCheckSum(column[14][0], column[15][0], column[16][0])

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
