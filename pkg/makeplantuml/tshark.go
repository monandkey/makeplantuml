package makeplantuml

import (
	"os"
	"fmt"
	"strings"
	"runtime"
	"os/exec"
	"encoding/csv"
)

type tsharkHeader struct {
	number   string
	time     string
	srcAddr  ipAddr
	srcPort  portNumber
	dstAddr  ipAddr
	dstPort  portNumber
	protocol string
	message  string
	checksum checksumProtocol
}

type ipAddr struct {
	v4Addr string
	v6Addr string
}

type portNumber struct {
	udpPort  string
	tcpPort  string
	sctpPort string
}

type checksumProtocol struct {
	udpChecksum  string
	tcpChecksum  string
	sctpChecksum string
}

func (t tsharkHeader) setHeader(out string) {
	arr := strings.Split(out, "\n")
	for _, v := range arr {
		arr2 := strings.Split(v, ",")
		t.number                = arr2[0]
		t.time                  = arr2[1]
		t.srcAddr.v4Addr        = arr2[2]
		t.srcAddr.v6Addr        = arr2[3]
		t.srcPort.udpPort       = arr2[4]
		t.srcPort.tcpPort       = arr2[5]
		t.srcPort.sctpPort      = arr2[6]
		t.dstAddr.v4Addr        = arr2[7]
		t.dstAddr.v6Addr        = arr2[8]
		t.dstPort.udpPort       = arr2[9]
		t.dstPort.tcpPort       = arr2[10]
		t.dstPort.sctpPort      = arr2[11]
		t.protocol              = arr2[12]
		t.message               = arr2[13]
		t.checksum.udpChecksum  = arr2[14]
		t.checksum.tcpChecksum  = arr2[15]
		t.checksum.sctpChecksum = arr2[16]
	}
}

func createCSV(out string) (string, error) {
	csvFile := ".tmp.csv"
	file, err := os.OpenFile(csvFile, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return csvFile, err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{out})
	writer.Flush()

	return csvFile, nil
}

func RunTshark() {
	var cmd string
	switch(runtime.GOOS) {
		case "windows":
			cmd = "C:/Program Files/Wireshark-3.2.3/tshark.exe"
		case "linux":
			cmd = "tshark"
		default:
			fmt.Println("Your OS not support.")
			return
	}

	out, err := exec.Command(cmd,
		"-r", "./sample/3g_4g_nokia.pcap",
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
		return
	}

	csv, err := createCSV(string(out))
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println(csv)
}
