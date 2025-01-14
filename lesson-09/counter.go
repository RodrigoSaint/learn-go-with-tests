package main

import (
	"fmt"
	"io"
	"iter"
	"math/rand"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

func countDownFrom(count int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := count; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func Countdown(writter io.Writer, sleeper Sleeper) {
	for i := range countDownFrom(3) {
		fmt.Fprintln(writter, i)
		sleeper.Sleep()
	}
}

type DefaultSleeper struct{}

func (sleeper *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func randomGenerator(amount int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for range amount {
			randomText := fmt.Sprintf("Random Number: %d", rand.Int())
			if !yield(randomText) {
				return
			}
		}
	}
}

func main() {
	for randomValue := range randomGenerator(10) {
		fmt.Println(randomValue)
	}
	sleeper := DefaultSleeper{}
	Countdown(os.Stdout, &sleeper)
}
