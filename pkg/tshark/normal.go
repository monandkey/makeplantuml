package tshark

import (
	"os/exec"
)

func (n NormalTshark) new() TsharkMethod {
	return NormalTshark{}
}

func (n NormalTshark) SetTsharkCommand() {
	n.tshark = getTsharkCommand()
}

func (n NormalTshark) CreateCommand() {
	n.cmd = []string{
		"-r", "test",
		"-t", "ad",
		"-T", "fields",
		"-E", "separator=,",
		"-E", "quote=d",
		"-e", "frame.number",
		"-e", "_ws.col.Time",
		"-e", "ip.src",
		"-e", "ipv6.src",
		"-e", "udp.srcport",
		"-e", "tcp.srcport",
		"-e", "sctp.srcport",
		"-e", "ip.dst",
		"-e", "ipv6.dst",
		"-e", "udp.dstport",
		"-e", "tcp.dstport",
		"-e", "sctp.dstport",
		"-e", "_ws.col.Protocol",
		"-e", "_ws.col.Info",
		"-e", "ip.len",
		"-e", "ipv6.plen",
	}
}

func (n NormalTshark) RunE() error {
	var err error
	n.out, err = exec.Command(n.tshark, n.cmd...).Output()

	if err != nil {
		return err
	}

	return nil
}

func (n NormalTshark) SetHeader() {
	return
}
