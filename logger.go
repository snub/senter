package senter

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "senter: ", log.LstdFlags)
