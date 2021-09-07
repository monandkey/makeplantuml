package tshark

import (
	"regexp"
	"strings"
	"strconv"
)

func SetAddress(v4 string, v6 string, lenv4 string, lenv6 string) string {
	v4 = regexp.MustCompile("\"").ReplaceAllString(v4, "")
	v6 = regexp.MustCompile("\"").ReplaceAllString(v6, "")
	lenv4 = regexp.MustCompile("\"").ReplaceAllString(lenv4, "")
	lenv6 = regexp.MustCompile("\"").ReplaceAllString(lenv6, "")

	var (
		v4s     []string
		v6s     []string
		linesv4 []string
		linesv6 []string
		lines   []map[string]string
	)

	if regexp.MustCompile(",").Match([]byte(lenv4)) {
		linesv4 = strings.Split(lenv4, ",")
		v4s = strings.Split(v4, ",")

		for i, _ := range linesv4 {
			column := map[string]string{
				"address": v4s[i],
				"length":  linesv4[i],
			}
			lines = append(lines, []map[string]string{column}...)
		}
	
	} else if v4 != "" {
		column := map[string]string{
			"address": v4,
			"length":  lenv4,
		}
		lines = append(lines, []map[string]string{column}...)
	}

	if regexp.MustCompile(",").Match([]byte(lenv6)) {
		linesv6 = strings.Split(lenv6, ",")
		v6s = strings.Split(v6, ",")

		for i, _ := range linesv6 {
			column := map[string]string{
				"address": v6s[i],
				"length":  linesv6[i],
			}
			lines = append(lines, []map[string]string{column}...)
		}

	} else if v6 != "" {
		column := map[string]string{
			"address": v6,
			"length":  lenv6,
		}
		lines = append(lines, []map[string]string{column}...)
	}

	for i, _ := range lines {
		for j, _ := range lines {
			right, _ := strconv.Atoi(lines[i]["length"])
			left, _ := strconv.Atoi(lines[j]["length"])

			if right < left {
				tmp := lines[i]
				lines[i] = lines[j]
				lines[j] = tmp
			}
		}
	}
	return lines[0]["address"]
}

func SetPortAndCheckSum(u string, t string, s string) string {
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

func SetMessage(msg string, protocol string) string {
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

	if regexp.MustCompile("(HTTP)|(HTTP2)").Match([]byte(protocol)) {
		msg = regexp.MustCompile("^.*, PDU").ReplaceAllString(msg, "PDU")
		msg = regexp.MustCompile("^.*: ").ReplaceAllString(msg, "")
		msg = regexp.MustCompile(",.*$").ReplaceAllString(msg, "")
		msg = regexp.MustCompile("\\?.*$").ReplaceAllString(msg, "")
		msg = regexp.MustCompile("nf-instances.*$").ReplaceAllString(msg, "nf-instances")

		if regexp.MustCompile("(GET)|(HEAD)|(POST)|(PUT)|(DELETE)|(CONNECT)|(OPTIONS)|(TRACE)|(PATCH)|(PDU)|[0-9]{3}").Match([]byte(msg)) {
			return msg
		}
		return ""
	}

	if protocol == "TCP" || protocol == "SCTP" {
		return ""
	}
	return msg
}

func (t TsharkHeaders) SetHeader(out string) TsharkHeaders {
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
		tSrcAddr := SetAddress(column[2][0], column[3][0], column[17][0], column[18][0])
		tSrcPort := SetPortAndCheckSum(column[4][0], column[5][0], column[6][0])
		tDstAddr := SetAddress(column[7][0], column[8][0], column[17][0], column[18][0])
		tDstPort := SetPortAndCheckSum(column[9][0], column[10][0], column[11][0])
		tProtocol := regexp.MustCompile("\"").ReplaceAllString(column[12][0], "")
		tMessage := SetMessage(column[13][0], tProtocol)
		tChecksum := SetPortAndCheckSum(column[14][0], column[15][0], column[16][0])

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
