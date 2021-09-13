package tshark

import (
	"regexp"
	"strings"
	"strconv"
)

func (t TsharkArgs) SetAddress(v4 string, v6 string, lenv4 string, lenv6 string) string {
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

	if len(lines) == 1 {
		return lines[0]["address"]
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

func (t TsharkArgs) SetPortAndCheckSum(udp string, tcp string, sctp string) string {
	udp = regexp.MustCompile("\"").ReplaceAllString(udp, "")
	tcp = regexp.MustCompile("\"").ReplaceAllString(tcp, "")
	sctp = regexp.MustCompile("\"").ReplaceAllString(sctp, "")

	if udp != "" && tcp == "" && sctp == "" {
		return udp
	}

	if udp == "" && tcp != "" && sctp == "" {
		return tcp
	}

	if udp == "" && tcp == "" && sctp != "" {
		return sctp
	}

	return ""
}

func (t TsharkArgs) SetMessage(msg string, protocol string) string {
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
