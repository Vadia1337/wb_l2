package main

import (
	"testing"
)

func Test_NTPtimeCorrectTimeServer(t *testing.T) {
	err := time()
	if err != nil {
		t.Error("Error main.time")
	}
}

func Test_NTPtimeInCorrectTimeServer(t *testing.T) {
	timeServer = "dedm!"
	err := time()
	if err == nil {
		t.Error("Error. time func did not report an error of an incorrect server")
	}
}
