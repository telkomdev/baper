package main

import (
	"fmt"
	"os"

	"github.com/telkomdev/baper"
)

func main() {
	sender := baper.StdoutSender{}

	if err := baper.Collect(sender, 2); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
