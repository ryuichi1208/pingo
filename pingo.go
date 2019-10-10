package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type walker interface {
	walk(func(faild interface{}))
}

type Ping struct {
	icmpID  uint16
	icmpSeq uint16
}

func NewPing(icmpID uint16) *Ping {
	return &Ping{icmpID: icmpID, icmpSeq: uint16(1)}
}

func (ping *Ping) Send(hostname string) error {
	conn, err := net.Dial("ip4:icmp", hostname)
	if err != nil {
		return err
	}
	rawMsg, err := ping.buildEchoMessage()
	if err != nil {
		return err
	}

	_, err = conn.Write(rawMsg)
	return err
}

func (ping *Ping) buildEchoMessage() ([]byte, error) {
	curr := time.Now()
	data, err := curr.GobEncode()
	if err != nil {
		return nil, err
	}

	msg := ICMPEchoMessage{
		ICMPMessage: ICMPMessage{
			icmpType: 0x8,
			icmpCode: 0x0,
			checkSum: 0x0,
			data:     data,
		},
		imcpID:  ping.icmpID,
		icmpSeq: ping.icmpSeq,
	}
	ping.icmpSeq++
	return msg.Pack(), nil

}

func main() {
	hostname := "golang.org"
	if len(os.Args) >= 2 {
		hostname = os.Args[1]
	}

	pid := uint16(os.Getegid())
	ping := NewPing(pid)

	fmt.Printf("%d %d %s", pid, ping, hostname)

	go func() {
		for {
			if err := ping.Send(hostname); err != nil {
				log.Printf("faild to send icmp message : %v", err)
			}
		}
	}()
	
}
