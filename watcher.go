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
func (s *Watcher) SetState(button Button, state State) {
	s.access.Lock()
	defer s.access.Unlock()

	s.states[button] = state
}

// States returns an copy of the internal mouse button state map used by this watcher.
func (s *Watcher) States() map[Button]State {
	s.access.RLock()
	defer s.access.RUnlock()

	copy := make(map[Button]State)
	for button, state := range s.states {
		copy[button] = state
	}
	return copy
}

// State returns the current state of the specified mouse button.
func (s *Watcher) State(button Button) State {
	s.access.Lock()
	defer s.access.Unlock()

	state, ok := s.states[button]
	if !ok {
		s.states[button] = Up
	}
	return state
}

// Down tells whether the specified mouse button is currently in the down state.
func (s *Watcher) Down(button Button) bool {
	return s.State(button) == Down
}

// Up tells whether the specified mouse button is currently in the up state.
func (s *Watcher) Up(button Button) bool {
	return s.State(button) == Up
}

// NewWatcher returns a new, initialized, mouse watcher.
func NewWatcher() *Watcher {
	s := new(Watcher)
	s.states = make(map[Button]State)
	return s
}
