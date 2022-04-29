package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//аргумент таймаута на подключение к серверу (по умолчанию 10 секунд)
	timeoutArg := flag.Int("timeout", 10, "timeout")
	flag.Parse()

	//если не были указаны хост и порт кидаем ошибку и сообщаем о ней
	if flag.NArg() < 2 {
		log.Printf("usage %s host port\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	timeout := time.Duration(*timeoutArg) * time.Second
	//создаем адрес из хоста и порта
	addr := net.JoinHostPort(host, port)

	fmt.Println(addr)

	conn, err := net.DialTimeout("tcp", addr, timeout)

	if err != nil {
		time.Sleep(timeout)
	}

	fmt.Println(conn.LocalAddr())
	//создадим каналы для того, чтобы закрывать соединение\программу
	osSignals := make(chan os.Signal, 1)
	listenErr := make(chan error, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	go req(listenErr, osSignals)

	//если был получен сигнал в канал для сигналов - закрываем соединение
	select {
	case <-osSignals:
		conn.Close()
	//если в канал ошибок пришла ошибка - заканчиваем программу
	case err = <-listenErr:
		if err != nil {
			return
		}
	}
}

func req(listenErr chan<- error, osSignals chan<- os.Signal) {
	for {
		//читаем из консоли
		reader := bufio.NewReader(os.Stdin)
		out, err := reader.ReadString('\n')
		//если читать больше не получается отправляем сигнал о завершении
		if err != nil {
			if err == io.EOF {
				osSignals <- syscall.SIGQUIT
				return
			}
			//если была ошибка - передаем в канал ошибку
			listenErr <- err
		}
		//печатаем полученную строку
		fmt.Fprintf(os.Stdout, out+"\n")
	}
}
