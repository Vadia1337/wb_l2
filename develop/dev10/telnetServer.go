package main

import "net"

func main() {
	listener, err := net.Listen("tcp", ":13")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			defer conn.Close()
			buffer := make([]byte, 1024)
			for {
				_, err := conn.Read(buffer)
				if err != nil {
					return
				}

				buffer = append([]byte("Сервер отвечает: "), buffer...)

				_, err = conn.Write(buffer)
				if err != nil {
					return
				}

				clear(buffer)
			}
		}()
	}
}
