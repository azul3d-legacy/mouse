// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mouse implements various mouse related data types.
package mouse

import (
	"fmt"
)

// State represents an single mouse state, such as Up, Down, or a scroll
// direction.
type State int

const (
	InvalidState State = iota
	Down
	Up
)

// String returns an string representation of the mouse state.
func (s State) String() string {
	switch s {
	case InvalidState:
		return "InvalidState"
	case Down:
		return "Down"
	case Up:
		return "Up"
	}
	return fmt.Sprintf("State(%d)", s)
}

// Button represents an single mouse button.
type Button int

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

// String returns an string representation of the mouse button.
func (b Button) String() string {
	switch b {
	case Invalid:
		return "Invalid"
	case Left:
		return "Left"
	case Right:
		return "Right"
	case Wheel:
		return "Wheel"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	}
	return fmt.Sprintf("Button(%d)", b)
}

const (
	Left  = One
	Right = Two
	Wheel = Three
)
