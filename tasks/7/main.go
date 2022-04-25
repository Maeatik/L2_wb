package main

import (
	"fmt"
	"time"
)

func main() {
	//объявим функцию OR
	var or func(channels ...<-chan interface{}) <-chan interface{}

	or = func(channels ...<-chan interface{}) <-chan interface{} {
		//по количеству заданных каналов, если только один канал - возвращаем его. Если ничего не задано - выходим
		fmt.Println(len(channels))
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}
		//создаем канал интерфейсов
		orDone := make(chan interface{})
		//создаем горутину, которая будет ждать сообщения из незаблокированных каналов
		go func() {
			//отложено закрываем канал интерфейсов
			defer close(orDone)

			//если указано 2 ждем какой из них первый закроется
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			//если их больше двух - объединяем и когда закроется один канал - закроются и остальные вместе с ним
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		//возвращаем объединенные каналы
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(5*time.Second),
		sig(100*time.Hour),
		sig(2*time.Minute),
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(3*time.Minute),
		sig(2*time.Second),
		sig(100*time.Hour),
		sig(2*time.Minute),
		sig(2*time.Hour),
		sig(5*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))

}
