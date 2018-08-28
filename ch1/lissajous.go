// lissajous 产生随机利萨如图形的 gif 图
package main

import (
	"image/color"
	"math/rand"
	"time"
	"os"
	"io"
	"image/gif"
	"image"
	"math"
	"net/http"
	"log"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // 画板中的第一种颜色
	blackIndex = 1 // 画板中的下一种颜色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) >= 2 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe(":2333", nil))
		return
	}
	lissajous(os.Stdout) // go run lissajous.go > lissajous.gif
}

func lissajous(outer io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nFrames = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0
	// 生成 64 帧的 gif
	for i := 0; i < nFrames; i++ {
		// 生成 201*201 的画板
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		// 生成黑线图像
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(outer, &anim)
}
