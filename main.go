package main

import (
	"github.com/pkg/errors"
	"log"
)

func main() { log.Println(errors.New("error")) }
