package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.
*/

// слияние каналов (fan-in chan)
func main() {

	or := func(channels ...<-chan interface{}) <-chan interface{} {
		wg := sync.WaitGroup{}
		singleChan := make(chan interface{})

		wg.Add(len(channels))

		doneChan := func(ch <-chan interface{}) {
			for v := range ch {
				singleChan <- v
			}
			wg.Done()
		}

		for _, ch := range channels {
			go doneChan(ch)
		}

		go func() {
			wg.Wait()
			close(singleChan)
		}()

		return singleChan
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
		//sig(2*time.Hour),
		//sig(5*time.Minute),
		//sig(1*time.Hour),
		//sig(1*time.Minute),
		sig(1*time.Second),
		sig(5*time.Second),
		sig(10*time.Second),
		sig(15*time.Second),
	)

	fmt.Printf("fone after %v", time.Since(start))
}
