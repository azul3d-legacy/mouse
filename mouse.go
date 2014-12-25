// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mouse

// State represents a single mouse state, such as Up or Down.
type State int

const (
	InvalidState State = iota
	Down
	Up
)

// Button represents a single mouse button.
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

const (
	Left  = One
	Right = Two
	Wheel = Three
)
