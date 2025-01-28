package util

import (
	"log"
)

func Log(a ...any) {
	log.Println(a...)
}

func Logf(fmt string, a ...any) {
	log.Printf(fmt, a)
}
