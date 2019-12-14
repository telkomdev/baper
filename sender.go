package baper

import (
	"bytes"
	"io"
)

// Sender interface
type Sender interface {
	Send([]byte) error
}

// IOWriterSender struct
// you can create your own implementation
type IOWriterSender struct {
	Writer io.Writer
}

func (sender IOWriterSender) Send(line []byte) error {
	reader := bytes.NewReader(line)
	_, err := reader.WriteTo(sender.Writer)

	if err != nil {
		return err
	}

	return nil
}
