package physarum

import (
	"fmt"
	"image"
	"image/png"
	"os"

	// "math/rand"
	"time"
)

const (
	width      = 1 << 9
	height     = 1 << 9
	particles  = 1 << 21
	iterations = 2000
	blurRadius = 1
	blurPasses = 2
	zoomFactor = 1
	foodPath   = "../food.png"
)

// just produce the final frame
func one(model *Model, iterations int) {
	now := time.Now().UTC().UnixNano() / 1000
	path := fmt.Sprintf("out%d.png", now)
	fmt.Println()
	fmt.Println(path)
	fmt.Println(len(model.Particles), "particles")
	// PrintConfigs(model.Configs, model.AttractionTable)
	// SummarizeConfigs(model.Configs)
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
	for i := 0; i < iterations; i++ {
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

// Read our food map
// filename is food png path
// 1.Read image
// 2.Convert to grayscale
// 3.Noramlize to range between 0 and normMax
func readFood(filename string, normMax float32) []float32 {
	// Open the PNG file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Decode the PNG image
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	// Convert the image to grayscale
	gray := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			gray.Set(x, y, img.At(x, y))
		}
	}

	// Convert the grayscale image to a 2D array and normalize values between 0 and 5
	width := gray.Bounds().Dx()
	height := gray.Bounds().Dy()
	normalizedArray := make([]float32, width*height)
	idx := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grayColor := gray.GrayAt(x, y)
			// Normalize grayscale value between 0 and 5
			normalizedValue := float32(grayColor.Y) * normMax / 255.0
			normalizedArray[idx] = normalizedValue
			idx++
		}
	}

	return normalizedArray
}

func Run() {

	foodMap := readFood(foodPath, 5.0)
	// n := 2 + rand.Intn(4)
	n := 3
	configs := RandomConfigs(n)
	table := RandomAttractionTable(n)
	fmt.Print(table[0])
	fmt.Println()
	model := NewModel(
		width, height, particles, blurRadius, blurPasses, zoomFactor,
		configs, table, foodMap)
	frames(model, 10)
}
