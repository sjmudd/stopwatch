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
	"time"
)

// NamedStopwatch holds a map of string named stopwatches. Intended to be used when several
// Stopwatches are being used at once, and easy to use as they are name based.
type NamedStopwatch struct {
	stopwatches map[string]*Stopwatch
}

// 
func (ns *NamedStopwatch) AddStopwatch(name string, s *Stopwatch) {
	if ns.stopwatches == nil {
		ns.stopwatches = make(map[string](*Stopwatch))
	}
	ns.stopwatches[name] = s
}

func (ns *NamedStopwatch) Start(name string) {
	if s, ok := ns.stopwatches[name]; ok {
		s.Start()
	}
}
func (ns *NamedStopwatch) Stop(name string) {
	if s, ok := ns.stopwatches[name]; ok {
		s.Stop()
	}
}

func (ns *NamedStopwatch) Reset(name string) {
	if s, ok := ns.stopwatches[name]; ok {
		s.Reset()
	}
}

func (ns *NamedStopwatch) Keys() []string {
	keys := []string{}
	for k, _ := range ns.stopwatches {
		keys = append(keys, k)
	}
	return keys
}

func (ns *NamedStopwatch) Elapsed(name string) time.Duration {
	if s, ok := ns.stopwatches[name]; ok {
		return s.Elapsed()
	}
	return time.Duration(0)
}

func (ns *NamedStopwatch) ElapsedSeconds(name string) float64 {
	if s, ok := ns.stopwatches[name]; ok {
		return s.ElapsedSeconds()
	}
	return float64(0)
}

func (ns *NamedStopwatch) ElapsedMilliSeconds(name string) float64 {
	if s, ok := ns.stopwatches[name]; ok {
		return s.ElapsedMilliSeconds()
	}
	return float64(0)
}
