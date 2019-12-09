package baper

import (
	"bytes"
	"os"
)

// Sender interface
type Sender interface {
	Send([]byte) error
}

// StdoutSender struct
// you can create your own implementation
type StdoutSender struct {
}

func (sender StdoutSender) Send(line []byte) error {
	reader := bytes.NewReader(line)
	_, err := reader.WriteTo(os.Stdout)

	if err != nil {
		return err
	}

	return nil
}
