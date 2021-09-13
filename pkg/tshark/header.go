package tshark

import (
	"regexp"
	"strings"
	"strconv"
	"local.packages/s1ap"
	"local.packages/ngap"
	"local.packages/gtpv2"
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

func delStrConvInt(str string) int {
	str = regexp.MustCompile("\"").ReplaceAllString(str, "")
	str = regexp.MustCompile(",.*$").ReplaceAllString(str, "")
	i, _ := strconv.Atoi(str)
	return i
}

func (t TsharkArgs) SetAnnotation(column [][]string) string {
	tmpProtocol := column[12][0]

	type s1apAnnotation struct {
		type_of_id int
		dcnr_cap   int
	}

	type ngapAnnotation struct {
		type_id int
	}

	type gtpv2Annotation struct {
		oi    int
		si    int
		dcnr  int
		cause int
	}

	type pfcpAnnotation struct {
		cause int
	}

	type diameterAnnotation struct {
		resultCode int
		CcReqType  int
	}

	s1apValue := s1apAnnotation{
		type_of_id: delStrConvInt(column[16][0]),
		dcnr_cap:   delStrConvInt(column[17][0]),
	}

	ngapValue := ngapAnnotation{
		type_id: delStrConvInt(column[18][0]),
	}

	gtpv2Value := gtpv2Annotation{
		oi:    delStrConvInt(column[19][0]),
		si:    delStrConvInt(column[20][0]),
		dcnr:  delStrConvInt(column[21][0]),
		cause: delStrConvInt(column[22][0]),
	}

	// pfcpValue := pfcpAnnotation{
	// 	cause: delStrConvInt(column[23][0]),
	// }

	// diameterValue := diameterAnnotation{
	// 	resultCode: delStrConvInt(column[24][0]),
	// 	CcReqType:  delStrConvInt(column[25][0]),
	// }

	if regexp.MustCompile("S1AP").Match([]byte(tmpProtocol)) {
		typeOfId := s1ap.GetTypeOfId(s1apValue.type_of_id)
		nasDcnr := s1ap.GetDcnr(s1apValue.dcnr_cap)
		linking := typeOfId + nasDcnr
		return linking
	
	} else if regexp.MustCompile("NGAP").Match([]byte(tmpProtocol)) {
		typeOfId := ngap.GetTypeOfId(ngapValue.type_id)
		linking := typeOfId
		return linking
	
	} else if regexp.MustCompile("GTPv2").Match([]byte(tmpProtocol)) {
		oi := gtpv2.GetOiIndication(gtpv2Value.oi)
		si := gtpv2.GetSiIndication(gtpv2Value.si)
		dcnr := gtpv2.GetDcnr(gtpv2Value.dcnr)
		cause := gtpv2.GetCause(gtpv2Value.cause)
		linking := oi + si + dcnr + cause
		return linking
	}
	return ""
}

