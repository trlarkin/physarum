package physarum

import (
	"fmt"
	"image/png"
	// "math/rand"
	"time"
)

const (
	width      = 1 << 9
	height     = 1 << 9
	particles  = 1 << 22
	iterations = 400
	blurRadius = 1
	blurPasses = 2
	zoomFactor = 1
)

// just produce the final frame
func one(model *Model, iterations int) {
	now := time.Now().UTC().UnixNano() / 1000
	path := fmt.Sprintf("out%d.png", now)
	fmt.Println()
	fmt.Println(path)
	fmt.Println(len(model.Particles), "particles")
	PrintConfigs(model.Configs, model.AttractionTable)
	SummarizeConfigs(model.Configs)
	for i := 0; i < iterations; i++ {
		model.Step()
	}
	palette := RandomPalette()
	im := Image(model.W, model.H, model.Data(), palette, 0, 0, 1/2.2)
	SavePNG(path, im, png.DefaultCompression)
}

// produce multiple frames
func frames(model *Model, rate int) {
	palette := RandomPalette()

	saveImage := func(path string, w, h int, grids [][]float32, ch chan bool) {
		max := particles / float32(width*height) * 20
		im := Image(w, h, grids, palette, 0, max, 1/2.2)
		SavePNG(path, im, png.BestSpeed)
		if ch != nil {
			ch <- true
		}
	}

	ch := make(chan bool, 1)
	ch <- true
	for i := 0; ; i++ {
		fmt.Println(i)
		model.Step()
		if i%rate == 0 {
			<-ch
			path := fmt.Sprintf("frame%08d.png", i/rate)
			go saveImage(path, model.W, model.H, model.Data(), ch)
		}
	}
}

func ConfigVaryRotationAngle(as []float32) []Config {
	n := len(as)
	configs := make([]Config, n)
	for index, a := range as {
		configs[index] = Config{
			SensorAngle:      Radians(45),
			SensorDistance:   32,
			RotationAngle:    Radians(a),
			StepDistance:     1,
			DepositionAmount: 5,
			DecayFactor:      0.1,
		}
	}
	return configs
}

func Run() {
	// if false {
	// 	n := 2 + rand.Intn(4)
	// 	configs := RandomConfigs(n)
	// 	table := RandomAttractionTable(n)
	// 	model := NewModel(
	// 		width, height, particles, blurRadius, blurPasses, zoomFactor,
	// 		configs, table)
	// 	frames(model, 3)
	// }

	for _, f := range []float32{3} {

		// n := 2 + rand.Intn(4)
		configs := ConfigVaryRotationAngle([]float32{f})
		table := RandomAttractionTable(4)
		model := NewModel(
			width, height, particles, blurRadius, blurPasses, zoomFactor,
			configs, table)
		start := time.Now()
		frames(model, 10)
		// one(model, iterations)
		fmt.Println(time.Since(start))
	}
}
