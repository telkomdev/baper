package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/telkomdev/baper"
)

// Version number
const Version = "0.0.0"

func main() {

	var (
		showVersion bool
		interval    int
	)

	flag.IntVar(&interval, "interval", 1, "the interval between execution")
	flag.IntVar(&interval, "i", 1, "the interval between execution")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.BoolVar(&showVersion, "v", false, "show version")

	flag.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: ")
		fmt.Println("baper -[options]")
		fmt.Println()
		fmt.Println("-i | -interval (the interval between execution)")
		fmt.Println("-h | -help (show help)")
		fmt.Println("-v | -version (show version)")
		fmt.Println("---------------------------")
		fmt.Println()
	}

	flag.Parse()

	if showVersion {
		fmt.Printf("  baper version %s\n", Version)
		fmt.Println()
		os.Exit(0)
	}

	sender := baper.IOWriterSender{Writer: os.Stdout}
	collector := baper.New(sender, interval)

	kill := make(chan os.Signal, 1)
	// notify when user interrupt the process
	signal.Notify(kill, syscall.SIGINT, syscall.SIGTERM)

	go waitNotify(kill, collector)

	if err := collector.Collect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func waitNotify(kill chan os.Signal, collector *baper.Collector) {
	select {
	case <-kill:
		err := collector.Kill()
		if err != nil {
			fmt.Println("error kill process, ", err)
			os.Exit(1)
		}

		fmt.Println("kill process")
	}
}
