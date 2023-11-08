package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/
var (
	timeout int
)

func init() {
	flag.IntVar(&timeout, "timeout", 10, "set timeout to connect")
}

func main() {
	flag.Parse()
	err := telnet(flag.Arg(0) + ":" + flag.Arg(1))
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func telnet(addr string) error {
	timeOut := time.Duration(timeout) * time.Second
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()

	conn, err := net.DialTimeout("tcp", addr, timeOut)
	if err != nil {
		return err
	}
	defer conn.Close()

	go func() { // рутинка для постоянной проверки соединения по таймауту
		for {
			_, err = conn.Read([]byte{})
			if err != nil {
				stop()
				return
			}
			time.Sleep(timeOut)
		}
	}()

	go func() {
		io.Copy(os.Stdout, conn)
	}()

	go func() {
		io.Copy(conn, os.Stdin)
	}()

	<-ctx.Done()
	return nil
}
