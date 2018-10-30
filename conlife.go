package main

import (
	"strconv"
	"syscall/js"

	"github.com/beejjorgensen/conlife/life"
)

const cWidth = 400  // Canvas/game width
const cHeight = 300 // Canvas/game height

var conlife *life.Life // Reference to game data
var running bool       // If the game is running or not

var ctx js.Value         // Canvas drawing context
var imageData js.Value   // Reference to canvas's imageData
var newPixelData []uint8 // RGBA data for the canvas

var animFrameCb js.Callback // requestAnimationFrame callback

// updateLife single-steps the simulation and updates the canvas
func updateLife() {
	conlife.Step()
	drawLife()

	setInnerHTML("generation", strconv.FormatInt(int64(conlife.Generation), 10))
}

// drawLife renders the current game state
func drawLife() {
	// lifeData is read-only!
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

// requestAnimFrame requests another anim frame
func requestAnimFrame() {
	js.Global().Call("requestAnimationFrame", animFrameCb)
}

// onAnimFrame is called each animation frame
func onAnimFrame(args []js.Value) {
	updateLife()

	if running {
		requestAnimFrame()
	}
}

// startRun starts the simulation
func startRun() {
	if !running {
		setInnerHTML("run-button", "Stop")
		running = true

		requestAnimFrame()
	}
}

// stopRun stops the simulation
func stopRun() {
	if running {
		setInnerHTML("run-button", "Run")
		running = false
	}
}

// onStepButton is called when the step button is clicked
func onStepButton(args []js.Value) {
	stopRun()
	updateLife()
}

// onRunButton is called when the Run/Stop button is clicked
func onRunButton(args []js.Value) {
	if running {
		stopRun()
	} else {
		startRun()
	}
}

// setInnerHTML sets an HTML button label
func setInnerHTML(id, label string) {
	document := js.Global().Get("document")
	document.Call("getElementById", id).Set("innerHTML", label)
}

// initJs initializes all the JS stuff
func initJs() {
	document := js.Global().Get("document")

	// Get a reference to the canvas
	canvas := document.Call("getElementById", "lifecanvas")

	// Set canvas size
	canvas.Call("setAttribute", "width", cWidth)
	canvas.Call("setAttribute", "height", cHeight)

	// Get the context
	ctx = canvas.Call("getContext", "2d")

	// Get image data
	imageData = ctx.Call("getImageData", 0, 0, cWidth, cHeight)

	// Set up the button event listeners
	cb := js.NewCallback(onStepButton)
	document.Call("getElementById", "step-button").Call("addEventListener", "click", cb)

	cb = js.NewCallback(onRunButton)
	document.Call("getElementById", "run-button").Call("addEventListener", "click", cb)

	animFrameCb = js.NewCallback(onAnimFrame)
}

// Main
func main() {
	// Make a new Game
	conlife = life.New(cWidth, cHeight)
	conlife.Randomize()

	// Allocate an RGBA array for drawing to the canvas
	indexCount := cWidth * cHeight * 4
	newPixelData = make([]uint8, indexCount, indexCount)

	// Initialize JS and add the event listeners
	initJs()

	// Show the story so far
	drawLife()

	// Block forever
	done := make(chan struct{}, 0)
	<-done
}
