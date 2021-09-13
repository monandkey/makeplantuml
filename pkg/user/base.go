package user

import (
	"fmt"
	"regexp"
	"os/exec"
)

func (b *baseUser) SetCmd() {
	b.Cmd = b.GetTsharkCommand()
}

func (b *baseUser) SetArgs(fileName string) {
	b.Args = []string{
		"-r", fileName,
		"-t", "ad",
		"-T", "fields",
		"-E", "separator=,",
		"-E", "quote=d",
		"-e", "frame.number",     //  0
		"-e", "_ws.col.Time",     //  1
		"-e", "ip.src",           //  2
		"-e", "ipv6.src",         //  3
		"-e", "udp.srcport",      //  4
		"-e", "tcp.srcport",      //  5
		"-e", "sctp.srcport",     //  6
		"-e", "ip.dst",           //  7
		"-e", "ipv6.dst",         //  8
		"-e", "udp.dstport",      //  9
		"-e", "tcp.dstport",      // 10
		"-e", "sctp.dstport",     // 11
		"-e", "_ws.col.Protocol", // 12
		"-e", "_ws.col.Info",     // 13
		"-e", "ip.len",           // 14
		"-e", "ipv6.plen",        // 15
	}
}

func (b *baseUser) RunE() error {
	var err error
	b.Out, err = exec.Command(b.Cmd, b.Args...).Output()

	if err != nil {
		return err
	}
	return nil
}

func (b *baseUser) Parse() {
	lines := b.ConvertOutputResultIntoArray(string(b.Out))

	for _, line := range lines {
		line = b.CleanLineStrings(line)

		column := regexp.MustCompile("\".*?\"").FindAllStringSubmatch(line, -1)
		if len(column) == 0 {
			continue
		}

		tmpNumber := regexp.MustCompile("\"").ReplaceAllString(column[0][0], "")
		tmpTime := regexp.MustCompile("\"").ReplaceAllString(column[1][0], "")
		tmpSrcAddr := b.SetAddress(column[2][0], column[3][0], column[14][0], column[15][0])
		tmpSrcPort := b.SetPortAndCheckSum(column[4][0], column[5][0], column[6][0])
		tmpDstAddr := b.SetAddress(column[7][0], column[8][0], column[14][0], column[15][0])
		tmpDstPort := b.SetPortAndCheckSum(column[9][0], column[10][0], column[11][0])
		tmpProtocol := regexp.MustCompile("\"").ReplaceAllString(column[12][0], "")
		tmpMessage := b.SetMessage(column[13][0], tmpProtocol)

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
		b.Header = append(b.Header, []map[string]string{tmph}...)
	}
}

func (b *baseUser) NameResE(fileName string) error {
	if err := b.NameResolution(b.Header, fileName); err != nil {
		return err
	}
	return nil
}

func (b *baseUser) Display() {
	fmt.Println(b.Header)
}

func (b *baseUser) CreateE(title string) error {
	if err := b.CreateTemplate(title); err != nil {
		return err
	}
	return nil
}

func (b *baseUser) WritingE(timestamp bool) error {
	if err := b.WriteUml(b.Header, timestamp); err != nil {
		return err
	}
	return nil
}

func (b *baseUser) RenderingE(fileName string) error {
	if err := b.RenderingUml(fileName); err != nil {
		return err
	}
	return nil
}
