/*
Необходимо код примера 1 изменить так, чтобы tcp-сервер обрабатывал подключения параллельно
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}

	log.Println("завершение работы")
}

func handleConn(c net.Conn) {
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			log.Println(err)
		}
	}(c)

	var i int
	for {
		_, err := io.WriteString(c, fmt.Sprintf("%d\n", i))
		if err != nil {
			log.Println(err)
			return
		}
		i++
		time.Sleep(time.Second)
	}
}
