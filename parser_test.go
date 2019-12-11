package baper

import (
	"testing"
)

func TestIsNumber(t *testing.T) {

	t.Run("should return true with valid number string params", func(t *testing.T) {
		res := isNumber("5")
		if !res {
			t.Error("error, should return true")
		}
	})

	t.Run("should return false with invalid number string params", func(t *testing.T) {
		res := isNumber("v")
		if res == true {
			t.Error("error, should return false")
		}
	})
}

func TestTrim(t *testing.T) {

	input := " hello world   "
	expected := "hello world"

	t.Run("should success trim string", func(t *testing.T) {
		res := trim(input)

		if res != expected {
			t.Error("error trim string")
		}
	})
}

func TestSplit(t *testing.T) {

	input := " hello world   "
	expected := []string{"hello", "world"}

	t.Run("should success split string", func(t *testing.T) {
		s := trim(input)
		ss := split(s)

		if expected[0] != ss[0] {
			t.Error("error split string")
		}
	})
}

func TestParse(t *testing.T) {
	input := "  318.44    9  2.79   4  1 95  3.02 2.01 1.73"
	var expectedKBPerTime float64 = 318.44

	t.Run("should success parse string to stat", func(t *testing.T) {
		cpuStat := parse(input)
		if cpuStat == nil {
			t.Error("error parse input")
		}

		if cpuStat.Disk.KBPerTime != expectedKBPerTime {
			t.Error("error disk KB/t should equal to ", expectedKBPerTime)
		}
	})
}
