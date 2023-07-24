package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(15, 15)

	if total != 30 {
		t.Errorf("Wrong result. Actual %d, expected %d.", total, 30)
	}
}
