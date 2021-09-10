package tshark

// import (
// 	"regexp"
// 	"os/exec"
// )

// func (h HandsonTshark) new() TsharkMethod {
// 	return HandsonTshark{}
// }

// func (h HandsonTshark) SetTsharkCommand() {
// 	h.tshark = getTsharkCommand()
// }

// func (h HandsonTshark) CreateCommand() {
// 	h.cmd = []string{
// 		"-r", "test",
// 		"-t", "ad",
// 		"-T", "fields",
// 		"-E", "separator=,",
// 		"-E", "quote=d",
// 		"-e", "frame.number",
// 		"-e", "_ws.col.Time",
// 		"-e", "ip.src",
// 		"-e", "ipv6.src",
// 		"-e", "udp.srcport",
// 		"-e", "tcp.srcport",
// 		"-e", "sctp.srcport",
// 		"-e", "ip.dst",
// 		"-e", "ipv6.dst",
// 		"-e", "udp.dstport",
// 		"-e", "tcp.dstport",
// 		"-e", "sctp.dstport",
// 		"-e", "_ws.col.Protocol",
// 		"-e", "_ws.col.Info",
// 		"-e", "ip.len",
// 		"-e", "ipv6.plen",
// 		"-d", "tcp.port==29000-30000,http2",
// 	}
// }

// func (h HandsonTshark) RunE() error {
// 	var err error
// 	h.out, err = exec.Command(h.tshark, h.cmd...).Output()

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (h HandsonTshark) SetHeader() {
// 	lines := convertOutputResultIntoArray(string(h.out))

// 	for _, line := range lines {
// 		if regexp.MustCompile(",\r$").Match([]byte(line)) {
// 			line = regexp.MustCompile(",\r$").ReplaceAllString(line, ",\"\"")
// 		}

// 		if regexp.MustCompile(",$").Match([]byte(line)) {
// 			line = regexp.MustCompile(",$").ReplaceAllString(line, ",\"\"")
// 		}

// 		column := regexp.MustCompile("\".*?\"").FindAllStringSubmatch(line, -1)
// 		if len(column) == 0 {
// 			continue
// 		}

// 		tNumber := regexp.MustCompile("\"").ReplaceAllString(column[0][0], "")
// 		tTime := regexp.MustCompile("\"").ReplaceAllString(column[1][0], "")
// 		tSrcAddr := SetAddress(column[2][0], column[3][0], column[17][0], column[18][0])
// 		tSrcPort := SetPortAndCheckSum(column[4][0], column[5][0], column[6][0])
// 		tDstAddr := SetAddress(column[7][0], column[8][0], column[17][0], column[18][0])
// 		tDstPort := SetPortAndCheckSum(column[9][0], column[10][0], column[11][0])
// 		tProtocol := regexp.MustCompile("\"").ReplaceAllString(column[12][0], "")
// 		tMessage := SetMessage(column[13][0], tProtocol)
// 		tChecksum := SetPortAndCheckSum(column[14][0], column[15][0], column[16][0])

// 		if tMessage == "" {
// 			continue
// 		}

// 		th := TsharkHeader{
// 			Number:   tNumber,
// 			Time:     tTime,
// 			SrcAddr:  tSrcAddr,
// 			SrcPort:  tSrcPort,
// 			DstAddr:  tDstAddr,
// 			DstPort:  tDstPort,
// 			Protocol: tProtocol,
// 			Message:  tMessage,
// 			Checksum: tChecksum,
// 		}
// 		h.res = append(h.res, th)
// 	}
// }
