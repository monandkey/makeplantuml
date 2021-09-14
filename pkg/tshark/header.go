package tshark

import (
	"regexp"
)

func (t TsharkArgs) ParserOutput(kinds string) []map[string]string {
	switch(kinds) {
		case "normal":
			return t.ParsreOutputNormal()
		
		case "annotation":
			return t.ParserOutputAnnotation()
	}
	return t.Header
}

func (t TsharkArgs) ParsreOutputNormal() []map[string]string {
	lines := t.ConvertOutputResultIntoArray(string(t.Out))

	for _, line := range lines {
		line = t.CleanLineStrings(line)

		column := regexp.MustCompile("\".*?\"").FindAllStringSubmatch(line, -1)
		if len(column) == 0 {
			continue
		}

		tmpNumber := regexp.MustCompile("\"").ReplaceAllString(column[0][0], "")
		tmpTime := regexp.MustCompile("\"").ReplaceAllString(column[1][0], "")
		tmpSrcAddr := t.SetAddress(column[2][0], column[3][0], column[14][0], column[15][0])
		tmpSrcPort := t.SetPortAndCheckSum(column[4][0], column[5][0], column[6][0])
		tmpDstAddr := t.SetAddress(column[7][0], column[8][0], column[14][0], column[15][0])
		tmpDstPort := t.SetPortAndCheckSum(column[9][0], column[10][0], column[11][0])
		tmpProtocol := regexp.MustCompile("\"").ReplaceAllString(column[12][0], "")
		tmpMessage := t.SetMessage(column[13][0], tmpProtocol)

		tmph := map[string]string{
			"number":   tmpNumber,
			"time":     tmpTime,
			"srcAddr":  tmpSrcAddr,
			"srcPort":  tmpSrcPort,
			"dstAddr":  tmpDstAddr,
			"dstPort":  tmpDstPort,
			"protocol": tmpProtocol,
			"message":  tmpMessage,
		}
		t.Header = append(t.Header, []map[string]string{tmph}...)
	}
	return t.Header
}

func (t TsharkArgs) ParserOutputAnnotation() []map[string]string {
	lines := t.ConvertOutputResultIntoArray(string(t.Out))

	for _, line := range lines {
		line = t.CleanLineStrings(line)

		column := regexp.MustCompile("\".*?\"").FindAllStringSubmatch(line, -1)
		if len(column) == 0 {
			continue
		}

		tmpNumber := regexp.MustCompile("\"").ReplaceAllString(column[0][0], "")
		tmpTime := regexp.MustCompile("\"").ReplaceAllString(column[1][0], "")
		tmpSrcAddr := t.SetAddress(column[2][0], column[3][0], column[14][0], column[15][0])
		tmpSrcPort := t.SetPortAndCheckSum(column[4][0], column[5][0], column[6][0])
		tmpDstAddr := t.SetAddress(column[7][0], column[8][0], column[14][0], column[15][0])
		tmpDstPort := t.SetPortAndCheckSum(column[9][0], column[10][0], column[11][0])
		tmpProtocol := regexp.MustCompile("\"").ReplaceAllString(column[12][0], "")
		tmpMessage := t.SetMessage(column[13][0], tmpProtocol)
		tmpAnnotation := t.SetAnnotation(column)

		tmph := map[string]string{
			"number":     tmpNumber,
			"time":       tmpTime,
			"srcAddr":    tmpSrcAddr,
			"srcPort":    tmpSrcPort,
			"dstAddr":    tmpDstAddr,
			"dstPort":    tmpDstPort,
			"protocol":   tmpProtocol,
			"message":    tmpMessage,
			"annotation": tmpAnnotation,
		}
		t.Header = append(t.Header, []map[string]string{tmph}...)
	}
	return t.Header
}


