// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mouse

import (
	"fmt"
	"time"
)

// Event represents an single mouse event, such as pushing a button.
type Event struct {
	T      time.Time
	Button Button
	State  State
}

// Time returns the time at which this event occured.
func (e Event) Time() time.Time {
	return e.T
}

// String returns an string representation of this event.
func (e Event) String() string {
	return fmt.Sprintf("Event(Button=%v, State=%v, Time=%v)", e.Button, e.State, e.T)
}

// Scrolled is an event where the user has scrolled their mouse wheel.
type Scrolled struct {
	T time.Time

	// Amount of scrolling in horizontal (X) and vertical (Y) directions.
	X, Y float64
}

// Time implements the Event interface.
func (s Scrolled) Time() time.Time {
	return s.T
}

// String returns a string representation of this event.
func (s Scrolled) String() string {
	return fmt.Sprintf("Scrolled(X=%f, Y=%f, Time=%v)", s.X, s.Y, s.T)
}
