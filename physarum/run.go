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
	iterations = 20000
	blurRadius = 1
	blurPasses = 2
	zoomFactor = 1
)

// Paths to food files you want to use
var foodPaths = []string{"../foodBigNode2.png", "../foodNoBigNode2.png"}

// Highest iteration number food file will take place for
var foodIters = []int{10000, 20000}

// just produce the final frame
func one(model *Model, iterations int, path0 string) {
	now := time.Now().UTC().UnixNano() / 1000
	path := fmt.Sprintf("out%d.png", now)
	if path0 != "" {
		path = fmt.Sprintf("out_%s.png", path0)
	}
	fmt.Println()
	fmt.Println(path)
	fmt.Println(len(model.Particles), "particles")
	// PrintConfigs(model.Configs, model.AttractionTable)
	// SummarizeConfigs(model.Configs)
	for i := 0; i < iterations; i++ {
		if i%(iterations/10) == 0 {
			fmt.Println(i/(iterations/10)*10, "%")
		}
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
	for index, _ := range as {
		configs[index] = Config{
			SensorAngle:      Radians(45),
			SensorDistance:   8,
			RotationAngle:    Radians(45),
			StepDistance:     1,
			DepositionAmount: 2,
			DecayFactor:      0.05,
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

	var foodMap [][]float32
	for _, path := range foodPaths {
		fmt.Println("Reading food map: ")
		fmt.Println(path)
		foodArr := readFood(path, 5.0)
		foodMap = append(foodMap, foodArr)
	}

	// n := 2 + rand.Intn(4)
	n := 2
	configs := RandomConfigs(n)
	// config := &Config{
	// 	SensorAngle:      45,
	// 	SensorDistance:   50,
	// 	RotationAngle:    10,
	// 	StepDistance:     1,
	// 	DepositionAmount: 5,
	// 	DecayFactor:      0.1,
	// }
	// configs := []Config{*config}
	table := RandomAttractionTable(n)
	// table := [][]float32{{1}}
	fmt.Print(len(table))
	fmt.Print(len(table[0]))
	fmt.Println()
	model := NewModel(
		width, height, particles, blurRadius, blurPasses, zoomFactor,
		configs, table, foodMap, foodIters)

	frames(model, 20)
}

func Tristan() {
	var foodMap [][]float32
	foodMap = append(foodMap, readFood("../foodED.png", 100.0))
	foodIters = []int{20000}
	// n := 2 + rand.Intn(4)
	n := 3
	table := RandomAttractionTable(n)
	configs := ConfigVaryRotationAngle([]float32{35})
	// for _, p := range []int{17, 18, 19, 20, 21, 22} {
	fmt.Print(table[0])
	fmt.Println()
	model := NewModel(
		1<<9,      // width
		1<<9,      // height
		1<<17,     // numParticles
		1,         // blurRadius
		2,         // blurPasses
		1,         // zoomFactor
		configs,   // configs
		table,     // attractionTable
		foodMap,   // foodMap
		foodIters) // foodIters
	frames(model, 100)
	// one(model, 10000, "neighborhood2")
	// }
}

func Edgar() {

	var foodMap [][]float32
	for _, path := range foodPaths {
		fmt.Println("Reading food map: ")
		fmt.Println(path)
		foodArr := readFood(path, 5.0)
		foodMap = append(foodMap, foodArr)
	}

	// n := 2 + rand.Intn(4)
	n := 2

	// configs := RandomConfigs(n)
	configs := ConfigVaryRotationAngle([]float32{35})

	table := RandomAttractionTable(n)
	// table := [][]float32{{1}}
	fmt.Print(len(table))
	fmt.Print(len(table[0]))
	fmt.Println()
	model := NewModel(
		width, height, particles, blurRadius, blurPasses, zoomFactor,
		configs, table, foodMap, foodIters)

	frames(model, 100)
}
