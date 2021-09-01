package makeplantuml

import (
	"os"
	"fmt"
	"runtime"
	"os/exec"
)

func RunTshark() tsharkHeaders {
	var cmd string
	switch(runtime.GOOS) {
		case "windows":
			cmd = "C:/Program Files/Wireshark-3.2.3/tshark.exe"
		case "linux":
			cmd = "tshark"
		default:
			fmt.Println("Your OS not support.")
			os.Exit(0)
	}

	out, err := exec.Command(cmd,
		"-r", "./sample/440110000001519.pcap",
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
		"-e", "udp.checksum",
		"-e", "sctp.checksum",
		"-e", "tcp.checksum",
	).Output()

	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(0)
	}

	var fmtOut tsharkHeaders
	return fmtOut.setHeader(string(out))
}
