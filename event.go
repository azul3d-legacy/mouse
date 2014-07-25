// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mouse

import (
	"fmt"
	"time"
)

// Event represents an single mouse event, such as pushing an button, or using
// the scroll-wheel, etc.
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
