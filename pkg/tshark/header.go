package tshark

import (
	"regexp"
	"strings"
)

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

	if protocol == "TCP" || protocol == "SCTP" || regexp.MustCompile("ICMP").Match([]byte(protocol)) {
		return ""
	}
	return msg
}

func (t TsharkHeaders) setHeader(out string) TsharkHeaders {
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

		th := TsharkHeader{
			Number:   tNumber,
			Time:     tTime,
			SrcAddr:  tSrcAddr,
			SrcPort:  tSrcPort,
			DstAddr:  tDstAddr,
			DstPort:  tDstPort,
			Protocol: tProtocol,
			Message:  tMessage,
			Checksum: tChecksum,
		}
		t = append(t, th)
	}
	return t
}
