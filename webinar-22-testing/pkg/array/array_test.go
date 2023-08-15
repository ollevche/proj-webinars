package array_test

import (
	"testing"
	"webinar-22/pkg/array"
)

func TestHasArrayItem(t *testing.T) {
	type testCase struct {
		name        string
		sourceInput []string
		itemsInput  []string
		wantRes     bool
	}

	testCases := []testCase{
		{
			name:        "basic true test",
			sourceInput: []string{"a", "b", "c"},
			itemsInput:  []string{"v", "w", "a"},
			wantRes:     true,
		},
		{
			name:        "basic false test",
			sourceInput: []string{"a", "b", "c"},
			itemsInput:  []string{"v", "w", "aa"},
			wantRes:     false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			haveRes := array.HasArrayItem(testCase.sourceInput, testCase.itemsInput)

			if haveRes != testCase.wantRes {
				t.Logf("Have mismatched result: %v instead of %v", haveRes, testCase.wantRes)
				t.Fail()
			}
		})
	}
}
