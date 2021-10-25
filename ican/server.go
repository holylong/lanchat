//#
//#
//# ican主程序
//#	监听端口->接收文件基本信息->接收文件->返回成功
//# 监听端口：10093
//# 文件基本信息：文件名，md5支持，文件大小等
//# 接收文件：网络库接受文件
//# 验证文件是否正确，返回成功或者失败，失败根据用户选择是否重传或者断开
package main

import (
	"fmt"
	"net"
	"os"

	"github.com/holylong/ican/lib/protocol"
)

func headinfo(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			Log(string(data))
		}
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}

func handleConn(conn net.Conn) {
	fmt.Printf("client:%s connet ok\n", conn.RemoteAddr().String())
	defer conn.Close()
	//buffer := make([]byte, 1024)
	//n, err := conn.Read(buffer)
	//ckError(err)
	//fmt.Println("recv:", string(buffer[:n]))
	//timestramp := time.Now().String()
	//conn.Write([]byte(timestramp))

	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0)

	//声明一个管道用于接收解包的数据
	readerChannel := make(chan []byte, 16)
	go headinfo(readerChannel)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		} else {
			fmt.Println("recv:", n)
		}

		tmpBuffer = protocol.Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ican server start error please input argument")
		os.Exit(1)
	}

	fmt.Printf("listen:%s\n", os.Args[1])

	port := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	ckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	ckError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
	os.Exit(0)
}

func ckError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal Error:", err.Error())
		os.Exit(-1)
	}
}
