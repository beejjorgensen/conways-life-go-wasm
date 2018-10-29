package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
	"time"
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

	// Make a random array of 1's and 0's
	life := [cWidth * cHeight]int{}

	fmt.Println("A")
	for i := range life {
		life[i] = rand.Int() & 1
	}

	fmt.Println("B")
	// Convert it into an RGBA array
	/*
			for i := range life {
				var c int

				if life[i] == 0 {
					c = 0
				} else {
					c = 255
				}

				j := i * 4

				pixelData.SetIndex(j+0, c)
				pixelData.SetIndex(j+1, c)
				pixelData.SetIndex(j+2, c)
				pixelData.SetIndex(j+3, 255)
			}
		imageData.Set("data", js.TypedArrayOf(newPixelData))
	*/

	indexCount := cWidth * cHeight * 4
	newPixelData := make([]uint8, indexCount, indexCount)

	for i := range life {
		var c uint8

		if life[i] == 0 {
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
