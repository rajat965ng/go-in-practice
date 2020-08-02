package main

import (
	"testing"
)

func TestName(t *testing.T) {
	a, _ := Names()

	if a!="rajat" {
		t.Error("The function is not returning proper values !!")
	}
}
