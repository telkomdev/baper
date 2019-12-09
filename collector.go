package baper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strconv"
)

// Collect function will collect information from iostat
func Collect(sender Sender, interval int) error {
	const IOStatCommand = "iostat"

	intervalStr := strconv.Itoa(interval)

	cmd := exec.Command(IOStatCommand, "-w", intervalStr)

	outPipe, _ := cmd.StdoutPipe()
	err := cmd.Start()
	if err != nil {
		return err
	}

	fmt.Println("process running on PID = ", cmd.Process.Pid)

	// time.AfterFunc(10*time.Second, func() {
	// 	err = cmd.Process.Kill()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("kill process")
	// })

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

		cpuStat := Parse(line)
		if cpuStat != nil {
			payload, err := json.Marshal(cpuStat)
			if err != nil {
				fmt.Println("error marshal data : ", err.Error())
			}

			if err := sender.Send(payload); err != nil {
				return err
			}
		}
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
