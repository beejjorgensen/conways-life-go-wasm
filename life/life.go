package life

import (
	"math/rand"
	"time"
)

// Life contains all data about the current game
type Life struct {
	Generation int
	Width      int
	Height     int
	data       [2][]uint8
	readBuf    int
}

// New returns a new Life structure
func New(width, height int) *Life {
	size := width * height

	return &Life{
		Generation: 0,
		data:       [2][]uint8{make([]uint8, size, size), make([]uint8, size, size)},
		Width:      width,
		Height:     height,
	}
}

// Randomize fills the life grid randomly
func (l *Life) Randomize() {
	rand.Seed(time.Now().UTC().UnixNano())

	buf := l.data[l.readBuf]

	for i := range buf {
		buf[i] = uint8(rand.Uint32() & 1)
	}
}

// Return the value at the given x, y coordinate on the read buffer
func (l *Life) readVal(x, y int) uint8 {
	return l.data[l.readBuf][y*l.Width+x]
}

// Steps runs a number of steps
func (l *Life) Steps(n int) {
	for i := 0; i < n; i++ {
		l.Step()
	}
}

// Step runs a single generation
func (l *Life) Step() {
	var writeBuf []uint8

	if l.readBuf == 0 {
		writeBuf = l.data[1]
	} else {
		writeBuf = l.data[0]
	}

	// Loop through all the pixels
	for y := 0; y < l.Height; y++ {
		for x := 0; x < l.Width; x++ {

			// Compute the wraparound pixels
			xm1 := x - 1
			if xm1 < 0 {
				xm1 = l.Width - 1
			}

			xp1 := x + 1
			if xp1 >= l.Width {
				xp1 = 0
			}

			ym1 := y - 1
			if ym1 < 0 {
				ym1 = l.Height - 1
			}

			yp1 := y + 1
			if yp1 >= l.Height {
				yp1 = 0
			}

			// Count neighbors
			nCount := l.readVal(xm1, ym1) +
				l.readVal(x, ym1) +
				l.readVal(xp1, ym1) +
				l.readVal(xm1, y) +
				l.readVal(xp1, y) +
				l.readVal(xm1, yp1) +
				l.readVal(x, yp1) +
				l.readVal(xp1, yp1)

			// Rules of life and death
			var newVal uint8

			if l.readVal(x, y) == 0 {
				// If dead, comes to life if n=3
				if nCount == 3 {
					newVal = 1
				} else {
					newVal = 0
				}

			} else {
				// If alive, dies if n<2 or n>3
				if nCount < 2 || nCount > 3 {
					newVal = 0
				} else {
					newVal = 1
				}
			}

			// Write the new value
			writeBuf[y*l.Width+x] = newVal
		}
	}

	// Pageflip
	if l.readBuf == 0 {
		l.readBuf = 1
	} else {
		l.readBuf = 0
	}
}

// Get returns a reference to the current buffer. Treat this as read-only!
func (l *Life) Get() []uint8 {
	return l.data[l.readBuf]
}
