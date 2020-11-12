//args: -Eparalleltest
package testdata

import (
	"fmt"
	"testing"
)

func NoATestFunction()                                  {}
func TestingFunctionLooksLikeATestButIsNotWithParam()   {}
func TestingFunctionLooksLikeATestButIsWithParam(i int) {}
func AbcFunctionSuccessful(t *testing.T)                {}

func TestFunctionSuccessfulRangeTest(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
	}{{name: "foo"}}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fmt.Println(tc.name)
		})
	}
}

func TestFunctionSuccessfulNoRangeTest(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
	}{{name: "foo"}, {name: "bar"}}

	t.Run(testCases[0].name, func(t *testing.T) {
		t.Parallel()
		fmt.Println(testCases[0].name)
	})
	t.Run(testCases[1].name, func(t *testing.T) {
		t.Parallel()
		fmt.Println(testCases[1].name)
	})

}

func TestFunctionMissingCallToParallel(t *testing.T) {} // ERROR "Function TestFunctionMissingCallToParallel missing the call to method parallel"
func TestFunctionRangeMissingCallToParallel(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
	}{{name: "foo"}}

	// this range loop should be okay as it does not have test Run
	for _, tc := range testCases {
		fmt.Println(tc.name)
	}

	for _, tc := range testCases { // ERROR "Range statement for test TestFunctionRangeMissingCallToParallel missing the call to method parallel in test Run"
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
		})
	}
}

func TestFunctionMissingCallToParallelAndRangeNotUsingRangeValueInTDotRun(t *testing.T) { // ERROR "Function TestFunctionMissingCallToParallelAndRangeNotUsingRangeValueInTDotRun missing the call to method parallel"
	testCases := []struct {
		name string
	}{{name: "foo"}}

	for _, tc := range testCases { // ERROR "Range statement for test TestFunctionMissingCallToParallelAndRangeNotUsingRangeValueInTDotRun missing the call to method parallel in test Run"
		t.Run(tc.name, func(t *testing.T) {
			fmt.Println(tc.name)
		})
	}
}

func TestFunctionRangeNotUsingRangeValueInTDotRun(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
	}{{name: "foo"}}
	for _, tc := range testCases { // ERROR "Range statement for test TestFunctionRangeNotUsingRangeValueInTDotRun does not use range value in test Run"
		t.Run("tc.name", func(t *testing.T) {
			t.Parallel()
			fmt.Println(tc.name)
		})
	}
}

func TestFunctionRangeNotReInitialisingVariable(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
	}{{name: "foo"}}
	for _, tc := range testCases { // ERROR "Range statement for test TestFunctionRangeNotReInitialisingVariable does not reinitialise the variable tc"
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fmt.Println(tc.name)
		})
	}
}

func TestFunctionTwoTestRunMissingCallToParallel(t *testing.T) {
	t.Parallel()

	t.Run("1", func(t *testing.T) { // ERROR "Function TestFunctionTwoTestRunMissingCallToParallel has missing the call to method parallel in the test run"
		fmt.Println("1")
	})
	t.Run("2", func(t *testing.T) { // ERROR "Function TestFunctionTwoTestRunMissingCallToParallel has missing the call to method parallel in the test run"
		fmt.Println("2")
	})
}

func TestFunctionFirstOneTestRunMissingCallToParallel(t *testing.T) {
	t.Parallel()

	t.Run("1", func(t *testing.T) { // ERROR "Function TestFunctionFirstOneTestRunMissingCallToParallel has missing the call to method parallel in the test run"
		fmt.Println("1")
	})
	t.Run("2", func(t *testing.T) {
		t.Parallel()
		fmt.Println("2")
	})
}

func TestFunctionSecondOneTestRunMissingCallToParallel(t *testing.T) {
	t.Parallel()

	t.Run("1", func(t *testing.T) {
		t.Parallel()
		fmt.Println("1")
	})
	t.Run("2", func(t *testing.T) { // ERROR "Function TestFunctionSecondOneTestRunMissingCallToParallel has missing the call to method parallel in the test run"
		fmt.Println("2")
	})
}
