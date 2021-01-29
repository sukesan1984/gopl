package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1(args []string) {
	for _, s := range args {
		fmt.Print(s)
		fmt.Print(" ")
	}
}

func echo2(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func main() {
	args := os.Args[1:]
	s1 := time.Now()
	echo1(args)
	a1 := time.Since(s1).Seconds()
	s2 := time.Now()
	echo2(args)
	a2 := time.Since(s2).Seconds()
	fmt.Printf("not join: %fs\n", a1)
	fmt.Printf("    join: %fs\n", a2)
}
