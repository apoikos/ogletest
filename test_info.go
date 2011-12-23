// Copyright 2011 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ogletest

import (
	"github.com/jacobsa/oglemock"
)

// failureRecord represents a single failed expectation for a test.
type failureRecord struct {
	// The file name within which the expectation failed, e.g. "foo_test.go".
	FileName string

	// The line number at which the expectation failed.
	LineNumber int

	// The error generated by the testing framework. For example:
	//
	//     Expected: 17
	//     Actual:   "taco", which is not numeric
	//
	GeneratedError string

	// A user-specified string to print out with the error, if any.
	UserError string
}

// TestInfo represents information about a currently running or previously-run
// test.
type TestInfo struct {
	// A mock controller that is set up to report errors to the ogletest test
	// runner. This can be used for setting up mock expectations and handling
	// mock calls. The Finish method should not be run by the user; ogletest will
	// do that automatically after the test's TearDown method is run.
	MockController oglemock.Controller

	// A set of failure records that the test has produced.
	failureRecords []*failureRecord
}

// newTestInfo creates a valid but empty TestInfo struct.
func newTestInfo() *TestInfo {
	return &TestInfo{failureRecords: make([]*failureRecord, 0)}
}

// currentlyRunningTest is the state for the currently running test, if any.
var currentlyRunningTest *TestInfo
