package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/RengokuryuuHonokaCrimsonFlame/PRR-Lab1/constantes"
	"io"
	"log"
	"net"
	"os"
	"runtime"

	"golang.org/x/net/ipv4"

	_ "github.com/RengokuryuuHonokaCrimsonFlame/PRR-Lab1/message"
	_ "github.com/RengokuryuuHonokaCrimsonFlame/PRR-Lab1/constantes"
)

// debut, OMIT

func main() {
	go udpReader()
	conn, err := net.Dial("udp", constantes.MulticastAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(conn, os.Stdin)
}

// milieu, OMIT
func udpReader() {
	conn, err := net.ListenPacket("udp", constantes.MulticastAddr) // listen on port
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	p := ipv4.NewPacketConn(conn) // convert to ipv4 packetConn
	addr, err := net.ResolveUDPAddr("udp", constantes.MulticastAddr)
	if err != nil {
		log.Fatal(err)
	}
	var interf *net.Interface
	if runtime.GOOS == "darwin" {
		interf, _ = net.InterfaceByName("en0")
	}

	if err = p.JoinGroup(interf, addr); err != nil { // listen on ip multicast
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFrom(buf) // n, _, addr, err := p.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}
		s := bufio.NewScanner(bytes.NewReader(buf[0:n]))
		for s.Scan() {
			fmt.Printf("%s from %v\n", s.Text(), addr)
			mess := message.Creates.Text()
		}
	}
}

// fin, OMIT
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
