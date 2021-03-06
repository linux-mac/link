package main

import (
	"flag"
	"fmt"
	"github.com/funny/binary"
	"github.com/funny/link"
)

var (
	addr = flag.String("addr", "127.0.0.1:10010", "echo server address")
)

func main() {
	flag.Parse()

	session, err := link.Connect("tcp", *addr)
	if err != nil {
		panic(err)
	}

	go func() {
		var msg Message
		for {
			if err := session.Receive(msg); err != nil {
				break
			}
		}
	}()

	for {
		var msg Message
		if _, err := fmt.Scanf("%s\n", &msg); err != nil {
			break
		}
		session.Send(msg)
	}

	session.Close()
	println("bye")
}

type Message string

func (msg Message) Send(conn *binary.Writer) error {
	fmt.Printf("send: %s\n", msg)
	conn.WritePacket([]byte(msg), binary.SplitByUint16BE)
	return nil
}

func (msg Message) Receive(conn *binary.Reader) error {
	m := conn.ReadPacket(binary.SplitByUint16BE)
	fmt.Printf("recv: %s\n", m)
	return nil
}
