package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestParseExecuteResponse(t *testing.T) {

	currentWorkingDir, _ := os.Getwd()

	exampleSucceeded := filepath.Join(currentWorkingDir, "testdata", "example_succeeded.xml")
	exampleFailed := filepath.Join(currentWorkingDir, "testdata", "example_failed.xml")

	var tests = []struct {
		path         string
		status       int
		creationTime string
		identifier   string
		version      string
		title        string
	}{
		{exampleSucceeded,
			0,
			"2016-04-26T09:08:06Z",
			"com.terradue.wps_oozie.process.OozieAbstractAlgorithm",
			"1.0.0",
			"SRTM Digital Elevation Model",
		},
		{exampleFailed,
			1,
			"2016-04-26T09:16:58Z",
			"com.terradue.wps_oozie.process.OozieAbstractAlgorithm",
			"1.0.0",
			"SRTM Digital Elevation Model",
		},
	}

	loc, _ := time.LoadLocation("UTC")

	for _, test := range tests {
		response, err := parseExecuteResponse(test.path)

		if err != nil {
			t.Errorf("parseExecuteResponse failed using %s", test.path)
		}

		if response.Status.ProcessStatus != test.status {
			t.Errorf("parseExecuteResponse returned status %d instead of %d", response.Status.ProcessStatus, test.status)
		}

		parsedCreationTime := response.Status.CreationTime.In(loc).Format(time.RFC3339)

		if parsedCreationTime != test.creationTime {
			t.Errorf("parseExecuteResponse returned creationTime %s instead of %s", parsedCreationTime, test.creationTime)
		}

		if response.Process.Identifier != test.identifier {
			t.Errorf("parseExecuteResponse returned identifier %s instead of %s", response.Process.Identifier, test.identifier)
		}

		if response.Process.Version != test.version {
			t.Errorf("parseExecuteResponse returned creationTime %s instead of %s", response.Process.Version, test.version)
		}

		if response.Process.Title != test.title {
			t.Errorf("parseExecuteResponse returned creationTime %s instead of %s", response.Process.Title, test.title)
		}
	}
}
