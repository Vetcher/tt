package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"time"
)

const (
	Version = "1.0"
	Welcome = "Time-tracker " + Version
)

func main() {
	fmt.Println(Welcome)
	s := make(chan os.Signal)
	signal.Notify(s, os.Interrupt, os.Kill)
	last := time.Now()
	start := last
	fmt.Print(logTime(last))
	go func() {
		_ = <-s
		fmt.Println()
		fmt.Print(logTime(last))
		fmt.Printf("\nTracker works for %s\n", since(start))
		os.Exit(0)
	}()
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		fmt.Print(logTime(last))
		last = time.Now()
	}
}

func logTime(last time.Time) string {
	return fmt.Sprintf("%s | %s ", time.Now().Format("15:04:05"), since(last))
}

func since(last time.Time) string {
	now := time.Now()
	since := now.Sub(last)
	sinceH := truncateDuration(since, time.Hour)
	sinceM := truncateDuration(since-sinceH, time.Minute)
	sinceS := truncateDuration(since-sinceH-sinceM, time.Second)
	return fmt.Sprintf("%dh %dm %ds",
		trDurToInt(sinceH, time.Hour),
		trDurToInt(sinceM, time.Minute),
		trDurToInt(sinceS, time.Second),
	)
}

func truncateDuration(d1, d2 time.Duration) time.Duration {
	return d1 - d1%d2
}

func trDurToInt(d1, d2 time.Duration) int64 {
	return (d1 - d1%d2).Nanoseconds() / d2.Nanoseconds()
}
