package baper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strconv"
)

// Collector struct
type Collector struct {
	cmd    *exec.Cmd
	sender Sender
}

// New function, Collector's constructor
func New(sender Sender, interval int) *Collector {
	const IOStatCommand = "iostat"

	intervalStr := strconv.Itoa(interval)
	cmd := exec.Command(IOStatCommand, "-w", intervalStr)
	return &Collector{cmd: cmd, sender: sender}
}

// Collect function will collect information from iostat
func (collector *Collector) Collect() error {

	outPipe, _ := collector.cmd.StdoutPipe()
	err := collector.cmd.Start()
	if err != nil {
		return err
	}

	fmt.Println("process running on PID = ", collector.cmd.Process.Pid)

	reader := bufio.NewReader(outPipe)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		cpuStat := parse(line)
		if cpuStat != nil {
			payload, err := json.Marshal(cpuStat)
			if err != nil {
				fmt.Println("error marshal data : ", err.Error())
			}

			if err := collector.sender.Send(payload); err != nil {
				return err
			}
		}
	}

	err = collector.cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

// Kill function will kill cmd process
func (collector *Collector) Kill() error {
	// time.AfterFunc(10*time.Second, func() {
	// 	err = cmd.Process.Kill()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("kill process")
	// })
	return collector.cmd.Process.Kill()
}
