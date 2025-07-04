package testreport

import (
	"encoding/xml"
)

type TestReport struct {
	XMLName    xml.Name    `xml:"testsuites"`
	TestSuites []TestSuite `xml:"testsuite"`
}

type TestSuite struct {
	XMLName   xml.Name   `xml:"testsuite"`
	Name      string     `xml:"name,attr"`
	Tests     int        `xml:"tests,attr"`
	Failures  int        `xml:"failures,attr"`
	Skipped   int        `xml:"skipped,attr"`
	Errors    int        `xml:"errors,attr"`
	Time      float64    `xml:"time,attr"`
	TestCases []TestCase `xml:"testcase"`
}

type TestCase struct {
	XMLName           xml.Name `xml:"testcase"`
	ConfigurationHash string   `xml:"configuration-hash,attr"`
	Name              string   `xml:"name,attr"`
	ClassName         string   `xml:"classname,attr"`
	Time              float64  `xml:"time,attr"`
	Failure           *Failure `xml:"failure,omitempty"`
	Skipped           *Skipped `xml:"skipped,omitempty"`
	Error             *Error   `xml:"error,omitempty"`
	SystemErr         string   `xml:"system-err,omitempty"`
}

type Failure struct {
	XMLName xml.Name `xml:"failure,omitempty"`
	Message string   `xml:"message,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

type Skipped struct {
	XMLName xml.Name `xml:"skipped,omitempty"`
}

type Error struct {
	XMLName xml.Name `xml:"error,omitempty"`
	Message string   `xml:"message,attr,omitempty"`
	Value   string   `xml:",chardata"`
}
