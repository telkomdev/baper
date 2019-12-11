package baper

import (
	"fmt"
	"regexp"
	"strconv"
)

// CPUStat struct
type CPUStat struct {
	Disk        Disk        `json:"disk"`
	CPU         CPU         `json:"cpu"`
	LoadAverage LoadAverage `json:"loadAverage"`
}

// Disk struct
type Disk struct {
	KBPerTime   float64 `json:"kbPerTime"`
	TPS         float64 `json:"tps"`
	MBPerSecond float64 `json:"mbPerSecond"`
}

// CPU struct
type CPU struct {
	User   float64 `json:"user"`
	System float64 `json:"system"`
	Idle   float64 `json:"idle"`
}

// LoadAverage struct
type LoadAverage struct {
	OneMinute     float64 `json:"oneMinute"`
	FiveMinute    float64 `json:"fiveMinute"`
	FifteenMinute float64 `json:"fifteenMinute"`
}

// trim function
func trim(text string) string {
	re := regexp.MustCompile(`^\s+|\s+$`)
	result := re.ReplaceAllString(text, "")
	return result
}

// split function
func split(text string) []string {
	re := regexp.MustCompile(`\s+|\s+$`)
	result := re.Split(text, -1)
	return result
}

// isNumber function
func isNumber(v string) bool {
	_, err := strconv.ParseFloat(v, 64)
	return err == nil
}

// parse function
func parse(text string) *CPUStat {
	trimmed := trim(text)
	splited := split(trimmed)

	var datas []float64
	for _, v := range splited {
		if !isNumber(v) {
			continue
		}

		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println("error parsing data : ", err.Error())
		}

		datas = append(datas, f)
	}
	if len(datas) <= 0 {
		return nil
	}

	return &CPUStat{
		Disk:        Disk{KBPerTime: datas[0], TPS: datas[1], MBPerSecond: datas[2]},
		CPU:         CPU{User: datas[3], System: datas[4], Idle: datas[5]},
		LoadAverage: LoadAverage{OneMinute: datas[6], FiveMinute: datas[7], FifteenMinute: datas[8]},
	}
}
