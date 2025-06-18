package flakydetector

import (
	"testing"

	"github.com/bitrise-steplib/bitrise-step-detect-flaky-tests/testreport"
	"github.com/stretchr/testify/require"
)

func TestConvertToTestReportWithRepetition(t *testing.T) {
	tests := []struct {
		name       string
		testReport testreport.TestReport
		want       TestReportWithRepetition
	}{
		{
			name: "It groups repeated test cases",
			testReport: testreport.TestReport{
				TestSuites: []testreport.TestSuite{
					{
						Name: "Test Suite 1",
						TestCases: []testreport.TestCase{
							{
								ClassName: "Test Suite 1",
								Name:      "Test Case 1",
							},
							{
								ClassName: "Test Suite 1",
								Name:      "Test Case 2",
							},
							{
								ClassName: "Test Suite 1",
								Name:      "Test Case 2",
							},
							{
								ClassName: "Test Suite 1",
								Name:      "Test Case 2",
							},
							{
								ClassName: "Test Suite 1",
								Name:      "Test Case 3",
							},
						},
					},
				},
			},
			want: TestReportWithRepetition{
				TestSuites: []TestSuite{
					{
						Name:          "Test Suite 1",
						RepeatedTests: 1,
						TestCases: []TestCase{
							{
								TestCase: testreport.TestCase{
									ClassName: "Test Suite 1",
									Name:      "Test Case 1",
								},
							},
							{
								TestCase: testreport.TestCase{
									ClassName: "Test Suite 1",
									Name:      "Test Case 2",
								},
								Repetition: []testreport.TestCase{
									{
										ClassName: "Test Suite 1",
										Name:      "Test Case 2",
									},
									{
										ClassName: "Test Suite 1",
										Name:      "Test Case 2",
									},
									{
										ClassName: "Test Suite 1",
										Name:      "Test Case 2",
									},
								},
							},
							{
								TestCase: testreport.TestCase{
									ClassName: "Test Suite 1",
									Name:      "Test Case 3",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "It calculates properties of repeated test cases",
			testReport: testreport.TestReport{
				TestSuites: []testreport.TestSuite{
					{
						Name: "Test Suite 1",
						TestCases: []testreport.TestCase{
							{
								ClassName: "Test Suite 1",
								Name:      "Test Case 1",
								Time:      1.0,
								Failure: &testreport.Failure{
									Value: "Failure message 1",
								},
							},
							{
								ClassName: "Test Suite 1",
								Name:      "Test Case 1",
								Time:      1.0,
								Failure: &testreport.Failure{
									Value: "Failure message 2",
								},
							},
							{
								ClassName: "Test Suite 1",
								Name:      "Test Case 1",
								Time:      1.0,
								Failure: &testreport.Failure{
									Value: "Failure message 3",
								},
							},
						},
					},
				},
			},
			want: TestReportWithRepetition{
				TestSuites: []TestSuite{
					{
						Name:          "Test Suite 1",
						RepeatedTests: 1,
						TestCases: []TestCase{
							{
								TestCase: testreport.TestCase{
									ClassName: "Test Suite 1",
									Name:      "Test Case 1",
									Time:      3.0,
									Failure: &testreport.Failure{
										Value: "Failure message 3",
									},
								},
								Repetition: []testreport.TestCase{
									{
										ClassName: "Test Suite 1",
										Name:      "Test Case 1",
										Time:      1.0,
										Failure: &testreport.Failure{
											Value: "Failure message 1",
										},
									},
									{
										ClassName: "Test Suite 1",
										Name:      "Test Case 1",
										Time:      1.0,
										Failure: &testreport.Failure{
											Value: "Failure message 2",
										},
									},
									{
										ClassName: "Test Suite 1",
										Name:      "Test Case 1",
										Time:      1.0,
										Failure: &testreport.Failure{
											Value: "Failure message 3",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertToTestReportWithRepetition(tt.testReport)
			require.EqualValues(t, tt.want, got)
		})
	}
}
