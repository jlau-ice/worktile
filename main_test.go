package main

import (
	"github.com/jlau-ice/gotils/log"
	"github.com/jlau-ice/gotils/str"
	"testing"
)

func TestGotils(t *testing.T) {
	log.Warn("dsadasdas")
	log.Error("dasdasdasd")
	str.RandomString(15)
}
