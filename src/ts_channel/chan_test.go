package ts_chan

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Rules of thumb that should make things feel much more straightforward.
//
//  1. prefer using formal arguments for the channels you pass to go-routines
//     instead of accessing channels in global scope.
//     You can get more compiler checking this way, and better modularity too.
//
//  2. avoid both reading and writing on the same channel
//     in a particular go-routine (including the 'main' one).
//     Otherwise, deadlock is a much greater risk.
//
// ref: https://stackoverflow.com/questions/15715605/multiple-goroutines-listening-on-one-channel
//
// It creates the five go-routines writing to a single channel, each one writing five times.
// The main go-routine reads all twenty five messages -
// you may notice that the order they appear in is often not sequential (i.e. the concurrency is evident).
// This example demonstrates a feature of Go channels:
// it is possible to have multiple writers sharing one channel;
// Go will interleave the messages automatically.
func TestMultiWriterOneReader(t *testing.T) {
	c := make(chan string)

	for i := 1; i <= 5; i++ {
		go func(i int, co chan<- string) {
			for j := 1; j <= 5; j++ {
				co <- fmt.Sprintf("hi from %d.%d", i, j)
			}
		}(i, c)
	}

	for i := 1; i <= 25; i++ {
		fmt.Println(<-c)
	}
}

// 利用 channel 多读一写，实际上表现为竞态
// In both examples, no buffering was needed. It is generally a good principle to view buffering as a performance enhancer only.
// If your program does not deadlock without buffers, it won't deadlock with buffers either (but the converse is not always true).
// So, as another rule of thumb, start without buffering then add it later as needed.
func TestMultiReaderOneWriter(t *testing.T) {
	c := make(chan int)
	var w sync.WaitGroup
	w.Add(5)

	for i := 1; i <= 5; i++ {
		go func(i int, ci <-chan int) {
			j := 1
			for v := range ci {
				time.Sleep(time.Millisecond)
				fmt.Printf("%d.%d got %d\n", i, j, v)
				j += 1
			}
			w.Done()
		}(i, c)
	}

	for i := 1; i <= 25; i++ {
		c <- i
	}
	close(c)
	w.Wait()
}

// 空 chan 是否可读？
// 结果：可读不报错，只是会阻塞。
func TestNilChangReading(t *testing.T) {
	fmt.Printf("%s %d\n", "hello", 42)
	var a <-chan struct{} = nil
	fmt.Println("come in")

	go func() {
		a := 0
		b := `-\|/`
		for {
			fmt.Printf("\r%c", b[a])
			a += 1
			a %= 4
			time.Sleep(250 * time.Millisecond)
		}
	}()

	if b, ok := <-a; ok {
		fmt.Println("OK", b)
	} else {
		fmt.Println("NOT OK", b)
	}
	fmt.Println("go out")
}
