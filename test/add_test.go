package main

import (
	"testing"

	"github.com/sanjayyepuri/Akira/number"
)

func TestAdd(t *testing.T) {
	// t.Fatal("not implemented")
	var num1 float64 = 5
	var num2 float64 = -5

	num3 := num1 + num2

	if num3 != 0 {
		t.Errorf("5 + -5 = %g; want 0", num3)
	}

	var command = "5 - 5"
	out, err := number.CalculateCommand(command)
	if err != nil {
		t.Errorf("add command failed with %s", err)
	}

	if out != 0 {
		t.Errorf("5 - 5 = %g; want 0", out)
	}

}
