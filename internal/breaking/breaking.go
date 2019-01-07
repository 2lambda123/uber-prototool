// Copyright (c) 2019 Uber Technologies, Inc.
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

package breaking

import (
	"github.com/uber/prototool/internal/extract"
	"github.com/uber/prototool/internal/text"
	"go.uber.org/zap"
)

var (
	// AllCheckers are all known Checkers.
	AllCheckers = []Checker{
		Checker{
			ID:      "MESSAGES_NOT_DELETED",
			Purpose: "Checks that no messages have been deleted.",
			Check:   checkMessagesNotDeleted,
		},
		Checker{
			ID:      "MESSAGE_FIELDS_NOT_DELETED",
			Purpose: "Checks that no message fields have been deleted. We require fields to be deprecated instead of deleted and added to the reserved list.",
			Check:   checkMessageFieldsNotDeleted,
		},
		Checker{
			ID:      "MESSAGE_FIELDS_HAVE_SAME_TYPE",
			Purpose: "Checks that message fields have the same type.",
			Check:   checkMessageFieldsHaveSameType,
		},
	}
)

// Checker checks compatibility.
type Checker struct {
	// The ID of this Checker. This should be all UPPER_SNAKE_CASE.
	ID string
	// The purpose of this Checker. This should be a human-readable string.
	Purpose string
	// Check the compatibility of from and to.
	//
	// Returns an error only if there is a system error.
	Check func(addFailure func(*text.Failure), from *extract.PackageSet, to *extract.PackageSet) error
}

// Runner runs a series of Checkers.
type Runner interface {
	// Run runs Check on all the associated Checkers.
	//
	// Returns Failures if there are incompatibilities, or error if there is
	// a system error
	Run(from *extract.PackageSet, to *extract.PackageSet) ([]*text.Failure, error)
}

// RunnerOption is an option for a new Runner.
type RunnerOption func(*runner)

// RunnerWithLogger returns a RunnerOption that uses the given logger.
//
// The default is to use zap.NewNop().
func RunnerWithLogger(logger *zap.Logger) RunnerOption {
	return func(runner *runner) {
		runner.logger = logger
	}
}

// RunnerWithCheckers returns a RunnerOption that uses the given checkers.
//
// The default is to use AllCheckers.
func RunnerWithCheckers(checkers ...Checker) RunnerOption {
	return func(runner *runner) {
		runner.checkers = checkers
	}
}

// NewRunner returns a new Runner.
func NewRunner(options ...RunnerOption) Runner {
	return newRunner(options...)
}
