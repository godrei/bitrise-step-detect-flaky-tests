package flakydetector

import (
	"github.com/bitrise-steplib/bitrise-step-detect-flaky-tests/testreport"
)

type Detector struct {
}

func NewDetector() *Detector {
	return &Detector{}
}

func (d *Detector) DetectFlakyTests(testReport testreport.TestReport) ([]testreport.TestSuite, error) {
	for _, suite := range testReport.TestSuites {
		testCaseRepetition := make(map[string]int)
		for _, testCase := range suite.TestCases {
			testID := testCase.ClassName + "." + testCase.Name
			_, ok := testCaseRepetition[testID]
			if ok {
				
			}
			testCaseRepetition[testID]++
		}
	}
	return nil, nil
}
