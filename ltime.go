package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	s := New()
	err := s.LTime(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

type Stat struct {
	Start time.Time
	Prev  time.Time
}

func New() *Stat {
	curr := time.Now()
	return &Stat{
		Start: curr,
		Prev:  curr,
	}
}

func (stat *Stat) LTime(r io.Reader, w io.Writer) error {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		curr := time.Now()
		fmt.Fprintf(w, "%s\t%s\t%s\n", curr.Sub(stat.Start), curr.Sub(stat.Prev), sc.Text())
		stat.Prev = curr
	}
	return sc.Err()
}
