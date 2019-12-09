package main

import (
	"flag"
	"fmt"
	"os"

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

	sender := baper.StdoutSender{}

	if err := baper.Collect(sender, interval); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
