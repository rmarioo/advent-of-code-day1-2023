package main

import "testing"

func TestSingleNumberInString(t *testing.T) {

	assertEquals(DigitsIn("2dfgd"), 22, t)
}

func TestTwoNumbersInString(t *testing.T) {

	assertEquals(DigitsIn("dfg3vfd4g"), 34, t)
}

func TestMoreThanTwoNumbersInString(t *testing.T) {

	assertEquals(DigitsIn("dfg3vfd4g8sdf7sd"), 37, t)
}

func TestTrebuchet(t *testing.T) {

	linesSupplier := StubLinesSupplier{lines: []string{"1g2", "5h6"}}

	assertEquals(Trebuchet(linesSupplier), 12+56, t)
}

func TestTrebuchetAcceptance(t *testing.T) {

	linesSupplier := FromFileLinesSupplier{fileName: "input.txt"}

	assertEquals(Trebuchet(linesSupplier), 55130, t)
}

type StubLinesSupplier struct {
	lines []string
}

func (f StubLinesSupplier) Lines() []string {
	return f.lines
}

func assertEquals(res int, expected int, t *testing.T) {
	if res != expected {
		t.Errorf("expected %d got %d", expected, res)
	}
}
