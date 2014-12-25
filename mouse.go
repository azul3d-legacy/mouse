// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mouse

// State represents a single mouse state, such as Up or Down.
type State uint8

const (
	InvalidState State = iota
	Down
	Up
)

// Button represents a single mouse button.
type Button uint8

const (
	Invalid Button = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
)

// Left, Right, Middle and Wheel are simply aliases. Their true names are mouse
// button One, Two, and Three (for both Middle and Wheel, respectively).
const (
	Left  = One
	Right = Two
	Wheel = Three
	Middle = Three
)
