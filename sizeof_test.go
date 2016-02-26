package main

import (
	"os"
	"strings"
	"testing"
)

func TestReaderSize(t *testing.T) {
	r := strings.NewReader("test")
	n, err := readerSize(r)
	if err != nil {
		t.Error(err)
	}
	if n != 4 {
		t.Error("Expected 4 received %d\n", n)
	}

	r = strings.NewReader("test\n")
	n, err = readerSize(r)
	if err != nil {
		t.Error(err)
	}
	if n != 4 {
		t.Error("Expected 4 recieved %d\n")
	}
	if n == 5 {
		t.Error("Failed to trim\n")
	}
}

type ConvertTest struct {
	n    int64
	u    int64
	e    float64
	name string
}

var convertTests = []ConvertTest{
	{0, units["B"], 0, "zero"},
	{1, units["b"], 8.0, "bits"},
	{16, units["B"], 16.0, "bytes"},
	{512, units["Kb"], 4.0, "kilobits"},
	{24576, units["KB"], 24.0, "kilobytes"},
	{2097152, units["Mb"], 16.0, "megabits"},
	{16777216, units["MB"], 16.0, "megabytes"},
	{2147483648, units["Gb"], 16.0, "gigabits"},
	{17179869184, units["GB"], 16.0, "gigabytes"},
	{2199023255552, units["Tb"], 16.0, "terabits"},
	{17592186044416, units["TB"], 16.0, "terabytes"},
}

func TestConvert(t *testing.T) {
	var result float64
	for _, test := range convertTests {
		result = convert(test.n, test.u)
		if result != test.e {
			t.Errorf("Conversion test for %s failed\nExpected %f recieved %f\n", test.name, test.e, result)
		}
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
