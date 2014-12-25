// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mouse

import (
	"bytes"
	"fmt"
	"sync"
)

// Watcher watches the state of various mouse buttons and their states.
type Watcher struct {
	access sync.RWMutex
	states map[Button]State
}

// String returns a multi-line string representation of this mouse watcher and
// it's associated states.
func (w *Watcher) String() string {
	w.access.RLock()
	defer w.access.RUnlock()

	bb := new(bytes.Buffer)
	fmt.Fprintf(bb, "Watcher(\n")
	for b, s := range w.states {
		fmt.Fprintf(bb, "    %v: %v\n", b, s)
	}
	fmt.Fprintf(bb, ")")
	return bb.String()
}

// SetState specifies the current state of the specified mouse button.
func (w *Watcher) SetState(button Button, state State) {
	w.access.Lock()
	defer w.access.Unlock()

	w.states[button] = state
}

// States returns an copy of the internal mouse button state map used by this
// watcher.
func (w *Watcher) States() map[Button]State {
	w.access.RLock()
	defer w.access.RUnlock()

	copy := make(map[Button]State)
	for button, state := range w.states {
		copy[button] = state
	}
	return copy
}

// EachState calls f with each known button to this watcher and it's current
// button state. It does so until the function returns false or there are no
// more buttons known to the watcher.
func (w *Watcher) EachState(f func(b Button, s State) bool) {
	w.access.RLock()
	defer w.access.RUnlock()

	for button, state := range w.states {
		// Call the function without the lock being held, so they can access
		// methods on this watcher still.
		w.access.RUnlock()
		cont := f(button, state)
		w.access.RLock()

		if !cont {
			return
		}
	}
}

// State returns the current state of the specified mouse button.
func (w *Watcher) State(button Button) State {
	w.access.Lock()
	defer w.access.Unlock()

	state, ok := w.states[button]
	if !ok {
		w.states[button] = Up
	}
	return state
}

// Down tells whether the specified mouse button is currently in the down
// state.
func (w *Watcher) Down(button Button) bool {
	return w.State(button) == Down
}

// Up tells whether the specified mouse button is currently in the up state.
func (w *Watcher) Up(button Button) bool {
	return w.State(button) == Up
}

// NewWatcher returns a new, initialized, mouse watcher.
func NewWatcher() *Watcher {
	w := new(Watcher)
	w.states = make(map[Button]State)
	return w
}
