package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	prev := time.Now()
	r := os.Stdin
	w := os.Stdout
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		curr := time.Now()
		fmt.Fprintf(w, "%s\t%s\t%s\n", curr.Sub(start), curr.Sub(prev), sc.Text())
		prev = curr
	}
	if sc.Err() != nil {
		log.Fatal(sc.Err())
	}
}
