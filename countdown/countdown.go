package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type ConfigurationSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationslept time.Duration
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
	return
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyTime) SpyTimeSleep(duration time.Duration) {
	s.durationslept = duration
}

func (c *ConfigurationSleeper) Sleep() {
	c.sleep(c.duration)
}

const write = "write"
const sleep = "sleep"

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintf(out, strconv.Itoa(i)+"\n")
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}

func main() {
	sleeper := &ConfigurationSleeper{5 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
