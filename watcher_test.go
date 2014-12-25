// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mouse

import "testing"

func TestWatcher(t *testing.T) {
	m := NewWatcher()
	m.SetState(Left, Down)
	m.SetState(Right, Up)
	if !m.Down(Left) {
		t.Fatal("expect mouse.Left in state mouse.Down")
	}
	if !m.Up(Right) {
		t.Fatal("expect mouse.Right in state mouse.Up")
	}
	if !m.Up(Wheel) {
		t.Fatal("expect mouse.Wheel in state mouse.Up")
	}
}
