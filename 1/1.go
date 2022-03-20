package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
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

// TODO: 测试f1()和echo的性能差异
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

var palette = []color.Color{color.White, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}}

// Windows 环境下使用cmd执行
func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1+uint8(i/(nframes/len(palette))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main() {
	// f1()
	// f2()
	// f3()
	// f4()
	lissajous(os.Stdout)
}
