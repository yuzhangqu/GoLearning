package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func f1() {
	fmt.Println(strings.Join(os.Args, " "))
}

func f2() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
}

func echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// TODO: 测试f1_1()和echo的性能差异
func f3() {
}

type IndexCount struct {
	n   int
	idx int
}

func f4() {
	counts := make(map[string]IndexCount)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file1> <file2> ...\n", os.Args[0])
	} else {
		for i, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, i, counts)
			f.Close()
		}
	}
	for line, ic := range counts {
		if ic.n > 1 {
			fmt.Printf("%d\t%s\t%s\n", ic.n, line, files[ic.idx])
		}
	}
}

func countLines(f *os.File, idx int, counts map[string]IndexCount) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = IndexCount{n: counts[input.Text()].n + 1, idx: idx}
	}
}

func main() {
	f1()
	f2()
	f3()
	f4()
}
