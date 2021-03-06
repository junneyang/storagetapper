// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package metrics

import "time"

//Timer is a wrapper around metrics.Timer
type Timer struct {
	backend timer
}

//TimerInit is a constructor for Timer
func TimerInit(s scope, name string) *Timer {
	var t Timer
	t.backend = s.InitTimer(name)
	return &t
}

//Start starts the timer
func (t *Timer) Start() *Timer {
	t.backend.Start()
	return t
}

//Stop stops the timer
func (t *Timer) Stop() {
	t.backend.Stop()
}

//Record sets the value of the timer to a specific value
func (t *Timer) Record(v time.Duration) {
	t.backend.Record(v)
}
