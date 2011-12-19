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
	"testing"
)

////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////

type fakeMatcher struct {
	matchFunc func(interface{}) (MatchResult, string)
	description string
}

func (m *fakeMatcher) Matches(c interface{}) (MatchResult, string) {
	return m.matchFunc(c)
}

func (m *fakeMatcher) Description() string {
	return m.description
}

////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////

func TestCallsWrapped(t *testing.T) {
	var suppliedCandidate interface{}
	matchFunc := func(c interface{}) (MatchResult, string) {
		suppliedCandidate = c
		return MATCH_TRUE, ""
	}

	wrapped := &fakeMatcher{matchFunc, ""}
	matcher := Not(wrapped)

	matcher.Matches(17)
	if suppliedCandidate != 17 {
		t.Errorf("Expected 17, got %v", suppliedCandidate)
	}
}

func TestTrueMatchFromWrapped(t *testing.T) {
	matchFunc := func(c interface{}) (MatchResult, string) {
		return MATCH_TRUE, ""
	}

	wrapped := &fakeMatcher{matchFunc, ""}
	matcher := Not(wrapped)

	res, _ := matcher.Matches(0)
	if res != MATCH_FALSE {
		t.Errorf("Expected MATCH_FALSE, got %v", res)
	}
}

func TestFalseMatchFromWrapped(t *testing.T) {
	matchFunc := func(c interface{}) (MatchResult, string) {
		return MATCH_FALSE, "taco"
	}

	wrapped := &fakeMatcher{matchFunc, ""}
	matcher := Not(wrapped)

	res, err := matcher.Matches(0)
	if res != MATCH_TRUE {
		t.Errorf("Expected MATCH_TRUE, got %v", res)
	}

	if err != "" {
		t.Errorf("Expected empty string, got %v", err)
	}
}

func TestUndefinedMatchFromWrapped(t *testing.T) {
	matchFunc := func(c interface{}) (MatchResult, string) {
		return MATCH_UNDEFINED, "taco"
	}

	wrapped := &fakeMatcher{matchFunc, ""}
	matcher := Not(wrapped)

	res, err := matcher.Matches(0)
	if res != MATCH_UNDEFINED {
		t.Errorf("Expected MATCH_UNDEFINED, got %v", res)
	}

	if err != "taco" {
		t.Errorf("Expected taco, got %v", err)
	}
}

func TestDescription(t *testing.T) {
}

