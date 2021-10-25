package main

/**
	0. read file / chal sha256
	1. send file name
	2. send file sha256
	3. send file
	4. send eof
**/

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/holylong/ican/lib/protocol"
	"github.com/holylong/ican/lib/util"
)

func handleConn(conn net.Conn) {
	//conn.Write([]byte("hello server"))
	fmt.Println("handle conn")
	for i := 0; i < 10000; i++ {
		fmt.Printf("send message %d\n", i)
		//msg := "{\"status\":\"0000\",\"message\":\"success\",\"data\":{\"title\":{\"id\":\"001\",\"name\":\"白菜\"},\"content\":[{\"id\":\"001\",\"value\":\"你好白菜\"},{\"id\":\"002\",\"value\":\"你好萝卜\"}]}}"
		msg := "words"
		conn.Write(protocol.Packet([]byte(msg)))
	}
}

func ckError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal Error:", err.Error())
		os.Exit(-1)
	}
}

func startClient(address string) {

	addr := net.ParseIP(address)
	if addr == nil {
		fmt.Println("Invalid Address")
	} else {
		fmt.Println("address ", addr.String(), "is ok")
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	ckError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	ckError(err)
	defer conn.Close()
	fmt.Println("client app start ok")

	go handleConn(conn)

	for {
		time.Sleep(1 * 1e8)
	}

	os.Exit(0)
}

func main() {
	fmt.Println("ican client start")

	if len(os.Args) != 2 {
		fmt.Println("ican server start error please input argument")
		os.Exit(1)
	}

	fmt.Printf("%s\n", os.Args[1])

	address := os.Args[1]

	//startClient(address)
	util.DoSha256(address)
}
