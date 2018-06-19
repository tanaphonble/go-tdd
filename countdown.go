package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// Sleeper -
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper -
type ConfigurableSleeper struct {
	duration time.Duration
}

// Sleep -
func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

// Countdown -
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
