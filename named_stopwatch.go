/*
Copyright (c) 2016, Simon J Mudd
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// Package stopwatch implements simple stopwatch functionality
package stopwatch

import (
	"errors"
	"fmt"
	"time"
)

// NamedStopwatch holds a map of string named stopwatches. Intended to be used when several
// Stopwatches are being used at once, and easy to use as they are name based.
type NamedStopwatch struct {
	stopwatches map[string](*Stopwatch)
}

// NewNamedStopwatch creates an empty Stopwatch list
func NewNamedStopwatch() *NamedStopwatch {
	return new(NamedStopwatch)
}

// AddStopwatch adds a Stopwatch with the given name.
func (ns *NamedStopwatch) Add(name string) error {
	if ns.stopwatches == nil {
		ns.stopwatches = make(map[string](*Stopwatch))
	}
	if _, ok := ns.stopwatches[name]; ok {
		return errors.New(fmt.Sprintf("Stopwatch name %q already exists", name))
	}
	ns.stopwatches[name] = New(nil)
	return nil
}

// Delete removes a Stopwatch with the given name (if it exists)
func (ns *NamedStopwatch) Delete(name string) {
	if ns.stopwatches == nil {
		return
	}

	delete(ns.stopwatches, name) // check if it exists in case the user did the wrong thing
}

// Exists returns true if the NamedStopwatch exists
func (ns *NamedStopwatch) Exists(name string) bool {
	if ns == nil {
		return false
	}

	_, found := ns.stopwatches[name]

	return found
}

// Start starts a NamedStopwatch if it exists
func (ns *NamedStopwatch) Start(name string) {
	if ns == nil {
		return
	}
	if s, ok := ns.stopwatches[name]; ok {
		s.Start()
	}
}

// Stop stops a NamedStopwatch if it exists
func (ns *NamedStopwatch) Stop(name string) {
	if ns == nil {
		return
	}
	if s, ok := ns.stopwatches[name]; ok {
		s.Stop()
	}
}

// Reset resets a NamedStopwatch if it exists
func (ns *NamedStopwatch) Reset(name string) {
	if ns == nil {
		return
	}
	if s, ok := ns.stopwatches[name]; ok {
		s.Reset()
	}
}

// Keys returns the known names of Stopwatches
func (ns *NamedStopwatch) Keys() []string {
	if ns == nil {
		return nil
	}
	keys := []string{}
	for k, _ := range ns.stopwatches {
		keys = append(keys, k)
	}
	return keys
}

// Elapsed returns the elapsed time.Duration of the named stopwatch if it exists or 0
func (ns *NamedStopwatch) Elapsed(name string) time.Duration {
	if s, ok := ns.stopwatches[name]; ok {
		return s.Elapsed()
	}
	return time.Duration(0)
}

// Elapsed returns the elapsed time in seconds of the named stopwatch if it exists or 0
func (ns *NamedStopwatch) ElapsedSeconds(name string) float64 {
	if s, ok := ns.stopwatches[name]; ok {
		return s.ElapsedSeconds()
	}
	return float64(0)
}

// Elapsed returns the elapsed time in milliseconds of the named stopwatch if it exists or 0
func (ns *NamedStopwatch) ElapsedMilliSeconds(name string) float64 {
	if s, ok := ns.stopwatches[name]; ok {
		return s.ElapsedMilliSeconds()
	}
	return float64(0)
}
