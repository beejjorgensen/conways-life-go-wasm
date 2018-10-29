package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
	"time"

	"github.com/beejjorgensen/conlife/life"
)

const cWidth = 400
const cHeight = 300

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	jsGlobal := js.Global()
	document := jsGlobal.Get("document")

	// Get a reference to the canvas
	canvas := document.Call("getElementById", "lifecanvas")

	// Set canvas size
	canvas.Call("setAttribute", "width", cWidth)
	canvas.Call("setAttribute", "height", cHeight)

	// Get the context
	ctx := canvas.Call("getContext", "2d")

	// Get image data
	imageData := ctx.Call("getImageData", 0, 0, cWidth, cHeight)
	//pixelData := imageData.Get("data")

	fmt.Println("A")

	// Make a new Game
	conlife := life.New(cWidth, cHeight)
	conlife.Randomize()

	conlife.Steps(50)

	fmt.Println("B")

	indexCount := cWidth * cHeight * 4
	newPixelData := make([]uint8, indexCount, indexCount)

	lifeData := conlife.Get()

	for i := range lifeData {
		var c uint8

		if lifeData[i] == 0 {
			c = 0
		} else {
			c = 255
		}

		j := i * 4

		newPixelData[j+0] = c
		newPixelData[j+1] = c
		newPixelData[j+2] = c
		newPixelData[j+3] = 255
	}

	newPixelDataArray := js.TypedArrayOf(newPixelData)

	fmt.Println("X")
	imageData.Get("data").Call("set", newPixelDataArray)

	newPixelDataArray.Release()

	fmt.Println("C")
	//fmt.Printf("Context: %v\n", ctx)
	//fmt.Printf("RGBA: %v\n", rgba)
	ctx.Call("putImageData", imageData, 0, 0)
	fmt.Println("D")
}
