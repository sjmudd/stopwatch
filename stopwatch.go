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
// This package is only thread safe for the AddElapsed() call.
// Other usage assumes that usage is single threaded.
package stopwatch

import (
	"fmt"
	"sync"
	"time"
)

// DefaultFormat allows the Stopwatch.String() function to be
// configured differently to time.Duration if needed.  This is done
// at the global package level to avoid having to do on each Stopwatch
// instance.
var DefaultFormat = func(t time.Duration) string { return t.String() }

// Stopwatch is a structure to hold information about a stopwatch
type Stopwatch struct {
	sync.RWMutex
	format  func(time.Duration) string
	elapsed time.Duration
	refTime time.Time
}

// Start returns a pointer to a new Stopwatch struct and indicates
// that the stopwatch has started.
func Start(f func(time.Duration) string) *Stopwatch {
	s := New(f)
	s.Start()

	return s
}

// New returns a pointer to a new Stopwatch struct.
func New(f func(time.Duration) string) *Stopwatch {
	s := new(Stopwatch)
	s.format = f

	return s
}

// Start records that we are now running. If called previously this
// is a no-op.
func (s *Stopwatch) Start() {
	s.Lock()
	defer s.Unlock()
	if s.IsRunning() {
		fmt.Printf("WARNING: Stopwatch.Start() IsRunning is true\n")
	} else {
		s.refTime = time.Now()
	}
}

// Stop collects the elapsed time if running and remembers we are
// not running.
func (s *Stopwatch) Stop() {
	s.Lock()
	defer s.Unlock()
	if s.IsRunning() {
		s.elapsed += time.Since(s.refTime)
		s.refTime = time.Time{}
	} else {
		fmt.Printf("WARNING: Stopwatch.Stop() IsRunning is false\n")
	}
}

// Reset resets the counters.
func (s *Stopwatch) Reset() {
	s.Lock()
	defer s.Unlock()
	if s.IsRunning() {
		fmt.Printf("WARNING: Stopwatch.Reset() IsRunning is true\n")
	}
	s.refTime = time.Time{}
	s.elapsed = 0
}

// String gives the string representation of the duration.
func (s *Stopwatch) String() string {
	s.RLock()
	defer s.Unlock()
	// display using local formatting if possible
	if s.format != nil {
		return s.format(s.elapsed)
	}
	// display using package DefaultFormat
	return DefaultFormat(s.elapsed)
}

// SetStringFormat allows the String() function to be configured
// differently to time.Duration for the specific Stopwatch.
func (s *Stopwatch) SetStringFormat(f func(time.Duration) string) {
	s.Lock()
	defer s.Unlock()
	s.format = f
}

// IsRunning is a helper function to indicate if in theory the
// stopwatch is working.
func (s *Stopwatch) IsRunning() bool {
	s.RLock()
	defer s.Unlock()
	return !s.refTime.IsZero()
}

// Elapsed returns the elapsed time since starting (in time.Duration).
func (s *Stopwatch) Elapsed() time.Duration {
	s.RLock()
	defer s.Unlock()
	if s.IsRunning() {
		return time.Since(s.refTime)
	}
	return s.elapsed
}

// ElapsedSeconds is a helper function returns the number of seconds
// since starting.
func (s *Stopwatch) ElapsedSeconds() float64 {
	s.RLock()
	defer s.Unlock()
	return s.Elapsed().Seconds()
}

// ElapsedMilliSeconds is a helper function returns the number of
// milliseconds since starting.
func (s *Stopwatch) ElapsedMilliSeconds() float64 {
	s.RLock()
	defer s.Unlock()
	return float64(s.Elapsed() / time.Millisecond)
}

// AddElapsed just adds an elapsed time to the value that's been stored.
func (s *Stopwatch) AddElapsed(t time.Duration) {
	s.Lock()
	defer s.Unlock()
	s.elapsed += t
}
