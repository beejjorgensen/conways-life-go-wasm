package main

import (
	"syscall/js"

	"github.com/beejjorgensen/conlife/life"
)

const cWidth = 400
const cHeight = 300

var conlife *life.Life
var ctx js.Value
var imageData js.Value

// updateLife single-steps the simulation and updates the canvas
func updateLife() {
	conlife.Step()

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

	imageData.Get("data").Call("set", newPixelDataArray)

	newPixelDataArray.Release()

	ctx.Call("putImageData", imageData, 0, 0)
}

// onStepButton is called when the step button is clicked
func onStepButton(args []js.Value) {
	updateLife()
}

// initJs initializes all the JS stuff
func initJs() {
	jsGlobal := js.Global()
	document := jsGlobal.Get("document")

	// Get a reference to the canvas
	canvas := document.Call("getElementById", "lifecanvas")

	// Set canvas size
	canvas.Call("setAttribute", "width", cWidth)
	canvas.Call("setAttribute", "height", cHeight)

	// Get the context
	ctx = canvas.Call("getContext", "2d")

	// Get image data
	imageData = ctx.Call("getImageData", 0, 0, cWidth, cHeight)

	// Set up the button event listener
	stepCb := js.NewCallback(onStepButton)
	document.Call("getElementById", "step-button").Call("addEventListener", "click", stepCb)
}

// Main
func main() {
	// Make a new Game
	conlife = life.New(cWidth, cHeight)
	conlife.Randomize()

	// Initialize JS and add the event listeners
	initJs()

	done := make(chan struct{}, 0)

	<-done
}
