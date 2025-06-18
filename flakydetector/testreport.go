package flakydetector

import (
	"github.com/bitrise-steplib/bitrise-step-detect-flaky-tests/testreport"
)

type TestReportWithRepetition struct {
	TestSuites []TestSuite
}

type TestSuite struct {
	Name          string
	Tests         int
	RepeatedTests int
	Failures      int
	Skipped       int
	Errors        int
	Time          float64
	TestCases     []TestCase
}

type TestCase struct {
	testreport.TestCase
	Repetition []testreport.TestCase
}

func ConvertToTestReportWithRepetition(testReport testreport.TestReport) TestReportWithRepetition {
	testReportWithRepetition := TestReportWithRepetition{}

	for _, suite := range testReport.TestSuites {
		// Map to track unique test cases by their ID. The ID is a combination of ClassName and Name.
		// The value in the map is the index of the test case in the testCases slice.
		testCasesToIdx := map[string]int{}
		var testCases []TestCase

		// Count how many test cases are repeated
		repeatedTests := 0

		for _, testCase := range suite.TestCases {
			testID := testCase.ClassName + "." + testCase.Name

			if idx, ok := testCasesToIdx[testID]; !ok {
				// Store the index (in the testCases slice) of the new test case
				testCasesToIdx[testID] = len(testCases)
				// TODO: this should be changed: the root test case should be the aggregated result of the repetitions
				// and repetition should include the first run too

				testCases = append(testCases, TestCase{
					TestCase: testCase,
				})
			} else {
				repeatedTestCase := testCases[idx]

				if len(repeatedTestCase.Repetition) == 0 {
					// If this is the first repetition, count it as a repeated test
					repeatedTests++

					// Add the first run to the repetition
					repeatedTestCase.Repetition = append(repeatedTestCase.Repetition, repeatedTestCase.TestCase)
				}

				// Add the current test case to the repetition
				repeatedTestCase.Repetition = append(repeatedTestCase.Repetition, testCase)

				// Calculate the properties of the repeated test case based on the repetitions
				var totalTime float64
				for _, repetition := range repeatedTestCase.Repetition {
					totalTime += repetition.Time
				}

				lastRepetition := repeatedTestCase.Repetition[len(repeatedTestCase.Repetition)-1]

				repeatedTestCase.TestCase = testreport.TestCase{
					Name:      lastRepetition.Name,
					ClassName: lastRepetition.ClassName,
					Time:      totalTime,
					Failure:   lastRepetition.Failure,
					Skipped:   lastRepetition.Skipped,
					Error:     lastRepetition.Error,
					SystemErr: lastRepetition.SystemErr,
				}

				testCases[idx] = repeatedTestCase
			}
		}

		testReportWithRepetition.TestSuites = append(testReportWithRepetition.TestSuites, TestSuite{
			Name:          suite.Name,
			Tests:         suite.Tests,
			RepeatedTests: repeatedTests,
			Failures:      suite.Failures,
			Skipped:       suite.Skipped,
			Errors:        suite.Errors,
			Time:          suite.Time,
			TestCases:     testCases,
		})
	}

	return testReportWithRepetition
}
